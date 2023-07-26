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
	"github.com/marianozunino/cc-backend-go/ent/applicantprofileskill"
	"github.com/marianozunino/cc-backend-go/ent/predicate"
	"github.com/marianozunino/cc-backend-go/ent/skill"
)

// ApplicantProfileSkillUpdate is the builder for updating ApplicantProfileSkill entities.
type ApplicantProfileSkillUpdate struct {
	config
	hooks    []Hook
	mutation *ApplicantProfileSkillMutation
}

// Where appends a list predicates to the ApplicantProfileSkillUpdate builder.
func (apsu *ApplicantProfileSkillUpdate) Where(ps ...predicate.ApplicantProfileSkill) *ApplicantProfileSkillUpdate {
	apsu.mutation.Where(ps...)
	return apsu
}

// SetApplicantProfileID sets the "applicant_profile_id" field.
func (apsu *ApplicantProfileSkillUpdate) SetApplicantProfileID(u uuid.UUID) *ApplicantProfileSkillUpdate {
	apsu.mutation.SetApplicantProfileID(u)
	return apsu
}

// SetLevel sets the "level" field.
func (apsu *ApplicantProfileSkillUpdate) SetLevel(a applicantprofileskill.Level) *ApplicantProfileSkillUpdate {
	apsu.mutation.SetLevel(a)
	return apsu
}

// SetSkillID sets the "skill_id" field.
func (apsu *ApplicantProfileSkillUpdate) SetSkillID(u uuid.UUID) *ApplicantProfileSkillUpdate {
	apsu.mutation.SetSkillID(u)
	return apsu
}

// SetNillableSkillID sets the "skill_id" field if the given value is not nil.
func (apsu *ApplicantProfileSkillUpdate) SetNillableSkillID(u *uuid.UUID) *ApplicantProfileSkillUpdate {
	if u != nil {
		apsu.SetSkillID(*u)
	}
	return apsu
}

// ClearSkillID clears the value of the "skill_id" field.
func (apsu *ApplicantProfileSkillUpdate) ClearSkillID() *ApplicantProfileSkillUpdate {
	apsu.mutation.ClearSkillID()
	return apsu
}

// SetCreatedAt sets the "created_at" field.
func (apsu *ApplicantProfileSkillUpdate) SetCreatedAt(t time.Time) *ApplicantProfileSkillUpdate {
	apsu.mutation.SetCreatedAt(t)
	return apsu
}

// SetUpdatedAt sets the "updated_at" field.
func (apsu *ApplicantProfileSkillUpdate) SetUpdatedAt(t time.Time) *ApplicantProfileSkillUpdate {
	apsu.mutation.SetUpdatedAt(t)
	return apsu
}

