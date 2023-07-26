// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}), field.String("email").Unique(), field.String("password"), field.Enum("role").Values("ADMIN", "APPLICANT", "RECRUITER"), field.String("reset_password_token").Optional(), field.Time("reset_password_expires").Optional(), field.Time("created_at").Optional(), field.Time("updated_at").Optional(), field.Time("deleted_at").Optional(), field.String("full_name"), field.String("salt"), field.Int32("external_id").Optional()}
}
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("applicant_profiles", ApplicantProfile.Type),
		edge.To("interviews", Interview.Type),
		edge.To("posts", Post.Type),
	}
}
func (User) Annotations() []schema.Annotation {
	return nil
}
