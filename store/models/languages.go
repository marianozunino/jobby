package models

import (
	"database/sql"
	"time"
)

type Languages struct {
	ID                 int            `db:"id"`
	Name               string         `db:"name"`
	Level              string         `db:"level"`
	Comments           sql.NullString `db:"comments"`
	ApplicantProfileID string         `db:"applicant_profile_id"`
	CreatedAt          time.Time      `db:"created_at"`
	UpdatedAt          time.Time      `db:"updated_at"`
	DeletedAt          sql.NullTime   `db:"deleted_at"`
}
