package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/store"
	"github.com/marianozunino/cc-backend-go/store/models"
)

type categoryService struct {
	Store store.Store
}

var _ CategoryService = &categoryService{}

// CategoriesWithChildrens implements CategoryService
func (c *categoryService) ChildCategoriesFor(parentIDs []uuid.UUID) ([]*dtos.Category, error) {
	categories, err := c.Store.ChildCategoriesFor(parentIDs)
	if err != nil {
		return nil, err
	}

	return c.BuildFromEntities(categories), nil
}

// CategoriesWithParents implements CategoryService.
func (c *categoryService) ParentCategoriesFor(childIDs []uuid.UUID) ([]*dtos.Category, error) {
	categories, err := c.Store.ParentCategoriesFor(childIDs)
	if err != nil {
		return nil, err
	}
	return c.BuildFromEntities(categories), nil
}

// CreateCategory implements CategoryService
func (c *categoryService) CreateCategory(input dtos.CategoryCreateInput) (*dtos.Category, error) {
	slug, err := c.FindUniqueSlug(input.Name)
	if err != nil {
		return nil, fmt.Errorf("error getting count of categories: %w", err)
	}
	model := models.Categories{
		Name:     input.Name,
		Slug:     slug,
		IsRoot:   input.IsRoot,
		ParentID: input.ParentID,
	}

	category, err := c.Store.CreateCategory(model)
	if err != nil {
		return nil, err
	}

	return c.BuildFromEntity(category), nil
}

// DeleteCategory implements CategoryService
func (c *categoryService) DeleteCategory(id uuid.UUID) (*dtos.Category, error) {
	category, err := c.Store.DeleteCategory(id)

	if err != nil {
		return nil, err
	}

	return c.BuildFromEntity(category), nil
}

// GetCategories implements CategoryService
func (c *categoryService) GetCategories() ([]dtos.Category, error) {
	panic("unimplemented")
}

// GetCategory implements CategoryService
func (c *categoryService) GetCategory(id uuid.UUID) (*dtos.Category, error) {
	category, err := c.Store.Category(id)
	if err != nil {
		return nil, err
	}

	return c.BuildFromEntity(category), nil
}

// PaginatedCategories implements CategoryService
func (c *categoryService) PaginatedCategories(orderBy *dtos.CategoryAggregationInput, take *int, skip *int, where *dtos.CategoryWhereInput) (*dtos.PaginatedCategoryResponse, error) {
	data, err := c.Store.PaginatedCategories(orderBy, take, skip, where)
	if err != nil {
		return nil, err
	}

	categories := c.BuildFromEntities(data)

	count, err := c.Store.CountCategories()
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
func (c *categoryService) UpdateCategory(id uuid.UUID, input dtos.CategoryUpdateInput) (*dtos.Category, error) {
	model := models.Categories{
		ID:       id,
		Name:     input.Name,
		IsRoot:   input.IsRoot,
		ParentID: input.ParentID,
	}

	category, err := c.Store.UpdateCategory(model)
	if err != nil {
		return nil, err
	}
	return c.BuildFromEntity(category), nil
}

func (s *categoryService) BuildFromEntity(entity *models.Categories) *dtos.Category {
	dto := &dtos.Category{
		ID:        entity.ID,
		Name:      entity.Name,
		Slug:      entity.Slug,
		IsRoot:    entity.IsRoot,
		ParentID:  entity.ParentID,
		CreatedAt: entity.CreatedAt.Time,
		UpdatedAt: entity.UpdatedAt.Time,
		DeletedAt: getTimeOrNil(entity.DeletedAt),
	}
	return dto
}

func (s *categoryService) BuildFromEntities(entities []*models.Categories) []*dtos.Category {
	dtos := make([]*dtos.Category, len(entities))
	for i, entity := range entities {
		dtos[i] = s.BuildFromEntity(entity)
	}
	return dtos
}

func createSlugForCategory(name string) string {
	return slug.Make(name)
}

func (s *categoryService) FindUniqueSlug(name string) (string, error) {
	initialSlug := createSlugForCategory(name)
	slug := initialSlug

	counter := 1
	maxIterations := 10
	for counter < maxIterations {
		isTaken, err := s.Store.IsSlugTaken(slug)

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
