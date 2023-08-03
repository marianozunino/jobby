// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/marianozunino/jobby/ent/application"
	"github.com/marianozunino/jobby/ent/interview"
	"github.com/marianozunino/jobby/ent/predicate"
	"github.com/marianozunino/jobby/ent/user"
)

// InterviewQuery is the builder for querying Interview entities.
type InterviewQuery struct {
	config
	ctx             *QueryContext
	order           []interview.OrderOption
	inters          []Interceptor
	predicates      []predicate.Interview
	withApplication *ApplicationQuery
	withUser        *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InterviewQuery builder.
func (iq *InterviewQuery) Where(ps ...predicate.Interview) *InterviewQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *InterviewQuery) Limit(limit int) *InterviewQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *InterviewQuery) Offset(offset int) *InterviewQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InterviewQuery) Unique(unique bool) *InterviewQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *InterviewQuery) Order(o ...interview.OrderOption) *InterviewQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryApplication chains the current query on the "application" edge.
func (iq *InterviewQuery) QueryApplication() *ApplicationQuery {
	query := (&ApplicationClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(interview.Table, interview.FieldID, selector),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, interview.ApplicationTable, interview.ApplicationColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (iq *InterviewQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(interview.Table, interview.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, interview.UserTable, interview.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Interview entity from the query.
// Returns a *NotFoundError when no Interview was found.
func (iq *InterviewQuery) First(ctx context.Context) (*Interview, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{interview.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InterviewQuery) FirstX(ctx context.Context) *Interview {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Interview ID from the query.
// Returns a *NotFoundError when no Interview ID was found.
func (iq *InterviewQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{interview.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InterviewQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Interview entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Interview entity is found.
// Returns a *NotFoundError when no Interview entities are found.
func (iq *InterviewQuery) Only(ctx context.Context) (*Interview, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{interview.Label}
	default:
		return nil, &NotSingularError{interview.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InterviewQuery) OnlyX(ctx context.Context) *Interview {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Interview ID in the query.
// Returns a *NotSingularError when more than one Interview ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InterviewQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{interview.Label}
	default:
		err = &NotSingularError{interview.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InterviewQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Interviews.
func (iq *InterviewQuery) All(ctx context.Context) ([]*Interview, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Interview, *InterviewQuery]()
	return withInterceptors[[]*Interview](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *InterviewQuery) AllX(ctx context.Context) []*Interview {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Interview IDs.
func (iq *InterviewQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(interview.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InterviewQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InterviewQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*InterviewQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InterviewQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InterviewQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InterviewQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InterviewQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InterviewQuery) Clone() *InterviewQuery {
	if iq == nil {
		return nil
	}
	return &InterviewQuery{
		config:          iq.config,
		ctx:             iq.ctx.Clone(),
		order:           append([]interview.OrderOption{}, iq.order...),
		inters:          append([]Interceptor{}, iq.inters...),
		predicates:      append([]predicate.Interview{}, iq.predicates...),
		withApplication: iq.withApplication.Clone(),
		withUser:        iq.withUser.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithApplication tells the query-builder to eager-load the nodes that are connected to
// the "application" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InterviewQuery) WithApplication(opts ...func(*ApplicationQuery)) *InterviewQuery {
	query := (&ApplicationClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withApplication = query
	return iq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InterviewQuery) WithUser(opts ...func(*UserQuery)) *InterviewQuery {
	query := (&UserClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withUser = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Comment string `json:"comment,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Interview.Query().
//		GroupBy(interview.FieldComment).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *InterviewQuery) GroupBy(field string, fields ...string) *InterviewGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InterviewGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = interview.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Comment string `json:"comment,omitempty"`
//	}
//
//	client.Interview.Query().
//		Select(interview.FieldComment).
//		Scan(ctx, &v)
func (iq *InterviewQuery) Select(fields ...string) *InterviewSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &InterviewSelect{InterviewQuery: iq}
	sbuild.label = interview.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InterviewSelect configured with the given aggregations.
func (iq *InterviewQuery) Aggregate(fns ...AggregateFunc) *InterviewSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *InterviewQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !interview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InterviewQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Interview, error) {
	var (
		nodes       = []*Interview{}
		_spec       = iq.querySpec()
		loadedTypes = [2]bool{
			iq.withApplication != nil,
			iq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Interview).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Interview{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withApplication; query != nil {
		if err := iq.loadApplication(ctx, query, nodes, nil,
			func(n *Interview, e *Application) { n.Edges.Application = e }); err != nil {
			return nil, err
		}
	}
	if query := iq.withUser; query != nil {
		if err := iq.loadUser(ctx, query, nodes, nil,
			func(n *Interview, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *InterviewQuery) loadApplication(ctx context.Context, query *ApplicationQuery, nodes []*Interview, init func(*Interview), assign func(*Interview, *Application)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Interview)
	for i := range nodes {
		if nodes[i].ApplicationID == nil {
			continue
		}
		fk := *nodes[i].ApplicationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(application.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "application_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (iq *InterviewQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Interview, init func(*Interview), assign func(*Interview, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Interview)
	for i := range nodes {
		fk := nodes[i].InterviewerID
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
			return fmt.Errorf(`unexpected foreign-key "interviewer_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (iq *InterviewQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InterviewQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(interview.Table, interview.Columns, sqlgraph.NewFieldSpec(interview.FieldID, field.TypeUUID))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, interview.FieldID)
		for i := range fields {
			if fields[i] != interview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if iq.withApplication != nil {
			_spec.Node.AddColumnOnce(interview.FieldApplicationID)
		}
		if iq.withUser != nil {
			_spec.Node.AddColumnOnce(interview.FieldInterviewerID)
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InterviewQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(interview.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = interview.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InterviewGroupBy is the group-by builder for Interview entities.
type InterviewGroupBy struct {
	selector
	build *InterviewQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InterviewGroupBy) Aggregate(fns ...AggregateFunc) *InterviewGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *InterviewGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InterviewQuery, *InterviewGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *InterviewGroupBy) sqlScan(ctx context.Context, root *InterviewQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InterviewSelect is the builder for selecting fields of Interview entities.
type InterviewSelect struct {
	*InterviewQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *InterviewSelect) Aggregate(fns ...AggregateFunc) *InterviewSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *InterviewSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InterviewQuery, *InterviewSelect](ctx, is.InterviewQuery, is, is.inters, v)
}

func (is *InterviewSelect) sqlScan(ctx context.Context, root *InterviewQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
