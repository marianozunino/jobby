// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marianozunino/jobby/ent/contactusmessage"
	"github.com/marianozunino/jobby/ent/predicate"
)

// ContactUsMessageUpdate is the builder for updating ContactUsMessage entities.
type ContactUsMessageUpdate struct {
	config
	hooks    []Hook
	mutation *ContactUsMessageMutation
}

// Where appends a list predicates to the ContactUsMessageUpdate builder.
func (cumu *ContactUsMessageUpdate) Where(ps ...predicate.ContactUsMessage) *ContactUsMessageUpdate {
	cumu.mutation.Where(ps...)
	return cumu
}

// SetName sets the "name" field.
func (cumu *ContactUsMessageUpdate) SetName(s string) *ContactUsMessageUpdate {
	cumu.mutation.SetName(s)
	return cumu
}

// SetEmail sets the "email" field.
func (cumu *ContactUsMessageUpdate) SetEmail(s string) *ContactUsMessageUpdate {
	cumu.mutation.SetEmail(s)
	return cumu
}

// SetMessage sets the "message" field.
func (cumu *ContactUsMessageUpdate) SetMessage(s string) *ContactUsMessageUpdate {
	cumu.mutation.SetMessage(s)
	return cumu
}

// SetPhone sets the "phone" field.
func (cumu *ContactUsMessageUpdate) SetPhone(s string) *ContactUsMessageUpdate {
	cumu.mutation.SetPhone(s)
	return cumu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cumu *ContactUsMessageUpdate) SetNillablePhone(s *string) *ContactUsMessageUpdate {
	if s != nil {
		cumu.SetPhone(*s)
	}
	return cumu
}

// ClearPhone clears the value of the "phone" field.
func (cumu *ContactUsMessageUpdate) ClearPhone() *ContactUsMessageUpdate {
	cumu.mutation.ClearPhone()
	return cumu
}

