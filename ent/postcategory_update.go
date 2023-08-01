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
	"github.com/marianozunino/cc-backend-go/ent/postcategory"
	"github.com/marianozunino/cc-backend-go/ent/predicate"
)

// PostCategoryUpdate is the builder for updating PostCategory entities.
type PostCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *PostCategoryMutation
}

// Where appends a list predicates to the PostCategoryUpdate builder.
func (pcu *PostCategoryUpdate) Where(ps ...predicate.PostCategory) *PostCategoryUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetName sets the "name" field.
func (pcu *PostCategoryUpdate) SetName(s string) *PostCategoryUpdate {
	pcu.mutation.SetName(s)
	return pcu
}

// SetCreatedAt sets the "created_at" field.
func (pcu *PostCategoryUpdate) SetCreatedAt(t time.Time) *PostCategoryUpdate {
	pcu.mutation.SetCreatedAt(t)
	return pcu
}

// SetUpdatedAt sets the "updated_at" field.
func (pcu *PostCategoryUpdate) SetUpdatedAt(t time.Time) *PostCategoryUpdate {
	pcu.mutation.SetUpdatedAt(t)
	return pcu
}

// SetDeletedAt sets the "deleted_at" field.
func (pcu *PostCategoryUpdate) SetDeletedAt(t time.Time) *PostCategoryUpdate {
	pcu.mutation.SetDeletedAt(t)
	return pcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pcu *PostCategoryUpdate) SetNillableDeletedAt(t *time.Time) *PostCategoryUpdate {
	if t != nil {
		pcu.SetDeletedAt(*t)
	}
	return pcu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pcu *PostCategoryUpdate) ClearDeletedAt() *PostCategoryUpdate {
	pcu.mutation.ClearDeletedAt()
	return pcu
}

// AddPostCategoryIDs adds the "post_categories" edge to the PostCategory entity by IDs.
func (pcu *PostCategoryUpdate) AddPostCategoryIDs(ids ...uuid.UUID) *PostCategoryUpdate {
	pcu.mutation.AddPostCategoryIDs(ids...)
	return pcu
}

// AddPostCategories adds the "post_categories" edges to the PostCategory entity.
func (pcu *PostCategoryUpdate) AddPostCategories(p ...*PostCategory) *PostCategoryUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcu.AddPostCategoryIDs(ids...)
}

// Mutation returns the PostCategoryMutation object of the builder.
func (pcu *PostCategoryUpdate) Mutation() *PostCategoryMutation {
	return pcu.mutation
}

// ClearPostCategories clears all "post_categories" edges to the PostCategory entity.
func (pcu *PostCategoryUpdate) ClearPostCategories() *PostCategoryUpdate {
	pcu.mutation.ClearPostCategories()
	return pcu
}

// RemovePostCategoryIDs removes the "post_categories" edge to PostCategory entities by IDs.
func (pcu *PostCategoryUpdate) RemovePostCategoryIDs(ids ...uuid.UUID) *PostCategoryUpdate {
	pcu.mutation.RemovePostCategoryIDs(ids...)
	return pcu
}

