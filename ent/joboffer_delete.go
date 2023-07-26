// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/marianozunino/cc-backend-go/ent/joboffer"
	"github.com/marianozunino/cc-backend-go/ent/predicate"
)

// JobOfferDelete is the builder for deleting a JobOffer entity.
type JobOfferDelete struct {
	config
	hooks    []Hook
	mutation *JobOfferMutation
}

// Where appends a list predicates to the JobOfferDelete builder.
func (jod *JobOfferDelete) Where(ps ...predicate.JobOffer) *JobOfferDelete {
	jod.mutation.Where(ps...)
	return jod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jod *JobOfferDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, jod.sqlExec, jod.mutation, jod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (jod *JobOfferDelete) ExecX(ctx context.Context) int {
	n, err := jod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jod *JobOfferDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(joboffer.Table, sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeUUID))
	if ps := jod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, jod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	jod.mutation.done = true
	return affected, err
}

// JobOfferDeleteOne is the builder for deleting a single JobOffer entity.
type JobOfferDeleteOne struct {
	jod *JobOfferDelete
}

// Where appends a list predicates to the JobOfferDelete builder.
func (jodo *JobOfferDeleteOne) Where(ps ...predicate.JobOffer) *JobOfferDeleteOne {
	jodo.jod.mutation.Where(ps...)
	return jodo
}

// Exec executes the deletion query.
func (jodo *JobOfferDeleteOne) Exec(ctx context.Context) error {
	n, err := jodo.jod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{joboffer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jodo *JobOfferDeleteOne) ExecX(ctx context.Context) {
	if err := jodo.Exec(ctx); err != nil {
		panic(err)
	}
}
