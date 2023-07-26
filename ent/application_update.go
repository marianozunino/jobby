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
	"github.com/marianozunino/cc-backend-go/ent/applicantprofile"
	"github.com/marianozunino/cc-backend-go/ent/application"
	"github.com/marianozunino/cc-backend-go/ent/interview"
	"github.com/marianozunino/cc-backend-go/ent/joboffer"
	"github.com/marianozunino/cc-backend-go/ent/predicate"
)

// ApplicationUpdate is the builder for updating Application entities.
type ApplicationUpdate struct {
	config
	hooks    []Hook
	mutation *ApplicationMutation
}

// Where appends a list predicates to the ApplicationUpdate builder.
func (au *ApplicationUpdate) Where(ps ...predicate.Application) *ApplicationUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetJobOfferID sets the "job_offer_id" field.
func (au *ApplicationUpdate) SetJobOfferID(u uuid.UUID) *ApplicationUpdate {
	au.mutation.SetJobOfferID(u)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *ApplicationUpdate) SetCreatedAt(t time.Time) *ApplicationUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableCreatedAt(t *time.Time) *ApplicationUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// ClearCreatedAt clears the value of the "created_at" field.
func (au *ApplicationUpdate) ClearCreatedAt() *ApplicationUpdate {
	au.mutation.ClearCreatedAt()
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ApplicationUpdate) SetUpdatedAt(t time.Time) *ApplicationUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableUpdatedAt(t *time.Time) *ApplicationUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (au *ApplicationUpdate) ClearUpdatedAt() *ApplicationUpdate {
	au.mutation.ClearUpdatedAt()
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *ApplicationUpdate) SetDeletedAt(t time.Time) *ApplicationUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *ApplicationUpdate) SetNillableDeletedAt(t *time.Time) *ApplicationUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (au *ApplicationUpdate) ClearDeletedAt() *ApplicationUpdate {
	au.mutation.ClearDeletedAt()
	return au
}

// SetApplicantProfileID sets the "applicant_profile_id" field.
func (au *ApplicationUpdate) SetApplicantProfileID(u uuid.UUID) *ApplicationUpdate {
	au.mutation.SetApplicantProfileID(u)
	return au
}

// SetApplicantProfile sets the "applicant_profile" edge to the ApplicantProfile entity.
func (au *ApplicationUpdate) SetApplicantProfile(a *ApplicantProfile) *ApplicationUpdate {
	return au.SetApplicantProfileID(a.ID)
}

// SetJobOffer sets the "job_offer" edge to the JobOffer entity.
func (au *ApplicationUpdate) SetJobOffer(j *JobOffer) *ApplicationUpdate {
	return au.SetJobOfferID(j.ID)
}

// AddInterviewIDs adds the "interviews" edge to the Interview entity by IDs.
func (au *ApplicationUpdate) AddInterviewIDs(ids ...uuid.UUID) *ApplicationUpdate {
	au.mutation.AddInterviewIDs(ids...)
	return au
}

// AddInterviews adds the "interviews" edges to the Interview entity.
func (au *ApplicationUpdate) AddInterviews(i ...*Interview) *ApplicationUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return au.AddInterviewIDs(ids...)
}

// Mutation returns the ApplicationMutation object of the builder.
func (au *ApplicationUpdate) Mutation() *ApplicationMutation {
	return au.mutation
}

// ClearApplicantProfile clears the "applicant_profile" edge to the ApplicantProfile entity.
func (au *ApplicationUpdate) ClearApplicantProfile() *ApplicationUpdate {
	au.mutation.ClearApplicantProfile()
	return au
}

// ClearJobOffer clears the "job_offer" edge to the JobOffer entity.
func (au *ApplicationUpdate) ClearJobOffer() *ApplicationUpdate {
	au.mutation.ClearJobOffer()
	return au
}

// ClearInterviews clears all "interviews" edges to the Interview entity.
func (au *ApplicationUpdate) ClearInterviews() *ApplicationUpdate {
	au.mutation.ClearInterviews()
	return au
}

// RemoveInterviewIDs removes the "interviews" edge to Interview entities by IDs.
func (au *ApplicationUpdate) RemoveInterviewIDs(ids ...uuid.UUID) *ApplicationUpdate {
	au.mutation.RemoveInterviewIDs(ids...)
	return au
}

// RemoveInterviews removes "interviews" edges to Interview entities.
func (au *ApplicationUpdate) RemoveInterviews(i ...*Interview) *ApplicationUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return au.RemoveInterviewIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ApplicationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ApplicationUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ApplicationUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ApplicationUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ApplicationUpdate) check() error {
	if _, ok := au.mutation.ApplicantProfileID(); au.mutation.ApplicantProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Application.applicant_profile"`)
	}
	if _, ok := au.mutation.JobOfferID(); au.mutation.JobOfferCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Application.job_offer"`)
	}
	return nil
}

