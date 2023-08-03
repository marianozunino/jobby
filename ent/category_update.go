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
	"github.com/marianozunino/jobby/ent/applicantinterest"
	"github.com/marianozunino/jobby/ent/category"
	"github.com/marianozunino/jobby/ent/joboffercategory"
	"github.com/marianozunino/jobby/ent/predicate"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *CategoryUpdate) SetName(s string) *CategoryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetSlug sets the "slug" field.
func (cu *CategoryUpdate) SetSlug(s string) *CategoryUpdate {
	cu.mutation.SetSlug(s)
	return cu
}

// SetParentID sets the "parent_id" field.
func (cu *CategoryUpdate) SetParentID(u uuid.UUID) *CategoryUpdate {
	cu.mutation.SetParentID(u)
	return cu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableParentID(u *uuid.UUID) *CategoryUpdate {
	if u != nil {
		cu.SetParentID(*u)
	}
	return cu
}

// ClearParentID clears the value of the "parent_id" field.
func (cu *CategoryUpdate) ClearParentID() *CategoryUpdate {
	cu.mutation.ClearParentID()
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CategoryUpdate) SetCreatedAt(t time.Time) *CategoryUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableCreatedAt(t *time.Time) *CategoryUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (cu *CategoryUpdate) ClearCreatedAt() *CategoryUpdate {
	cu.mutation.ClearCreatedAt()
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CategoryUpdate) SetUpdatedAt(t time.Time) *CategoryUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableUpdatedAt(t *time.Time) *CategoryUpdate {
	if t != nil {
		cu.SetUpdatedAt(*t)
	}
	return cu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cu *CategoryUpdate) ClearUpdatedAt() *CategoryUpdate {
	cu.mutation.ClearUpdatedAt()
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CategoryUpdate) SetDeletedAt(t time.Time) *CategoryUpdate {
	cu.mutation.SetDeletedAt(t)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CategoryUpdate) SetNillableDeletedAt(t *time.Time) *CategoryUpdate {
	if t != nil {
		cu.SetDeletedAt(*t)
	}
	return cu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cu *CategoryUpdate) ClearDeletedAt() *CategoryUpdate {
	cu.mutation.ClearDeletedAt()
	return cu
}

// SetIsRoot sets the "is_root" field.
func (cu *CategoryUpdate) SetIsRoot(b bool) *CategoryUpdate {
	cu.mutation.SetIsRoot(b)
	return cu
}

// AddApplicantInterestIDs adds the "applicant_interests" edge to the ApplicantInterest entity by IDs.
func (cu *CategoryUpdate) AddApplicantInterestIDs(ids ...uuid.UUID) *CategoryUpdate {
	cu.mutation.AddApplicantInterestIDs(ids...)
	return cu
}

// AddApplicantInterests adds the "applicant_interests" edges to the ApplicantInterest entity.
func (cu *CategoryUpdate) AddApplicantInterests(a ...*ApplicantInterest) *CategoryUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddApplicantInterestIDs(ids...)
}

// AddChildCategoryIDs adds the "child_categories" edge to the Category entity by IDs.
func (cu *CategoryUpdate) AddChildCategoryIDs(ids ...uuid.UUID) *CategoryUpdate {
	cu.mutation.AddChildCategoryIDs(ids...)
	return cu
}

// AddChildCategories adds the "child_categories" edges to the Category entity.
func (cu *CategoryUpdate) AddChildCategories(c ...*Category) *CategoryUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddChildCategoryIDs(ids...)
}

// AddJobOfferCategoryIDs adds the "job_offer_categories" edge to the JobOfferCategory entity by IDs.
func (cu *CategoryUpdate) AddJobOfferCategoryIDs(ids ...uuid.UUID) *CategoryUpdate {
	cu.mutation.AddJobOfferCategoryIDs(ids...)
	return cu
}

