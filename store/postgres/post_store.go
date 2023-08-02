package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/ent/post"
	"github.com/marianozunino/cc-backend-go/store"
)

// assert that StatusStore implements store.StatusStore
var _ store.PostStore = &PostStore{}

type PostStore struct {
	*ent.Client
}

func (p *PostStore) Post(ctx context.Context, id uuid.UUID) (*ent.Post, error) {
	post, err := p.Client.Post.Query().Where(post.IDEQ(id)).First(ctx)

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *PostStore) PaginatedPosts(ctx context.Context, orderBy *dtos.PostAggregationInput, take *int, skip *int, where *dtos.PostWhereInput) (ent.Posts, error) {
	query := p.Client.Post.Query().Where(post.DeletedAtIsNil())

	rawOrderOptions := createRawOrderOptions(orderBy)

	var orderOptions []post.OrderOption = make([]post.OrderOption, len(rawOrderOptions))
	for i, v := range rawOrderOptions {
		orderOptions[i] = v.(post.OrderOption)
	}

	if len(rawOrderOptions) == 0 {
		rawOrderOptions = append(rawOrderOptions, ent.Asc(post.FieldID))
	}

	query = query.Order(orderOptions...)

	if take != nil {
		query = query.Limit(*take)
	}
	if skip != nil {
		query = query.Offset(*skip)
	}

	return query.All(ctx)
}

func (p *PostStore) CountPosts(ctx context.Context) (int, error) {
	return p.Client.Post.Query().Where(post.DeletedAtIsNil()).Count(ctx)
}

func (p *PostStore) CreatePost(ctx context.Context, post *ent.Post) (*ent.Post, error) {
	return p.Client.Post.Create().
		SetTitle(post.Title).
		SetSlug(post.Slug).
		SetContent(post.Content).
		SetIsHighlighted(post.IsHighlighted).
		SetIsPublished(post.IsPublished).
		SetNillablePreviewImage(post.PreviewImage).
		SetNillableAuthorID(post.AuthorID).
		Save(ctx)
}

func (p *PostStore) UpdatePost(ctx context.Context, post *ent.Post) (*ent.Post, error) {
	return p.Client.Post.UpdateOneID(post.ID).
		SetTitle(post.Title).
		SetContent(post.Content).
		SetNillablePreviewImage(post.PreviewImage).
		SetIsPublished(post.IsPublished).
		SetOrClearAuthorID(post.AuthorID).
		SetIsHighlighted(post.IsHighlighted).
		Save(ctx)
}

func (p *PostStore) DeletePost(ctx context.Context, id uuid.UUID) (*ent.Post, error) {
	return p.Client.Post.UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(ctx)
}

// IsPostSlugTaken implements store.PostStore.
func (p *PostStore) IsPostSlugTaken(ctx context.Context, slug string) (bool, error) {
	return p.Client.Post.Query().Where(post.SlugEQ(slug)).Exist(ctx)
}

// PublishPost implements store.PostStore.
func (p *PostStore) PublishPost(ctx context.Context, id uuid.UUID) (*ent.Post, error) {
	return p.Client.Post.UpdateOneID(id).
		SetIsPublished(true).
		SetPublishedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
}
