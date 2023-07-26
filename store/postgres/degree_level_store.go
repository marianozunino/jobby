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

// assert that DegreeLevelStore implements store.DegreeLevelStore
var _ store.DegreeLevelStore = &DegreeLevelStore{}

type DegreeLevelStore struct {
	*sqlx.DB
}

// CountDegreeLevels implements store.DegreeLevelStore.
func (c *DegreeLevelStore) CountDegreeLevels() (int, error) {
	var count int
	err := c.Get(&count, "SELECT COUNT(*) FROM degree_levels")
	if err != nil {
		return 0, fmt.Errorf("error counting degree_levels: %w", err)
	}
	return count, err
}

// CreateDegreeLevel implements store.DegreeLevelStore.
func (c *DegreeLevelStore) CreateDegreeLevel(degreeLevel models.DegreeLevels) (*models.DegreeLevels, error) {
	query := `INSERT INTO degree_levels (name) VALUES ($1) RETURNING *`
	createdDegreeLevel := &models.DegreeLevels{}
	if err := c.Get(createdDegreeLevel, query, degreeLevel.Name); err != nil {
		return nil, fmt.Errorf("error creating degree_level: %w", err)
	}
	return createdDegreeLevel, nil
}

// DegreeLevel implements store.DegreeLevelStore.
func (c *DegreeLevelStore) DegreeLevel(id uuid.UUID) (*models.DegreeLevels, error) {
	degreeLevel := &models.DegreeLevels{}
	if err := c.Get(degreeLevel, "SELECT * FROM degree_levels WHERE id = $1", id); err != nil {
		return nil, fmt.Errorf("error getting degree_level: %w", err)
	}
	return degreeLevel, nil
}

// DeleteDegreeLevel implements store.DegreeLevelStore.
func (c *DegreeLevelStore) DeleteDegreeLevel(id uuid.UUID) (*models.DegreeLevels, error) {
	degreeLevel := &models.DegreeLevels{}
	if err := c.Get(degreeLevel, "DELETE FROM degree_levels WHERE id = $1 RETURNING *", id); err != nil {
		return nil, fmt.Errorf("error deleting degree_level: %w", err)
	}
	return degreeLevel, nil
}

// PaginatedDegreeLevels implements store.DegreeLevelStore.
func (c *DegreeLevelStore) PaginatedDegreeLevels(orderBy *dtos.DegreeLevelAggregationInput, take *int, skip *int, where *dtos.DegreeLevelWhereInput) ([]*models.DegreeLevels, error) {
	query := "SELECT * FROM degree_levels"

	if where != nil {
		query += " WHERE deleted_at IS NULL AND "
		var whereFields []string
		switch {
		case where.ID != nil:
			whereFields = append(whereFields, fmt.Sprintf("id = '%s'", where.ID))
		case where.Name != nil:
			whereFields = append(whereFields, fmt.Sprintf("name ilike '%%%s%%'", *where.Name))
		}
		query += strings.Join(whereFields, " AND ")
	}

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

	var degreeLevels []*models.DegreeLevels

	if err := c.DB.Select(&degreeLevels, query); err != nil {
		return nil, fmt.Errorf("error getting degree_levels: %w", err)
	}

	return degreeLevels, nil

}

// UpdateDegreeLevel implements store.DegreeLevelStore.
func (c *DegreeLevelStore) UpdateDegreeLevel(degreeLevel models.DegreeLevels) (*models.DegreeLevels, error) {
	query := `UPDATE degree_levels SET name = $1 WHERE id = $2 RETURNING *`
	updatedDegreeLevel := &models.DegreeLevels{}
	if err := c.Get(updatedDegreeLevel, query,
		degreeLevel.Name,
		degreeLevel.ID,
	); err != nil {
		return nil, fmt.Errorf("error updating degree_level: %w", err)
	}
	return updatedDegreeLevel, nil
}
