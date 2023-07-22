package store

import (
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/store/models"
)

type ErrNotFound struct {
	Err error
}

func (r *ErrNotFound) Error() string {
	return r.Err.Error()
}

type StatusStore interface {
	Status(id uuid.UUID) (*models.Status, error)
	Statuses() ([]*models.Status, error)
	CreateStatus(status models.Status) (*models.Status, error)
	UpdateStatus(status models.Status) (*models.Status, error)
	DeleteStatus(id uuid.UUID) (*models.Status, error)
	PaginatedStatuses(orderBy *dtos.StatusAggregationInput, take *int, skip *int) ([]*models.Status, error)
	CountStatuses() (int, error)
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

type Store interface {
	StatusStore
	JobOfferStore
	MessageStore
	CategoryStore
}
