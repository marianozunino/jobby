package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type ApplicantInterests struct {
	ID                 uuid.UUID    `db:"id"`
	CategoryID         uuid.UUID    `db:"category_id"`
	CreatedAt          sql.NullTime `db:"created_at"`
	UpdatedAt          sql.NullTime `db:"updated_at"`
	DeletedAt          sql.NullTime `db:"deleted_at"`
	ApplicantProfileID uuid.UUID    `db:"applicant_profile_id"`
}
