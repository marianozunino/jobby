package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type ApplicantProfileSkills struct {
	ID                 uuid.UUID     `db:"id"`
	ApplicantProfileID uuid.NullUUID `db:"applicant_profile_id"`
	Level              string        `db:"level"`
	SkillID            uuid.NullUUID `db:"skill_id"`

	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
