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
	Message(id uuid.UUID) (*models.ContactUsMessages, error)
	PaginatedMessages(orderBy *dtos.MessageAggregationInput, take *int, skip *int) ([]*models.ContactUsMessages, error)
	CountMessages() (int, error)
	SendMessage(message models.ContactUsMessages) (*models.ContactUsMessages, error)
	DeleteMessage(id uuid.UUID) (*models.ContactUsMessages, error)
	UpdateMessage(message models.ContactUsMessages) (*models.ContactUsMessages, error)
}

type CategoryStore interface {
	Category(id uuid.UUID) (*models.Categories, error)
	PaginatedCategories(orderBy *dtos.CategoryAggregationInput, take *int, skip *int, where *dtos.CategoryWhereInput) ([]*models.Categories, error)
	CountCategories() (int, error)
	CreateCategory(category models.Categories) (*models.Categories, error)
	UpdateCategory(category models.Categories) (*models.Categories, error)
	DeleteCategory(id uuid.UUID) (*models.Categories, error)
	ChildCategoriesFor(parentIDs []uuid.UUID) ([]*models.Categories, error)
	ParentCategoriesFor(childIDs []uuid.UUID) ([]*models.Categories, error)
	IsSlugTaken(slug string) (bool, error)
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
