package postgres

import (
	"fmt"

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
func (j *JobOfferStore) JobOffersWithStatus(statusID []uuid.UUID) ([]*models.JobOffers, error) {

	query, args, err := sqlx.In("SELECT * FROM job_offers WHERE status_id IN (?)", statusID)

	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}
	query = j.Rebind(query)

	var jobOffers []*models.JobOffers

	err = j.Select(&jobOffers, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying job offers: %w", err)
	}

	return jobOffers, nil
}
