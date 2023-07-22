package postgres

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/store"
	"github.com/marianozunino/cc-backend-go/store/models"
)

// assert that StatusStore implements store.StatusStore
var _ store.CategoryStore = &CategoryStore{}

type CategoryStore struct {
	*sqlx.DB
}

// FindCategoryBySlug implements store.CategoryStore.
func (c *CategoryStore) IsSlugTaken(slug string) (bool, error) {
	var count int
	if err := c.DB.Get(&count, "SELECT COUNT(*) FROM categories WHERE slug=$1", slug); err != nil {
		return false, fmt.Errorf("error getting count of categories: %w", err)
	}
	return count > 0, nil

}

// Category implements store.CategoryStore
func (c *CategoryStore) Category(id uuid.UUID) (*models.Categories, error) {
	query := "SELECT * FROM categories WHERE id=$1"
	category := &models.Categories{}
	if err := c.DB.Get(category, query, id); err != nil {
		return nil, fmt.Errorf("error getting category: %w", err)
	}
	return category, nil
}

// CountCategories implements store.CategoryStore
func (c *CategoryStore) CountCategories() (int, error) {
	var count int
	if err := c.DB.Get(&count, "SELECT COUNT(*) FROM categories"); err != nil {
		return 0, fmt.Errorf("error getting count of categories: %w", err)
	}
	return count, nil
}

// CreateCategory implements store.CategoryStore
func (c *CategoryStore) CreateCategory(category models.Categories) (*models.Categories, error) {
	var output *models.Categories = &models.Categories{}
	if err := c.DB.Get(
		output,
		"INSERT INTO categories (name, slug, is_root, parent_id) VALUES ($1, $2, $3, $4) RETURNING *",
		category.Name,
		category.Slug,
		category.IsRoot,
		category.ParentID,
	); err != nil {
		return nil, fmt.Errorf("error creating category: %w", err)
	}
	return output, nil
}

// DeleteCategory implements store.CategoryStore
func (c *CategoryStore) DeleteCategory(id uuid.UUID) (*models.Categories, error) {
	category := &models.Categories{}
	if err := c.DB.Get(category, "DELETE FROM categories WHERE id=$1 RETURNING *", id); err != nil {
		return nil, fmt.Errorf("error deleting category: %w", err)
	}
	return category, nil
}

// PaginatedCategories implements store.CategoryStore

func (c *CategoryStore) PaginatedCategories(orderBy *dtos.CategoryAggregationInput, take *int, skip *int, where *dtos.CategoryWhereInput) ([]*models.Categories, error) {
	query := "SELECT * FROM categories"

	if where != nil {
		query += " WHERE deleted_at IS NULL AND "
		var whereFields []string
		switch {
		case where.ID != nil:
			whereFields = append(whereFields, fmt.Sprintf("id = '%s'", where.ID))
		case where.Name != nil:
			whereFields = append(whereFields, fmt.Sprintf("name ilike '%%%s%%'", *where.Name))
		case where.Slug != nil:
			whereFields = append(whereFields, fmt.Sprintf("slug ilike '%%%s%%'", *where.Slug))
		case where.IsRoot != nil:
			whereFields = append(whereFields, fmt.Sprintf("is_root = %t", *where.IsRoot))
		}
		query += strings.Join(whereFields, " AND ")
	}

	if orderBy != nil {
		query += " ORDER BY "
		var sortFields []string

		if orderBy.ID != nil {
			sortFields = append(sortFields, fmt.Sprintf("id %s", orderBy.ID))
		}
		if orderBy.Name != nil {
			sortFields = append(sortFields, fmt.Sprintf("name %s", orderBy.Name))
		}
		if orderBy.Slug != nil {
			sortFields = append(sortFields, fmt.Sprintf("slug %s", orderBy.Slug))
		}
		if orderBy.IsRoot != nil {
			sortFields = append(sortFields, fmt.Sprintf("is_root %s", orderBy.IsRoot))
		}
		if orderBy.CreatedAt != nil {
			sortFields = append(sortFields, fmt.Sprintf("created_at %s", orderBy.CreatedAt))
		}
		if orderBy.UpdatedAt != nil {
			sortFields = append(sortFields, fmt.Sprintf("updated_at %s", orderBy.UpdatedAt))
		}
		if orderBy.DeletedAt != nil {
			sortFields = append(sortFields, fmt.Sprintf("deleted_at %s", orderBy.DeletedAt))
		}
		query += strings.Join(sortFields, ", ")
	}

	if take != nil {
		query += fmt.Sprintf(" LIMIT %d", *take)
	}
	if skip != nil {
		query += fmt.Sprintf(" OFFSET %d", *skip)
	}

	categories := []*models.Categories{}

	if err := c.DB.Select(&categories, query); err != nil {
		return nil, fmt.Errorf("error getting paginated categories: %w", err)
	}

	return categories, nil
}

// UpdateCategory implements store.CategoryStore
func (c *CategoryStore) UpdateCategory(category models.Categories) (*models.Categories, error) {
	var output *models.Categories = &models.Categories{}
	if err := c.DB.Get(
		output,
		"UPDATE categories SET name=$1, slug=$2, is_root=$3, parent_id=$4 WHERE id=$5 RETURNING *",
		category.Name,
		category.Slug,
		category.IsRoot,
		category.ParentID,
		category.ID,
	); err != nil {
		return nil, fmt.Errorf("error updating category: %w", err)
	}
	return output, nil

}

// CategoriesWithChildrens implements store.CategoryStore
func (c *CategoryStore) ChildCategoriesFor(parentIDs []uuid.UUID) ([]*models.Categories, error) {
	query, args, err := sqlx.In("SELECT * FROM categories WHERE parent_id IN (?)", parentIDs)

	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}
	query = c.Rebind(query)

	var categories []*models.Categories

	err = c.Select(&categories, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying job offers: %w", err)
	}

	return categories, nil
}

// CategoriesWithParents implements store.CategoryStore.
func (c *CategoryStore) ParentCategoriesFor(childIDs []uuid.UUID) ([]*models.Categories, error) {
	query, args, err := sqlx.In("SELECT * FROM categories WHERE id IN (?)", childIDs)

	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}
	query = c.Rebind(query)

	var categories []*models.Categories

	err = c.Select(&categories, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying job offers: %w", err)
	}

	return categories, nil
}
