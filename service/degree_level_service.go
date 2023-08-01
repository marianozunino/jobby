package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/store"
)

type degreeLevelService struct {
	Store store.Store
}

var _ DegreeLevelService = &degreeLevelService{}

func (d *degreeLevelService) BuildFromEntities(entities ent.DegreeLevels) []*dtos.DegreeLevel {
	dtos := make([]*dtos.DegreeLevel, len(entities))
	for i, entity := range entities {
		dtos[i] = d.BuildFromEntity(entity)
	}
	return dtos
}

func (d *degreeLevelService) BuildFromEntity(entity *ent.DegreeLevel) *dtos.DegreeLevel {
	dto := &dtos.DegreeLevel{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
	return dto
}

func (d degreeLevelService) GetDegreeLevel(ctx context.Context, id uuid.UUID) (*dtos.DegreeLevel, error) {
	degreeLevel, err := d.Store.DegreeLevel(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting degree_level: %w", err)
	}
	return d.BuildFromEntity(degreeLevel), nil
}

func (d degreeLevelService) PaginatedDegreeLevels(ctx context.Context, orderBy *dtos.DegreeLevelAggregationInput, take *int, skip *int, where *dtos.DegreeLevelWhereInput) (*dtos.PaginatedDegreeLevelResponse, error) {
	degreeLevels, err := d.Store.PaginatedDegreeLevels(ctx, orderBy, take, skip, where)
	if err != nil {
		return nil, fmt.Errorf("error getting paginated degree_levels: %w", err)
	}
	total, err := d.Store.CountDegreeLevels(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting total degree_levels: %w", err)
	}
	return &dtos.PaginatedDegreeLevelResponse{
		Edges: d.BuildFromEntities(degreeLevels),
		Take:  take,
		Skip:  skip,
		Total: total,
	}, nil
}

func (d degreeLevelService) CreateDegreeLevel(ctx context.Context, degreeLevel dtos.DegreeLevelCreateInput) (*dtos.DegreeLevel, error) {
	model := &ent.DegreeLevel{}
	model.Name = degreeLevel.Name

	createdDegreeLevel, err := d.Store.CreateDegreeLevel(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("error creating degree_level: %w", err)
	}

	return d.BuildFromEntity(createdDegreeLevel), nil
}

func (d degreeLevelService) UpdateDegreeLevel(ctx context.Context, id uuid.UUID, degreeLevel dtos.DegreeLevelUpdateInput) (*dtos.DegreeLevel, error) {
	model := &ent.DegreeLevel{}
	model.ID = id
	model.Name = degreeLevel.Name

	updatedDegreeLevel, err := d.Store.UpdateDegreeLevel(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("error updating degree_level: %w", err)
	}

	return d.BuildFromEntity(updatedDegreeLevel), nil
}

func (d degreeLevelService) DeleteDegreeLevel(ctx context.Context, id uuid.UUID) (*dtos.DegreeLevel, error) {
	degreeLevel, err := d.Store.DeleteDegreeLevel(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error deleting degree_level: %w", err)
	}
	return d.BuildFromEntity(degreeLevel), nil
}
