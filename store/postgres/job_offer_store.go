package postgres

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/marianozunino/cc-backend-go/store"
	"github.com/marianozunino/cc-backend-go/store/models"
)

var _ store.JobOfferStore = &JobOfferStore{}

type JobOfferStore struct {
	*sqlx.DB
}

// JobOffersWithStatus implements store.JobOfferStore.
func (j *JobOfferStore) JobOffersWithStatus(statusID uuid.UUID) ([]*models.JobOffers, error) {
	var jobOffers []*models.JobOffers

	if err := j.Select(&jobOffers, "SELECT * FROM job_offers WHERE status_id = $1", statusID); err != nil {
		return nil, err
	}

	return jobOffers, nil
}
