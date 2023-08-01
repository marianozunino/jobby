package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/ent/joboffer"
	"github.com/marianozunino/cc-backend-go/store"
)

var _ store.JobOfferStore = &JobOfferStore{}

type JobOfferStore struct {
	*ent.Client
}

// JobOffersWithStatus implements store.JobOfferStore.
func (j JobOfferStore) JobOffersWithStatus(ctx context.Context, statusID []uuid.UUID) (ent.JobOffers, error) {
	return j.Client.JobOffer.Query().Where(joboffer.StatusIDIn(statusID...)).All(ctx)
}
