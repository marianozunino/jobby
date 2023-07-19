package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Interviews struct {
	ID            uuid.UUID `db:"id"`
	Comment       string    `db:"comment"`
	Status        string    `db:"status"`
	InterviewDate time.Time `db:"interview_date"`
	ApplicationID uuid.UUID `db:"application_id"`
	InterviewerID uuid.UUID `db:"interviewer_id"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
