package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/store"
)

type categoryService struct {
	Store store.Store
}

var _ CategoryService = &categoryService{}

// CategoriesWithChildrens implements CategoryService
func (c *categoryService) ChildCategoriesFor(ctx context.Context, parentIDs []uuid.UUID) ([]*dtos.Category, error) {
	categories, err := c.Store.ChildCategoriesFor(ctx, parentIDs)
	if err != nil {
		return nil, err
	}

	return c.BuildFromEntities(categories), nil
}

// CategoriesWithParents implements CategoryService.
func (c *categoryService) ParentCategoriesFor(ctx context.Context, childIDs []uuid.UUID) ([]*dtos.Category, error) {
	categories, err := c.Store.ParentCategoriesFor(ctx, childIDs)
	if err != nil {
		return nil, err
	}
	return c.BuildFromEntities(categories), nil
}

// CreateCategory implements CategoryService
func (c *categoryService) CreateCategory(ctx context.Context, input dtos.CategoryCreateInput) (*dtos.Category, error) {
	slug, err := c.FindUniqueSlug(ctx, input.Name)
	if err != nil {
		return nil, fmt.Errorf("error getting count of categories: %w", err)
	}

	model := &ent.Category{
		Name:     input.Name,
		Slug:     slug,
		IsRoot:   input.IsRoot,
		ParentID: input.ParentID,
	}

	category, err := c.Store.CreateCategory(ctx, model)
	if err != nil {
		return nil, err
	}

	return c.BuildFromEntity(category), nil
}

// DeleteCategory implements CategoryService
func (c *categoryService) DeleteCategory(ctx context.Context, id uuid.UUID) (*dtos.Category, error) {
	category, err := c.Store.DeleteCategory(ctx, id)

	if err != nil {
		return nil, err
	}

	return c.BuildFromEntity(category), nil
}

// GetCategories implements CategoryService
func (c *categoryService) GetCategories(ctx context.Context) ([]dtos.Category, error) {
	panic("unimplemented")
}

// GetCategory implements CategoryService
func (c *categoryService) GetCategory(ctx context.Context, id uuid.UUID) (*dtos.Category, error) {
	category, err := c.Store.Category(ctx, id)
	if err != nil {
		return nil, err
	}

	return c.BuildFromEntity(category), nil
}

// PaginatedCategories implements CategoryService
func (c *categoryService) PaginatedCategories(ctx context.Context, orderBy *dtos.CategoryAggregationInput, take *int, skip *int, where *dtos.CategoryWhereInput) (*dtos.PaginatedCategoryResponse, error) {
	data, err := c.Store.PaginatedCategories(ctx, orderBy, take, skip, where)
	if err != nil {
		return nil, err
	}

	categories := c.BuildFromEntities(data)

	count, err := c.Store.CountCategories(ctx)
	if err != nil {
		return nil, err
	}

	return &dtos.PaginatedCategoryResponse{
		Edges: categories,
		Total: count,
		Take:  take,
		Skip:  skip,
	}, nil
}

// UpdateCategory implements CategoryService
func (c *categoryService) UpdateCategory(ctx context.Context, id uuid.UUID, input dtos.CategoryUpdateInput) (*dtos.Category, error) {
	// if isRoot is true, then parentID should be null

	if input.IsRoot && input.ParentID != nil {
		return nil, fmt.Errorf("parentID should be null if isRoot is true")
	}

	if !input.IsRoot && input.ParentID == nil {
		return nil, fmt.Errorf("parentID should not be null if isRoot is false")
	}

	if input.ParentID != nil {
		if *input.ParentID == id {
			return nil, fmt.Errorf("parentID should not be equal to id")
		}

		parent, err := c.Store.Category(ctx, *input.ParentID)

		if err != nil {
			return nil, err
		}

		if !parent.IsRoot {
			return nil, fmt.Errorf("parentID should be a root category")
		}

	}

	model := &ent.Category{
		ID:       id,
		Name:     input.Name,
		IsRoot:   input.IsRoot,
		ParentID: input.ParentID,
	}

	category, err := c.Store.UpdateCategory(ctx, model)
	if err != nil {
		return nil, err
	}
	return c.BuildFromEntity(category), nil
}

func (s *categoryService) BuildFromEntity(entity *ent.Category) *dtos.Category {
	dto := &dtos.Category{
		ID:        entity.ID,
		Name:      entity.Name,
		Slug:      entity.Slug,
		IsRoot:    entity.IsRoot,
		ParentID:  entity.ParentID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
	return dto
}

func (s *categoryService) BuildFromEntities(entities ent.Categories) []*dtos.Category {
	dtos := make([]*dtos.Category, len(entities))
	for i, entity := range entities {
		dtos[i] = s.BuildFromEntity(entity)
	}
	return dtos
}

func createSlugForCategory(name string) string {
	return slug.Make(name)
}

func (s *categoryService) FindUniqueSlug(ctx context.Context, name string) (string, error) {
	initialSlug := createSlugForCategory(name)
	slug := initialSlug

	counter := 1
	maxIterations := 10
	for counter < maxIterations {
		isTaken, err := s.Store.IsSlugTaken(ctx, slug)

		if err != nil {
			return "", err
		}

		if !isTaken {
			return slug, nil
		}

		slug = fmt.Sprintf("%s-%d", initialSlug, counter)
		counter++
	}

	return "", fmt.Errorf("could not find unique slug for category %s", name)
}
