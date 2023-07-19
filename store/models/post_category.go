package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type PostCategory struct {
	ID         uuid.UUID `db:"id"`
	PostID     uuid.UUID `db:"post_id"`
	CategoryID uuid.UUID `db:"category_id"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
