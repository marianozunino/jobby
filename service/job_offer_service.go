package service

import (
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/store"
)

type jobOfferService struct {
	Store store.Store
}

// JobOffersWithStatus implements JobOfferService.
func (s *jobOfferService) JobOffersWithStatus(statusID string) ([]*dtos.JobOffer, error) {
	parsedId, err := uuid.Parse(statusID)
	if err != nil {
		return nil, err
	}
	jobOffers, err := s.Store.JobOffersWithStatus(parsedId)
	if err != nil {
		return nil, err
	}

	result := make([]*dtos.JobOffer, len(jobOffers))
	for i, jobOffer := range jobOffers {
		result[i] = &dtos.JobOffer{
			ID:        jobOffer.ID.String(),
			Slug:      jobOffer.Slug,
			StatusID:  jobOffer.StatusID.String(),
			CreatedAt: jobOffer.CreatedAt.Time,
			UpdatedAt: jobOffer.UpdatedAt.Time,
			DeletedAt: getTimeOrNil(jobOffer.DeletedAt),
			Title:     jobOffer.Title,
		}
	}

	return result, nil

}

var _ JobOfferService = &jobOfferService{}
