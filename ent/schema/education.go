// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Education struct {
	ent.Schema
}

func (Education) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}), field.String("title"), field.String("institution").Optional(), field.Time("date_obtained").Optional(), field.UUID("degree_level_id", uuid.UUID{}).Optional().Nillable(), field.Time("created_at").Optional(), field.Time("updated_at").Optional(), field.Time("deleted_at").Optional(), field.String("comments").Optional(), field.UUID("applicant_profile_id", uuid.UUID{}).Optional()}
}
func (Education) Edges() []ent.Edge {
	return []ent.Edge{edge.From("degree_level", DegreeLevel.Type).Ref("educations").Unique().Field("degree_level_id") /* edge.From("applicant_profile", ApplicantProfile.Type).Ref("educations").Unique().Field("applicant_profile_id") */}
}
func (Education) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "education"}}
}
