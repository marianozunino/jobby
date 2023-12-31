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
	"github.com/marianozunino/jobby/ent/application"
	"github.com/marianozunino/jobby/ent/joboffer"
	"github.com/marianozunino/jobby/ent/joboffercategory"
	"github.com/marianozunino/jobby/ent/predicate"
	"github.com/marianozunino/jobby/ent/status"
)

// JobOfferUpdate is the builder for updating JobOffer entities.
type JobOfferUpdate struct {
	config
	hooks    []Hook
	mutation *JobOfferMutation
}

// Where appends a list predicates to the JobOfferUpdate builder.
func (jou *JobOfferUpdate) Where(ps ...predicate.JobOffer) *JobOfferUpdate {
	jou.mutation.Where(ps...)
	return jou
}

// SetTitle sets the "title" field.
func (jou *JobOfferUpdate) SetTitle(s string) *JobOfferUpdate {
	jou.mutation.SetTitle(s)
	return jou
}

// SetReference sets the "reference" field.
func (jou *JobOfferUpdate) SetReference(i int32) *JobOfferUpdate {
	jou.mutation.ResetReference()
	jou.mutation.SetReference(i)
	return jou
}

// AddReference adds i to the "reference" field.
func (jou *JobOfferUpdate) AddReference(i int32) *JobOfferUpdate {
	jou.mutation.AddReference(i)
	return jou
}

// SetStartDate sets the "start_date" field.
func (jou *JobOfferUpdate) SetStartDate(t time.Time) *JobOfferUpdate {
	jou.mutation.SetStartDate(t)
	return jou
}

// SetEndDate sets the "end_date" field.
func (jou *JobOfferUpdate) SetEndDate(t time.Time) *JobOfferUpdate {
	jou.mutation.SetEndDate(t)
	return jou
}

// SetAddress1 sets the "address1" field.
func (jou *JobOfferUpdate) SetAddress1(s string) *JobOfferUpdate {
	jou.mutation.SetAddress1(s)
	return jou
}

// SetNillableAddress1 sets the "address1" field if the given value is not nil.
func (jou *JobOfferUpdate) SetNillableAddress1(s *string) *JobOfferUpdate {
	if s != nil {
		jou.SetAddress1(*s)
	}
	return jou
}

// ClearAddress1 clears the value of the "address1" field.
func (jou *JobOfferUpdate) ClearAddress1() *JobOfferUpdate {
	jou.mutation.ClearAddress1()
	return jou
}

// SetAddress2 sets the "address2" field.
func (jou *JobOfferUpdate) SetAddress2(s string) *JobOfferUpdate {
	jou.mutation.SetAddress2(s)
	return jou
}

// SetNillableAddress2 sets the "address2" field if the given value is not nil.
func (jou *JobOfferUpdate) SetNillableAddress2(s *string) *JobOfferUpdate {
	if s != nil {
		jou.SetAddress2(*s)
	}
	return jou
}

// ClearAddress2 clears the value of the "address2" field.
func (jou *JobOfferUpdate) ClearAddress2() *JobOfferUpdate {
	jou.mutation.ClearAddress2()
	return jou
}

// SetDepartment sets the "department" field.
func (jou *JobOfferUpdate) SetDepartment(s string) *JobOfferUpdate {
	jou.mutation.SetDepartment(s)
	return jou
}

// SetDescription sets the "description" field.
func (jou *JobOfferUpdate) SetDescription(s string) *JobOfferUpdate {
	jou.mutation.SetDescription(s)
	return jou
}

// SetWorkingHours sets the "working_hours" field.
func (jou *JobOfferUpdate) SetWorkingHours(s string) *JobOfferUpdate {
	jou.mutation.SetWorkingHours(s)
	return jou
}

// SetSalary sets the "salary" field.
func (jou *JobOfferUpdate) SetSalary(s string) *JobOfferUpdate {
	jou.mutation.SetSalary(s)
	return jou
}

// SetSlug sets the "slug" field.
func (jou *JobOfferUpdate) SetSlug(s string) *JobOfferUpdate {
	jou.mutation.SetSlug(s)
	return jou
}

// SetIsFeatured sets the "is_featured" field.
func (jou *JobOfferUpdate) SetIsFeatured(b bool) *JobOfferUpdate {
	jou.mutation.SetIsFeatured(b)
	return jou
}