// SetDeletedAt sets the "deleted_at" field.
func (apsu *ApplicantProfileSkillUpdate) SetDeletedAt(t time.Time) *ApplicantProfileSkillUpdate {
	apsu.mutation.SetDeletedAt(t)
	return apsu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (apsu *ApplicantProfileSkillUpdate) SetNillableDeletedAt(t *time.Time) *ApplicantProfileSkillUpdate {
	if t != nil {
		apsu.SetDeletedAt(*t)
	}
	return apsu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (apsu *ApplicantProfileSkillUpdate) ClearDeletedAt() *ApplicantProfileSkillUpdate {
	apsu.mutation.ClearDeletedAt()
	return apsu
}

// SetApplicantProfile sets the "applicant_profile" edge to the ApplicantProfile entity.
func (apsu *ApplicantProfileSkillUpdate) SetApplicantProfile(a *ApplicantProfile) *ApplicantProfileSkillUpdate {
	return apsu.SetApplicantProfileID(a.ID)
}

// SetSkill sets the "skill" edge to the Skill entity.
func (apsu *ApplicantProfileSkillUpdate) SetSkill(s *Skill) *ApplicantProfileSkillUpdate {
	return apsu.SetSkillID(s.ID)
}

// Mutation returns the ApplicantProfileSkillMutation object of the builder.
func (apsu *ApplicantProfileSkillUpdate) Mutation() *ApplicantProfileSkillMutation {
	return apsu.mutation
}

// ClearApplicantProfile clears the "applicant_profile" edge to the ApplicantProfile entity.
func (apsu *ApplicantProfileSkillUpdate) ClearApplicantProfile() *ApplicantProfileSkillUpdate {
	apsu.mutation.ClearApplicantProfile()
	return apsu
}

// ClearSkill clears the "skill" edge to the Skill entity.
func (apsu *ApplicantProfileSkillUpdate) ClearSkill() *ApplicantProfileSkillUpdate {
	apsu.mutation.ClearSkill()
	return apsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (apsu *ApplicantProfileSkillUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, apsu.sqlSave, apsu.mutation, apsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apsu *ApplicantProfileSkillUpdate) SaveX(ctx context.Context) int {
	affected, err := apsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (apsu *ApplicantProfileSkillUpdate) Exec(ctx context.Context) error {
	_, err := apsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apsu *ApplicantProfileSkillUpdate) ExecX(ctx context.Context) {
	if err := apsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apsu *ApplicantProfileSkillUpdate) check() error {
	if v, ok := apsu.mutation.Level(); ok {
		if err := applicantprofileskill.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "ApplicantProfileSkill.level": %w`, err)}
		}
	}
	if _, ok := apsu.mutation.ApplicantProfileID(); apsu.mutation.ApplicantProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ApplicantProfileSkill.applicant_profile"`)
	}
	return nil
}

func (apsu *ApplicantProfileSkillUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := apsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(applicantprofileskill.Table, applicantprofileskill.Columns, sqlgraph.NewFieldSpec(applicantprofileskill.FieldID, field.TypeUUID))
	if ps := apsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apsu.mutation.Level(); ok {
		_spec.SetField(applicantprofileskill.FieldLevel, field.TypeEnum, value)
	}
	if value, ok := apsu.mutation.CreatedAt(); ok {
		_spec.SetField(applicantprofileskill.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := apsu.mutation.UpdatedAt(); ok {
		_spec.SetField(applicantprofileskill.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := apsu.mutation.DeletedAt(); ok {
		_spec.SetField(applicantprofileskill.FieldDeletedAt, field.TypeTime, value)
	}
	if apsu.mutation.DeletedAtCleared() {
		_spec.ClearField(applicantprofileskill.FieldDeletedAt, field.TypeTime)
	}
	if apsu.mutation.ApplicantProfileCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apsu.mutation.ApplicantProfileIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apsu.mutation.SkillCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apsu.mutation.SkillIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, apsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicantprofileskill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	apsu.mutation.done = true
	return n, nil
}

// ApplicantProfileSkillUpdateOne is the builder for updating a single ApplicantProfileSkill entity.
type ApplicantProfileSkillUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApplicantProfileSkillMutation
}

// SetApplicantProfileID sets the "applicant_profile_id" field.
func (apsuo *ApplicantProfileSkillUpdateOne) SetApplicantProfileID(u uuid.UUID) *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.SetApplicantProfileID(u)
	return apsuo
}

// SetLevel sets the "level" field.
func (apsuo *ApplicantProfileSkillUpdateOne) SetLevel(a applicantprofileskill.Level) *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.SetLevel(a)
	return apsuo
}

// SetSkillID sets the "skill_id" field.
func (apsuo *ApplicantProfileSkillUpdateOne) SetSkillID(u uuid.UUID) *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.SetSkillID(u)
	return apsuo
}

// SetNillableSkillID sets the "skill_id" field if the given value is not nil.
func (apsuo *ApplicantProfileSkillUpdateOne) SetNillableSkillID(u *uuid.UUID) *ApplicantProfileSkillUpdateOne {
	if u != nil {
		apsuo.SetSkillID(*u)
	}
	return apsuo
}

// ClearSkillID clears the value of the "skill_id" field.
func (apsuo *ApplicantProfileSkillUpdateOne) ClearSkillID() *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.ClearSkillID()
	return apsuo
}

// SetCreatedAt sets the "created_at" field.
func (apsuo *ApplicantProfileSkillUpdateOne) SetCreatedAt(t time.Time) *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.SetCreatedAt(t)
	return apsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (apsuo *ApplicantProfileSkillUpdateOne) SetUpdatedAt(t time.Time) *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.SetUpdatedAt(t)
	return apsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (apsuo *ApplicantProfileSkillUpdateOne) SetDeletedAt(t time.Time) *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.SetDeletedAt(t)
	return apsuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (apsuo *ApplicantProfileSkillUpdateOne) SetNillableDeletedAt(t *time.Time) *ApplicantProfileSkillUpdateOne {
	if t != nil {
		apsuo.SetDeletedAt(*t)
	}
	return apsuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (apsuo *ApplicantProfileSkillUpdateOne) ClearDeletedAt() *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.ClearDeletedAt()
	return apsuo
}

// SetApplicantProfile sets the "applicant_profile" edge to the ApplicantProfile entity.
func (apsuo *ApplicantProfileSkillUpdateOne) SetApplicantProfile(a *ApplicantProfile) *ApplicantProfileSkillUpdateOne {
	return apsuo.SetApplicantProfileID(a.ID)
}

// SetSkill sets the "skill" edge to the Skill entity.
func (apsuo *ApplicantProfileSkillUpdateOne) SetSkill(s *Skill) *ApplicantProfileSkillUpdateOne {
	return apsuo.SetSkillID(s.ID)
}

// Mutation returns the ApplicantProfileSkillMutation object of the builder.
func (apsuo *ApplicantProfileSkillUpdateOne) Mutation() *ApplicantProfileSkillMutation {
	return apsuo.mutation
}

// ClearApplicantProfile clears the "applicant_profile" edge to the ApplicantProfile entity.
func (apsuo *ApplicantProfileSkillUpdateOne) ClearApplicantProfile() *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.ClearApplicantProfile()
	return apsuo
}

// ClearSkill clears the "skill" edge to the Skill entity.
func (apsuo *ApplicantProfileSkillUpdateOne) ClearSkill() *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.ClearSkill()
	return apsuo
}

