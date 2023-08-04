package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/dtos/mapper"
	"github.com/marianozunino/jobby/ent"
	"github.com/marianozunino/jobby/store"
)

type postService struct {
	Store store.Store
}

var _ PostService = &postService{}

func createSlugForPost(name string) string {
	return slug.Make(name)
}

func (p *postService) CreatePost(ctx context.Context, input dtos.PostCreateInput, authorID uuid.UUID) (*dtos.Post, error) {

	slug, err := p.FindUniqueSlug(ctx, createSlugForPost(input.Title))
	if err != nil {
		return nil, err
	}

	model := &ent.Post{
		Title:         input.Title,
		Content:       input.Content,
		IsHighlighted: input.IsHighlighted,
		IsPublished:   input.IsPublished,
		PreviewImage:  input.PreviewImage,
		AuthorID:      &authorID,
		Slug:          slug,
	}

	if post, err := p.Store.CreatePost(ctx, model); err != nil {
		return nil, err
	} else {
		return mapper.PostToDto(post), nil
	}

}

func (p *postService) DeletePost(ctx context.Context, id uuid.UUID) (*dtos.Post, error) {
	post, err := p.Store.DeletePost(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.PostToDto(post), nil
}

func (p *postService) UpdatePost(ctx context.Context, id uuid.UUID, input dtos.PostUpdateInput) (*dtos.Post, error) {

	slug, err := p.FindUniqueSlug(ctx, createSlugForPost(input.Title))
	if err != nil {
		return nil, err
	}

	model := &ent.Post{
		ID:            id,
		Title:         input.Title,
		Content:       input.Content,
		IsHighlighted: input.IsHighlighted,
		PreviewImage:  input.PreviewImage,
		AuthorID:      nil,
		Slug:          slug,
	}

	post, err := p.Store.UpdatePost(ctx, model)
	if err != nil {
		return nil, err
	}

	return mapper.PostToDto(post), nil
}

func (p *postService) GetPost(ctx context.Context, id uuid.UUID) (*dtos.Post, error) {

	post, err := p.Store.Post(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.PostToDto(post), nil
}

func (p *postService) GetPosts(ctx context.Context) ([]dtos.Post, error) {
	panic("not implemented") // TODO: Implement

}

func (p *postService) PaginatedPosts(ctx context.Context, orderBy *dtos.PostAggregationInput, take *int, skip *int, where *dtos.PostWhereInput) (*dtos.PaginatedPostResponse, error) {

	data, err := p.Store.PaginatedPosts(ctx, orderBy, take, skip, where)
	if err != nil {
		return nil, err
	}

	posts := mapper.PostsToDtos(data)

	count, err := p.Store.CountPosts(ctx)
	if err != nil {
		return nil, err
	}

	return &dtos.PaginatedPostResponse{
		Edges: posts,
		Total: count,
		Take:  take,
		Skip:  skip,
	}, nil
}

// PublishPost implements PostService.
func (p *postService) PublishPost(ctx context.Context, id uuid.UUID) (*dtos.Post, error) {

	post, err := p.Store.PublishPost(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.PostToDto(post), nil
}

func (s *postService) FindUniqueSlug(ctx context.Context, name string) (string, error) {
	initialSlug := createSlugForPost(name)
	slug := initialSlug

	counter := 1
	maxIterations := 10
	for counter < maxIterations {
		isTaken, err := s.Store.IsPostSlugTaken(ctx, slug)

		if err != nil {
			return "", err
		}

		if !isTaken {
			return slug, nil
		}

		slug = fmt.Sprintf("%s-%d", initialSlug, counter)
		counter++
	}

	return "", fmt.Errorf("could not find unique slug for post %s", name)
}

// PostAuthorFor implements PostService.
func (p *postService) PostAuthorFor(ctx context.Context, postIDs []uuid.UUID) (map[uuid.UUID]*dtos.User, error) {

	postIdWithAuthor, err := p.Store.PostAuthorFor(ctx, postIDs)
	if err != nil {
		return nil, err
	}

	result := make(map[uuid.UUID]*dtos.User)
	for postID, author := range postIdWithAuthor {
		result[postID] = mapper.UserToDto(author)
	}

	return result, nil
}