// AddJobOfferCategories adds the "job_offer_categories" edges to the JobOfferCategory entity.
func (cu *CategoryUpdate) AddJobOfferCategories(j ...*JobOfferCategory) *CategoryUpdate {
	ids := make([]uuid.UUID, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cu.AddJobOfferCategoryIDs(ids...)
}

// SetParentCategoryID sets the "parent_category" edge to the Category entity by ID.
func (cu *CategoryUpdate) SetParentCategoryID(id uuid.UUID) *CategoryUpdate {
	cu.mutation.SetParentCategoryID(id)
	return cu
}

// SetNillableParentCategoryID sets the "parent_category" edge to the Category entity by ID if the given value is not nil.
func (cu *CategoryUpdate) SetNillableParentCategoryID(id *uuid.UUID) *CategoryUpdate {
	if id != nil {
		cu = cu.SetParentCategoryID(*id)
	}
	return cu
}

// SetParentCategory sets the "parent_category" edge to the Category entity.
func (cu *CategoryUpdate) SetParentCategory(c *Category) *CategoryUpdate {
	return cu.SetParentCategoryID(c.ID)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearApplicantInterests clears all "applicant_interests" edges to the ApplicantInterest entity.
func (cu *CategoryUpdate) ClearApplicantInterests() *CategoryUpdate {
	cu.mutation.ClearApplicantInterests()
	return cu
}

// RemoveApplicantInterestIDs removes the "applicant_interests" edge to ApplicantInterest entities by IDs.
func (cu *CategoryUpdate) RemoveApplicantInterestIDs(ids ...uuid.UUID) *CategoryUpdate {
	cu.mutation.RemoveApplicantInterestIDs(ids...)
	return cu
}

// RemoveApplicantInterests removes "applicant_interests" edges to ApplicantInterest entities.
func (cu *CategoryUpdate) RemoveApplicantInterests(a ...*ApplicantInterest) *CategoryUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveApplicantInterestIDs(ids...)
}

// ClearChildCategories clears all "child_categories" edges to the Category entity.
func (cu *CategoryUpdate) ClearChildCategories() *CategoryUpdate {
	cu.mutation.ClearChildCategories()
	return cu
}

// RemoveChildCategoryIDs removes the "child_categories" edge to Category entities by IDs.
func (cu *CategoryUpdate) RemoveChildCategoryIDs(ids ...uuid.UUID) *CategoryUpdate {
	cu.mutation.RemoveChildCategoryIDs(ids...)
	return cu
}

// RemoveChildCategories removes "child_categories" edges to Category entities.
func (cu *CategoryUpdate) RemoveChildCategories(c ...*Category) *CategoryUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveChildCategoryIDs(ids...)
}

// ClearJobOfferCategories clears all "job_offer_categories" edges to the JobOfferCategory entity.
func (cu *CategoryUpdate) ClearJobOfferCategories() *CategoryUpdate {
	cu.mutation.ClearJobOfferCategories()
	return cu
}

// RemoveJobOfferCategoryIDs removes the "job_offer_categories" edge to JobOfferCategory entities by IDs.
func (cu *CategoryUpdate) RemoveJobOfferCategoryIDs(ids ...uuid.UUID) *CategoryUpdate {
	cu.mutation.RemoveJobOfferCategoryIDs(ids...)
	return cu
}