// Where appends a list predicates to the ApplicantProfileSkillUpdate builder.
func (apsuo *ApplicantProfileSkillUpdateOne) Where(ps ...predicate.ApplicantProfileSkill) *ApplicantProfileSkillUpdateOne {
	apsuo.mutation.Where(ps...)
	return apsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (apsuo *ApplicantProfileSkillUpdateOne) Select(field string, fields ...string) *ApplicantProfileSkillUpdateOne {
	apsuo.fields = append([]string{field}, fields...)
	return apsuo
}

// Save executes the query and returns the updated ApplicantProfileSkill entity.
func (apsuo *ApplicantProfileSkillUpdateOne) Save(ctx context.Context) (*ApplicantProfileSkill, error) {
	return withHooks(ctx, apsuo.sqlSave, apsuo.mutation, apsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apsuo *ApplicantProfileSkillUpdateOne) SaveX(ctx context.Context) *ApplicantProfileSkill {
	node, err := apsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (apsuo *ApplicantProfileSkillUpdateOne) Exec(ctx context.Context) error {
	_, err := apsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apsuo *ApplicantProfileSkillUpdateOne) ExecX(ctx context.Context) {
	if err := apsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apsuo *ApplicantProfileSkillUpdateOne) check() error {
	if v, ok := apsuo.mutation.Level(); ok {
		if err := applicantprofileskill.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "ApplicantProfileSkill.level": %w`, err)}
		}
	}
	if _, ok := apsuo.mutation.ApplicantProfileID(); apsuo.mutation.ApplicantProfileCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ApplicantProfileSkill.applicant_profile"`)
	}
	return nil
}

func (apsuo *ApplicantProfileSkillUpdateOne) sqlSave(ctx context.Context) (_node *ApplicantProfileSkill, err error) {
	if err := apsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(applicantprofileskill.Table, applicantprofileskill.Columns, sqlgraph.NewFieldSpec(applicantprofileskill.FieldID, field.TypeUUID))
	id, ok := apsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ApplicantProfileSkill.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := apsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicantprofileskill.FieldID)
		for _, f := range fields {
			if !applicantprofileskill.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != applicantprofileskill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := apsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apsuo.mutation.Level(); ok {
		_spec.SetField(applicantprofileskill.FieldLevel, field.TypeEnum, value)
	}
	if value, ok := apsuo.mutation.CreatedAt(); ok {
		_spec.SetField(applicantprofileskill.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := apsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(applicantprofileskill.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := apsuo.mutation.DeletedAt(); ok {
		_spec.SetField(applicantprofileskill.FieldDeletedAt, field.TypeTime, value)
	}
	if apsuo.mutation.DeletedAtCleared() {
		_spec.ClearField(applicantprofileskill.FieldDeletedAt, field.TypeTime)
	}
	if apsuo.mutation.ApplicantProfileCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apsuo.mutation.ApplicantProfileIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apsuo.mutation.SkillCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apsuo.mutation.SkillIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ApplicantProfileSkill{config: apsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, apsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applicantprofileskill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	apsuo.mutation.done = true
	return _node, nil
}
