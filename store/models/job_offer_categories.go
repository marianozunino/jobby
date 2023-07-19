package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type JobOfferCategories struct {
	ID         uuid.UUID `db:"id"`
	JobOfferID uuid.UUID `db:"job_offer_id"`
	CategoryID uuid.UUID `db:"category_id"`

	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
