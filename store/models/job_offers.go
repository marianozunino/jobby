package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type JobOffers struct {
	ID             uuid.UUID      `db:"id"`
	Title          string         `db:"title"`
	Reference      int            `db:"reference"`
	StartDate      time.Time      `db:"start_date"`
	EndDate        time.Time      `db:"end_date"`
	Address1       sql.NullString `db:"address1"`
	Address2       sql.NullString `db:"address2"`
	Department     string         `db:"department"`
	Description    string         `db:"description"`
	WorkingHours   string         `db:"working_hours"`
	Salary         string         `db:"salary"`
	Slug           string         `db:"slug"`
	IsFeatured     sql.NullBool   `db:"is_featured"`
	HasBeenEmailed sql.NullBool   `db:"has_been_emailed"`

	StatusID uuid.UUID `db:"status_id"`

	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
