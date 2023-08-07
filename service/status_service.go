package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/ent"
	"github.com/marianozunino/jobby/store"
)

type statusService struct {
	Store store.Store
}

var _ StatusService = &statusService{}

// GetStatuses implements StatusService.
func (*statusService) GetStatuses(ctx context.Context) ([]dtos.Status, error) {
	panic("unimplemented")
}

// DeleteStatus implements StatusService.
func (s *statusService) DeleteStatus(ctx context.Context, id string) (*dtos.Status, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	status, err := s.Store.DeleteStatus(ctx, parsedId)

	if err != nil {
		return nil, err
	}

	return s.BuildFromEntity(status), nil
}

// CreateStatus implements StatusService.
func (s *statusService) CreateStatus(ctx context.Context,
	input dtos.StatusCreateInput) (*dtos.Status, error) {

	model := &ent.Status{
		Name: input.Name,
	}

	status, err := s.Store.CreateStatus(ctx, model)
	if err != nil {
		return nil, err
	}

	return s.BuildFromEntity(status), nil
}

// GetStatus implements StatusService.
func (s *statusService) GetStatus(
	ctx context.Context,
	id string) (*dtos.Status, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	status, err := s.Store.Status(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	return s.BuildFromEntity(status), nil
}

// PaginatedStatuses implements StatusService.
func (s *statusService) PaginatedStatuses(
	ctx context.Context,
	orderBy *dtos.StatusAggregationInput, take *int, skip *int) (*dtos.PaginatedStatusResponse, error) {
	data, err := s.Store.PaginatedStatuses(ctx, orderBy, take, skip)
	if err != nil {
		return nil, err
	}

	statuses := s.BuildFromEntities(data)

	count, err := s.Store.CountStatuses(ctx)
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
func (s *statusService) UpdateStatus(
	ctx context.Context,
	id string, input dtos.StatusUpdateInput) (*dtos.Status, error) {
	parsedId, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	model := &ent.Status{
		ID:   parsedId,
		Name: input.Name,
	}

	status, err := s.Store.UpdateStatus(ctx, model)
	if err != nil {
		return nil, err
	}
	return s.BuildFromEntity(status), nil

}

func (s *statusService) BuildFromEntity(entity *ent.Status) *dtos.Status {
	dto := &dtos.Status{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
	return dto
}

func (s *statusService) BuildFromEntities(entities ent.StatusSlice) []*dtos.Status {
	dtos := make([]*dtos.Status, len(entities))
	for i, entity := range entities {
		dtos[i] = s.BuildFromEntity(entity)
	}
	return dtos
}

// StatusForJobOffers implements StatusService.
func (s *statusService) StatusForJobOffers(ctx context.Context, jobOfferIDs []uuid.UUID) (map[uuid.UUID]*dtos.Status, error) {
	statuses, err := s.Store.StatusForJobOffers(ctx, jobOfferIDs)
	if err != nil {
		return nil, err
	}

	result := make(map[uuid.UUID]*dtos.Status)
	for jobOfferID, status := range statuses {
		result[jobOfferID] = s.BuildFromEntity(status)
	}

	return result, nil
}
