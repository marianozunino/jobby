package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/ent"
	"github.com/marianozunino/jobby/ent/post"
	"github.com/marianozunino/jobby/ent/postcategory"
	"github.com/marianozunino/jobby/store"
)

// assert that StatusStore implements store.StatusStore
var _ store.PostCategoryStore = &PostCategoryStore{}

type PostCategoryStore struct {
	*ent.Client
}

// Category implements store.PostCategoryStore
func (c *PostCategoryStore) PostCategory(ctx context.Context, id uuid.UUID) (*ent.PostCategory, error) {
	return c.Client.PostCategory.Query().Where(postcategory.ID(id), postcategory.DeletedAtIsNil()).First(ctx)
}

// CountCategories implements store.PostCategoryStore
func (c *PostCategoryStore) CountPostCategories(ctx context.Context) (int, error) {
	return c.Client.PostCategory.Query().Where(postcategory.DeletedAtIsNil()).Count(ctx)
}

// CreateCategory implements store.PostCategoryStore
func (c *PostCategoryStore) CreatePostCategory(ctx context.Context, category *ent.PostCategory) (*ent.PostCategory, error) {
	return c.Client.PostCategory.Create().
		SetName(category.Name).
		Save(ctx)

}

// DeleteCategory implements store.PostCategoryStore
func (c *PostCategoryStore) DeletePostCategory(ctx context.Context, id uuid.UUID) (*ent.PostCategory, error) {
	postcategory, err := c.Client.PostCategory.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := c.Client.PostCategory.DeleteOne(postcategory).Exec(ctx); err != nil {
		return nil, err
	}
	return postcategory, nil

}

// PaginatedCategories implements store.PostCategoryStore
func (c *PostCategoryStore) PaginatedPostCategories(ctx context.Context, orderBy *dtos.PostCategoryAggregationInput, take *int, skip *int, where *dtos.PostCategoryWhereInput) (ent.PostCategories, error) {
	query := c.Client.PostCategory.Query().Where(postcategory.DeletedAtIsNil())

	if where != nil {
		if where.ID != nil {
			query = query.Where(postcategory.ID(*where.ID))
		}
		if where.Name != nil {
			query = query.Where(postcategory.NameContains(*where.Name))
		}
	}

	order := []postcategory.OrderOption{}

	if orderBy != nil {
		if orderBy.ID != nil {
			if *orderBy.ID == dtos.SortOrderAsc {
				order = append(order, ent.Asc(postcategory.FieldID))
			} else {
				order = append(order, ent.Desc(postcategory.FieldID))
			}
		}
		if orderBy.Name != nil {
			if *orderBy.Name == dtos.SortOrderAsc {
				order = append(order, ent.Asc(postcategory.FieldName))
			} else {
				order = append(order, ent.Desc(postcategory.FieldName))
			}
		}
		if orderBy.CreatedAt != nil {
			if *orderBy.CreatedAt == dtos.SortOrderAsc {
				order = append(order, ent.Asc(postcategory.FieldCreatedAt))
			} else {
				order = append(order, ent.Desc(postcategory.FieldCreatedAt))
			}
		}
		if orderBy.UpdatedAt != nil {
			if *orderBy.UpdatedAt == dtos.SortOrderAsc {
				order = append(order, ent.Asc(postcategory.FieldUpdatedAt))
			} else {
				order = append(order, ent.Desc(postcategory.FieldUpdatedAt))
			}
		}
		if orderBy.DeletedAt != nil {
			if *orderBy.DeletedAt == dtos.SortOrderAsc {
				order = append(order, ent.Asc(postcategory.FieldDeletedAt))
			} else {
				order = append(order, ent.Desc(postcategory.FieldDeletedAt))
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

// UpdatePostCategory implements store.PostCategoryStore
func (c *PostCategoryStore) UpdatePostCategory(ctx context.Context, category *ent.PostCategory) (*ent.PostCategory, error) {
	return c.Client.PostCategory.UpdateOneID(category.ID).SetName(category.Name).Save(ctx)
}

// PostCategoriesFor implements store.PostStore.

func (p *PostStore) PostCategoriesFor(ctx context.Context, postIDs []uuid.UUID) (map[uuid.UUID]ent.PostCategories, error) {
	posts, err := p.Client.Debug().Post.Query().Where(post.IDIn(postIDs...)).WithPostCategory().All(ctx)

	if err != nil {
		return nil, err
	}

	result := make(map[uuid.UUID]ent.PostCategories)

	for _, post := range posts {
		result[post.ID] = post.Edges.PostCategory
	}

	return result, nil
}
