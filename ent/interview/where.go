// Code generated by ent, DO NOT EDIT.

package interview

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldLTE(FieldID, id))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldComment, v))
}

// InterviewDate applies equality check predicate on the "interview_date" field. It's identical to InterviewDateEQ.
func InterviewDate(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldInterviewDate, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldStatus, v))
}

// ApplicationID applies equality check predicate on the "application_id" field. It's identical to ApplicationIDEQ.
func ApplicationID(v uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldApplicationID, v))
}

// InterviewerID applies equality check predicate on the "interviewer_id" field. It's identical to InterviewerIDEQ.
func InterviewerID(v uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldInterviewerID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldDeletedAt, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Interview {
	return predicate.Interview(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Interview {
	return predicate.Interview(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Interview {
	return predicate.Interview(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Interview {
	return predicate.Interview(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Interview {
	return predicate.Interview(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Interview {
	return predicate.Interview(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Interview {
	return predicate.Interview(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Interview {
	return predicate.Interview(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Interview {
	return predicate.Interview(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Interview {
	return predicate.Interview(sql.FieldHasSuffix(FieldComment, v))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Interview {
	return predicate.Interview(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Interview {
	return predicate.Interview(sql.FieldContainsFold(FieldComment, v))
}

// InterviewDateEQ applies the EQ predicate on the "interview_date" field.
func InterviewDateEQ(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldInterviewDate, v))
}

// InterviewDateNEQ applies the NEQ predicate on the "interview_date" field.
func InterviewDateNEQ(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldNEQ(FieldInterviewDate, v))
}

// InterviewDateIn applies the In predicate on the "interview_date" field.
func InterviewDateIn(vs ...time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldIn(FieldInterviewDate, vs...))
}

// InterviewDateNotIn applies the NotIn predicate on the "interview_date" field.
func InterviewDateNotIn(vs ...time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldNotIn(FieldInterviewDate, vs...))
}

// InterviewDateGT applies the GT predicate on the "interview_date" field.
func InterviewDateGT(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldGT(FieldInterviewDate, v))
}

// InterviewDateGTE applies the GTE predicate on the "interview_date" field.
func InterviewDateGTE(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldGTE(FieldInterviewDate, v))
}

// InterviewDateLT applies the LT predicate on the "interview_date" field.
func InterviewDateLT(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldLT(FieldInterviewDate, v))
}

// InterviewDateLTE applies the LTE predicate on the "interview_date" field.
func InterviewDateLTE(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldLTE(FieldInterviewDate, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Interview {
	return predicate.Interview(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Interview {
	return predicate.Interview(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Interview {
	return predicate.Interview(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Interview {
	return predicate.Interview(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Interview {
	return predicate.Interview(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Interview {
	return predicate.Interview(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Interview {
	return predicate.Interview(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Interview {
	return predicate.Interview(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Interview {
	return predicate.Interview(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Interview {
	return predicate.Interview(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Interview {
	return predicate.Interview(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Interview {
	return predicate.Interview(sql.FieldContainsFold(FieldStatus, v))
}

// ApplicationIDEQ applies the EQ predicate on the "application_id" field.
func ApplicationIDEQ(v uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldApplicationID, v))
}

// ApplicationIDNEQ applies the NEQ predicate on the "application_id" field.
func ApplicationIDNEQ(v uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldNEQ(FieldApplicationID, v))
}

// ApplicationIDIn applies the In predicate on the "application_id" field.
func ApplicationIDIn(vs ...uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldIn(FieldApplicationID, vs...))
}

// ApplicationIDNotIn applies the NotIn predicate on the "application_id" field.
func ApplicationIDNotIn(vs ...uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldNotIn(FieldApplicationID, vs...))
}

// ApplicationIDIsNil applies the IsNil predicate on the "application_id" field.
func ApplicationIDIsNil() predicate.Interview {
	return predicate.Interview(sql.FieldIsNull(FieldApplicationID))
}

// ApplicationIDNotNil applies the NotNil predicate on the "application_id" field.
func ApplicationIDNotNil() predicate.Interview {
	return predicate.Interview(sql.FieldNotNull(FieldApplicationID))
}

// InterviewerIDEQ applies the EQ predicate on the "interviewer_id" field.
func InterviewerIDEQ(v uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldInterviewerID, v))
}

// InterviewerIDNEQ applies the NEQ predicate on the "interviewer_id" field.
func InterviewerIDNEQ(v uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldNEQ(FieldInterviewerID, v))
}

// InterviewerIDIn applies the In predicate on the "interviewer_id" field.
func InterviewerIDIn(vs ...uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldIn(FieldInterviewerID, vs...))
}

// InterviewerIDNotIn applies the NotIn predicate on the "interviewer_id" field.
func InterviewerIDNotIn(vs ...uuid.UUID) predicate.Interview {
	return predicate.Interview(sql.FieldNotIn(FieldInterviewerID, vs...))
}

// InterviewerIDIsNil applies the IsNil predicate on the "interviewer_id" field.
func InterviewerIDIsNil() predicate.Interview {
	return predicate.Interview(sql.FieldIsNull(FieldInterviewerID))
}

// InterviewerIDNotNil applies the NotNil predicate on the "interviewer_id" field.
func InterviewerIDNotNil() predicate.Interview {
	return predicate.Interview(sql.FieldNotNull(FieldInterviewerID))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Interview {
	return predicate.Interview(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Interview {
	return predicate.Interview(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Interview {
	return predicate.Interview(sql.FieldNotNull(FieldDeletedAt))
}

// HasApplication applies the HasEdge predicate on the "application" edge.
func HasApplication() predicate.Interview {
	return predicate.Interview(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApplicationTable, ApplicationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApplicationWith applies the HasEdge predicate on the "application" edge with a given conditions (other predicates).
func HasApplicationWith(preds ...predicate.Application) predicate.Interview {
	return predicate.Interview(func(s *sql.Selector) {
		step := newApplicationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Interview {
	return predicate.Interview(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Interview {
	return predicate.Interview(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Interview) predicate.Interview {
	return predicate.Interview(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Interview) predicate.Interview {
	return predicate.Interview(func(s *sql.Selector) {
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
func Not(p predicate.Interview) predicate.Interview {
	return predicate.Interview(func(s *sql.Selector) {
		p(s.Not())
	})
}