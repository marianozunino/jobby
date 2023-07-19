package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type Users struct {
	ID       uuid.UUID `db:"id"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Role     string    `db:"role"`

	ResetPasswordToken   sql.NullString `db:"reset_password_token"`
	ResetPasswordExpires sql.NullTime   `db:"reset_password_expires"`

	FullName   string        `db:"full_name"`
	Salt       string        `db:"salt"`
	ExternalID sql.NullInt64 `db:"external_id"`

	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
