// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/marianozunino/jobby/ent/applicantinterest"
	"github.com/marianozunino/jobby/ent/applicantprofile"
	"github.com/marianozunino/jobby/ent/category"
)

// ApplicantInterestCreate is the builder for creating a ApplicantInterest entity.
type ApplicantInterestCreate struct {
	config
	mutation *ApplicantInterestMutation
	hooks    []Hook
}

// SetCategoryID sets the "category_id" field.
func (aic *ApplicantInterestCreate) SetCategoryID(u uuid.UUID) *ApplicantInterestCreate {
	aic.mutation.SetCategoryID(u)
	return aic
}

// SetCreatedAt sets the "created_at" field.
func (aic *ApplicantInterestCreate) SetCreatedAt(t time.Time) *ApplicantInterestCreate {
	aic.mutation.SetCreatedAt(t)
	return aic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aic *ApplicantInterestCreate) SetNillableCreatedAt(t *time.Time) *ApplicantInterestCreate {
	if t != nil {
		aic.SetCreatedAt(*t)
	}
	return aic
}

// SetUpdatedAt sets the "updated_at" field.
func (aic *ApplicantInterestCreate) SetUpdatedAt(t time.Time) *ApplicantInterestCreate {
	aic.mutation.SetUpdatedAt(t)
	return aic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aic *ApplicantInterestCreate) SetNillableUpdatedAt(t *time.Time) *ApplicantInterestCreate {
	if t != nil {
		aic.SetUpdatedAt(*t)
	}
	return aic
}

// SetDeletedAt sets the "deleted_at" field.
func (aic *ApplicantInterestCreate) SetDeletedAt(t time.Time) *ApplicantInterestCreate {
	aic.mutation.SetDeletedAt(t)
	return aic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aic *ApplicantInterestCreate) SetNillableDeletedAt(t *time.Time) *ApplicantInterestCreate {
	if t != nil {
		aic.SetDeletedAt(*t)
	}
	return aic
}

// SetApplicantProfileID sets the "applicant_profile_id" field.
func (aic *ApplicantInterestCreate) SetApplicantProfileID(u uuid.UUID) *ApplicantInterestCreate {
	aic.mutation.SetApplicantProfileID(u)
	return aic
}

// SetID sets the "id" field.
func (aic *ApplicantInterestCreate) SetID(u uuid.UUID) *ApplicantInterestCreate {
	aic.mutation.SetID(u)
	return aic
}

// SetApplicantProfile sets the "applicant_profile" edge to the ApplicantProfile entity.
func (aic *ApplicantInterestCreate) SetApplicantProfile(a *ApplicantProfile) *ApplicantInterestCreate {
	return aic.SetApplicantProfileID(a.ID)
}

// SetCategory sets the "category" edge to the Category entity.
func (aic *ApplicantInterestCreate) SetCategory(c *Category) *ApplicantInterestCreate {
	return aic.SetCategoryID(c.ID)
}

// Mutation returns the ApplicantInterestMutation object of the builder.
func (aic *ApplicantInterestCreate) Mutation() *ApplicantInterestMutation {
	return aic.mutation
}

// Save creates the ApplicantInterest in the database.
func (aic *ApplicantInterestCreate) Save(ctx context.Context) (*ApplicantInterest, error) {
	return withHooks(ctx, aic.sqlSave, aic.mutation, aic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (aic *ApplicantInterestCreate) SaveX(ctx context.Context) *ApplicantInterest {
	v, err := aic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aic *ApplicantInterestCreate) Exec(ctx context.Context) error {
	_, err := aic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aic *ApplicantInterestCreate) ExecX(ctx context.Context) {
	if err := aic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aic *ApplicantInterestCreate) check() error {
	if _, ok := aic.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category_id", err: errors.New(`ent: missing required field "ApplicantInterest.category_id"`)}
	}
	if _, ok := aic.mutation.ApplicantProfileID(); !ok {
		return &ValidationError{Name: "applicant_profile_id", err: errors.New(`ent: missing required field "ApplicantInterest.applicant_profile_id"`)}
	}
	if _, ok := aic.mutation.ApplicantProfileID(); !ok {
		return &ValidationError{Name: "applicant_profile", err: errors.New(`ent: missing required edge "ApplicantInterest.applicant_profile"`)}
	}
	if _, ok := aic.mutation.CategoryID(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required edge "ApplicantInterest.category"`)}
	}
	return nil
}

func (aic *ApplicantInterestCreate) sqlSave(ctx context.Context) (*ApplicantInterest, error) {
	if err := aic.check(); err != nil {
		return nil, err
	}
	_node, _spec := aic.createSpec()
	if err := sqlgraph.CreateNode(ctx, aic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	aic.mutation.id = &_node.ID
	aic.mutation.done = true
	return _node, nil
}

func (aic *ApplicantInterestCreate) createSpec() (*ApplicantInterest, *sqlgraph.CreateSpec) {
	var (
		_node = &ApplicantInterest{config: aic.config}
		_spec = sqlgraph.NewCreateSpec(applicantinterest.Table, sqlgraph.NewFieldSpec(applicantinterest.FieldID, field.TypeUUID))
	)
	if id, ok := aic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := aic.mutation.CreatedAt(); ok {
		_spec.SetField(applicantinterest.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := aic.mutation.UpdatedAt(); ok {
		_spec.SetField(applicantinterest.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := aic.mutation.DeletedAt(); ok {
		_spec.SetField(applicantinterest.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := aic.mutation.ApplicantProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   applicantinterest.ApplicantProfileTable,
			Columns: []string{applicantinterest.ApplicantProfileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(applicantprofile.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ApplicantProfileID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := aic.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   applicantinterest.CategoryTable,
			Columns: []string{applicantinterest.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ApplicantInterestCreateBulk is the builder for creating many ApplicantInterest entities in bulk.
type ApplicantInterestCreateBulk struct {
	config
	builders []*ApplicantInterestCreate
}

// Save creates the ApplicantInterest entities in the database.
func (aicb *ApplicantInterestCreateBulk) Save(ctx context.Context) ([]*ApplicantInterest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aicb.builders))
	nodes := make([]*ApplicantInterest, len(aicb.builders))
	mutators := make([]Mutator, len(aicb.builders))
	for i := range aicb.builders {
		func(i int, root context.Context) {
			builder := aicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicantInterestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aicb *ApplicantInterestCreateBulk) SaveX(ctx context.Context) []*ApplicantInterest {
	v, err := aicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aicb *ApplicantInterestCreateBulk) Exec(ctx context.Context) error {
	_, err := aicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aicb *ApplicantInterestCreateBulk) ExecX(ctx context.Context) {
	if err := aicb.Exec(ctx); err != nil {
		panic(err)
	}
}
