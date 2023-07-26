// Code generated by ent, DO NOT EDIT.

package joboffercategory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldLTE(FieldID, id))
}

// JobOfferID applies equality check predicate on the "job_offer_id" field. It's identical to JobOfferIDEQ.
func JobOfferID(v uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldJobOfferID, v))
}

// CategoryID applies equality check predicate on the "category_id" field. It's identical to CategoryIDEQ.
func CategoryID(v uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldCategoryID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldDeletedAt, v))
}

// JobOfferIDEQ applies the EQ predicate on the "job_offer_id" field.
func JobOfferIDEQ(v uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldJobOfferID, v))
}

// JobOfferIDNEQ applies the NEQ predicate on the "job_offer_id" field.
func JobOfferIDNEQ(v uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNEQ(FieldJobOfferID, v))
}

// JobOfferIDIn applies the In predicate on the "job_offer_id" field.
func JobOfferIDIn(vs ...uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldIn(FieldJobOfferID, vs...))
}

// JobOfferIDNotIn applies the NotIn predicate on the "job_offer_id" field.
func JobOfferIDNotIn(vs ...uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNotIn(FieldJobOfferID, vs...))
}

// CategoryIDEQ applies the EQ predicate on the "category_id" field.
func CategoryIDEQ(v uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldCategoryID, v))
}

// CategoryIDNEQ applies the NEQ predicate on the "category_id" field.
func CategoryIDNEQ(v uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNEQ(FieldCategoryID, v))
}

// CategoryIDIn applies the In predicate on the "category_id" field.
func CategoryIDIn(vs ...uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldIn(FieldCategoryID, vs...))
}

// CategoryIDNotIn applies the NotIn predicate on the "category_id" field.
func CategoryIDNotIn(vs ...uuid.UUID) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNotIn(FieldCategoryID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.JobOfferCategory {
	return predicate.JobOfferCategory(sql.FieldNotNull(FieldDeletedAt))
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.JobOfferCategory {
	return predicate.JobOfferCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasJobOffer applies the HasEdge predicate on the "job_offer" edge.
func HasJobOffer() predicate.JobOfferCategory {
	return predicate.JobOfferCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, JobOfferTable, JobOfferColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJobOfferWith applies the HasEdge predicate on the "job_offer" edge with a given conditions (other predicates).
func HasJobOfferWith(preds ...predicate.JobOffer) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(func(s *sql.Selector) {
		step := newJobOfferStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JobOfferCategory) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JobOfferCategory) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(func(s *sql.Selector) {
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
func Not(p predicate.JobOfferCategory) predicate.JobOfferCategory {
	return predicate.JobOfferCategory(func(s *sql.Selector) {
		p(s.Not())
	})
}
