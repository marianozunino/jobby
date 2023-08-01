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
	"github.com/marianozunino/cc-backend-go/ent/joboffer"
	"github.com/marianozunino/cc-backend-go/ent/status"
)

// StatusCreate is the builder for creating a Status entity.
type StatusCreate struct {
	config
	mutation *StatusMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *StatusCreate) SetName(s string) *StatusCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *StatusCreate) SetCreatedAt(t time.Time) *StatusCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StatusCreate) SetNillableCreatedAt(t *time.Time) *StatusCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StatusCreate) SetUpdatedAt(t time.Time) *StatusCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StatusCreate) SetNillableUpdatedAt(t *time.Time) *StatusCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *StatusCreate) SetDeletedAt(t time.Time) *StatusCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *StatusCreate) SetNillableDeletedAt(t *time.Time) *StatusCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StatusCreate) SetID(u uuid.UUID) *StatusCreate {
	sc.mutation.SetID(u)
	return sc
}

// AddJobOfferIDs adds the "job_offers" edge to the JobOffer entity by IDs.
func (sc *StatusCreate) AddJobOfferIDs(ids ...uuid.UUID) *StatusCreate {
	sc.mutation.AddJobOfferIDs(ids...)
	return sc
}

// AddJobOffers adds the "job_offers" edges to the JobOffer entity.
func (sc *StatusCreate) AddJobOffers(j ...*JobOffer) *StatusCreate {
	ids := make([]uuid.UUID, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return sc.AddJobOfferIDs(ids...)
}

// Mutation returns the StatusMutation object of the builder.
func (sc *StatusCreate) Mutation() *StatusMutation {
	return sc.mutation
}

// Save creates the Status in the database.
func (sc *StatusCreate) Save(ctx context.Context) (*Status, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StatusCreate) SaveX(ctx context.Context) *Status {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StatusCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StatusCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StatusCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Status.name"`)}
	}
	return nil
}

func (sc *StatusCreate) sqlSave(ctx context.Context) (*Status, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StatusCreate) createSpec() (*Status, *sqlgraph.CreateSpec) {
	var (
		_node = &Status{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(status.Table, sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(status.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(status.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(status.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.SetField(status.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := sc.mutation.JobOffersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   status.JobOffersTable,
			Columns: []string{status.JobOffersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StatusCreateBulk is the builder for creating many Status entities in bulk.
type StatusCreateBulk struct {
	config
	builders []*StatusCreate
}

// Save creates the Status entities in the database.
func (scb *StatusCreateBulk) Save(ctx context.Context) ([]*Status, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Status, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StatusMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StatusCreateBulk) SaveX(ctx context.Context) []*Status {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StatusCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StatusCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}