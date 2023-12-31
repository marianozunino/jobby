// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Interview struct {
	ent.Schema
}

func (Interview) Fields() []ent.Field {
	return []ent.Field{field.UUID("id", uuid.UUID{}), field.String("comment"), field.Time("interview_date"), field.String("status"), field.UUID("application_id", uuid.UUID{}).Optional().Nillable(), field.UUID("interviewer_id", uuid.UUID{}).Optional(), field.Time("created_at"), field.Time("updated_at"), field.Time("deleted_at").Optional()}
}
func (Interview) Edges() []ent.Edge {
	return []ent.Edge{edge.From("application", Application.Type).Ref("interviews").Unique().Field("application_id"), edge.From("user", User.Type).Ref("interviews").Unique().Field("interviewer_id")}
}
func (Interview) Annotations() []schema.Annotation {
	return nil
}
