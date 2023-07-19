package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Posts struct {
	ID            uuid.UUID      `db:"id"`
	Title         string         `db:"title"`
	Content       string         `db:"content"`
	Slug          string         `db:"slug"`
	IsHighlighted bool           `db:"is_highlighted"`
	IsPublished   bool           `db:"is_published"`
	PublishedAt   sql.NullTime   `db:"published_at"`
	PreviewImage  sql.NullString `db:"preview_image"`
	AuthorID      uuid.UUID      `db:"author_id"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
