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
	"github.com/marianozunino/cc-backend-go/ent/applicantprofile"
	"github.com/marianozunino/cc-backend-go/ent/applicantprofileskill"
	"github.com/marianozunino/cc-backend-go/ent/skill"
)

// ApplicantProfileSkillCreate is the builder for creating a ApplicantProfileSkill entity.
type ApplicantProfileSkillCreate struct {
	config
	mutation *ApplicantProfileSkillMutation
	hooks    []Hook
}

// SetApplicantProfileID sets the "applicant_profile_id" field.
func (apsc *ApplicantProfileSkillCreate) SetApplicantProfileID(u uuid.UUID) *ApplicantProfileSkillCreate {
	apsc.mutation.SetApplicantProfileID(u)
	return apsc
}

// SetLevel sets the "level" field.
func (apsc *ApplicantProfileSkillCreate) SetLevel(a applicantprofileskill.Level) *ApplicantProfileSkillCreate {
	apsc.mutation.SetLevel(a)
	return apsc
}

// SetSkillID sets the "skill_id" field.
func (apsc *ApplicantProfileSkillCreate) SetSkillID(u uuid.UUID) *ApplicantProfileSkillCreate {
	apsc.mutation.SetSkillID(u)
	return apsc
}

// SetNillableSkillID sets the "skill_id" field if the given value is not nil.
func (apsc *ApplicantProfileSkillCreate) SetNillableSkillID(u *uuid.UUID) *ApplicantProfileSkillCreate {
	if u != nil {
		apsc.SetSkillID(*u)
	}
	return apsc
}

// SetCreatedAt sets the "created_at" field.
func (apsc *ApplicantProfileSkillCreate) SetCreatedAt(t time.Time) *ApplicantProfileSkillCreate {
	apsc.mutation.SetCreatedAt(t)
	return apsc
}

// SetUpdatedAt sets the "updated_at" field.
func (apsc *ApplicantProfileSkillCreate) SetUpdatedAt(t time.Time) *ApplicantProfileSkillCreate {
	apsc.mutation.SetUpdatedAt(t)
	return apsc
}

// SetDeletedAt sets the "deleted_at" field.
func (apsc *ApplicantProfileSkillCreate) SetDeletedAt(t time.Time) *ApplicantProfileSkillCreate {
	apsc.mutation.SetDeletedAt(t)
	return apsc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (apsc *ApplicantProfileSkillCreate) SetNillableDeletedAt(t *time.Time) *ApplicantProfileSkillCreate {
	if t != nil {
		apsc.SetDeletedAt(*t)
	}
	return apsc
}

// SetID sets the "id" field.
func (apsc *ApplicantProfileSkillCreate) SetID(u uuid.UUID) *ApplicantProfileSkillCreate {
	apsc.mutation.SetID(u)
	return apsc
}

// SetApplicantProfile sets the "applicant_profile" edge to the ApplicantProfile entity.
func (apsc *ApplicantProfileSkillCreate) SetApplicantProfile(a *ApplicantProfile) *ApplicantProfileSkillCreate {
	return apsc.SetApplicantProfileID(a.ID)
}

// SetSkill sets the "skill" edge to the Skill entity.
func (apsc *ApplicantProfileSkillCreate) SetSkill(s *Skill) *ApplicantProfileSkillCreate {
	return apsc.SetSkillID(s.ID)
}

// Mutation returns the ApplicantProfileSkillMutation object of the builder.
func (apsc *ApplicantProfileSkillCreate) Mutation() *ApplicantProfileSkillMutation {
	return apsc.mutation
}

// Save creates the ApplicantProfileSkill in the database.
func (apsc *ApplicantProfileSkillCreate) Save(ctx context.Context) (*ApplicantProfileSkill, error) {
	return withHooks(ctx, apsc.sqlSave, apsc.mutation, apsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (apsc *ApplicantProfileSkillCreate) SaveX(ctx context.Context) *ApplicantProfileSkill {
	v, err := apsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apsc *ApplicantProfileSkillCreate) Exec(ctx context.Context) error {
	_, err := apsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apsc *ApplicantProfileSkillCreate) ExecX(ctx context.Context) {
	if err := apsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apsc *ApplicantProfileSkillCreate) check() error {
	if _, ok := apsc.mutation.ApplicantProfileID(); !ok {
		return &ValidationError{Name: "applicant_profile_id", err: errors.New(`ent: missing required field "ApplicantProfileSkill.applicant_profile_id"`)}
	}
	if _, ok := apsc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "ApplicantProfileSkill.level"`)}
	}
	if v, ok := apsc.mutation.Level(); ok {
		if err := applicantprofileskill.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "ApplicantProfileSkill.level": %w`, err)}
		}
	}
	if _, ok := apsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ApplicantProfileSkill.created_at"`)}
	}
	if _, ok := apsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ApplicantProfileSkill.updated_at"`)}
	}
	if _, ok := apsc.mutation.ApplicantProfileID(); !ok {
		return &ValidationError{Name: "applicant_profile", err: errors.New(`ent: missing required edge "ApplicantProfileSkill.applicant_profile"`)}
	}
	return nil
}

