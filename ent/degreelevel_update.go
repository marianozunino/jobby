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
	"github.com/google/uuid"
	"github.com/marianozunino/jobby/ent/degreelevel"
	"github.com/marianozunino/jobby/ent/education"
	"github.com/marianozunino/jobby/ent/predicate"
)

// DegreeLevelUpdate is the builder for updating DegreeLevel entities.
type DegreeLevelUpdate struct {
	config
	hooks    []Hook
	mutation *DegreeLevelMutation
}

// Where appends a list predicates to the DegreeLevelUpdate builder.
func (dlu *DegreeLevelUpdate) Where(ps ...predicate.DegreeLevel) *DegreeLevelUpdate {
	dlu.mutation.Where(ps...)
	return dlu
}

// SetName sets the "name" field.
func (dlu *DegreeLevelUpdate) SetName(s string) *DegreeLevelUpdate {
	dlu.mutation.SetName(s)
	return dlu
}

// SetCreatedAt sets the "created_at" field.
func (dlu *DegreeLevelUpdate) SetCreatedAt(t time.Time) *DegreeLevelUpdate {
	dlu.mutation.SetCreatedAt(t)
	return dlu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dlu *DegreeLevelUpdate) SetNillableCreatedAt(t *time.Time) *DegreeLevelUpdate {
	if t != nil {
		dlu.SetCreatedAt(*t)
	}
	return dlu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (dlu *DegreeLevelUpdate) ClearCreatedAt() *DegreeLevelUpdate {
	dlu.mutation.ClearCreatedAt()
	return dlu
}

// SetUpdatedAt sets the "updated_at" field.
func (dlu *DegreeLevelUpdate) SetUpdatedAt(t time.Time) *DegreeLevelUpdate {
	dlu.mutation.SetUpdatedAt(t)
	return dlu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dlu *DegreeLevelUpdate) SetNillableUpdatedAt(t *time.Time) *DegreeLevelUpdate {
	if t != nil {
		dlu.SetUpdatedAt(*t)
	}
	return dlu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (dlu *DegreeLevelUpdate) ClearUpdatedAt() *DegreeLevelUpdate {
	dlu.mutation.ClearUpdatedAt()
	return dlu
}

// SetDeletedAt sets the "deleted_at" field.
func (dlu *DegreeLevelUpdate) SetDeletedAt(t time.Time) *DegreeLevelUpdate {
	dlu.mutation.SetDeletedAt(t)
	return dlu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dlu *DegreeLevelUpdate) SetNillableDeletedAt(t *time.Time) *DegreeLevelUpdate {
	if t != nil {
		dlu.SetDeletedAt(*t)
	}
	return dlu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (dlu *DegreeLevelUpdate) ClearDeletedAt() *DegreeLevelUpdate {
	dlu.mutation.ClearDeletedAt()
	return dlu
}

// AddEducationIDs adds the "educations" edge to the Education entity by IDs.
func (dlu *DegreeLevelUpdate) AddEducationIDs(ids ...uuid.UUID) *DegreeLevelUpdate {
	dlu.mutation.AddEducationIDs(ids...)
	return dlu
}

// AddEducations adds the "educations" edges to the Education entity.
func (dlu *DegreeLevelUpdate) AddEducations(e ...*Education) *DegreeLevelUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return dlu.AddEducationIDs(ids...)
}

// Mutation returns the DegreeLevelMutation object of the builder.
func (dlu *DegreeLevelUpdate) Mutation() *DegreeLevelMutation {
	return dlu.mutation
}

// ClearEducations clears all "educations" edges to the Education entity.
func (dlu *DegreeLevelUpdate) ClearEducations() *DegreeLevelUpdate {
	dlu.mutation.ClearEducations()
	return dlu
}

// RemoveEducationIDs removes the "educations" edge to Education entities by IDs.
func (dlu *DegreeLevelUpdate) RemoveEducationIDs(ids ...uuid.UUID) *DegreeLevelUpdate {
	dlu.mutation.RemoveEducationIDs(ids...)
	return dlu
}

// RemoveEducations removes "educations" edges to Education entities.
func (dlu *DegreeLevelUpdate) RemoveEducations(e ...*Education) *DegreeLevelUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return dlu.RemoveEducationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dlu *DegreeLevelUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dlu.sqlSave, dlu.mutation, dlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dlu *DegreeLevelUpdate) SaveX(ctx context.Context) int {
	affected, err := dlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dlu *DegreeLevelUpdate) Exec(ctx context.Context) error {
	_, err := dlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dlu *DegreeLevelUpdate) ExecX(ctx context.Context) {
	if err := dlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dlu *DegreeLevelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(degreelevel.Table, degreelevel.Columns, sqlgraph.NewFieldSpec(degreelevel.FieldID, field.TypeUUID))
	if ps := dlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dlu.mutation.Name(); ok {
		_spec.SetField(degreelevel.FieldName, field.TypeString, value)
	}
	if value, ok := dlu.mutation.CreatedAt(); ok {
		_spec.SetField(degreelevel.FieldCreatedAt, field.TypeTime, value)
	}
	if dlu.mutation.CreatedAtCleared() {
		_spec.ClearField(degreelevel.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := dlu.mutation.UpdatedAt(); ok {
		_spec.SetField(degreelevel.FieldUpdatedAt, field.TypeTime, value)
	}
	if dlu.mutation.UpdatedAtCleared() {
		_spec.ClearField(degreelevel.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := dlu.mutation.DeletedAt(); ok {
		_spec.SetField(degreelevel.FieldDeletedAt, field.TypeTime, value)
	}
	if dlu.mutation.DeletedAtCleared() {
		_spec.ClearField(degreelevel.FieldDeletedAt, field.TypeTime)
	}
	if dlu.mutation.EducationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degreelevel.EducationsTable,
			Columns: []string{degreelevel.EducationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(education.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dlu.mutation.RemovedEducationsIDs(); len(nodes) > 0 && !dlu.mutation.EducationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degreelevel.EducationsTable,
			Columns: []string{degreelevel.EducationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(education.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dlu.mutation.EducationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degreelevel.EducationsTable,
			Columns: []string{degreelevel.EducationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(education.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{degreelevel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dlu.mutation.done = true
	return n, nil
}

// DegreeLevelUpdateOne is the builder for updating a single DegreeLevel entity.
type DegreeLevelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DegreeLevelMutation
}

// SetName sets the "name" field.
func (dluo *DegreeLevelUpdateOne) SetName(s string) *DegreeLevelUpdateOne {
	dluo.mutation.SetName(s)
	return dluo
}

// SetCreatedAt sets the "created_at" field.
func (dluo *DegreeLevelUpdateOne) SetCreatedAt(t time.Time) *DegreeLevelUpdateOne {
	dluo.mutation.SetCreatedAt(t)
	return dluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dluo *DegreeLevelUpdateOne) SetNillableCreatedAt(t *time.Time) *DegreeLevelUpdateOne {
	if t != nil {
		dluo.SetCreatedAt(*t)
	}
	return dluo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (dluo *DegreeLevelUpdateOne) ClearCreatedAt() *DegreeLevelUpdateOne {
	dluo.mutation.ClearCreatedAt()
	return dluo
}

// SetUpdatedAt sets the "updated_at" field.
func (dluo *DegreeLevelUpdateOne) SetUpdatedAt(t time.Time) *DegreeLevelUpdateOne {
	dluo.mutation.SetUpdatedAt(t)
	return dluo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dluo *DegreeLevelUpdateOne) SetNillableUpdatedAt(t *time.Time) *DegreeLevelUpdateOne {
	if t != nil {
		dluo.SetUpdatedAt(*t)
	}
	return dluo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (dluo *DegreeLevelUpdateOne) ClearUpdatedAt() *DegreeLevelUpdateOne {
	dluo.mutation.ClearUpdatedAt()
	return dluo
}

// SetDeletedAt sets the "deleted_at" field.
func (dluo *DegreeLevelUpdateOne) SetDeletedAt(t time.Time) *DegreeLevelUpdateOne {
	dluo.mutation.SetDeletedAt(t)
	return dluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dluo *DegreeLevelUpdateOne) SetNillableDeletedAt(t *time.Time) *DegreeLevelUpdateOne {
	if t != nil {
		dluo.SetDeletedAt(*t)
	}
	return dluo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (dluo *DegreeLevelUpdateOne) ClearDeletedAt() *DegreeLevelUpdateOne {
	dluo.mutation.ClearDeletedAt()
	return dluo
}

// AddEducationIDs adds the "educations" edge to the Education entity by IDs.
func (dluo *DegreeLevelUpdateOne) AddEducationIDs(ids ...uuid.UUID) *DegreeLevelUpdateOne {
	dluo.mutation.AddEducationIDs(ids...)
	return dluo
}

// AddEducations adds the "educations" edges to the Education entity.
func (dluo *DegreeLevelUpdateOne) AddEducations(e ...*Education) *DegreeLevelUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return dluo.AddEducationIDs(ids...)
}

// Mutation returns the DegreeLevelMutation object of the builder.
func (dluo *DegreeLevelUpdateOne) Mutation() *DegreeLevelMutation {
	return dluo.mutation
}

// ClearEducations clears all "educations" edges to the Education entity.
func (dluo *DegreeLevelUpdateOne) ClearEducations() *DegreeLevelUpdateOne {
	dluo.mutation.ClearEducations()
	return dluo
}

// RemoveEducationIDs removes the "educations" edge to Education entities by IDs.
func (dluo *DegreeLevelUpdateOne) RemoveEducationIDs(ids ...uuid.UUID) *DegreeLevelUpdateOne {
	dluo.mutation.RemoveEducationIDs(ids...)
	return dluo
}

// RemoveEducations removes "educations" edges to Education entities.
func (dluo *DegreeLevelUpdateOne) RemoveEducations(e ...*Education) *DegreeLevelUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return dluo.RemoveEducationIDs(ids...)
}

// Where appends a list predicates to the DegreeLevelUpdate builder.
func (dluo *DegreeLevelUpdateOne) Where(ps ...predicate.DegreeLevel) *DegreeLevelUpdateOne {
	dluo.mutation.Where(ps...)
	return dluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dluo *DegreeLevelUpdateOne) Select(field string, fields ...string) *DegreeLevelUpdateOne {
	dluo.fields = append([]string{field}, fields...)
	return dluo
}

// Save executes the query and returns the updated DegreeLevel entity.
func (dluo *DegreeLevelUpdateOne) Save(ctx context.Context) (*DegreeLevel, error) {
	return withHooks(ctx, dluo.sqlSave, dluo.mutation, dluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dluo *DegreeLevelUpdateOne) SaveX(ctx context.Context) *DegreeLevel {
	node, err := dluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dluo *DegreeLevelUpdateOne) Exec(ctx context.Context) error {
	_, err := dluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dluo *DegreeLevelUpdateOne) ExecX(ctx context.Context) {
	if err := dluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dluo *DegreeLevelUpdateOne) sqlSave(ctx context.Context) (_node *DegreeLevel, err error) {
	_spec := sqlgraph.NewUpdateSpec(degreelevel.Table, degreelevel.Columns, sqlgraph.NewFieldSpec(degreelevel.FieldID, field.TypeUUID))
	id, ok := dluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DegreeLevel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, degreelevel.FieldID)
		for _, f := range fields {
			if !degreelevel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != degreelevel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dluo.mutation.Name(); ok {
		_spec.SetField(degreelevel.FieldName, field.TypeString, value)
	}
	if value, ok := dluo.mutation.CreatedAt(); ok {
		_spec.SetField(degreelevel.FieldCreatedAt, field.TypeTime, value)
	}
	if dluo.mutation.CreatedAtCleared() {
		_spec.ClearField(degreelevel.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := dluo.mutation.UpdatedAt(); ok {
		_spec.SetField(degreelevel.FieldUpdatedAt, field.TypeTime, value)
	}
	if dluo.mutation.UpdatedAtCleared() {
		_spec.ClearField(degreelevel.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := dluo.mutation.DeletedAt(); ok {
		_spec.SetField(degreelevel.FieldDeletedAt, field.TypeTime, value)
	}
	if dluo.mutation.DeletedAtCleared() {
		_spec.ClearField(degreelevel.FieldDeletedAt, field.TypeTime)
	}
	if dluo.mutation.EducationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degreelevel.EducationsTable,
			Columns: []string{degreelevel.EducationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(education.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dluo.mutation.RemovedEducationsIDs(); len(nodes) > 0 && !dluo.mutation.EducationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degreelevel.EducationsTable,
			Columns: []string{degreelevel.EducationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(education.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dluo.mutation.EducationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   degreelevel.EducationsTable,
			Columns: []string{degreelevel.EducationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(education.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DegreeLevel{config: dluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{degreelevel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dluo.mutation.done = true
	return _node, nil
}
