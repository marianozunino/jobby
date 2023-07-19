package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Applications struct {
	ID                 uuid.UUID `db:"id"`
	JobOfferID         uuid.UUID `db:"job_offer_id"`
	ApplicantProfileID uuid.UUID `db:"applicant_profile_id"`

	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
