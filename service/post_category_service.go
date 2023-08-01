package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/store"
)

type postCategoryService struct {
	Store store.Store
}

// assert that PostCategoryService implements store.PostCategoryService
var _ PostCategoryService = &postCategoryService{}

func (s *postCategoryService) BuildFromEntity(entity *ent.PostCategory) *dtos.PostCategory {
	dto := &dtos.PostCategory{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
	return dto
}

func (s *postCategoryService) BuildFromEntities(entities ent.PostCategories) []*dtos.PostCategory {
	dtos := make([]*dtos.PostCategory, len(entities))
	for i, entity := range entities {
		dtos[i] = s.BuildFromEntity(entity)
	}
	return dtos
}

func (p *postCategoryService) CreatePostCategory(ctx context.Context, input dtos.PostCategoryCreateInput) (*dtos.PostCategory, error) {
	model := &ent.PostCategory{
		Name: input.Name,
	}

	category, err := p.Store.CreatePostCategory(ctx, model)
	if err != nil {
		return nil, err
	}

	return p.BuildFromEntity(category), nil
}

func (p *postCategoryService) DeletePostCategory(ctx context.Context, id uuid.UUID) (*dtos.PostCategory, error) {
	category, err := p.Store.DeletePostCategory(ctx, id)

	if err != nil {
		return nil, err
	}

	return p.BuildFromEntity(category), nil
}

func (p *postCategoryService) UpdatePostCategory(ctx context.Context, id uuid.UUID, input dtos.PostCategoryUpdateInput) (*dtos.PostCategory, error) {
	model := &ent.PostCategory{
		ID:   id,
		Name: input.Name,
	}

	category, err := p.Store.UpdatePostCategory(ctx, model)
	if err != nil {
		return nil, err
	}
	return p.BuildFromEntity(category), nil
}

func (p *postCategoryService) GetPostCategory(ctx context.Context, id uuid.UUID) (*dtos.PostCategory, error) {
	category, err := p.Store.PostCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	return p.BuildFromEntity(category), nil
}

func (p *postCategoryService) GetPostCategories(ctx context.Context) ([]dtos.PostCategory, error) {
	panic("not implemented") // TODO: Implement
}

func (p *postCategoryService) PaginatedPostCategories(ctx context.Context, orderBy *dtos.PostCategoryAggregationInput, take *int, skip *int, where *dtos.PostCategoryWhereInput) (*dtos.PaginatedPostCategoryResponse, error) {
	data, err := p.Store.PaginatedPostCategories(ctx, orderBy, take, skip, where)
	if err != nil {
		return nil, err
	}

	categories := p.BuildFromEntities(data)

	count, err := p.Store.CountPostCategories(ctx)
	if err != nil {
		return nil, err
	}

	return &dtos.PaginatedPostCategoryResponse{
		Edges: categories,
		Total: count,
		Take:  take,
		Skip:  skip,
	}, nil
}
