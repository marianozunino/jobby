package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Categories struct {
	ID       uuid.UUID  `db:"id"`
	Name     string     `db:"name"`
	Slug     string     `db:"slug"`
	ParentID *uuid.UUID `db:"parent_id"`
	IsRoot   bool       `db:"is_root"`

	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