// SetNillableIsFeatured sets the "is_featured" field if the given value is not nil.
func (jou *JobOfferUpdate) SetNillableIsFeatured(b *bool) *JobOfferUpdate {
	if b != nil {
		jou.SetIsFeatured(*b)
	}
	return jou
}

// ClearIsFeatured clears the value of the "is_featured" field.
func (jou *JobOfferUpdate) ClearIsFeatured() *JobOfferUpdate {
	jou.mutation.ClearIsFeatured()
	return jou
}

// SetHasBeenEmailed sets the "has_been_emailed" field.
func (jou *JobOfferUpdate) SetHasBeenEmailed(b bool) *JobOfferUpdate {
	jou.mutation.SetHasBeenEmailed(b)
	return jou
}

// SetNillableHasBeenEmailed sets the "has_been_emailed" field if the given value is not nil.
func (jou *JobOfferUpdate) SetNillableHasBeenEmailed(b *bool) *JobOfferUpdate {
	if b != nil {
		jou.SetHasBeenEmailed(*b)
	}
	return jou
}

// ClearHasBeenEmailed clears the value of the "has_been_emailed" field.
func (jou *JobOfferUpdate) ClearHasBeenEmailed() *JobOfferUpdate {
	jou.mutation.ClearHasBeenEmailed()
	return jou
}

