// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marianozunino/jobby/ent/predicate"
	"github.com/marianozunino/jobby/ent/workexperience"
)

// WorkExperienceDelete is the builder for deleting a WorkExperience entity.
type WorkExperienceDelete struct {
	config
	hooks    []Hook
	mutation *WorkExperienceMutation
}

// Where appends a list predicates to the WorkExperienceDelete builder.
func (wed *WorkExperienceDelete) Where(ps ...predicate.WorkExperience) *WorkExperienceDelete {
	wed.mutation.Where(ps...)
	return wed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wed *WorkExperienceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wed.sqlExec, wed.mutation, wed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wed *WorkExperienceDelete) ExecX(ctx context.Context) int {
	n, err := wed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wed *WorkExperienceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(workexperience.Table, sqlgraph.NewFieldSpec(workexperience.FieldID, field.TypeUUID))
	if ps := wed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wed.mutation.done = true
	return affected, err
}

// WorkExperienceDeleteOne is the builder for deleting a single WorkExperience entity.
type WorkExperienceDeleteOne struct {
	wed *WorkExperienceDelete
}

// Where appends a list predicates to the WorkExperienceDelete builder.
func (wedo *WorkExperienceDeleteOne) Where(ps ...predicate.WorkExperience) *WorkExperienceDeleteOne {
	wedo.wed.mutation.Where(ps...)
	return wedo
}

// Exec executes the deletion query.
func (wedo *WorkExperienceDeleteOne) Exec(ctx context.Context) error {
	n, err := wedo.wed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workexperience.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wedo *WorkExperienceDeleteOne) ExecX(ctx context.Context) {
	if err := wedo.Exec(ctx); err != nil {
		panic(err)
	}
}
