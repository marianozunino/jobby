package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/store"
	"github.com/marianozunino/cc-backend-go/store/models"
)

type degreeLevelService struct {
	Store store.Store
}

var _ DegreeLevelService = &degreeLevelService{}

// CreateDegreeLevel implements DegreeLevelService.
func (d *degreeLevelService) CreateDegreeLevel(degreeLevel dtos.DegreeLevelCreateInput) (*dtos.DegreeLevel, error) {
	model := models.DegreeLevels{
		Name: degreeLevel.Name,
	}

	createdDegreeLevel, err := d.Store.CreateDegreeLevel(model)
	if err != nil {
		return nil, fmt.Errorf("error creating degree_level: %w", err)
	}

	return d.BuildFromEntity(createdDegreeLevel), nil

}

// DegreeLevel implements DegreeLevelService.
func (d *degreeLevelService) GetDegreeLevel(id uuid.UUID) (*dtos.DegreeLevel, error) {
	degreeLevel, err := d.Store.DegreeLevel(id)
	if err != nil {
		return nil, fmt.Errorf("error getting degree_level: %w", err)
	}
	return d.BuildFromEntity(degreeLevel), nil
}

// DeleteDegreeLevel implements DegreeLevelService.
func (d *degreeLevelService) DeleteDegreeLevel(id uuid.UUID) (*dtos.DegreeLevel, error) {
	degreeLevel, err := d.Store.DeleteDegreeLevel(id)
	if err != nil {
		return nil, fmt.Errorf("error deleting degree_level: %w", err)
	}
	return d.BuildFromEntity(degreeLevel), nil
}

// PaginatedDegreeLevels implements DegreeLevelService.
func (d *degreeLevelService) PaginatedDegreeLevels(orderBy *dtos.DegreeLevelAggregationInput, take *int, skip *int, where *dtos.DegreeLevelWhereInput) (*dtos.PaginatedDegreeLevelResponse, error) {
	degreeLevels, err := d.Store.PaginatedDegreeLevels(orderBy, take, skip, where)
	if err != nil {
		return nil, fmt.Errorf("error getting paginated degree_levels: %w", err)
	}
	total, err := d.Store.CountDegreeLevels()
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

// UpdateDegreeLevel implements DegreeLevelService.
func (d *degreeLevelService) UpdateDegreeLevel(id uuid.UUID, degreeLevel dtos.DegreeLevelUpdateInput) (*dtos.DegreeLevel, error) {
	model := models.DegreeLevels{
		ID:   id,
		Name: degreeLevel.Name,
	}

	updatedDegreeLevel, err := d.Store.UpdateDegreeLevel(model)
	if err != nil {
		return nil, fmt.Errorf("error updating degree_level: %w", err)
	}

	return d.BuildFromEntity(updatedDegreeLevel), nil
}

func (d *degreeLevelService) BuildFromEntities(entities []*models.DegreeLevels) []*dtos.DegreeLevel {
	dtos := make([]*dtos.DegreeLevel, len(entities))
	for i, entity := range entities {
		dtos[i] = d.BuildFromEntity(entity)
	}
	return dtos
}

func (d *degreeLevelService) BuildFromEntity(entity *models.DegreeLevels) *dtos.DegreeLevel {
	dto := &dtos.DegreeLevel{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt.Time,
		UpdatedAt: entity.UpdatedAt.Time,
		DeletedAt: getTimeOrNil(entity.DeletedAt),
	}
	return dto
}