// SetStatusID sets the "status_id" field.
func (jou *JobOfferUpdate) SetStatusID(u uuid.UUID) *JobOfferUpdate {
	jou.mutation.SetStatusID(u)
	return jou
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (jou *JobOfferUpdate) SetNillableStatusID(u *uuid.UUID) *JobOfferUpdate {
	if u != nil {
		jou.SetStatusID(*u)
	}
	return jou
}

// ClearStatusID clears the value of the "status_id" field.
func (jou *JobOfferUpdate) ClearStatusID() *JobOfferUpdate {
	jou.mutation.ClearStatusID()
	return jou
}

// SetCreatedAt sets the "created_at" field.
func (jou *JobOfferUpdate) SetCreatedAt(t time.Time) *JobOfferUpdate {
	jou.mutation.SetCreatedAt(t)
	return jou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jou *JobOfferUpdate) SetNillableCreatedAt(t *time.Time) *JobOfferUpdate {
	if t != nil {
		jou.SetCreatedAt(*t)
	}
	return jou
}

// ClearCreatedAt clears the value of the "created_at" field.
func (jou *JobOfferUpdate) ClearCreatedAt() *JobOfferUpdate {
	jou.mutation.ClearCreatedAt()
	return jou
}

// SetUpdatedAt sets the "updated_at" field.
func (jou *JobOfferUpdate) SetUpdatedAt(t time.Time) *JobOfferUpdate {
	jou.mutation.SetUpdatedAt(t)
	return jou
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (jou *JobOfferUpdate) SetNillableUpdatedAt(t *time.Time) *JobOfferUpdate {
	if t != nil {
		jou.SetUpdatedAt(*t)
	}
	return jou
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (jou *JobOfferUpdate) ClearUpdatedAt() *JobOfferUpdate {
	jou.mutation.ClearUpdatedAt()
	return jou
}

// SetDeletedAt sets the "deleted_at" field.
func (jou *JobOfferUpdate) SetDeletedAt(t time.Time) *JobOfferUpdate {
	jou.mutation.SetDeletedAt(t)
	return jou
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (jou *JobOfferUpdate) SetNillableDeletedAt(t *time.Time) *JobOfferUpdate {
	if t != nil {
		jou.SetDeletedAt(*t)
	}
	return jou
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (jou *JobOfferUpdate) ClearDeletedAt() *JobOfferUpdate {
	jou.mutation.ClearDeletedAt()
	return jou
}

// AddApplicationIDs adds the "applications" edge to the Application entity by IDs.
func (jou *JobOfferUpdate) AddApplicationIDs(ids ...uuid.UUID) *JobOfferUpdate {
	jou.mutation.AddApplicationIDs(ids...)
	return jou
}

// AddApplications adds the "applications" edges to the Application entity.
func (jou *JobOfferUpdate) AddApplications(a ...*Application) *JobOfferUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return jou.AddApplicationIDs(ids...)
}

// AddJobOfferCategoryIDs adds the "job_offer_categories" edge to the JobOfferCategory entity by IDs.
func (jou *JobOfferUpdate) AddJobOfferCategoryIDs(ids ...uuid.UUID) *JobOfferUpdate {
	jou.mutation.AddJobOfferCategoryIDs(ids...)
	return jou
}

// AddJobOfferCategories adds the "job_offer_categories" edges to the JobOfferCategory entity.
func (jou *JobOfferUpdate) AddJobOfferCategories(j ...*JobOfferCategory) *JobOfferUpdate {
	ids := make([]uuid.UUID, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jou.AddJobOfferCategoryIDs(ids...)
}

// SetStatus sets the "status" edge to the Status entity.
func (jou *JobOfferUpdate) SetStatus(s *Status) *JobOfferUpdate {
	return jou.SetStatusID(s.ID)
}

// Mutation returns the JobOfferMutation object of the builder.
func (jou *JobOfferUpdate) Mutation() *JobOfferMutation {
	return jou.mutation
}

// ClearApplications clears all "applications" edges to the Application entity.
func (jou *JobOfferUpdate) ClearApplications() *JobOfferUpdate {
	jou.mutation.ClearApplications()
	return jou
}

// RemoveApplicationIDs removes the "applications" edge to Application entities by IDs.
func (jou *JobOfferUpdate) RemoveApplicationIDs(ids ...uuid.UUID) *JobOfferUpdate {
	jou.mutation.RemoveApplicationIDs(ids...)
	return jou
}

// RemoveApplications removes "applications" edges to Application entities.
func (jou *JobOfferUpdate) RemoveApplications(a ...*Application) *JobOfferUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return jou.RemoveApplicationIDs(ids...)
}

// ClearJobOfferCategories clears all "job_offer_categories" edges to the JobOfferCategory entity.
func (jou *JobOfferUpdate) ClearJobOfferCategories() *JobOfferUpdate {
	jou.mutation.ClearJobOfferCategories()
	return jou
}

// RemoveJobOfferCategoryIDs removes the "job_offer_categories" edge to JobOfferCategory entities by IDs.
func (jou *JobOfferUpdate) RemoveJobOfferCategoryIDs(ids ...uuid.UUID) *JobOfferUpdate {
	jou.mutation.RemoveJobOfferCategoryIDs(ids...)
	return jou
}

// RemoveJobOfferCategories removes "job_offer_categories" edges to JobOfferCategory entities.
func (jou *JobOfferUpdate) RemoveJobOfferCategories(j ...*JobOfferCategory) *JobOfferUpdate {
	ids := make([]uuid.UUID, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jou.RemoveJobOfferCategoryIDs(ids...)
}

// ClearStatus clears the "status" edge to the Status entity.
func (jou *JobOfferUpdate) ClearStatus() *JobOfferUpdate {
	jou.mutation.ClearStatus()
	return jou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jou *JobOfferUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, jou.sqlSave, jou.mutation, jou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jou *JobOfferUpdate) SaveX(ctx context.Context) int {
	affected, err := jou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jou *JobOfferUpdate) Exec(ctx context.Context) error {
	_, err := jou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jou *JobOfferUpdate) ExecX(ctx context.Context) {
	if err := jou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (jou *JobOfferUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(joboffer.Table, joboffer.Columns, sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeUUID))
	if ps := jou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jou.mutation.Title(); ok {
		_spec.SetField(joboffer.FieldTitle, field.TypeString, value)
	}
	if value, ok := jou.mutation.Reference(); ok {
		_spec.SetField(joboffer.FieldReference, field.TypeInt32, value)
	}
	if value, ok := jou.mutation.AddedReference(); ok {
		_spec.AddField(joboffer.FieldReference, field.TypeInt32, value)
	}
	if value, ok := jou.mutation.StartDate(); ok {
		_spec.SetField(joboffer.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := jou.mutation.EndDate(); ok {
		_spec.SetField(joboffer.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := jou.mutation.Address1(); ok {
		_spec.SetField(joboffer.FieldAddress1, field.TypeString, value)
	}
	if jou.mutation.Address1Cleared() {
		_spec.ClearField(joboffer.FieldAddress1, field.TypeString)
	}
	if value, ok := jou.mutation.Address2(); ok {
		_spec.SetField(joboffer.FieldAddress2, field.TypeString, value)
	}
	if jou.mutation.Address2Cleared() {
		_spec.ClearField(joboffer.FieldAddress2, field.TypeString)
	}
	if value, ok := jou.mutation.Department(); ok {
		_spec.SetField(joboffer.FieldDepartment, field.TypeString, value)
	}
	if value, ok := jou.mutation.Description(); ok {
		_spec.SetField(joboffer.FieldDescription, field.TypeString, value)
	}
	if value, ok := jou.mutation.WorkingHours(); ok {
		_spec.SetField(joboffer.FieldWorkingHours, field.TypeString, value)
	}
	if value, ok := jou.mutation.Salary(); ok {
		_spec.SetField(joboffer.FieldSalary, field.TypeString, value)
	}
	if value, ok := jou.mutation.Slug(); ok {
		_spec.SetField(joboffer.FieldSlug, field.TypeString, value)
	}
	if value, ok := jou.mutation.IsFeatured(); ok {
		_spec.SetField(joboffer.FieldIsFeatured, field.TypeBool, value)
	}
	if jou.mutation.IsFeaturedCleared() {
		_spec.ClearField(joboffer.FieldIsFeatured, field.TypeBool)
	}
	if value, ok := jou.mutation.HasBeenEmailed(); ok {
		_spec.SetField(joboffer.FieldHasBeenEmailed, field.TypeBool, value)
	}
	if jou.mutation.HasBeenEmailedCleared() {
		_spec.ClearField(joboffer.FieldHasBeenEmailed, field.TypeBool)
	}
	if value, ok := jou.mutation.CreatedAt(); ok {
		_spec.SetField(joboffer.FieldCreatedAt, field.TypeTime, value)
	}
	if jou.mutation.CreatedAtCleared() {
		_spec.ClearField(joboffer.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := jou.mutation.UpdatedAt(); ok {
		_spec.SetField(joboffer.FieldUpdatedAt, field.TypeTime, value)
	}
	if jou.mutation.UpdatedAtCleared() {
		_spec.ClearField(joboffer.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := jou.mutation.DeletedAt(); ok {
		_spec.SetField(joboffer.FieldDeletedAt, field.TypeTime, value)
	}
	if jou.mutation.DeletedAtCleared() {
		_spec.ClearField(joboffer.FieldDeletedAt, field.TypeTime)
	}
	if jou.mutation.ApplicationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.ApplicationsTable,
			Columns: []string{joboffer.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jou.mutation.RemovedApplicationsIDs(); len(nodes) > 0 && !jou.mutation.ApplicationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.ApplicationsTable,
			Columns: []string{joboffer.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jou.mutation.ApplicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.ApplicationsTable,
			Columns: []string{joboffer.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if jou.mutation.JobOfferCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.JobOfferCategoriesTable,
			Columns: []string{joboffer.JobOfferCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffercategory.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jou.mutation.RemovedJobOfferCategoriesIDs(); len(nodes) > 0 && !jou.mutation.JobOfferCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.JobOfferCategoriesTable,
			Columns: []string{joboffer.JobOfferCategoriesColumn},
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
	if nodes := jou.mutation.JobOfferCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.JobOfferCategoriesTable,
			Columns: []string{joboffer.JobOfferCategoriesColumn},
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
	if jou.mutation.StatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   joboffer.StatusTable,
			Columns: []string{joboffer.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jou.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   joboffer.StatusTable,
			Columns: []string{joboffer.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, jou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{joboffer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jou.mutation.done = true
	return n, nil
}

// JobOfferUpdateOne is the builder for updating a single JobOffer entity.
type JobOfferUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobOfferMutation
}

// SetTitle sets the "title" field.
func (jouo *JobOfferUpdateOne) SetTitle(s string) *JobOfferUpdateOne {
	jouo.mutation.SetTitle(s)
	return jouo
}

// SetReference sets the "reference" field.
func (jouo *JobOfferUpdateOne) SetReference(i int32) *JobOfferUpdateOne {
	jouo.mutation.ResetReference()
	jouo.mutation.SetReference(i)
	return jouo
}

// AddReference adds i to the "reference" field.
func (jouo *JobOfferUpdateOne) AddReference(i int32) *JobOfferUpdateOne {
	jouo.mutation.AddReference(i)
	return jouo
}

// SetStartDate sets the "start_date" field.
func (jouo *JobOfferUpdateOne) SetStartDate(t time.Time) *JobOfferUpdateOne {
	jouo.mutation.SetStartDate(t)
	return jouo
}

// SetEndDate sets the "end_date" field.
func (jouo *JobOfferUpdateOne) SetEndDate(t time.Time) *JobOfferUpdateOne {
	jouo.mutation.SetEndDate(t)
	return jouo
}

// SetAddress1 sets the "address1" field.
func (jouo *JobOfferUpdateOne) SetAddress1(s string) *JobOfferUpdateOne {
	jouo.mutation.SetAddress1(s)
	return jouo
}

// SetNillableAddress1 sets the "address1" field if the given value is not nil.
func (jouo *JobOfferUpdateOne) SetNillableAddress1(s *string) *JobOfferUpdateOne {
	if s != nil {
		jouo.SetAddress1(*s)
	}
	return jouo
}

// ClearAddress1 clears the value of the "address1" field.
func (jouo *JobOfferUpdateOne) ClearAddress1() *JobOfferUpdateOne {
	jouo.mutation.ClearAddress1()
	return jouo
}

// SetAddress2 sets the "address2" field.
func (jouo *JobOfferUpdateOne) SetAddress2(s string) *JobOfferUpdateOne {
	jouo.mutation.SetAddress2(s)
	return jouo
}

// SetNillableAddress2 sets the "address2" field if the given value is not nil.
func (jouo *JobOfferUpdateOne) SetNillableAddress2(s *string) *JobOfferUpdateOne {
	if s != nil {
		jouo.SetAddress2(*s)
	}
	return jouo
}

// ClearAddress2 clears the value of the "address2" field.
func (jouo *JobOfferUpdateOne) ClearAddress2() *JobOfferUpdateOne {
	jouo.mutation.ClearAddress2()
	return jouo
}

// SetDepartment sets the "department" field.
func (jouo *JobOfferUpdateOne) SetDepartment(s string) *JobOfferUpdateOne {
	jouo.mutation.SetDepartment(s)
	return jouo
}

// SetDescription sets the "description" field.
func (jouo *JobOfferUpdateOne) SetDescription(s string) *JobOfferUpdateOne {
	jouo.mutation.SetDescription(s)
	return jouo
}

// SetWorkingHours sets the "working_hours" field.
func (jouo *JobOfferUpdateOne) SetWorkingHours(s string) *JobOfferUpdateOne {
	jouo.mutation.SetWorkingHours(s)
	return jouo
}

// SetSalary sets the "salary" field.
func (jouo *JobOfferUpdateOne) SetSalary(s string) *JobOfferUpdateOne {
	jouo.mutation.SetSalary(s)
	return jouo
}

// SetSlug sets the "slug" field.
func (jouo *JobOfferUpdateOne) SetSlug(s string) *JobOfferUpdateOne {
	jouo.mutation.SetSlug(s)
	return jouo
}

// SetIsFeatured sets the "is_featured" field.
func (jouo *JobOfferUpdateOne) SetIsFeatured(b bool) *JobOfferUpdateOne {
	jouo.mutation.SetIsFeatured(b)
	return jouo
}

// SetNillableIsFeatured sets the "is_featured" field if the given value is not nil.
func (jouo *JobOfferUpdateOne) SetNillableIsFeatured(b *bool) *JobOfferUpdateOne {
	if b != nil {
		jouo.SetIsFeatured(*b)
	}
	return jouo
}

// ClearIsFeatured clears the value of the "is_featured" field.
func (jouo *JobOfferUpdateOne) ClearIsFeatured() *JobOfferUpdateOne {
	jouo.mutation.ClearIsFeatured()
	return jouo
}

// SetHasBeenEmailed sets the "has_been_emailed" field.
func (jouo *JobOfferUpdateOne) SetHasBeenEmailed(b bool) *JobOfferUpdateOne {
	jouo.mutation.SetHasBeenEmailed(b)
	return jouo
}

// SetNillableHasBeenEmailed sets the "has_been_emailed" field if the given value is not nil.
func (jouo *JobOfferUpdateOne) SetNillableHasBeenEmailed(b *bool) *JobOfferUpdateOne {
	if b != nil {
		jouo.SetHasBeenEmailed(*b)
	}
	return jouo
}

// ClearHasBeenEmailed clears the value of the "has_been_emailed" field.
func (jouo *JobOfferUpdateOne) ClearHasBeenEmailed() *JobOfferUpdateOne {
	jouo.mutation.ClearHasBeenEmailed()
	return jouo
}

// SetStatusID sets the "status_id" field.
func (jouo *JobOfferUpdateOne) SetStatusID(u uuid.UUID) *JobOfferUpdateOne {
	jouo.mutation.SetStatusID(u)
	return jouo
}

// SetNillableStatusID sets the "status_id" field if the given value is not nil.
func (jouo *JobOfferUpdateOne) SetNillableStatusID(u *uuid.UUID) *JobOfferUpdateOne {
	if u != nil {
		jouo.SetStatusID(*u)
	}
	return jouo
}

// ClearStatusID clears the value of the "status_id" field.
func (jouo *JobOfferUpdateOne) ClearStatusID() *JobOfferUpdateOne {
	jouo.mutation.ClearStatusID()
	return jouo
}

// SetCreatedAt sets the "created_at" field.
func (jouo *JobOfferUpdateOne) SetCreatedAt(t time.Time) *JobOfferUpdateOne {
	jouo.mutation.SetCreatedAt(t)
	return jouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jouo *JobOfferUpdateOne) SetNillableCreatedAt(t *time.Time) *JobOfferUpdateOne {
	if t != nil {
		jouo.SetCreatedAt(*t)
	}
	return jouo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (jouo *JobOfferUpdateOne) ClearCreatedAt() *JobOfferUpdateOne {
	jouo.mutation.ClearCreatedAt()
	return jouo
}

// SetUpdatedAt sets the "updated_at" field.
func (jouo *JobOfferUpdateOne) SetUpdatedAt(t time.Time) *JobOfferUpdateOne {
	jouo.mutation.SetUpdatedAt(t)
	return jouo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (jouo *JobOfferUpdateOne) SetNillableUpdatedAt(t *time.Time) *JobOfferUpdateOne {
	if t != nil {
		jouo.SetUpdatedAt(*t)
	}
	return jouo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (jouo *JobOfferUpdateOne) ClearUpdatedAt() *JobOfferUpdateOne {
	jouo.mutation.ClearUpdatedAt()
	return jouo
}

// SetDeletedAt sets the "deleted_at" field.
func (jouo *JobOfferUpdateOne) SetDeletedAt(t time.Time) *JobOfferUpdateOne {
	jouo.mutation.SetDeletedAt(t)
	return jouo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (jouo *JobOfferUpdateOne) SetNillableDeletedAt(t *time.Time) *JobOfferUpdateOne {
	if t != nil {
		jouo.SetDeletedAt(*t)
	}
	return jouo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (jouo *JobOfferUpdateOne) ClearDeletedAt() *JobOfferUpdateOne {
	jouo.mutation.ClearDeletedAt()
	return jouo
}

// AddApplicationIDs adds the "applications" edge to the Application entity by IDs.
func (jouo *JobOfferUpdateOne) AddApplicationIDs(ids ...uuid.UUID) *JobOfferUpdateOne {
	jouo.mutation.AddApplicationIDs(ids...)
	return jouo
}

// AddApplications adds the "applications" edges to the Application entity.
func (jouo *JobOfferUpdateOne) AddApplications(a ...*Application) *JobOfferUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return jouo.AddApplicationIDs(ids...)
}

// AddJobOfferCategoryIDs adds the "job_offer_categories" edge to the JobOfferCategory entity by IDs.
func (jouo *JobOfferUpdateOne) AddJobOfferCategoryIDs(ids ...uuid.UUID) *JobOfferUpdateOne {
	jouo.mutation.AddJobOfferCategoryIDs(ids...)
	return jouo
}

// AddJobOfferCategories adds the "job_offer_categories" edges to the JobOfferCategory entity.
func (jouo *JobOfferUpdateOne) AddJobOfferCategories(j ...*JobOfferCategory) *JobOfferUpdateOne {
	ids := make([]uuid.UUID, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jouo.AddJobOfferCategoryIDs(ids...)
}

// SetStatus sets the "status" edge to the Status entity.
func (jouo *JobOfferUpdateOne) SetStatus(s *Status) *JobOfferUpdateOne {
	return jouo.SetStatusID(s.ID)
}

// Mutation returns the JobOfferMutation object of the builder.
func (jouo *JobOfferUpdateOne) Mutation() *JobOfferMutation {
	return jouo.mutation
}

// ClearApplications clears all "applications" edges to the Application entity.
func (jouo *JobOfferUpdateOne) ClearApplications() *JobOfferUpdateOne {
	jouo.mutation.ClearApplications()
	return jouo
}

// RemoveApplicationIDs removes the "applications" edge to Application entities by IDs.
func (jouo *JobOfferUpdateOne) RemoveApplicationIDs(ids ...uuid.UUID) *JobOfferUpdateOne {
	jouo.mutation.RemoveApplicationIDs(ids...)
	return jouo
}

// RemoveApplications removes "applications" edges to Application entities.
func (jouo *JobOfferUpdateOne) RemoveApplications(a ...*Application) *JobOfferUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return jouo.RemoveApplicationIDs(ids...)
}

// ClearJobOfferCategories clears all "job_offer_categories" edges to the JobOfferCategory entity.
func (jouo *JobOfferUpdateOne) ClearJobOfferCategories() *JobOfferUpdateOne {
	jouo.mutation.ClearJobOfferCategories()
	return jouo
}

// RemoveJobOfferCategoryIDs removes the "job_offer_categories" edge to JobOfferCategory entities by IDs.
func (jouo *JobOfferUpdateOne) RemoveJobOfferCategoryIDs(ids ...uuid.UUID) *JobOfferUpdateOne {
	jouo.mutation.RemoveJobOfferCategoryIDs(ids...)
	return jouo
}

// RemoveJobOfferCategories removes "job_offer_categories" edges to JobOfferCategory entities.
func (jouo *JobOfferUpdateOne) RemoveJobOfferCategories(j ...*JobOfferCategory) *JobOfferUpdateOne {
	ids := make([]uuid.UUID, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jouo.RemoveJobOfferCategoryIDs(ids...)
}

// ClearStatus clears the "status" edge to the Status entity.
func (jouo *JobOfferUpdateOne) ClearStatus() *JobOfferUpdateOne {
	jouo.mutation.ClearStatus()
	return jouo
}

// Where appends a list predicates to the JobOfferUpdate builder.
func (jouo *JobOfferUpdateOne) Where(ps ...predicate.JobOffer) *JobOfferUpdateOne {
	jouo.mutation.Where(ps...)
	return jouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jouo *JobOfferUpdateOne) Select(field string, fields ...string) *JobOfferUpdateOne {
	jouo.fields = append([]string{field}, fields...)
	return jouo
}

// Save executes the query and returns the updated JobOffer entity.
func (jouo *JobOfferUpdateOne) Save(ctx context.Context) (*JobOffer, error) {
	return withHooks(ctx, jouo.sqlSave, jouo.mutation, jouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jouo *JobOfferUpdateOne) SaveX(ctx context.Context) *JobOffer {
	node, err := jouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jouo *JobOfferUpdateOne) Exec(ctx context.Context) error {
	_, err := jouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jouo *JobOfferUpdateOne) ExecX(ctx context.Context) {
	if err := jouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (jouo *JobOfferUpdateOne) sqlSave(ctx context.Context) (_node *JobOffer, err error) {
	_spec := sqlgraph.NewUpdateSpec(joboffer.Table, joboffer.Columns, sqlgraph.NewFieldSpec(joboffer.FieldID, field.TypeUUID))
	id, ok := jouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobOffer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, joboffer.FieldID)
		for _, f := range fields {
			if !joboffer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != joboffer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jouo.mutation.Title(); ok {
		_spec.SetField(joboffer.FieldTitle, field.TypeString, value)
	}
	if value, ok := jouo.mutation.Reference(); ok {
		_spec.SetField(joboffer.FieldReference, field.TypeInt32, value)
	}
	if value, ok := jouo.mutation.AddedReference(); ok {
		_spec.AddField(joboffer.FieldReference, field.TypeInt32, value)
	}
	if value, ok := jouo.mutation.StartDate(); ok {
		_spec.SetField(joboffer.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := jouo.mutation.EndDate(); ok {
		_spec.SetField(joboffer.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := jouo.mutation.Address1(); ok {
		_spec.SetField(joboffer.FieldAddress1, field.TypeString, value)
	}
	if jouo.mutation.Address1Cleared() {
		_spec.ClearField(joboffer.FieldAddress1, field.TypeString)
	}
	if value, ok := jouo.mutation.Address2(); ok {
		_spec.SetField(joboffer.FieldAddress2, field.TypeString, value)
	}
	if jouo.mutation.Address2Cleared() {
		_spec.ClearField(joboffer.FieldAddress2, field.TypeString)
	}
	if value, ok := jouo.mutation.Department(); ok {
		_spec.SetField(joboffer.FieldDepartment, field.TypeString, value)
	}
	if value, ok := jouo.mutation.Description(); ok {
		_spec.SetField(joboffer.FieldDescription, field.TypeString, value)
	}
	if value, ok := jouo.mutation.WorkingHours(); ok {
		_spec.SetField(joboffer.FieldWorkingHours, field.TypeString, value)
	}
	if value, ok := jouo.mutation.Salary(); ok {
		_spec.SetField(joboffer.FieldSalary, field.TypeString, value)
	}
	if value, ok := jouo.mutation.Slug(); ok {
		_spec.SetField(joboffer.FieldSlug, field.TypeString, value)
	}
	if value, ok := jouo.mutation.IsFeatured(); ok {
		_spec.SetField(joboffer.FieldIsFeatured, field.TypeBool, value)
	}
	if jouo.mutation.IsFeaturedCleared() {
		_spec.ClearField(joboffer.FieldIsFeatured, field.TypeBool)
	}
	if value, ok := jouo.mutation.HasBeenEmailed(); ok {
		_spec.SetField(joboffer.FieldHasBeenEmailed, field.TypeBool, value)
	}
	if jouo.mutation.HasBeenEmailedCleared() {
		_spec.ClearField(joboffer.FieldHasBeenEmailed, field.TypeBool)
	}
	if value, ok := jouo.mutation.CreatedAt(); ok {
		_spec.SetField(joboffer.FieldCreatedAt, field.TypeTime, value)
	}
	if jouo.mutation.CreatedAtCleared() {
		_spec.ClearField(joboffer.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := jouo.mutation.UpdatedAt(); ok {
		_spec.SetField(joboffer.FieldUpdatedAt, field.TypeTime, value)
	}
	if jouo.mutation.UpdatedAtCleared() {
		_spec.ClearField(joboffer.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := jouo.mutation.DeletedAt(); ok {
		_spec.SetField(joboffer.FieldDeletedAt, field.TypeTime, value)
	}
	if jouo.mutation.DeletedAtCleared() {
		_spec.ClearField(joboffer.FieldDeletedAt, field.TypeTime)
	}
	if jouo.mutation.ApplicationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.ApplicationsTable,
			Columns: []string{joboffer.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jouo.mutation.RemovedApplicationsIDs(); len(nodes) > 0 && !jouo.mutation.ApplicationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.ApplicationsTable,
			Columns: []string{joboffer.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jouo.mutation.ApplicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.ApplicationsTable,
			Columns: []string{joboffer.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if jouo.mutation.JobOfferCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.JobOfferCategoriesTable,
			Columns: []string{joboffer.JobOfferCategoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joboffercategory.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jouo.mutation.RemovedJobOfferCategoriesIDs(); len(nodes) > 0 && !jouo.mutation.JobOfferCategoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.JobOfferCategoriesTable,
			Columns: []string{joboffer.JobOfferCategoriesColumn},
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
	if nodes := jouo.mutation.JobOfferCategoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joboffer.JobOfferCategoriesTable,
			Columns: []string{joboffer.JobOfferCategoriesColumn},
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
	if jouo.mutation.StatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   joboffer.StatusTable,
			Columns: []string{joboffer.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jouo.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   joboffer.StatusTable,
			Columns: []string{joboffer.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(status.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &JobOffer{config: jouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{joboffer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jouo.mutation.done = true
	return _node, nil
}
