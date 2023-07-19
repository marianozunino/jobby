package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Education struct {
	ID                 uuid.UUID      `db:"id"`
	Title              string         `db:"title"`
	Institution        sql.NullString `db:"institution"`
	DateObtained       sql.NullTime   `db:"date_obtained"`
	Comments           sql.NullString `db:"comments"`
	DegreeLevelID      uuid.NullUUID  `db:"degree_level_id"`
	ApplicantProfileID uuid.NullUUID  `db:"applicant_profile_id"`

	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
