package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/jobby/ent"
	"github.com/marianozunino/jobby/ent/joboffer"
	"github.com/marianozunino/jobby/ent/joboffercategory"
	"github.com/marianozunino/jobby/store"
)

var _ store.JobOfferStore = &JobOfferStore{}

type JobOfferStore struct {
	*ent.Client
}

// JobOffersWithStatus implements store.JobOfferStore.
func (j JobOfferStore) JobOffersWithStatus(ctx context.Context, statusID []uuid.UUID) (ent.JobOffers, error) {
	return j.Client.JobOffer.Query().Where(joboffer.StatusIDIn(statusID...)).All(ctx)
}

// JobOffersForCategories implements store.JobOfferStore.
func (j JobOfferStore) JobOffersForCategories(ctx context.Context, categoryIDs []uuid.UUID) (map[uuid.UUID]ent.JobOffers, error) {
	mappedJobOffers := make(map[uuid.UUID]ent.JobOffers)

	categoriesWithJobOffers, err := j.Client.JobOfferCategory.Query().Where(joboffercategory.CategoryIDIn(categoryIDs...)).WithJobOffer().All(ctx)

	if err != nil {
		return nil, err
	}

	for _, relation := range categoriesWithJobOffers {
		mappedJobOffers[relation.CategoryID] = append(mappedJobOffers[relation.CategoryID], relation.Edges.JobOffer)
	}

	return mappedJobOffers, nil

}
