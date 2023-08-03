package postgres

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/ent"
	"github.com/marianozunino/jobby/ent/status"
	"github.com/marianozunino/jobby/store"
)

// CreateStatus implements store.StatusStore.
func (ss *StatusStore) CreateStatus(ctx context.Context, status *ent.Status) (*ent.Status, error) {
	return ss.Client.Status.Create().SetName(status.Name).Save(ctx)
}

// DeleteStatus implements store.StatusStore.
func (ss *StatusStore) DeleteStatus(ctx context.Context, id uuid.UUID) (*ent.Status, error) {
	return ss.Client.Status.UpdateOneID(id).SetDeletedAt(time.Now()).Save(ctx)
}

// Status implements store.StatusStore.
func (ss *StatusStore) Status(ctx context.Context, id uuid.UUID) (*ent.Status, error) {
	return ss.Client.Status.Get(ctx, id)
}

// Statuses implements store.StatusStore.
func (ss *StatusStore) Statuses(ctx context.Context) (ent.StatusSlice, error) {
	return ss.Client.Status.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull(status.FieldDeletedAt))
	}).All(ctx)
}

// UpdateStatus implements store.StatusStore.
func (ss *StatusStore) UpdateStatus(ctx context.Context, status *ent.Status) (*ent.Status, error) {
	return ss.Client.Status.
		UpdateOneID(status.ID).
		SetName(status.Name).
		Save(ctx)
}

func (ss *StatusStore) PaginatedStatuses(ctx context.Context, orderBy *dtos.StatusAggregationInput, take *int, skip *int) (ent.StatusSlice, error) {
	query := ss.Client.Status.Query().
		Where(func(s *sql.Selector) {
			s.Where(sql.IsNull(status.FieldDeletedAt))
		})

	order := []status.OrderOption{}

	if orderBy != nil {
		if orderBy.ID != nil {
			if *orderBy.ID == dtos.SortOrderAsc {
				order = append(order, ent.Asc(status.FieldID))
			} else {
				order = append(order, ent.Desc(status.FieldID))
			}
		}
		if orderBy.Name != nil {
			if *orderBy.Name == dtos.SortOrderAsc {
				order = append(order, ent.Asc(status.FieldName))
			} else {
				order = append(order, ent.Desc(status.FieldName))
			}
		}
		if orderBy.CreatedAt != nil {
			if *orderBy.CreatedAt == dtos.SortOrderAsc {
				order = append(order, ent.Asc(status.FieldCreatedAt))
			} else {
				order = append(order, ent.Desc(status.FieldCreatedAt))
			}
		}
		if orderBy.UpdatedAt != nil {
			if *orderBy.UpdatedAt == dtos.SortOrderAsc {
				order = append(order, ent.Asc(status.FieldUpdatedAt))
			} else {
				order = append(order, ent.Desc(status.FieldUpdatedAt))
			}
		}
		if orderBy.DeletedAt != nil {
			if *orderBy.DeletedAt == dtos.SortOrderAsc {
				order = append(order, ent.Asc(status.FieldDeletedAt))
			} else {
				order = append(order, ent.Desc(status.FieldDeletedAt))
			}
		}
	}

	if len(order) == 0 {
		order = append(order, ent.Asc(status.FieldID))
	}

	query = query.Order(order...)

	if take != nil {
		query = query.Limit(*take)
	}
	if skip != nil {
		query = query.Offset(*skip)
	}

	return query.All(ctx)
}

// CountStatuses implements store.StatusStore.
func (ss *StatusStore) CountStatuses(ctx context.Context) (int, error) {
	return ss.Client.Status.Query().Where(status.DeletedAtIsNil()).Count(ctx)
}

// assert that StatusStore implements store.StatusStore
var _ store.StatusStore = &StatusStore{}

type StatusStore struct {
	*ent.Client
}