func (au *ApplicationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(application.Table, application.Columns, sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(application.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.CreatedAtCleared() {
		_spec.ClearField(application.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(application.FieldUpdatedAt, field.TypeTime, value)
	}
	if au.mutation.UpdatedAtCleared() {
		_spec.ClearField(application.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(application.FieldDeletedAt, field.TypeTime, value)
	}
	if au.mutation.DeletedAtCleared() {
		_spec.ClearField(application.FieldDeletedAt, field.TypeTime)
	}
	if au.mutation.ApplicantProfileCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ApplicantProfileIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.JobOfferCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.JobOfferIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.InterviewsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedInterviewsIDs(); len(nodes) > 0 && !au.mutation.InterviewsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.InterviewsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{application.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ApplicationUpdateOne is the builder for updating a single Application entity.
type ApplicationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApplicationMutation
}

// SetJobOfferID sets the "job_offer_id" field.
func (auo *ApplicationUpdateOne) SetJobOfferID(u uuid.UUID) *ApplicationUpdateOne {
	auo.mutation.SetJobOfferID(u)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *ApplicationUpdateOne) SetCreatedAt(t time.Time) *ApplicationUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableCreatedAt(t *time.Time) *ApplicationUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (auo *ApplicationUpdateOne) ClearCreatedAt() *ApplicationUpdateOne {
	auo.mutation.ClearCreatedAt()
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ApplicationUpdateOne) SetUpdatedAt(t time.Time) *ApplicationUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableUpdatedAt(t *time.Time) *ApplicationUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (auo *ApplicationUpdateOne) ClearUpdatedAt() *ApplicationUpdateOne {
	auo.mutation.ClearUpdatedAt()
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *ApplicationUpdateOne) SetDeletedAt(t time.Time) *ApplicationUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *ApplicationUpdateOne) SetNillableDeletedAt(t *time.Time) *ApplicationUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (auo *ApplicationUpdateOne) ClearDeletedAt() *ApplicationUpdateOne {
	auo.mutation.ClearDeletedAt()
	return auo
}

// SetApplicantProfileID sets the "applicant_profile_id" field.
func (auo *ApplicationUpdateOne) SetApplicantProfileID(u uuid.UUID) *ApplicationUpdateOne {
	auo.mutation.SetApplicantProfileID(u)
	return auo
}

// SetApplicantProfile sets the "applicant_profile" edge to the ApplicantProfile entity.
func (auo *ApplicationUpdateOne) SetApplicantProfile(a *ApplicantProfile) *ApplicationUpdateOne {
	return auo.SetApplicantProfileID(a.ID)
}

// SetJobOffer sets the "job_offer" edge to the JobOffer entity.
func (auo *ApplicationUpdateOne) SetJobOffer(j *JobOffer) *ApplicationUpdateOne {
	return auo.SetJobOfferID(j.ID)
}

// AddInterviewIDs adds the "interviews" edge to the Interview entity by IDs.
func (auo *ApplicationUpdateOne) AddInterviewIDs(ids ...uuid.UUID) *ApplicationUpdateOne {
	auo.mutation.AddInterviewIDs(ids...)
	return auo
}

// AddInterviews adds the "interviews" edges to the Interview entity.
func (auo *ApplicationUpdateOne) AddInterviews(i ...*Interview) *ApplicationUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return auo.AddInterviewIDs(ids...)
}

// Mutation returns the ApplicationMutation object of the builder.
func (auo *ApplicationUpdateOne) Mutation() *ApplicationMutation {
	return auo.mutation
}

// ClearApplicantProfile clears the "applicant_profile" edge to the ApplicantProfile entity.
func (auo *ApplicationUpdateOne) ClearApplicantProfile() *ApplicationUpdateOne {
	auo.mutation.ClearApplicantProfile()
	return auo
}

// ClearJobOffer clears the "job_offer" edge to the JobOffer entity.
func (auo *ApplicationUpdateOne) ClearJobOffer() *ApplicationUpdateOne {
	auo.mutation.ClearJobOffer()
	return auo
}

// ClearInterviews clears all "interviews" edges to the Interview entity.
func (auo *ApplicationUpdateOne) ClearInterviews() *ApplicationUpdateOne {
	auo.mutation.ClearInterviews()
	return auo
}

// RemoveInterviewIDs removes the "interviews" edge to Interview entities by IDs.
func (auo *ApplicationUpdateOne) RemoveInterviewIDs(ids ...uuid.UUID) *ApplicationUpdateOne {
	auo.mutation.RemoveInterviewIDs(ids...)
	return auo
}

// RemoveInterviews removes "interviews" edges to Interview entities.
func (auo *ApplicationUpdateOne) RemoveInterviews(i ...*Interview) *ApplicationUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return auo.RemoveInterviewIDs(ids...)
}

// Where appends a list predicates to the ApplicationUpdate builder.
func (auo *ApplicationUpdateOne) Where(ps ...predicate.Application) *ApplicationUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ApplicationUpdateOne) Select(field string, fields ...string) *ApplicationUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Application entity.
func (auo *ApplicationUpdateOne) Save(ctx context.Context) (*Application, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ApplicationUpdateOne) SaveX(ctx context.Context) *Application {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ApplicationUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ApplicationUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ApplicationUpdateOne) check() error {
	if _, ok := auo.mutation.ApplicantProfileID(); auo.mutation.ApplicantProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Application.applicant_profile"`)
	}
	if _, ok := auo.mutation.JobOfferID(); auo.mutation.JobOfferCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Application.job_offer"`)
	}
	return nil
}

func (auo *ApplicationUpdateOne) sqlSave(ctx context.Context) (_node *Application, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(application.Table, application.Columns, sqlgraph.NewFieldSpec(application.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Application.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, application.FieldID)
		for _, f := range fields {
			if !application.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != application.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(application.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.CreatedAtCleared() {
		_spec.ClearField(application.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(application.FieldUpdatedAt, field.TypeTime, value)
	}
	if auo.mutation.UpdatedAtCleared() {
		_spec.ClearField(application.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(application.FieldDeletedAt, field.TypeTime, value)
	}
	if auo.mutation.DeletedAtCleared() {
		_spec.ClearField(application.FieldDeletedAt, field.TypeTime)
	}
	if auo.mutation.ApplicantProfileCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ApplicantProfileIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.JobOfferCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.JobOfferIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.InterviewsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedInterviewsIDs(); len(nodes) > 0 && !auo.mutation.InterviewsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.InterviewsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Application{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{application.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