// RemoveJobOfferCategories removes "job_offer_categories" edges to JobOfferCategory entities.
func (cu *CategoryUpdate) RemoveJobOfferCategories(j ...*JobOfferCategory) *CategoryUpdate {
	ids := make([]uuid.UUID, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cu.RemoveJobOfferCategoryIDs(ids...)
}

// ClearParentCategory clears the "parent_category" edge to the Category entity.
func (cu *CategoryUpdate) ClearParentCategory() *CategoryUpdate {
	cu.mutation.ClearParentCategory()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Slug(); ok {
		_spec.SetField(category.FieldSlug, field.TypeString, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(category.FieldCreatedAt, field.TypeTime, value)
	}
	if cu.mutation.CreatedAtCleared() {
		_spec.ClearField(category.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.UpdatedAtCleared() {
		_spec.ClearField(category.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.SetField(category.FieldDeletedAt, field.TypeTime, value)
	}
	if cu.mutation.DeletedAtCleared() {
		_spec.ClearField(category.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cu.mutation.IsRoot(); ok {
		_spec.SetField(category.FieldIsRoot, field.TypeBool, value)
	}
	if cu.mutation.ApplicantInterestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedApplicantInterestsIDs(); len(nodes) > 0 && !cu.mutation.ApplicantInterestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ApplicantInterestsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ChildCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildCategoriesTable,
			Columns: []string{category.ChildCategoriesColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedChildCategoriesIDs(); len(nodes) > 0 && !cu.mutation.ChildCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildCategoriesTable,
			Columns: []string{category.ChildCategoriesColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ChildCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildCategoriesTable,
			Columns: []string{category.ChildCategoriesColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.JobOfferCategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedJobOfferCategoriesIDs(); len(nodes) > 0 && !cu.mutation.JobOfferCategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.JobOfferCategoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ParentCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.ParentCategoryTable,
			Columns: []string{category.ParentCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ParentCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.ParentCategoryTable,
			Columns: []string{category.ParentCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetName sets the "name" field.
func (cuo *CategoryUpdateOne) SetName(s string) *CategoryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetSlug sets the "slug" field.
func (cuo *CategoryUpdateOne) SetSlug(s string) *CategoryUpdateOne {
	cuo.mutation.SetSlug(s)
	return cuo
}

// SetParentID sets the "parent_id" field.
func (cuo *CategoryUpdateOne) SetParentID(u uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.SetParentID(u)
	return cuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableParentID(u *uuid.UUID) *CategoryUpdateOne {
	if u != nil {
		cuo.SetParentID(*u)
	}
	return cuo
}

// ClearParentID clears the value of the "parent_id" field.
func (cuo *CategoryUpdateOne) ClearParentID() *CategoryUpdateOne {
	cuo.mutation.ClearParentID()
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CategoryUpdateOne) SetCreatedAt(t time.Time) *CategoryUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableCreatedAt(t *time.Time) *CategoryUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (cuo *CategoryUpdateOne) ClearCreatedAt() *CategoryUpdateOne {
	cuo.mutation.ClearCreatedAt()
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CategoryUpdateOne) SetUpdatedAt(t time.Time) *CategoryUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *CategoryUpdateOne {
	if t != nil {
		cuo.SetUpdatedAt(*t)
	}
	return cuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cuo *CategoryUpdateOne) ClearUpdatedAt() *CategoryUpdateOne {
	cuo.mutation.ClearUpdatedAt()
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CategoryUpdateOne) SetDeletedAt(t time.Time) *CategoryUpdateOne {
	cuo.mutation.SetDeletedAt(t)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableDeletedAt(t *time.Time) *CategoryUpdateOne {
	if t != nil {
		cuo.SetDeletedAt(*t)
	}
	return cuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cuo *CategoryUpdateOne) ClearDeletedAt() *CategoryUpdateOne {
	cuo.mutation.ClearDeletedAt()
	return cuo
}

// SetIsRoot sets the "is_root" field.
func (cuo *CategoryUpdateOne) SetIsRoot(b bool) *CategoryUpdateOne {
	cuo.mutation.SetIsRoot(b)
	return cuo
}

// AddApplicantInterestIDs adds the "applicant_interests" edge to the ApplicantInterest entity by IDs.
func (cuo *CategoryUpdateOne) AddApplicantInterestIDs(ids ...uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.AddApplicantInterestIDs(ids...)
	return cuo
}

// AddApplicantInterests adds the "applicant_interests" edges to the ApplicantInterest entity.
func (cuo *CategoryUpdateOne) AddApplicantInterests(a ...*ApplicantInterest) *CategoryUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddApplicantInterestIDs(ids...)
}

// AddChildCategoryIDs adds the "child_categories" edge to the Category entity by IDs.
func (cuo *CategoryUpdateOne) AddChildCategoryIDs(ids ...uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.AddChildCategoryIDs(ids...)
	return cuo
}

// AddChildCategories adds the "child_categories" edges to the Category entity.
func (cuo *CategoryUpdateOne) AddChildCategories(c ...*Category) *CategoryUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddChildCategoryIDs(ids...)
}

// AddJobOfferCategoryIDs adds the "job_offer_categories" edge to the JobOfferCategory entity by IDs.
func (cuo *CategoryUpdateOne) AddJobOfferCategoryIDs(ids ...uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.AddJobOfferCategoryIDs(ids...)
	return cuo
}

// AddJobOfferCategories adds the "job_offer_categories" edges to the JobOfferCategory entity.
func (cuo *CategoryUpdateOne) AddJobOfferCategories(j ...*JobOfferCategory) *CategoryUpdateOne {
	ids := make([]uuid.UUID, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cuo.AddJobOfferCategoryIDs(ids...)
}

// SetParentCategoryID sets the "parent_category" edge to the Category entity by ID.
func (cuo *CategoryUpdateOne) SetParentCategoryID(id uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.SetParentCategoryID(id)
	return cuo
}

// SetNillableParentCategoryID sets the "parent_category" edge to the Category entity by ID if the given value is not nil.
func (cuo *CategoryUpdateOne) SetNillableParentCategoryID(id *uuid.UUID) *CategoryUpdateOne {
	if id != nil {
		cuo = cuo.SetParentCategoryID(*id)
	}
	return cuo
}

// SetParentCategory sets the "parent_category" edge to the Category entity.
func (cuo *CategoryUpdateOne) SetParentCategory(c *Category) *CategoryUpdateOne {
	return cuo.SetParentCategoryID(c.ID)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearApplicantInterests clears all "applicant_interests" edges to the ApplicantInterest entity.
func (cuo *CategoryUpdateOne) ClearApplicantInterests() *CategoryUpdateOne {
	cuo.mutation.ClearApplicantInterests()
	return cuo
}

// RemoveApplicantInterestIDs removes the "applicant_interests" edge to ApplicantInterest entities by IDs.
func (cuo *CategoryUpdateOne) RemoveApplicantInterestIDs(ids ...uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.RemoveApplicantInterestIDs(ids...)
	return cuo
}

// RemoveApplicantInterests removes "applicant_interests" edges to ApplicantInterest entities.
func (cuo *CategoryUpdateOne) RemoveApplicantInterests(a ...*ApplicantInterest) *CategoryUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveApplicantInterestIDs(ids...)
}

// ClearChildCategories clears all "child_categories" edges to the Category entity.
func (cuo *CategoryUpdateOne) ClearChildCategories() *CategoryUpdateOne {
	cuo.mutation.ClearChildCategories()
	return cuo
}

// RemoveChildCategoryIDs removes the "child_categories" edge to Category entities by IDs.
func (cuo *CategoryUpdateOne) RemoveChildCategoryIDs(ids ...uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.RemoveChildCategoryIDs(ids...)
	return cuo
}

// RemoveChildCategories removes "child_categories" edges to Category entities.
func (cuo *CategoryUpdateOne) RemoveChildCategories(c ...*Category) *CategoryUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveChildCategoryIDs(ids...)
}

// ClearJobOfferCategories clears all "job_offer_categories" edges to the JobOfferCategory entity.
func (cuo *CategoryUpdateOne) ClearJobOfferCategories() *CategoryUpdateOne {
	cuo.mutation.ClearJobOfferCategories()
	return cuo
}

// RemoveJobOfferCategoryIDs removes the "job_offer_categories" edge to JobOfferCategory entities by IDs.
func (cuo *CategoryUpdateOne) RemoveJobOfferCategoryIDs(ids ...uuid.UUID) *CategoryUpdateOne {
	cuo.mutation.RemoveJobOfferCategoryIDs(ids...)
	return cuo
}

// RemoveJobOfferCategories removes "job_offer_categories" edges to JobOfferCategory entities.
func (cuo *CategoryUpdateOne) RemoveJobOfferCategories(j ...*JobOfferCategory) *CategoryUpdateOne {
	ids := make([]uuid.UUID, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return cuo.RemoveJobOfferCategoryIDs(ids...)
}

// ClearParentCategory clears the "parent_category" edge to the Category entity.
func (cuo *CategoryUpdateOne) ClearParentCategory() *CategoryUpdateOne {
	cuo.mutation.ClearParentCategory()
	return cuo
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cuo *CategoryUpdateOne) Where(ps ...predicate.Category) *CategoryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Slug(); ok {
		_spec.SetField(category.FieldSlug, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(category.FieldCreatedAt, field.TypeTime, value)
	}
	if cuo.mutation.CreatedAtCleared() {
		_spec.ClearField(category.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(category.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(category.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.SetField(category.FieldDeletedAt, field.TypeTime, value)
	}
	if cuo.mutation.DeletedAtCleared() {
		_spec.ClearField(category.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.IsRoot(); ok {
		_spec.SetField(category.FieldIsRoot, field.TypeBool, value)
	}
	if cuo.mutation.ApplicantInterestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedApplicantInterestsIDs(); len(nodes) > 0 && !cuo.mutation.ApplicantInterestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ApplicantInterestsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ChildCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildCategoriesTable,
			Columns: []string{category.ChildCategoriesColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedChildCategoriesIDs(); len(nodes) > 0 && !cuo.mutation.ChildCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildCategoriesTable,
			Columns: []string{category.ChildCategoriesColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ChildCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.ChildCategoriesTable,
			Columns: []string{category.ChildCategoriesColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.JobOfferCategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedJobOfferCategoriesIDs(); len(nodes) > 0 && !cuo.mutation.JobOfferCategoriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.JobOfferCategoriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ParentCategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.ParentCategoryTable,
			Columns: []string{category.ParentCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ParentCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   category.ParentCategoryTable,
			Columns: []string{category.ParentCategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
