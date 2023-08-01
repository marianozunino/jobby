// Code generated by ent, DO NOT EDIT.

package application

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldID, id))
}

// JobOfferID applies equality check predicate on the "job_offer_id" field. It's identical to JobOfferIDEQ.
func JobOfferID(v uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldJobOfferID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldDeletedAt, v))
}

// ApplicantProfileID applies equality check predicate on the "applicant_profile_id" field. It's identical to ApplicantProfileIDEQ.
func ApplicantProfileID(v uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldApplicantProfileID, v))
}

// JobOfferIDEQ applies the EQ predicate on the "job_offer_id" field.
func JobOfferIDEQ(v uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldJobOfferID, v))
}

// JobOfferIDNEQ applies the NEQ predicate on the "job_offer_id" field.
func JobOfferIDNEQ(v uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldJobOfferID, v))
}

// JobOfferIDIn applies the In predicate on the "job_offer_id" field.
func JobOfferIDIn(vs ...uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldJobOfferID, vs...))
}

// JobOfferIDNotIn applies the NotIn predicate on the "job_offer_id" field.
func JobOfferIDNotIn(vs ...uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldJobOfferID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Application {
	return predicate.Application(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Application {
	return predicate.Application(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Application {
	return predicate.Application(sql.FieldNotNull(FieldDeletedAt))
}

// ApplicantProfileIDEQ applies the EQ predicate on the "applicant_profile_id" field.
func ApplicantProfileIDEQ(v uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldEQ(FieldApplicantProfileID, v))
}

// ApplicantProfileIDNEQ applies the NEQ predicate on the "applicant_profile_id" field.
func ApplicantProfileIDNEQ(v uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldNEQ(FieldApplicantProfileID, v))
}

// ApplicantProfileIDIn applies the In predicate on the "applicant_profile_id" field.
func ApplicantProfileIDIn(vs ...uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldIn(FieldApplicantProfileID, vs...))
}

// ApplicantProfileIDNotIn applies the NotIn predicate on the "applicant_profile_id" field.
func ApplicantProfileIDNotIn(vs ...uuid.UUID) predicate.Application {
	return predicate.Application(sql.FieldNotIn(FieldApplicantProfileID, vs...))
}

// HasApplicantProfile applies the HasEdge predicate on the "applicant_profile" edge.
func HasApplicantProfile() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApplicantProfileTable, ApplicantProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApplicantProfileWith applies the HasEdge predicate on the "applicant_profile" edge with a given conditions (other predicates).
func HasApplicantProfileWith(preds ...predicate.ApplicantProfile) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		step := newApplicantProfileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasJobOffer applies the HasEdge predicate on the "job_offer" edge.
func HasJobOffer() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, JobOfferTable, JobOfferColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJobOfferWith applies the HasEdge predicate on the "job_offer" edge with a given conditions (other predicates).
func HasJobOfferWith(preds ...predicate.JobOffer) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		step := newJobOfferStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInterviews applies the HasEdge predicate on the "interviews" edge.
func HasInterviews() predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InterviewsTable, InterviewsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInterviewsWith applies the HasEdge predicate on the "interviews" edge with a given conditions (other predicates).
func HasInterviewsWith(preds ...predicate.Interview) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		step := newInterviewsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
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
func Not(p predicate.Application) predicate.Application {
	return predicate.Application(func(s *sql.Selector) {
		p(s.Not())
	})
}