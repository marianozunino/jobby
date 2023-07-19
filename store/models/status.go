package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Status struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`

	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