// SetCreatedAt sets the "created_at" field.
func (cumu *ContactUsMessageUpdate) SetCreatedAt(t time.Time) *ContactUsMessageUpdate {
	cumu.mutation.SetCreatedAt(t)
	return cumu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cumu *ContactUsMessageUpdate) SetNillableCreatedAt(t *time.Time) *ContactUsMessageUpdate {
	if t != nil {
		cumu.SetCreatedAt(*t)
	}
	return cumu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (cumu *ContactUsMessageUpdate) ClearCreatedAt() *ContactUsMessageUpdate {
	cumu.mutation.ClearCreatedAt()
	return cumu
}

// SetUpdatedAt sets the "updated_at" field.
func (cumu *ContactUsMessageUpdate) SetUpdatedAt(t time.Time) *ContactUsMessageUpdate {
	cumu.mutation.SetUpdatedAt(t)
	return cumu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cumu *ContactUsMessageUpdate) SetNillableUpdatedAt(t *time.Time) *ContactUsMessageUpdate {
	if t != nil {
		cumu.SetUpdatedAt(*t)
	}
	return cumu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cumu *ContactUsMessageUpdate) ClearUpdatedAt() *ContactUsMessageUpdate {
	cumu.mutation.ClearUpdatedAt()
	return cumu
}

// SetDeletedAt sets the "deleted_at" field.
func (cumu *ContactUsMessageUpdate) SetDeletedAt(t time.Time) *ContactUsMessageUpdate {
	cumu.mutation.SetDeletedAt(t)
	return cumu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cumu *ContactUsMessageUpdate) SetNillableDeletedAt(t *time.Time) *ContactUsMessageUpdate {
	if t != nil {
		cumu.SetDeletedAt(*t)
	}
	return cumu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cumu *ContactUsMessageUpdate) ClearDeletedAt() *ContactUsMessageUpdate {
	cumu.mutation.ClearDeletedAt()
	return cumu
}

// Mutation returns the ContactUsMessageMutation object of the builder.
func (cumu *ContactUsMessageUpdate) Mutation() *ContactUsMessageMutation {
	return cumu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cumu *ContactUsMessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cumu.sqlSave, cumu.mutation, cumu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cumu *ContactUsMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := cumu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cumu *ContactUsMessageUpdate) Exec(ctx context.Context) error {
	_, err := cumu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cumu *ContactUsMessageUpdate) ExecX(ctx context.Context) {
	if err := cumu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cumu *ContactUsMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(contactusmessage.Table, contactusmessage.Columns, sqlgraph.NewFieldSpec(contactusmessage.FieldID, field.TypeUUID))
	if ps := cumu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cumu.mutation.Name(); ok {
		_spec.SetField(contactusmessage.FieldName, field.TypeString, value)
	}
	if value, ok := cumu.mutation.Email(); ok {
		_spec.SetField(contactusmessage.FieldEmail, field.TypeString, value)
	}
	if value, ok := cumu.mutation.Message(); ok {
		_spec.SetField(contactusmessage.FieldMessage, field.TypeString, value)
	}
	if value, ok := cumu.mutation.Phone(); ok {
		_spec.SetField(contactusmessage.FieldPhone, field.TypeString, value)
	}
	if cumu.mutation.PhoneCleared() {
		_spec.ClearField(contactusmessage.FieldPhone, field.TypeString)
	}
	if value, ok := cumu.mutation.CreatedAt(); ok {
		_spec.SetField(contactusmessage.FieldCreatedAt, field.TypeTime, value)
	}
	if cumu.mutation.CreatedAtCleared() {
		_spec.ClearField(contactusmessage.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := cumu.mutation.UpdatedAt(); ok {
		_spec.SetField(contactusmessage.FieldUpdatedAt, field.TypeTime, value)
	}
	if cumu.mutation.UpdatedAtCleared() {
		_spec.ClearField(contactusmessage.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cumu.mutation.DeletedAt(); ok {
		_spec.SetField(contactusmessage.FieldDeletedAt, field.TypeTime, value)
	}
	if cumu.mutation.DeletedAtCleared() {
		_spec.ClearField(contactusmessage.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cumu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contactusmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cumu.mutation.done = true
	return n, nil
}

// ContactUsMessageUpdateOne is the builder for updating a single ContactUsMessage entity.
type ContactUsMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContactUsMessageMutation
}

// SetName sets the "name" field.
func (cumuo *ContactUsMessageUpdateOne) SetName(s string) *ContactUsMessageUpdateOne {
	cumuo.mutation.SetName(s)
	return cumuo
}

// SetEmail sets the "email" field.
func (cumuo *ContactUsMessageUpdateOne) SetEmail(s string) *ContactUsMessageUpdateOne {
	cumuo.mutation.SetEmail(s)
	return cumuo
}

// SetMessage sets the "message" field.
func (cumuo *ContactUsMessageUpdateOne) SetMessage(s string) *ContactUsMessageUpdateOne {
	cumuo.mutation.SetMessage(s)
	return cumuo
}

// SetPhone sets the "phone" field.
func (cumuo *ContactUsMessageUpdateOne) SetPhone(s string) *ContactUsMessageUpdateOne {
	cumuo.mutation.SetPhone(s)
	return cumuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cumuo *ContactUsMessageUpdateOne) SetNillablePhone(s *string) *ContactUsMessageUpdateOne {
	if s != nil {
		cumuo.SetPhone(*s)
	}
	return cumuo
}

// ClearPhone clears the value of the "phone" field.
func (cumuo *ContactUsMessageUpdateOne) ClearPhone() *ContactUsMessageUpdateOne {
	cumuo.mutation.ClearPhone()
	return cumuo
}

// SetCreatedAt sets the "created_at" field.
func (cumuo *ContactUsMessageUpdateOne) SetCreatedAt(t time.Time) *ContactUsMessageUpdateOne {
	cumuo.mutation.SetCreatedAt(t)
	return cumuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cumuo *ContactUsMessageUpdateOne) SetNillableCreatedAt(t *time.Time) *ContactUsMessageUpdateOne {
	if t != nil {
		cumuo.SetCreatedAt(*t)
	}
	return cumuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (cumuo *ContactUsMessageUpdateOne) ClearCreatedAt() *ContactUsMessageUpdateOne {
	cumuo.mutation.ClearCreatedAt()
	return cumuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cumuo *ContactUsMessageUpdateOne) SetUpdatedAt(t time.Time) *ContactUsMessageUpdateOne {
	cumuo.mutation.SetUpdatedAt(t)
	return cumuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cumuo *ContactUsMessageUpdateOne) SetNillableUpdatedAt(t *time.Time) *ContactUsMessageUpdateOne {
	if t != nil {
		cumuo.SetUpdatedAt(*t)
	}
	return cumuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cumuo *ContactUsMessageUpdateOne) ClearUpdatedAt() *ContactUsMessageUpdateOne {
	cumuo.mutation.ClearUpdatedAt()
	return cumuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cumuo *ContactUsMessageUpdateOne) SetDeletedAt(t time.Time) *ContactUsMessageUpdateOne {
	cumuo.mutation.SetDeletedAt(t)
	return cumuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cumuo *ContactUsMessageUpdateOne) SetNillableDeletedAt(t *time.Time) *ContactUsMessageUpdateOne {
	if t != nil {
		cumuo.SetDeletedAt(*t)
	}
	return cumuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cumuo *ContactUsMessageUpdateOne) ClearDeletedAt() *ContactUsMessageUpdateOne {
	cumuo.mutation.ClearDeletedAt()
	return cumuo
}

// Mutation returns the ContactUsMessageMutation object of the builder.
func (cumuo *ContactUsMessageUpdateOne) Mutation() *ContactUsMessageMutation {
	return cumuo.mutation
}

// Where appends a list predicates to the ContactUsMessageUpdate builder.
func (cumuo *ContactUsMessageUpdateOne) Where(ps ...predicate.ContactUsMessage) *ContactUsMessageUpdateOne {
	cumuo.mutation.Where(ps...)
	return cumuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cumuo *ContactUsMessageUpdateOne) Select(field string, fields ...string) *ContactUsMessageUpdateOne {
	cumuo.fields = append([]string{field}, fields...)
	return cumuo
}

// Save executes the query and returns the updated ContactUsMessage entity.
func (cumuo *ContactUsMessageUpdateOne) Save(ctx context.Context) (*ContactUsMessage, error) {
	return withHooks(ctx, cumuo.sqlSave, cumuo.mutation, cumuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cumuo *ContactUsMessageUpdateOne) SaveX(ctx context.Context) *ContactUsMessage {
	node, err := cumuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cumuo *ContactUsMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := cumuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cumuo *ContactUsMessageUpdateOne) ExecX(ctx context.Context) {
	if err := cumuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cumuo *ContactUsMessageUpdateOne) sqlSave(ctx context.Context) (_node *ContactUsMessage, err error) {
	_spec := sqlgraph.NewUpdateSpec(contactusmessage.Table, contactusmessage.Columns, sqlgraph.NewFieldSpec(contactusmessage.FieldID, field.TypeUUID))
	id, ok := cumuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ContactUsMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cumuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contactusmessage.FieldID)
		for _, f := range fields {
			if !contactusmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contactusmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cumuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cumuo.mutation.Name(); ok {
		_spec.SetField(contactusmessage.FieldName, field.TypeString, value)
	}
	if value, ok := cumuo.mutation.Email(); ok {
		_spec.SetField(contactusmessage.FieldEmail, field.TypeString, value)
	}
	if value, ok := cumuo.mutation.Message(); ok {
		_spec.SetField(contactusmessage.FieldMessage, field.TypeString, value)
	}
	if value, ok := cumuo.mutation.Phone(); ok {
		_spec.SetField(contactusmessage.FieldPhone, field.TypeString, value)
	}
	if cumuo.mutation.PhoneCleared() {
		_spec.ClearField(contactusmessage.FieldPhone, field.TypeString)
	}
	if value, ok := cumuo.mutation.CreatedAt(); ok {
		_spec.SetField(contactusmessage.FieldCreatedAt, field.TypeTime, value)
	}
	if cumuo.mutation.CreatedAtCleared() {
		_spec.ClearField(contactusmessage.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := cumuo.mutation.UpdatedAt(); ok {
		_spec.SetField(contactusmessage.FieldUpdatedAt, field.TypeTime, value)
	}
	if cumuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(contactusmessage.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cumuo.mutation.DeletedAt(); ok {
		_spec.SetField(contactusmessage.FieldDeletedAt, field.TypeTime, value)
	}
	if cumuo.mutation.DeletedAtCleared() {
		_spec.ClearField(contactusmessage.FieldDeletedAt, field.TypeTime)
	}
	_node = &ContactUsMessage{config: cumuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cumuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contactusmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cumuo.mutation.done = true
	return _node, nil
}
