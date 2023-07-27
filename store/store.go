package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/store/models"
)

type ErrNotFound struct {
	Err error
}

func (r *ErrNotFound) Error() string {
	return r.Err.Error()
}

type StatusStore interface {
	Status(ctx context.Context, id uuid.UUID) (*ent.Status, error)
	Statuses(ctx context.Context) (ent.StatusSlice, error)
	CreateStatus(ctx context.Context, status *ent.Status) (*ent.Status, error)
	UpdateStatus(ctx context.Context, status *ent.Status) (*ent.Status, error)
	DeleteStatus(ctx context.Context, id uuid.UUID) (*ent.Status, error)
	PaginatedStatuses(ctx context.Context, orderBy *dtos.StatusAggregationInput, take *int, skip *int) (ent.StatusSlice, error)
	CountStatuses(ctx context.Context) (int, error)
}

type JobOfferStore interface {
	JobOffersWithStatus(statusID []uuid.UUID) ([]*models.JobOffers, error)
}

type MessageStore interface {
	Message(ctx context.Context, id uuid.UUID) (*ent.ContactUsMessage, error)
	PaginatedMessages(ctx context.Context, orderBy *dtos.MessageAggregationInput, take *int, skip *int) (ent.ContactUsMessages, error)
	CountMessages(ctx context.Context) (int, error)
	SendMessage(ctx context.Context, message *ent.ContactUsMessage) (*ent.ContactUsMessage, error)
	UpdateMessage(ctx context.Context, message *ent.ContactUsMessage) (*ent.ContactUsMessage, error)
	DeleteMessage(ctx context.Context, id uuid.UUID) (*ent.ContactUsMessage, error)
}

type CategoryStore interface {
	Category(ctx context.Context, id uuid.UUID) (*ent.Category, error)
	PaginatedCategories(ctx context.Context, orderBy *dtos.CategoryAggregationInput, take *int, skip *int, where *dtos.CategoryWhereInput) (ent.Categories, error)
	CountCategories(ctx context.Context) (int, error)
	CreateCategory(ctx context.Context, category *ent.Category) (*ent.Category, error)
	UpdateCategory(ctx context.Context, category *ent.Category) (*ent.Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (*ent.Category, error)
	ChildCategoriesFor(ctx context.Context, parentIDs []uuid.UUID) (ent.Categories, error)
	ParentCategoriesFor(ctx context.Context, childIDs []uuid.UUID) (ent.Categories, error)
	IsSlugTaken(ctx context.Context, slug string) (bool, error)
}

type DegreeLevelStore interface {
	DegreeLevel(id uuid.UUID) (*models.DegreeLevels, error)
	PaginatedDegreeLevels(orderBy *dtos.DegreeLevelAggregationInput, take *int, skip *int, where *dtos.DegreeLevelWhereInput) ([]*models.DegreeLevels, error)
	CountDegreeLevels() (int, error)
	CreateDegreeLevel(degreeLevel models.DegreeLevels) (*models.DegreeLevels, error)
	UpdateDegreeLevel(degreeLevel models.DegreeLevels) (*models.DegreeLevels, error)
	DeleteDegreeLevel(id uuid.UUID) (*models.DegreeLevels, error)
}

type Store interface {
	StatusStore
	JobOfferStore
	MessageStore
	CategoryStore
	DegreeLevelStore
}
