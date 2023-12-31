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
	"github.com/marianozunino/jobby/ent/applicantprofile"
	"github.com/marianozunino/jobby/ent/application"
	"github.com/marianozunino/jobby/ent/interview"
	"github.com/marianozunino/jobby/ent/joboffer"
)

// ApplicationCreate is the builder for creating a Application entity.
type ApplicationCreate struct {
	config
	mutation *ApplicationMutation
	hooks    []Hook
}

// SetJobOfferID sets the "job_offer_id" field.
func (ac *ApplicationCreate) SetJobOfferID(u uuid.UUID) *ApplicationCreate {
	ac.mutation.SetJobOfferID(u)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *ApplicationCreate) SetCreatedAt(t time.Time) *ApplicationCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableCreatedAt(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *ApplicationCreate) SetUpdatedAt(t time.Time) *ApplicationCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableUpdatedAt(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *ApplicationCreate) SetDeletedAt(t time.Time) *ApplicationCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableDeletedAt(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetApplicantProfileID sets the "applicant_profile_id" field.
func (ac *ApplicationCreate) SetApplicantProfileID(u uuid.UUID) *ApplicationCreate {
	ac.mutation.SetApplicantProfileID(u)
	return ac
}

// SetID sets the "id" field.
func (ac *ApplicationCreate) SetID(u uuid.UUID) *ApplicationCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetApplicantProfile sets the "applicant_profile" edge to the ApplicantProfile entity.
func (ac *ApplicationCreate) SetApplicantProfile(a *ApplicantProfile) *ApplicationCreate {
	return ac.SetApplicantProfileID(a.ID)
}

// SetJobOffer sets the "job_offer" edge to the JobOffer entity.
func (ac *ApplicationCreate) SetJobOffer(j *JobOffer) *ApplicationCreate {
	return ac.SetJobOfferID(j.ID)
}

// AddInterviewIDs adds the "interviews" edge to the Interview entity by IDs.
func (ac *ApplicationCreate) AddInterviewIDs(ids ...uuid.UUID) *ApplicationCreate {
	ac.mutation.AddInterviewIDs(ids...)
	return ac
}

// AddInterviews adds the "interviews" edges to the Interview entity.
func (ac *ApplicationCreate) AddInterviews(i ...*Interview) *ApplicationCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ac.AddInterviewIDs(ids...)
}

// Mutation returns the ApplicationMutation object of the builder.
func (ac *ApplicationCreate) Mutation() *ApplicationMutation {
	return ac.mutation
}

// Save creates the Application in the database.
func (ac *ApplicationCreate) Save(ctx context.Context) (*Application, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ApplicationCreate) SaveX(ctx context.Context) *Application {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ApplicationCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ApplicationCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ApplicationCreate) check() error {
	if _, ok := ac.mutation.JobOfferID(); !ok {
		return &ValidationError{Name: "job_offer_id", err: errors.New(`ent: missing required field "Application.job_offer_id"`)}
	}
	if _, ok := ac.mutation.ApplicantProfileID(); !ok {
		return &ValidationError{Name: "applicant_profile_id", err: errors.New(`ent: missing required field "Application.applicant_profile_id"`)}
	}
	if _, ok := ac.mutation.ApplicantProfileID(); !ok {
		return &ValidationError{Name: "applicant_profile", err: errors.New(`ent: missing required edge "Application.applicant_profile"`)}
	}
	if _, ok := ac.mutation.JobOfferID(); !ok {
		return &ValidationError{Name: "job_offer", err: errors.New(`ent: missing required edge "Application.job_offer"`)}
	}
	return nil
}

func (ac *ApplicationCreate) sqlSave(ctx context.Context) (*Application, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ApplicationCreate) createSpec() (*Application, *sqlgraph.CreateSpec) {
	var (
		_node = &Application{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(application.Table, sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(application.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(application.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(application.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := ac.mutation.ApplicantProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.ApplicantProfileTable,
			Columns: []string{application.ApplicantProfileColumn},
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
	if nodes := ac.mutation.JobOfferIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   application.JobOfferTable,
			Columns: []string{application.JobOfferColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.JobOfferID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.InterviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.InterviewsTable,
			Columns: []string{application.InterviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(interview.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ApplicationCreateBulk is the builder for creating many Application entities in bulk.
type ApplicationCreateBulk struct {
	config
	builders []*ApplicationCreate
}

// Save creates the Application entities in the database.
func (acb *ApplicationCreateBulk) Save(ctx context.Context) ([]*Application, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Application, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ApplicationCreateBulk) SaveX(ctx context.Context) []*Application {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ApplicationCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ApplicationCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
