package service

import (
	"database/sql"
	"time"

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
	CreateStatus(input dtos.StatusCreateInput) (*dtos.Status, error)
	DeleteStatus(id string) (*dtos.Status, error)
	UpdateStatus(id string, input dtos.StatusUpdateInput) (*dtos.Status, error)
	GetStatus(id string) (*dtos.Status, error)
	GetStatuses() ([]dtos.Status, error)
	PaginatedStatuses(orderBy *dtos.StatusAggregationInput, take *int, skip *int) (*dtos.PaginatedStatusResponse, error)
}

type JobOfferService interface {
	JobOffersWithStatus(statusID string) ([]*dtos.JobOffer, error)
}

type MessageService interface {
	SendMessage(input dtos.MessageCreateInput) (*dtos.Message, error)
	UpdateMessage(id string, input dtos.MessageUpdateInput) (*dtos.Message, error)
	DeleteMessage(id string) (*dtos.Message, error)
	PaginatedMessages(orderBy *dtos.MessageAggregationInput, take *int, skip *int) (*dtos.PaginatedMessageResponse, error)
	GetMessage(id string) (*dtos.Message, error)
}

type Service interface {
	StatusService
	JobOfferService
	MessageService
}

type service struct {
	StatusService
	JobOfferService
	MessageService
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
	}
}
