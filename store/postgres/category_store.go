package postgres

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/ent/category"
	"github.com/marianozunino/cc-backend-go/store"
)

// assert that StatusStore implements store.StatusStore
var _ store.CategoryStore = &CategoryStore{}

type CategoryStore struct {
	*ent.Client
}

// FindCategoryBySlug implements store.CategoryStore.
func (c *CategoryStore) IsSlugTaken(ctx context.Context, slug string) (bool, error) {
	return c.Client.Category.Query().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("slug", slug))
	}).Exist(ctx)
}

// Category implements store.CategoryStore
func (c *CategoryStore) Category(ctx context.Context, id uuid.UUID) (*ent.Category, error) {
	return c.Client.Category.Query().Where(category.ID(id), category.DeletedAtIsNil()).First(ctx)
}

// CountCategories implements store.CategoryStore
func (c *CategoryStore) CountCategories(ctx context.Context) (int, error) {
	return c.Client.Category.Query().Where(category.DeletedAtIsNil()).Count(ctx)
}

// CreateCategory implements store.CategoryStore
func (c *CategoryStore) CreateCategory(ctx context.Context, category *ent.Category) (*ent.Category, error) {
	return c.Client.Category.Create().
		SetName(category.Name).
		SetSlug(category.Slug).
		SetIsRoot(category.IsRoot).
		SetNillableParentID(category.ParentID).
		Save(ctx)

}

// DeleteCategory implements store.CategoryStore
func (c *CategoryStore) DeleteCategory(ctx context.Context, id uuid.UUID) (*ent.Category, error) {
	return c.Client.Category.UpdateOneID(id).Where(category.DeletedAtIsNil()).SetDeletedAt(time.Now()).Save(ctx)
}

// PaginatedCategories implements store.CategoryStore
func (c *CategoryStore) PaginatedCategories(ctx context.Context, orderBy *dtos.CategoryAggregationInput, take *int, skip *int, where *dtos.CategoryWhereInput) (ent.Categories, error) {
	query := c.Client.Category.Query().Where(category.DeletedAtIsNil())

	if where != nil {
		if where.ID != nil {
			query = query.Where(category.ID(*where.ID))
		}
		if where.Name != nil {
			query = query.Where(category.NameContains(*where.Name))
		}
		if where.Slug != nil {
			query = query.Where(category.SlugContains(*where.Slug))
		}
		if where.IsRoot != nil {
			query = query.Where(category.IsRoot(*where.IsRoot))
		}
	}

	order := []category.OrderOption{}

	if orderBy != nil {
		if orderBy.ID != nil {
			if *orderBy.ID == dtos.SortOrderAsc {
				order = append(order, ent.Asc(category.FieldID))
			} else {
				order = append(order, ent.Desc(category.FieldID))
			}
		}
		if orderBy.Name != nil {
			if *orderBy.Name == dtos.SortOrderAsc {
				order = append(order, ent.Asc(category.FieldName))
			} else {
				order = append(order, ent.Desc(category.FieldName))
			}
		}
		if orderBy.Slug != nil {
			if *orderBy.Slug == dtos.SortOrderAsc {
				order = append(order, ent.Asc(category.FieldSlug))
			} else {
				order = append(order, ent.Desc(category.FieldSlug))
			}
		}
		if orderBy.IsRoot != nil {
			if *orderBy.IsRoot == dtos.SortOrderAsc {
				order = append(order, ent.Asc(category.FieldIsRoot))
			} else {
				order = append(order, ent.Desc(category.FieldIsRoot))
			}
		}
		if orderBy.CreatedAt != nil {
			if *orderBy.CreatedAt == dtos.SortOrderAsc {
				order = append(order, ent.Asc(category.FieldCreatedAt))
			} else {
				order = append(order, ent.Desc(category.FieldCreatedAt))
			}
		}
		if orderBy.UpdatedAt != nil {
			if *orderBy.UpdatedAt == dtos.SortOrderAsc {
				order = append(order, ent.Asc(category.FieldUpdatedAt))
			} else {
				order = append(order, ent.Desc(category.FieldUpdatedAt))
			}
		}
		if orderBy.DeletedAt != nil {
			if *orderBy.DeletedAt == dtos.SortOrderAsc {
				order = append(order, ent.Asc(category.FieldDeletedAt))
			} else {
				order = append(order, ent.Desc(category.FieldDeletedAt))
			}
		}
	}

	if take != nil {
		query = query.Limit(*take)
	}
	if skip != nil {
		query = query.Offset(*skip)
	}

	if len(order) > 0 {
		query = query.Order(order...)
	}

	return query.All(ctx)
}

// UpdateCategory implements store.CategoryStore
func (c *CategoryStore) UpdateCategory(ctx context.Context, category *ent.Category) (*ent.Category, error) {
	return c.Client.Debug().Category.UpdateOneID(category.ID).SetName(category.Name).SetIsRoot(category.IsRoot).SetNillableParentID(category.ParentID).Save(ctx)
}

// CategoriesWithChildrens implements store.CategoryStore
func (c *CategoryStore) ChildCategoriesFor(ctx context.Context, ids []uuid.UUID) (ent.Categories, error) {
	return c.Client.Category.Query().Where(category.ParentIDIn(ids...)).All(ctx)
}

// CategoriesWithParents implements store.CategoryStore.
func (c *CategoryStore) ParentCategoriesFor(ctx context.Context, ids []uuid.UUID) (ent.Categories, error) {
	return c.Client.Category.Query().Where(category.IDIn(ids...)).All(ctx)
}
