package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ApplicantProfiles struct {
	ID               uuid.UUID      `db:"id"`
	Birthday         time.Time      `db:"birthday"`
	Gender           string         `db:"gender"`
	Phone            string         `db:"phone"`
	Address1         string         `db:"address1"`
	Address2         string         `db:"address2"`
	Cv               sql.NullString `db:"cv"`
	InternalComments string         `db:"internal_comments"`
	ReceiveEmails    bool           `db:"receive_emails"`
	CreatedAt        sql.NullTime   `db:"created_at"`
	UpdatedAt        sql.NullTime   `db:"updated_at"`
	DeletedAt        sql.NullTime   `db:"deleted_at"`
	UserID           string         `db:"user_id"`
	ExtraSkills      sql.NullString `db:"extra_skills"`
}
