// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/ent/applicantinterest"
	"github.com/marianozunino/cc-backend-go/ent/applicantprofile"
	"github.com/marianozunino/cc-backend-go/ent/applicantprofileskill"
	"github.com/marianozunino/cc-backend-go/ent/application"
	"github.com/marianozunino/cc-backend-go/ent/education"
	"github.com/marianozunino/cc-backend-go/ent/language"
	"github.com/marianozunino/cc-backend-go/ent/predicate"
	"github.com/marianozunino/cc-backend-go/ent/user"
	"github.com/marianozunino/cc-backend-go/ent/workexperience"
)

// ApplicantProfileQuery is the builder for querying ApplicantProfile entities.
type ApplicantProfileQuery struct {
	config
	ctx                        *QueryContext
	order                      []applicantprofile.OrderOption
	inters                     []Interceptor
	predicates                 []predicate.ApplicantProfile
	withApplicantInterests     *ApplicantInterestQuery
	withApplicantProfileSkills *ApplicantProfileSkillQuery
	withUser                   *UserQuery
	withApplications           *ApplicationQuery
	withEducations             *EducationQuery
	withLanguages              *LanguageQuery
	withWorkExperiences        *WorkExperienceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ApplicantProfileQuery builder.
func (apq *ApplicantProfileQuery) Where(ps ...predicate.ApplicantProfile) *ApplicantProfileQuery {
	apq.predicates = append(apq.predicates, ps...)
	return apq
}

// Limit the number of records to be returned by this query.
func (apq *ApplicantProfileQuery) Limit(limit int) *ApplicantProfileQuery {
	apq.ctx.Limit = &limit
	return apq
}

// Offset to start from.
func (apq *ApplicantProfileQuery) Offset(offset int) *ApplicantProfileQuery {
	apq.ctx.Offset = &offset
	return apq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (apq *ApplicantProfileQuery) Unique(unique bool) *ApplicantProfileQuery {
	apq.ctx.Unique = &unique
	return apq
}

// Order specifies how the records should be ordered.
func (apq *ApplicantProfileQuery) Order(o ...applicantprofile.OrderOption) *ApplicantProfileQuery {
	apq.order = append(apq.order, o...)
	return apq
}

// QueryApplicantInterests chains the current query on the "applicant_interests" edge.
func (apq *ApplicantProfileQuery) QueryApplicantInterests() *ApplicantInterestQuery {
	query := (&ApplicantInterestClient{config: apq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(applicantprofile.Table, applicantprofile.FieldID, selector),
			sqlgraph.To(applicantinterest.Table, applicantinterest.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, applicantprofile.ApplicantInterestsTable, applicantprofile.ApplicantInterestsColumn),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryApplicantProfileSkills chains the current query on the "applicant_profile_skills" edge.
func (apq *ApplicantProfileQuery) QueryApplicantProfileSkills() *ApplicantProfileSkillQuery {
	query := (&ApplicantProfileSkillClient{config: apq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(applicantprofile.Table, applicantprofile.FieldID, selector),
			sqlgraph.To(applicantprofileskill.Table, applicantprofileskill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, applicantprofile.ApplicantProfileSkillsTable, applicantprofile.ApplicantProfileSkillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (apq *ApplicantProfileQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: apq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(applicantprofile.Table, applicantprofile.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, applicantprofile.UserTable, applicantprofile.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryApplications chains the current query on the "applications" edge.
func (apq *ApplicantProfileQuery) QueryApplications() *ApplicationQuery {
	query := (&ApplicationClient{config: apq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(applicantprofile.Table, applicantprofile.FieldID, selector),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, applicantprofile.ApplicationsTable, applicantprofile.ApplicationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEducations chains the current query on the "educations" edge.
func (apq *ApplicantProfileQuery) QueryEducations() *EducationQuery {
	query := (&EducationClient{config: apq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(applicantprofile.Table, applicantprofile.FieldID, selector),
			sqlgraph.To(education.Table, education.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, applicantprofile.EducationsTable, applicantprofile.EducationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLanguages chains the current query on the "languages" edge.
func (apq *ApplicantProfileQuery) QueryLanguages() *LanguageQuery {
	query := (&LanguageClient{config: apq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(applicantprofile.Table, applicantprofile.FieldID, selector),
			sqlgraph.To(language.Table, language.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, applicantprofile.LanguagesTable, applicantprofile.LanguagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWorkExperiences chains the current query on the "work_experiences" edge.
func (apq *ApplicantProfileQuery) QueryWorkExperiences() *WorkExperienceQuery {
	query := (&WorkExperienceClient{config: apq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(applicantprofile.Table, applicantprofile.FieldID, selector),
			sqlgraph.To(workexperience.Table, workexperience.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, applicantprofile.WorkExperiencesTable, applicantprofile.WorkExperiencesColumn),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ApplicantProfile entity from the query.
// Returns a *NotFoundError when no ApplicantProfile was found.
func (apq *ApplicantProfileQuery) First(ctx context.Context) (*ApplicantProfile, error) {
	nodes, err := apq.Limit(1).All(setContextOp(ctx, apq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{applicantprofile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (apq *ApplicantProfileQuery) FirstX(ctx context.Context) *ApplicantProfile {
	node, err := apq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ApplicantProfile ID from the query.
// Returns a *NotFoundError when no ApplicantProfile ID was found.
func (apq *ApplicantProfileQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = apq.Limit(1).IDs(setContextOp(ctx, apq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{applicantprofile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (apq *ApplicantProfileQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := apq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ApplicantProfile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ApplicantProfile entity is found.
// Returns a *NotFoundError when no ApplicantProfile entities are found.
func (apq *ApplicantProfileQuery) Only(ctx context.Context) (*ApplicantProfile, error) {
	nodes, err := apq.Limit(2).All(setContextOp(ctx, apq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{applicantprofile.Label}
	default:
		return nil, &NotSingularError{applicantprofile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (apq *ApplicantProfileQuery) OnlyX(ctx context.Context) *ApplicantProfile {
	node, err := apq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ApplicantProfile ID in the query.
// Returns a *NotSingularError when more than one ApplicantProfile ID is found.
// Returns a *NotFoundError when no entities are found.
func (apq *ApplicantProfileQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = apq.Limit(2).IDs(setContextOp(ctx, apq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{applicantprofile.Label}
	default:
		err = &NotSingularError{applicantprofile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (apq *ApplicantProfileQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := apq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ApplicantProfiles.
func (apq *ApplicantProfileQuery) All(ctx context.Context) ([]*ApplicantProfile, error) {
	ctx = setContextOp(ctx, apq.ctx, "All")
	if err := apq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ApplicantProfile, *ApplicantProfileQuery]()
	return withInterceptors[[]*ApplicantProfile](ctx, apq, qr, apq.inters)
}

// AllX is like All, but panics if an error occurs.
func (apq *ApplicantProfileQuery) AllX(ctx context.Context) []*ApplicantProfile {
	nodes, err := apq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ApplicantProfile IDs.
func (apq *ApplicantProfileQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if apq.ctx.Unique == nil && apq.path != nil {
		apq.Unique(true)
	}
	ctx = setContextOp(ctx, apq.ctx, "IDs")
	if err = apq.Select(applicantprofile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (apq *ApplicantProfileQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := apq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (apq *ApplicantProfileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, apq.ctx, "Count")
	if err := apq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, apq, querierCount[*ApplicantProfileQuery](), apq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (apq *ApplicantProfileQuery) CountX(ctx context.Context) int {
	count, err := apq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (apq *ApplicantProfileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, apq.ctx, "Exist")
	switch _, err := apq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (apq *ApplicantProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := apq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ApplicantProfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (apq *ApplicantProfileQuery) Clone() *ApplicantProfileQuery {
	if apq == nil {
		return nil
	}
	return &ApplicantProfileQuery{
		config:                     apq.config,
		ctx:                        apq.ctx.Clone(),
		order:                      append([]applicantprofile.OrderOption{}, apq.order...),
		inters:                     append([]Interceptor{}, apq.inters...),
		predicates:                 append([]predicate.ApplicantProfile{}, apq.predicates...),
		withApplicantInterests:     apq.withApplicantInterests.Clone(),
		withApplicantProfileSkills: apq.withApplicantProfileSkills.Clone(),
		withUser:                   apq.withUser.Clone(),
		withApplications:           apq.withApplications.Clone(),
		withEducations:             apq.withEducations.Clone(),
		withLanguages:              apq.withLanguages.Clone(),
		withWorkExperiences:        apq.withWorkExperiences.Clone(),
		// clone intermediate query.
		sql:  apq.sql.Clone(),
		path: apq.path,
	}
}

// WithApplicantInterests tells the query-builder to eager-load the nodes that are connected to
// the "applicant_interests" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *ApplicantProfileQuery) WithApplicantInterests(opts ...func(*ApplicantInterestQuery)) *ApplicantProfileQuery {
	query := (&ApplicantInterestClient{config: apq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	apq.withApplicantInterests = query
	return apq
}

// WithApplicantProfileSkills tells the query-builder to eager-load the nodes that are connected to
// the "applicant_profile_skills" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *ApplicantProfileQuery) WithApplicantProfileSkills(opts ...func(*ApplicantProfileSkillQuery)) *ApplicantProfileQuery {
	query := (&ApplicantProfileSkillClient{config: apq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	apq.withApplicantProfileSkills = query
	return apq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *ApplicantProfileQuery) WithUser(opts ...func(*UserQuery)) *ApplicantProfileQuery {
	query := (&UserClient{config: apq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	apq.withUser = query
	return apq
}

// WithApplications tells the query-builder to eager-load the nodes that are connected to
// the "applications" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *ApplicantProfileQuery) WithApplications(opts ...func(*ApplicationQuery)) *ApplicantProfileQuery {
	query := (&ApplicationClient{config: apq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	apq.withApplications = query
	return apq
}

// WithEducations tells the query-builder to eager-load the nodes that are connected to
// the "educations" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *ApplicantProfileQuery) WithEducations(opts ...func(*EducationQuery)) *ApplicantProfileQuery {
	query := (&EducationClient{config: apq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	apq.withEducations = query
	return apq
}

// WithLanguages tells the query-builder to eager-load the nodes that are connected to
// the "languages" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *ApplicantProfileQuery) WithLanguages(opts ...func(*LanguageQuery)) *ApplicantProfileQuery {
	query := (&LanguageClient{config: apq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	apq.withLanguages = query
	return apq
}

// WithWorkExperiences tells the query-builder to eager-load the nodes that are connected to
// the "work_experiences" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *ApplicantProfileQuery) WithWorkExperiences(opts ...func(*WorkExperienceQuery)) *ApplicantProfileQuery {
	query := (&WorkExperienceClient{config: apq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	apq.withWorkExperiences = query
	return apq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Birthday time.Time `json:"birthday,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ApplicantProfile.Query().
//		GroupBy(applicantprofile.FieldBirthday).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (apq *ApplicantProfileQuery) GroupBy(field string, fields ...string) *ApplicantProfileGroupBy {
	apq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ApplicantProfileGroupBy{build: apq}
	grbuild.flds = &apq.ctx.Fields
	grbuild.label = applicantprofile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Birthday time.Time `json:"birthday,omitempty"`
//	}
//
//	client.ApplicantProfile.Query().
//		Select(applicantprofile.FieldBirthday).
//		Scan(ctx, &v)
func (apq *ApplicantProfileQuery) Select(fields ...string) *ApplicantProfileSelect {
	apq.ctx.Fields = append(apq.ctx.Fields, fields...)
	sbuild := &ApplicantProfileSelect{ApplicantProfileQuery: apq}
	sbuild.label = applicantprofile.Label
	sbuild.flds, sbuild.scan = &apq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ApplicantProfileSelect configured with the given aggregations.
func (apq *ApplicantProfileQuery) Aggregate(fns ...AggregateFunc) *ApplicantProfileSelect {
	return apq.Select().Aggregate(fns...)
}

func (apq *ApplicantProfileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range apq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, apq); err != nil {
				return err
			}
		}
	}
	for _, f := range apq.ctx.Fields {
		if !applicantprofile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if apq.path != nil {
		prev, err := apq.path(ctx)
		if err != nil {
			return err
		}
		apq.sql = prev
	}
	return nil
}

func (apq *ApplicantProfileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ApplicantProfile, error) {
	var (
		nodes       = []*ApplicantProfile{}
		_spec       = apq.querySpec()
		loadedTypes = [7]bool{
			apq.withApplicantInterests != nil,
			apq.withApplicantProfileSkills != nil,
			apq.withUser != nil,
			apq.withApplications != nil,
			apq.withEducations != nil,
			apq.withLanguages != nil,
			apq.withWorkExperiences != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ApplicantProfile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ApplicantProfile{config: apq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, apq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := apq.withApplicantInterests; query != nil {
		if err := apq.loadApplicantInterests(ctx, query, nodes,
			func(n *ApplicantProfile) { n.Edges.ApplicantInterests = []*ApplicantInterest{} },
			func(n *ApplicantProfile, e *ApplicantInterest) {
				n.Edges.ApplicantInterests = append(n.Edges.ApplicantInterests, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := apq.withApplicantProfileSkills; query != nil {
		if err := apq.loadApplicantProfileSkills(ctx, query, nodes,
			func(n *ApplicantProfile) { n.Edges.ApplicantProfileSkills = []*ApplicantProfileSkill{} },
			func(n *ApplicantProfile, e *ApplicantProfileSkill) {
				n.Edges.ApplicantProfileSkills = append(n.Edges.ApplicantProfileSkills, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := apq.withUser; query != nil {
		if err := apq.loadUser(ctx, query, nodes, nil,
			func(n *ApplicantProfile, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := apq.withApplications; query != nil {
		if err := apq.loadApplications(ctx, query, nodes,
			func(n *ApplicantProfile) { n.Edges.Applications = []*Application{} },
			func(n *ApplicantProfile, e *Application) { n.Edges.Applications = append(n.Edges.Applications, e) }); err != nil {
			return nil, err
		}
	}
	if query := apq.withEducations; query != nil {
		if err := apq.loadEducations(ctx, query, nodes,
			func(n *ApplicantProfile) { n.Edges.Educations = []*Education{} },
			func(n *ApplicantProfile, e *Education) { n.Edges.Educations = append(n.Edges.Educations, e) }); err != nil {
			return nil, err
		}
	}
	if query := apq.withLanguages; query != nil {
		if err := apq.loadLanguages(ctx, query, nodes,
			func(n *ApplicantProfile) { n.Edges.Languages = []*Language{} },
			func(n *ApplicantProfile, e *Language) { n.Edges.Languages = append(n.Edges.Languages, e) }); err != nil {
			return nil, err
		}
	}
	if query := apq.withWorkExperiences; query != nil {
		if err := apq.loadWorkExperiences(ctx, query, nodes,
			func(n *ApplicantProfile) { n.Edges.WorkExperiences = []*WorkExperience{} },
			func(n *ApplicantProfile, e *WorkExperience) {
				n.Edges.WorkExperiences = append(n.Edges.WorkExperiences, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (apq *ApplicantProfileQuery) loadApplicantInterests(ctx context.Context, query *ApplicantInterestQuery, nodes []*ApplicantProfile, init func(*ApplicantProfile), assign func(*ApplicantProfile, *ApplicantInterest)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*ApplicantProfile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(applicantinterest.FieldApplicantProfileID)
	}
	query.Where(predicate.ApplicantInterest(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(applicantprofile.ApplicantInterestsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ApplicantProfileID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "applicant_profile_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (apq *ApplicantProfileQuery) loadApplicantProfileSkills(ctx context.Context, query *ApplicantProfileSkillQuery, nodes []*ApplicantProfile, init func(*ApplicantProfile), assign func(*ApplicantProfile, *ApplicantProfileSkill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*ApplicantProfile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(applicantprofileskill.FieldApplicantProfileID)
	}
	query.Where(predicate.ApplicantProfileSkill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(applicantprofile.ApplicantProfileSkillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ApplicantProfileID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "applicant_profile_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (apq *ApplicantProfileQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ApplicantProfile, init func(*ApplicantProfile), assign func(*ApplicantProfile, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ApplicantProfile)
	for i := range nodes {
		if nodes[i].UserID == nil {
			continue
		}
		fk := *nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (apq *ApplicantProfileQuery) loadApplications(ctx context.Context, query *ApplicationQuery, nodes []*ApplicantProfile, init func(*ApplicantProfile), assign func(*ApplicantProfile, *Application)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*ApplicantProfile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(application.FieldApplicantProfileID)
	}
	query.Where(predicate.Application(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(applicantprofile.ApplicationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ApplicantProfileID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "applicant_profile_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (apq *ApplicantProfileQuery) loadEducations(ctx context.Context, query *EducationQuery, nodes []*ApplicantProfile, init func(*ApplicantProfile), assign func(*ApplicantProfile, *Education)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*ApplicantProfile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Education(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(applicantprofile.EducationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.applicant_profile_educations
		if fk == nil {
			return fmt.Errorf(`foreign-key "applicant_profile_educations" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "applicant_profile_educations" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (apq *ApplicantProfileQuery) loadLanguages(ctx context.Context, query *LanguageQuery, nodes []*ApplicantProfile, init func(*ApplicantProfile), assign func(*ApplicantProfile, *Language)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*ApplicantProfile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(language.FieldApplicantProfileID)
	}
	query.Where(predicate.Language(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(applicantprofile.LanguagesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ApplicantProfileID
		if fk == nil {
			return fmt.Errorf(`foreign-key "applicant_profile_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "applicant_profile_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (apq *ApplicantProfileQuery) loadWorkExperiences(ctx context.Context, query *WorkExperienceQuery, nodes []*ApplicantProfile, init func(*ApplicantProfile), assign func(*ApplicantProfile, *WorkExperience)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*ApplicantProfile)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(workexperience.FieldApplicantProfileID)
	}
	query.Where(predicate.WorkExperience(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(applicantprofile.WorkExperiencesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ApplicantProfileID
		if fk == nil {
			return fmt.Errorf(`foreign-key "applicant_profile_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "applicant_profile_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (apq *ApplicantProfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := apq.querySpec()
	_spec.Node.Columns = apq.ctx.Fields
	if len(apq.ctx.Fields) > 0 {
		_spec.Unique = apq.ctx.Unique != nil && *apq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, apq.driver, _spec)
}

func (apq *ApplicantProfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(applicantprofile.Table, applicantprofile.Columns, sqlgraph.NewFieldSpec(applicantprofile.FieldID, field.TypeUUID))
	_spec.From = apq.sql
	if unique := apq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if apq.path != nil {
		_spec.Unique = true
	}
	if fields := apq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applicantprofile.FieldID)
		for i := range fields {
			if fields[i] != applicantprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if apq.withUser != nil {
			_spec.Node.AddColumnOnce(applicantprofile.FieldUserID)
		}
	}
	if ps := apq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := apq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := apq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := apq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (apq *ApplicantProfileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(apq.driver.Dialect())
	t1 := builder.Table(applicantprofile.Table)
	columns := apq.ctx.Fields
	if len(columns) == 0 {
		columns = applicantprofile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if apq.sql != nil {
		selector = apq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if apq.ctx.Unique != nil && *apq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range apq.predicates {
		p(selector)
	}
	for _, p := range apq.order {
		p(selector)
	}
	if offset := apq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := apq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ApplicantProfileGroupBy is the group-by builder for ApplicantProfile entities.
type ApplicantProfileGroupBy struct {
	selector
	build *ApplicantProfileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (apgb *ApplicantProfileGroupBy) Aggregate(fns ...AggregateFunc) *ApplicantProfileGroupBy {
	apgb.fns = append(apgb.fns, fns...)
	return apgb
}

// Scan applies the selector query and scans the result into the given value.
func (apgb *ApplicantProfileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, apgb.build.ctx, "GroupBy")
	if err := apgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApplicantProfileQuery, *ApplicantProfileGroupBy](ctx, apgb.build, apgb, apgb.build.inters, v)
}

func (apgb *ApplicantProfileGroupBy) sqlScan(ctx context.Context, root *ApplicantProfileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(apgb.fns))
	for _, fn := range apgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*apgb.flds)+len(apgb.fns))
		for _, f := range *apgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*apgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := apgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ApplicantProfileSelect is the builder for selecting fields of ApplicantProfile entities.
type ApplicantProfileSelect struct {
	*ApplicantProfileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aps *ApplicantProfileSelect) Aggregate(fns ...AggregateFunc) *ApplicantProfileSelect {
	aps.fns = append(aps.fns, fns...)
	return aps
}

// Scan applies the selector query and scans the result into the given value.
func (aps *ApplicantProfileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aps.ctx, "Select")
	if err := aps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApplicantProfileQuery, *ApplicantProfileSelect](ctx, aps.ApplicantProfileQuery, aps, aps.inters, v)
}

func (aps *ApplicantProfileSelect) sqlScan(ctx context.Context, root *ApplicantProfileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aps.fns))
	for _, fn := range aps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}