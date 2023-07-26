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
	"github.com/marianozunino/cc-backend-go/ent/contactusmessage"
)

// ContactUsMessageCreate is the builder for creating a ContactUsMessage entity.
type ContactUsMessageCreate struct {
	config
	mutation *ContactUsMessageMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cumc *ContactUsMessageCreate) SetName(s string) *ContactUsMessageCreate {
	cumc.mutation.SetName(s)
	return cumc
}

// SetEmail sets the "email" field.
func (cumc *ContactUsMessageCreate) SetEmail(s string) *ContactUsMessageCreate {
	cumc.mutation.SetEmail(s)
	return cumc
}

// SetMessage sets the "message" field.
func (cumc *ContactUsMessageCreate) SetMessage(s string) *ContactUsMessageCreate {
	cumc.mutation.SetMessage(s)
	return cumc
}

// SetCreatedAt sets the "created_at" field.
func (cumc *ContactUsMessageCreate) SetCreatedAt(t time.Time) *ContactUsMessageCreate {
	cumc.mutation.SetCreatedAt(t)
	return cumc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cumc *ContactUsMessageCreate) SetNillableCreatedAt(t *time.Time) *ContactUsMessageCreate {
	if t != nil {
		cumc.SetCreatedAt(*t)
	}
	return cumc
}

// SetUpdatedAt sets the "updated_at" field.
func (cumc *ContactUsMessageCreate) SetUpdatedAt(t time.Time) *ContactUsMessageCreate {
	cumc.mutation.SetUpdatedAt(t)
	return cumc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cumc *ContactUsMessageCreate) SetNillableUpdatedAt(t *time.Time) *ContactUsMessageCreate {
	if t != nil {
		cumc.SetUpdatedAt(*t)
	}
	return cumc
}

// SetDeletedAt sets the "deleted_at" field.
func (cumc *ContactUsMessageCreate) SetDeletedAt(t time.Time) *ContactUsMessageCreate {
	cumc.mutation.SetDeletedAt(t)
	return cumc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cumc *ContactUsMessageCreate) SetNillableDeletedAt(t *time.Time) *ContactUsMessageCreate {
	if t != nil {
		cumc.SetDeletedAt(*t)
	}
	return cumc
}

// SetPhone sets the "phone" field.
func (cumc *ContactUsMessageCreate) SetPhone(s string) *ContactUsMessageCreate {
	cumc.mutation.SetPhone(s)
	return cumc
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (cumc *ContactUsMessageCreate) SetNillablePhone(s *string) *ContactUsMessageCreate {
	if s != nil {
		cumc.SetPhone(*s)
	}
	return cumc
}

// SetID sets the "id" field.
func (cumc *ContactUsMessageCreate) SetID(u uuid.UUID) *ContactUsMessageCreate {
	cumc.mutation.SetID(u)
	return cumc
}

// Mutation returns the ContactUsMessageMutation object of the builder.
func (cumc *ContactUsMessageCreate) Mutation() *ContactUsMessageMutation {
	return cumc.mutation
}

// Save creates the ContactUsMessage in the database.
func (cumc *ContactUsMessageCreate) Save(ctx context.Context) (*ContactUsMessage, error) {
	return withHooks(ctx, cumc.sqlSave, cumc.mutation, cumc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cumc *ContactUsMessageCreate) SaveX(ctx context.Context) *ContactUsMessage {
	v, err := cumc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cumc *ContactUsMessageCreate) Exec(ctx context.Context) error {
	_, err := cumc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cumc *ContactUsMessageCreate) ExecX(ctx context.Context) {
	if err := cumc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cumc *ContactUsMessageCreate) check() error {
	if _, ok := cumc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ContactUsMessage.name"`)}
	}
	if _, ok := cumc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "ContactUsMessage.email"`)}
	}
	if _, ok := cumc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "ContactUsMessage.message"`)}
	}
	return nil
}

func (cumc *ContactUsMessageCreate) sqlSave(ctx context.Context) (*ContactUsMessage, error) {
	if err := cumc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cumc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cumc.driver, _spec); err != nil {
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
	cumc.mutation.id = &_node.ID
	cumc.mutation.done = true
	return _node, nil
}

func (cumc *ContactUsMessageCreate) createSpec() (*ContactUsMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &ContactUsMessage{config: cumc.config}
		_spec = sqlgraph.NewCreateSpec(contactusmessage.Table, sqlgraph.NewFieldSpec(contactusmessage.FieldID, field.TypeUUID))
	)
	if id, ok := cumc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cumc.mutation.Name(); ok {
		_spec.SetField(contactusmessage.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cumc.mutation.Email(); ok {
		_spec.SetField(contactusmessage.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := cumc.mutation.Message(); ok {
		_spec.SetField(contactusmessage.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if value, ok := cumc.mutation.CreatedAt(); ok {
		_spec.SetField(contactusmessage.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cumc.mutation.UpdatedAt(); ok {
		_spec.SetField(contactusmessage.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cumc.mutation.DeletedAt(); ok {
		_spec.SetField(contactusmessage.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cumc.mutation.Phone(); ok {
		_spec.SetField(contactusmessage.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	return _node, _spec
}

// ContactUsMessageCreateBulk is the builder for creating many ContactUsMessage entities in bulk.
type ContactUsMessageCreateBulk struct {
	config
	builders []*ContactUsMessageCreate
}

// Save creates the ContactUsMessage entities in the database.
func (cumcb *ContactUsMessageCreateBulk) Save(ctx context.Context) ([]*ContactUsMessage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cumcb.builders))
	nodes := make([]*ContactUsMessage, len(cumcb.builders))
	mutators := make([]Mutator, len(cumcb.builders))
	for i := range cumcb.builders {
		func(i int, root context.Context) {
			builder := cumcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContactUsMessageMutation)
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
					_, err = mutators[i+1].Mutate(root, cumcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cumcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cumcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cumcb *ContactUsMessageCreateBulk) SaveX(ctx context.Context) []*ContactUsMessage {
	v, err := cumcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cumcb *ContactUsMessageCreateBulk) Exec(ctx context.Context) error {
	_, err := cumcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cumcb *ContactUsMessageCreateBulk) ExecX(ctx context.Context) {
	if err := cumcb.Exec(ctx); err != nil {
		panic(err)
	}
}
