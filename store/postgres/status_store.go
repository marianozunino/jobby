package postgres

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/store"
	"github.com/marianozunino/cc-backend-go/store/models"
)

// CreateStatus implements store.StatusStore.
func (ss *StatusStore) CreateStatus(status models.Status) (*models.Status, error) {
	var output *models.Status = &models.Status{}
	if err := ss.DB.Get(
		output,
		"INSERT INTO status (name) VALUES ($1) RETURNING *",
	); err != nil {
		return nil, fmt.Errorf("error creating status: %w", err)
	}
	return output, nil
}

// DeleteStatus implements store.StatusStore.
func (ss *StatusStore) DeleteStatus(id uuid.UUID) (*models.Status, error) {
	var output *models.Status = &models.Status{}
	if err := ss.DB.Get(output, "DELETE FROM status WHERE id=$1 RETURNING *", id); err != nil {
		return nil, fmt.Errorf("error deleting status: %w", err)
	}
	return output, nil
}

// Status implements store.StatusStore.
func (ss *StatusStore) Status(id uuid.UUID) (*models.Status, error) {
	var status *models.Status = &models.Status{}
	if err := ss.DB.Get(status, "SELECT * FROM status WHERE id=$1", id); err != nil {
		return nil, fmt.Errorf("error getting status: %w", err)
	}
	return status, nil
}

// Statuses implements store.StatusStore.
func (ss *StatusStore) Statuses() ([]*models.Status, error) {
	var statuses []*models.Status = []*models.Status{}
	if err := ss.DB.Select(&statuses, "SELECT * FROM status"); err != nil {
		return nil, fmt.Errorf("error getting statuses: %w", err)
	}
	return statuses, nil
}

// UpdateStatus implements store.StatusStore.
func (ss *StatusStore) UpdateStatus(status models.Status) (*models.Status, error) {
	var output *models.Status = &models.Status{}
	if err := ss.DB.Get(
		output,
		"UPDATE status SET name=$2 WHERE id=$1 RETURNING *",
		status.ID,
		status.Name,
	); err != nil {
		return nil, fmt.Errorf("error updating status: %w", err)
	}
	return output, nil
}

func (ss *StatusStore) PaginatedStatuses(orderBy *dtos.StatusAggregationInput, take *int, skip *int) ([]*models.Status, error) {
	query := "SELECT * FROM status"

	if orderBy != nil {
		query += " ORDER BY "
		var sortFields []string

		if orderBy.ID != nil {
			sortFields = append(sortFields, fmt.Sprintf("id %s", orderBy.ID))
		}
		if orderBy.Name != nil {
			sortFields = append(sortFields, fmt.Sprintf("name %s", orderBy.Name))
		}
		if orderBy.CreatedAt != nil {
			sortFields = append(sortFields, fmt.Sprintf("created_at %s", orderBy.CreatedAt))
		}
		if orderBy.UpdatedAt != nil {
			sortFields = append(sortFields, fmt.Sprintf("updated_at %s", orderBy.UpdatedAt))
		}
		if orderBy.DeletedAt != nil {
			sortFields = append(sortFields, fmt.Sprintf("deleted_at %s", orderBy.DeletedAt))
		}
		query += strings.Join(sortFields, ", ")
	}

	if take != nil {
		query += fmt.Sprintf(" LIMIT %d", *take)
	}
	if skip != nil {
		query += fmt.Sprintf(" OFFSET %d", *skip)
	}

	statuses := []*models.Status{}

	if err := ss.DB.Select(&statuses, query); err != nil {
		return nil, fmt.Errorf("error getting paginated statuses: %w", err)
	}

	return statuses, nil
}

// assert that StatusStore implements store.StatusStore
var _ store.StatusStore = &StatusStore{}

type StatusStore struct {
	*sqlx.DB
}

// CountStatuses implements store.StatusStore.
func (ss *StatusStore) CountStatuses() (int, error) {
	var count int
	if err := ss.DB.Get(&count, "SELECT COUNT(*) FROM status"); err != nil {
		return 0, fmt.Errorf("error getting count of statuses: %w", err)
	}
	return count, nil
}
