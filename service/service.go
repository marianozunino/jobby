package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/store"
)

func getTimeOrNil(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

func getStringOrNil(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}

type StatusService interface {
	CreateStatus(ctx context.Context, input dtos.StatusCreateInput) (*dtos.Status, error)
	DeleteStatus(ctx context.Context, id string) (*dtos.Status, error)
	UpdateStatus(ctx context.Context, id string, input dtos.StatusUpdateInput) (*dtos.Status, error)
	GetStatus(ctx context.Context, id string) (*dtos.Status, error)
	GetStatuses(ctx context.Context) ([]dtos.Status, error)
	PaginatedStatuses(ctx context.Context, orderBy *dtos.StatusAggregationInput, take *int, skip *int) (*dtos.PaginatedStatusResponse, error)
}

type JobOfferService interface {
	JobOffersWithStatus(statusID []uuid.UUID) ([]*dtos.JobOffer, error)
}

type MessageService interface {
	SendMessage(ctx context.Context, input dtos.MessageCreateInput) (*dtos.Message, error)
	UpdateMessage(ctx context.Context, id string, input dtos.MessageUpdateInput) (*dtos.Message, error)
	DeleteMessage(ctx context.Context, id string) (*dtos.Message, error)
	PaginatedMessages(ctx context.Context, orderBy *dtos.MessageAggregationInput, take *int, skip *int) (*dtos.PaginatedMessageResponse, error)
	GetMessage(ctx context.Context, id string) (*dtos.Message, error)
}

type CategoryService interface {
	CreateCategory(ctx context.Context, input dtos.CategoryCreateInput) (*dtos.Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (*dtos.Category, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, input dtos.CategoryUpdateInput) (*dtos.Category, error)
	GetCategory(ctx context.Context, id uuid.UUID) (*dtos.Category, error)
	GetCategories(ctx context.Context) ([]dtos.Category, error)
	PaginatedCategories(ctx context.Context, orderBy *dtos.CategoryAggregationInput, take *int, skip *int, where *dtos.CategoryWhereInput) (*dtos.PaginatedCategoryResponse, error)
	ChildCategoriesFor(ctx context.Context, parentIDs []uuid.UUID) ([]*dtos.Category, error)
	ParentCategoriesFor(ctx context.Context, childIDs []uuid.UUID) ([]*dtos.Category, error)
}

type DegreeLevelService interface {
	GetDegreeLevel(id uuid.UUID) (*dtos.DegreeLevel, error)
	PaginatedDegreeLevels(orderBy *dtos.DegreeLevelAggregationInput, take *int, skip *int, where *dtos.DegreeLevelWhereInput) (*dtos.PaginatedDegreeLevelResponse, error)
	CreateDegreeLevel(degreeLevel dtos.DegreeLevelCreateInput) (*dtos.DegreeLevel, error)
	UpdateDegreeLevel(id uuid.UUID, degreeLevel dtos.DegreeLevelUpdateInput) (*dtos.DegreeLevel, error)
	DeleteDegreeLevel(id uuid.UUID) (*dtos.DegreeLevel, error)
}

type Service interface {
	StatusService
	JobOfferService
	MessageService
	CategoryService
	DegreeLevelService
}

type service struct {
	StatusService
	JobOfferService
	MessageService
	CategoryService
	DegreeLevelService
}

func NewService(store store.Store) Service {
	return &service{
		StatusService: &statusService{
			Store: store,
		},
		JobOfferService: &jobOfferService{
			Store: store,
		},
		MessageService: &messageService{
			Store: store,
		},
		CategoryService: &categoryService{
			Store: store,
		},
		DegreeLevelService: &degreeLevelService{
			Store: store,
		},
	}
}
