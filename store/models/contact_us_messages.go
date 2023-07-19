package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type ContactUsMessages struct {
	ID      uuid.UUID      `db:"id"`
	Name    string         `db:"name"`
	Email   string         `db:"email"`
	Message string         `db:"message"`
	Phone   sql.NullString `db:"phone"`

	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