// RemovePostCategories removes "post_categories" edges to PostCategory entities.
func (pcu *PostCategoryUpdate) RemovePostCategories(p ...*PostCategory) *PostCategoryUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcu.RemovePostCategoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *PostCategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pcu.sqlSave, pcu.mutation, pcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *PostCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *PostCategoryUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *PostCategoryUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcu *PostCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(postcategory.Table, postcategory.Columns, sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeUUID))
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.Name(); ok {
		_spec.SetField(postcategory.FieldName, field.TypeString, value)
	}
	if value, ok := pcu.mutation.CreatedAt(); ok {
		_spec.SetField(postcategory.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pcu.mutation.UpdatedAt(); ok {
		_spec.SetField(postcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pcu.mutation.DeletedAt(); ok {
		_spec.SetField(postcategory.FieldDeletedAt, field.TypeTime, value)
	}
	if pcu.mutation.DeletedAtCleared() {
		_spec.ClearField(postcategory.FieldDeletedAt, field.TypeTime)
	}
	if pcu.mutation.PostCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   postcategory.PostCategoriesTable,
			Columns: postcategory.PostCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RemovedPostCategoriesIDs(); len(nodes) > 0 && !pcu.mutation.PostCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   postcategory.PostCategoriesTable,
			Columns: postcategory.PostCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.PostCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   postcategory.PostCategoriesTable,
			Columns: postcategory.PostCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{postcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pcu.mutation.done = true
	return n, nil
}

// PostCategoryUpdateOne is the builder for updating a single PostCategory entity.
type PostCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostCategoryMutation
}

// SetName sets the "name" field.
func (pcuo *PostCategoryUpdateOne) SetName(s string) *PostCategoryUpdateOne {
	pcuo.mutation.SetName(s)
	return pcuo
}

// SetCreatedAt sets the "created_at" field.
func (pcuo *PostCategoryUpdateOne) SetCreatedAt(t time.Time) *PostCategoryUpdateOne {
	pcuo.mutation.SetCreatedAt(t)
	return pcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pcuo *PostCategoryUpdateOne) SetUpdatedAt(t time.Time) *PostCategoryUpdateOne {
	pcuo.mutation.SetUpdatedAt(t)
	return pcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (pcuo *PostCategoryUpdateOne) SetDeletedAt(t time.Time) *PostCategoryUpdateOne {
	pcuo.mutation.SetDeletedAt(t)
	return pcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pcuo *PostCategoryUpdateOne) SetNillableDeletedAt(t *time.Time) *PostCategoryUpdateOne {
	if t != nil {
		pcuo.SetDeletedAt(*t)
	}
	return pcuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pcuo *PostCategoryUpdateOne) ClearDeletedAt() *PostCategoryUpdateOne {
	pcuo.mutation.ClearDeletedAt()
	return pcuo
}

// AddPostCategoryIDs adds the "post_categories" edge to the PostCategory entity by IDs.
func (pcuo *PostCategoryUpdateOne) AddPostCategoryIDs(ids ...uuid.UUID) *PostCategoryUpdateOne {
	pcuo.mutation.AddPostCategoryIDs(ids...)
	return pcuo
}

// AddPostCategories adds the "post_categories" edges to the PostCategory entity.
func (pcuo *PostCategoryUpdateOne) AddPostCategories(p ...*PostCategory) *PostCategoryUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcuo.AddPostCategoryIDs(ids...)
}

// Mutation returns the PostCategoryMutation object of the builder.
func (pcuo *PostCategoryUpdateOne) Mutation() *PostCategoryMutation {
	return pcuo.mutation
}

// ClearPostCategories clears all "post_categories" edges to the PostCategory entity.
func (pcuo *PostCategoryUpdateOne) ClearPostCategories() *PostCategoryUpdateOne {
	pcuo.mutation.ClearPostCategories()
	return pcuo
}

// RemovePostCategoryIDs removes the "post_categories" edge to PostCategory entities by IDs.
func (pcuo *PostCategoryUpdateOne) RemovePostCategoryIDs(ids ...uuid.UUID) *PostCategoryUpdateOne {
	pcuo.mutation.RemovePostCategoryIDs(ids...)
	return pcuo
}

// RemovePostCategories removes "post_categories" edges to PostCategory entities.
func (pcuo *PostCategoryUpdateOne) RemovePostCategories(p ...*PostCategory) *PostCategoryUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcuo.RemovePostCategoryIDs(ids...)
}

// Where appends a list predicates to the PostCategoryUpdate builder.
func (pcuo *PostCategoryUpdateOne) Where(ps ...predicate.PostCategory) *PostCategoryUpdateOne {
	pcuo.mutation.Where(ps...)
	return pcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *PostCategoryUpdateOne) Select(field string, fields ...string) *PostCategoryUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated PostCategory entity.
func (pcuo *PostCategoryUpdateOne) Save(ctx context.Context) (*PostCategory, error) {
	return withHooks(ctx, pcuo.sqlSave, pcuo.mutation, pcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *PostCategoryUpdateOne) SaveX(ctx context.Context) *PostCategory {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *PostCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *PostCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcuo *PostCategoryUpdateOne) sqlSave(ctx context.Context) (_node *PostCategory, err error) {
	_spec := sqlgraph.NewUpdateSpec(postcategory.Table, postcategory.Columns, sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeUUID))
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PostCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, postcategory.FieldID)
		for _, f := range fields {
			if !postcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != postcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.Name(); ok {
		_spec.SetField(postcategory.FieldName, field.TypeString, value)
	}
	if value, ok := pcuo.mutation.CreatedAt(); ok {
		_spec.SetField(postcategory.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(postcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pcuo.mutation.DeletedAt(); ok {
		_spec.SetField(postcategory.FieldDeletedAt, field.TypeTime, value)
	}
	if pcuo.mutation.DeletedAtCleared() {
		_spec.ClearField(postcategory.FieldDeletedAt, field.TypeTime)
	}
	if pcuo.mutation.PostCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   postcategory.PostCategoriesTable,
			Columns: postcategory.PostCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RemovedPostCategoriesIDs(); len(nodes) > 0 && !pcuo.mutation.PostCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   postcategory.PostCategoriesTable,
			Columns: postcategory.PostCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.PostCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   postcategory.PostCategoriesTable,
			Columns: postcategory.PostCategoriesPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PostCategory{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{postcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pcuo.mutation.done = true
	return _node, nil
}