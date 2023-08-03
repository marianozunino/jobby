package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/jobby/ent"
	"github.com/marianozunino/jobby/ent/joboffer"
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
