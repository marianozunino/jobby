// Code generated by ent, DO NOT EDIT.

package contactusmessage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/marianozunino/jobby/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldEmail, v))
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldMessage, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldPhone, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldDeletedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldContainsFold(FieldEmail, v))
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldContainsFold(FieldMessage, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldContainsFold(FieldPhone, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.ContactUsMessage {
	return predicate.ContactUsMessage(sql.FieldNotNull(FieldDeletedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ContactUsMessage) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ContactUsMessage) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(func(s *sql.Selector) {
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
func Not(p predicate.ContactUsMessage) predicate.ContactUsMessage {
	return predicate.ContactUsMessage(func(s *sql.Selector) {
		p(s.Not())
	})
}