func (apsc *ApplicantProfileSkillCreate) sqlSave(ctx context.Context) (*ApplicantProfileSkill, error) {
	if err := apsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := apsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, apsc.driver, _spec); err != nil {
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
	apsc.mutation.id = &_node.ID
	apsc.mutation.done = true
	return _node, nil
}

func (apsc *ApplicantProfileSkillCreate) createSpec() (*ApplicantProfileSkill, *sqlgraph.CreateSpec) {
	var (
		_node = &ApplicantProfileSkill{config: apsc.config}
		_spec = sqlgraph.NewCreateSpec(applicantprofileskill.Table, sqlgraph.NewFieldSpec(applicantprofileskill.FieldID, field.TypeUUID))
	)
	if id, ok := apsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := apsc.mutation.Level(); ok {
		_spec.SetField(applicantprofileskill.FieldLevel, field.TypeEnum, value)
		_node.Level = value
	}
	if value, ok := apsc.mutation.CreatedAt(); ok {
		_spec.SetField(applicantprofileskill.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := apsc.mutation.UpdatedAt(); ok {
		_spec.SetField(applicantprofileskill.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := apsc.mutation.DeletedAt(); ok {
		_spec.SetField(applicantprofileskill.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := apsc.mutation.ApplicantProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   applicantprofileskill.ApplicantProfileTable,
			Columns: []string{applicantprofileskill.ApplicantProfileColumn},
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
	if nodes := apsc.mutation.SkillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   applicantprofileskill.SkillTable,
			Columns: []string{applicantprofileskill.SkillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SkillID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ApplicantProfileSkillCreateBulk is the builder for creating many ApplicantProfileSkill entities in bulk.
type ApplicantProfileSkillCreateBulk struct {
	config
	builders []*ApplicantProfileSkillCreate
}

// Save creates the ApplicantProfileSkill entities in the database.
func (apscb *ApplicantProfileSkillCreateBulk) Save(ctx context.Context) ([]*ApplicantProfileSkill, error) {
	specs := make([]*sqlgraph.CreateSpec, len(apscb.builders))
	nodes := make([]*ApplicantProfileSkill, len(apscb.builders))
	mutators := make([]Mutator, len(apscb.builders))
	for i := range apscb.builders {
		func(i int, root context.Context) {
			builder := apscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicantProfileSkillMutation)
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
					_, err = mutators[i+1].Mutate(root, apscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, apscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, apscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (apscb *ApplicantProfileSkillCreateBulk) SaveX(ctx context.Context) []*ApplicantProfileSkill {
	v, err := apscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apscb *ApplicantProfileSkillCreateBulk) Exec(ctx context.Context) error {
	_, err := apscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apscb *ApplicantProfileSkillCreateBulk) ExecX(ctx context.Context) {
	if err := apscb.Exec(ctx); err != nil {
		panic(err)
	}
}