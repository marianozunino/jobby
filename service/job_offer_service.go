package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/ent"
	"github.com/marianozunino/jobby/store"
)

type jobOfferService struct {
	Store store.Store
}

var _ JobOfferService = &jobOfferService{}

// JobOffersWithStatus implements JobOfferService.
func (j jobOfferService) JobOffersWithStatus(ctx context.Context, statusID []uuid.UUID) ([]*dtos.JobOffer, error) {
	jobOffers, err := j.Store.JobOffersWithStatus(ctx, statusID)
	if err != nil {
		return nil, err
	}
	return j.BuildFromEntities(jobOffers), nil
}

func (j jobOfferService) BuildFromEntity(entity *ent.JobOffer) *dtos.JobOffer {
	dto := &dtos.JobOffer{
		ID:             entity.ID,
		Slug:           entity.Slug,
		Title:          entity.Title,
		StatusID:       entity.StatusID,
		Salary:         entity.Salary,
		Description:    entity.Description,
		EndDate:        entity.EndDate,
		StartDate:      entity.StartDate,
		Department:     entity.Department,
		WorkingHours:   entity.WorkingHours,
		IsFeatured:     entity.IsFeatured,
		HasBeenEmailed: entity.HasBeenEmailed,
		Address1:       entity.Address1,
		Address2:       entity.Address2,
		Reference:      int(entity.Reference),
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
		DeletedAt:      entity.DeletedAt,
	}
	return dto
}

func (j *jobOfferService) BuildFromEntities(entities ent.JobOffers) []*dtos.JobOffer {
	dtos := make([]*dtos.JobOffer, len(entities))
	for i, entity := range entities {
		dtos[i] = j.BuildFromEntity(entity)
	}
	return dtos
}

// JobOffersForCategories implements JobOfferService.
func (j *jobOfferService) JobOffersForCategories(ctx context.Context, categoryIDs []uuid.UUID) (map[uuid.UUID][]*dtos.JobOffer, error) {
	mappedJobOffers := make(map[uuid.UUID][]*dtos.JobOffer)

	categoriesWithJobOffers, err := j.Store.JobOffersForCategories(ctx, categoryIDs)

	if err != nil {
		return nil, err
	}

	for categoryID, jobOffers := range categoriesWithJobOffers {
		mappedJobOffers[categoryID] = j.BuildFromEntities(jobOffers)
	}

	return mappedJobOffers, nil
}
