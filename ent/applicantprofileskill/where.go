// Code generated by ent, DO NOT EDIT.

package applicantprofileskill

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/marianozunino/jobby/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldLTE(FieldID, id))
}

// ApplicantProfileID applies equality check predicate on the "applicant_profile_id" field. It's identical to ApplicantProfileIDEQ.
func ApplicantProfileID(v uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldApplicantProfileID, v))
}

// SkillID applies equality check predicate on the "skill_id" field. It's identical to SkillIDEQ.
func SkillID(v uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldSkillID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldDeletedAt, v))
}

// ApplicantProfileIDEQ applies the EQ predicate on the "applicant_profile_id" field.
func ApplicantProfileIDEQ(v uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldApplicantProfileID, v))
}

// ApplicantProfileIDNEQ applies the NEQ predicate on the "applicant_profile_id" field.
func ApplicantProfileIDNEQ(v uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNEQ(FieldApplicantProfileID, v))
}

// ApplicantProfileIDIn applies the In predicate on the "applicant_profile_id" field.
func ApplicantProfileIDIn(vs ...uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldIn(FieldApplicantProfileID, vs...))
}

// ApplicantProfileIDNotIn applies the NotIn predicate on the "applicant_profile_id" field.
func ApplicantProfileIDNotIn(vs ...uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNotIn(FieldApplicantProfileID, vs...))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v Level) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v Level) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...Level) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...Level) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNotIn(FieldLevel, vs...))
}

// SkillIDEQ applies the EQ predicate on the "skill_id" field.
func SkillIDEQ(v uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldSkillID, v))
}

// SkillIDNEQ applies the NEQ predicate on the "skill_id" field.
func SkillIDNEQ(v uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNEQ(FieldSkillID, v))
}

// SkillIDIn applies the In predicate on the "skill_id" field.
func SkillIDIn(vs ...uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldIn(FieldSkillID, vs...))
}

// SkillIDNotIn applies the NotIn predicate on the "skill_id" field.
func SkillIDNotIn(vs ...uuid.UUID) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNotIn(FieldSkillID, vs...))
}

// SkillIDIsNil applies the IsNil predicate on the "skill_id" field.
func SkillIDIsNil() predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldIsNull(FieldSkillID))
}

// SkillIDNotNil applies the NotNil predicate on the "skill_id" field.
func SkillIDNotNil() predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNotNull(FieldSkillID))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(sql.FieldNotNull(FieldDeletedAt))
}

// HasApplicantProfile applies the HasEdge predicate on the "applicant_profile" edge.
func HasApplicantProfile() predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApplicantProfileTable, ApplicantProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApplicantProfileWith applies the HasEdge predicate on the "applicant_profile" edge with a given conditions (other predicates).
func HasApplicantProfileWith(preds ...predicate.ApplicantProfile) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(func(s *sql.Selector) {
		step := newApplicantProfileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSkill applies the HasEdge predicate on the "skill" edge.
func HasSkill() predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SkillTable, SkillColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSkillWith applies the HasEdge predicate on the "skill" edge with a given conditions (other predicates).
func HasSkillWith(preds ...predicate.Skill) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(func(s *sql.Selector) {
		step := newSkillStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ApplicantProfileSkill) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ApplicantProfileSkill) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ApplicantProfileSkill) predicate.ApplicantProfileSkill {
	return predicate.ApplicantProfileSkill(func(s *sql.Selector) {
		p(s.Not())
	})
}
