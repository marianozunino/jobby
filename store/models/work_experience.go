package models

import (
	"database/sql"
	"time"
)

type WorkExperience struct {
	ID                 string         `db:"id"`
	Company            sql.NullString `db:"company"`
	Position           sql.NullString `db:"position"`
	Description        sql.NullString `db:"description"`
	StartDate          sql.NullTime   `db:"start_date"`
	EndDate            sql.NullTime   `db:"end_date"`
	ApplicantProfileID sql.NullString `db:"applicant_profile_id"`
	CreatedAt          time.Time      `db:"created_at"`
	UpdatedAt          time.Time      `db:"updated_at"`
	DeletedAt          sql.NullTime   `db:"deleted_at"`
	InternalComments   sql.NullString `db:"internal_comments"`
}
