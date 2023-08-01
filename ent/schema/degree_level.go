// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type DegreeLevel struct {
	ent.Schema
}

func (DegreeLevel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}),
		field.String("name"),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional().Nillable(),
	}
}
func (DegreeLevel) Edges() []ent.Edge {
	return []ent.Edge{edge.To("educations", Education.Type)}
}
func (DegreeLevel) Annotations() []schema.Annotation {
	return nil
}
