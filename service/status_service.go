package service

import (
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/store"
	"github.com/marianozunino/cc-backend-go/store/models"
)

type statusService struct {
	Store store.Store
}

// GetStatuses implements StatusService.
func (*statusService) GetStatuses() ([]dtos.Status, error) {
	panic("unimplemented")
}

// DeleteStatus implements StatusService.
func (s *statusService) DeleteStatus(id string) (*dtos.Status, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	status, err := s.Store.DeleteStatus(parsedId)

	if err != nil {
		return nil, err
	}

	return s.BuildFromEntity(status), nil
}

// CreateStatus implements StatusService.
func (s *statusService) CreateStatus(input dtos.StatusCreateInput) (*dtos.Status, error) {

	model := models.Status{
		Name: input.Name,
	}

	status, err := s.Store.CreateStatus(model)
	if err != nil {
		return nil, err
	}

	return s.BuildFromEntity(status), nil
}

// GetStatus implements StatusService.
func (s *statusService) GetStatus(id string) (*dtos.Status, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	status, err := s.Store.Status(parsedId)
	if err != nil {
		return nil, err
	}

	return s.BuildFromEntity(status), nil
}

// PaginatedStatuses implements StatusService.
func (s *statusService) PaginatedStatuses(orderBy *dtos.StatusAggregationInput, take *int, skip *int) (*dtos.PaginatedStatusResponse, error) {
	data, err := s.Store.PaginatedStatuses(orderBy, take, skip)
	if err != nil {
		return nil, err
	}

	statuses := s.BuildFromEntities(data)

	count, err := s.Store.CountStatuses()
	if err != nil {
		return nil, err
	}

	return &dtos.PaginatedStatusResponse{
		Edges: statuses,
		Total: count,
		Take:  take,
		Skip:  skip,
	}, nil

}

// UpdateStatus implements StatusService.
func (s *statusService) UpdateStatus(id string, input dtos.StatusUpdateInput) (*dtos.Status, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	model := models.Status{
		ID:   parsedId,
		Name: input.Name,
	}

	status, err := s.Store.UpdateStatus(model)
	if err != nil {
		return nil, err
	}
	return s.BuildFromEntity(status), nil

}

func (s *statusService) BuildFromEntity(entity *models.Status) *dtos.Status {
	dto := &dtos.Status{
		ID:        entity.ID.String(),
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt.Time,
		UpdatedAt: entity.UpdatedAt.Time,
		DeletedAt: getTimeOrNil(entity.DeletedAt),
	}
	return dto
}

func (s *statusService) BuildFromEntities(entities []*models.Status) []dtos.Status {
	dtos := make([]dtos.Status, len(entities))
	for i, entity := range entities {
		dtos[i] = *s.BuildFromEntity(entity)
	}
	return dtos
}

var _ StatusService = &statusService{}
