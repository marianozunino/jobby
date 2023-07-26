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
	"github.com/marianozunino/cc-backend-go/ent/applicantinterest"
	"github.com/marianozunino/cc-backend-go/ent/category"
	"github.com/marianozunino/cc-backend-go/ent/joboffercategory"
)

// CategoryCreate is the builder for creating a Category entity.
type CategoryCreate struct {
	config
	mutation *CategoryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CategoryCreate) SetName(s string) *CategoryCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetSlug sets the "slug" field.
func (cc *CategoryCreate) SetSlug(s string) *CategoryCreate {
	cc.mutation.SetSlug(s)
	return cc
}

// SetParentID sets the "parent_id" field.
func (cc *CategoryCreate) SetParentID(u uuid.UUID) *CategoryCreate {
	cc.mutation.SetParentID(u)
	return cc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableParentID(u *uuid.UUID) *CategoryCreate {
	if u != nil {
		cc.SetParentID(*u)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CategoryCreate) SetCreatedAt(t time.Time) *CategoryCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableCreatedAt(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CategoryCreate) SetUpdatedAt(t time.Time) *CategoryCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableUpdatedAt(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CategoryCreate) SetDeletedAt(t time.Time) *CategoryCreate {
	cc.mutation.SetDeletedAt(t)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CategoryCreate) SetNillableDeletedAt(t *time.Time) *CategoryCreate {
	if t != nil {
		cc.SetDeletedAt(*t)
	}
	return cc
}

// SetIsRoot sets the "is_root" field.
func (cc *CategoryCreate) SetIsRoot(b bool) *CategoryCreate {
	cc.mutation.SetIsRoot(b)
	return cc
}

// SetID sets the "id" field.
func (cc *CategoryCreate) SetID(u uuid.UUID) *CategoryCreate {
	cc.mutation.SetID(u)
	return cc
}

// AddApplicantInterestIDs adds the "applicant_interests" edge to the ApplicantInterest entity by IDs.
func (cc *CategoryCreate) AddApplicantInterestIDs(ids ...uuid.UUID) *CategoryCreate {
	cc.mutation.AddApplicantInterestIDs(ids...)
	return cc
}

// AddApplicantInterests adds the "applicant_interests" edges to the ApplicantInterest entity.
func (cc *CategoryCreate) AddApplicantInterests(a ...*ApplicantInterest) *CategoryCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddApplicantInterestIDs(ids...)
}

// AddJobOfferCategoryIDs adds the "job_offer_categories" edge to the JobOfferCategory entity by IDs.
func (cc *CategoryCreate) AddJobOfferCategoryIDs(ids ...uuid.UUID) *CategoryCreate {
	cc.mutation.AddJobOfferCategoryIDs(ids...)
	return cc
}

// AddJobOfferCategories adds the "job_offer_categories" edges to the JobOfferCategory entity.
func (cc *CategoryCreate) AddJobOfferCategories(j ...*JobOfferCategory) *CategoryCreate {
	ids := make([]uuid.UUID, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cc.AddJobOfferCategoryIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cc *CategoryCreate) Mutation() *CategoryMutation {
	return cc.mutation
}

// Save creates the Category in the database.
func (cc *CategoryCreate) Save(ctx context.Context) (*Category, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoryCreate) SaveX(ctx context.Context) *Category {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoryCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Category.name"`)}
	}
	if _, ok := cc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Category.slug"`)}
	}
	if _, ok := cc.mutation.IsRoot(); !ok {
		return &ValidationError{Name: "is_root", err: errors.New(`ent: missing required field "Category.is_root"`)}
	}
	return nil
}

func (cc *CategoryCreate) sqlSave(ctx context.Context) (*Category, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CategoryCreate) createSpec() (*Category, *sqlgraph.CreateSpec) {
	var (
		_node = &Category{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(category.Table, sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Slug(); ok {
		_spec.SetField(category.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := cc.mutation.ParentID(); ok {
		_spec.SetField(category.FieldParentID, field.TypeUUID, value)
		_node.ParentID = &value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(category.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.SetField(category.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.IsRoot(); ok {
		_spec.SetField(category.FieldIsRoot, field.TypeBool, value)
		_node.IsRoot = value
	}
	if nodes := cc.mutation.ApplicantInterestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ApplicantInterestsTable,
			Columns: []string{category.ApplicantInterestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(applicantinterest.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.JobOfferCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.JobOfferCategoriesTable,
			Columns: []string{category.JobOfferCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffercategory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CategoryCreateBulk is the builder for creating many Category entities in bulk.
type CategoryCreateBulk struct {
	config
	builders []*CategoryCreate
}

// Save creates the Category entities in the database.
func (ccb *CategoryCreateBulk) Save(ctx context.Context) ([]*Category, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Category, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CategoryCreateBulk) SaveX(ctx context.Context) []*Category {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
