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
	"github.com/marianozunino/jobby/ent/degreelevel"
	"github.com/marianozunino/jobby/ent/education"
	"github.com/marianozunino/jobby/ent/predicate"
)

// DegreeLevelQuery is the builder for querying DegreeLevel entities.
type DegreeLevelQuery struct {
	config
	ctx            *QueryContext
	order          []degreelevel.OrderOption
	inters         []Interceptor
	predicates     []predicate.DegreeLevel
	withEducations *EducationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DegreeLevelQuery builder.
func (dlq *DegreeLevelQuery) Where(ps ...predicate.DegreeLevel) *DegreeLevelQuery {
	dlq.predicates = append(dlq.predicates, ps...)
	return dlq
}

// Limit the number of records to be returned by this query.
func (dlq *DegreeLevelQuery) Limit(limit int) *DegreeLevelQuery {
	dlq.ctx.Limit = &limit
	return dlq
}

// Offset to start from.
func (dlq *DegreeLevelQuery) Offset(offset int) *DegreeLevelQuery {
	dlq.ctx.Offset = &offset
	return dlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dlq *DegreeLevelQuery) Unique(unique bool) *DegreeLevelQuery {
	dlq.ctx.Unique = &unique
	return dlq
}

// Order specifies how the records should be ordered.
func (dlq *DegreeLevelQuery) Order(o ...degreelevel.OrderOption) *DegreeLevelQuery {
	dlq.order = append(dlq.order, o...)
	return dlq
}

// QueryEducations chains the current query on the "educations" edge.
func (dlq *DegreeLevelQuery) QueryEducations() *EducationQuery {
	query := (&EducationClient{config: dlq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dlq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(degreelevel.Table, degreelevel.FieldID, selector),
			sqlgraph.To(education.Table, education.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, degreelevel.EducationsTable, degreelevel.EducationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dlq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DegreeLevel entity from the query.
// Returns a *NotFoundError when no DegreeLevel was found.
func (dlq *DegreeLevelQuery) First(ctx context.Context) (*DegreeLevel, error) {
	nodes, err := dlq.Limit(1).All(setContextOp(ctx, dlq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{degreelevel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dlq *DegreeLevelQuery) FirstX(ctx context.Context) *DegreeLevel {
	node, err := dlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DegreeLevel ID from the query.
// Returns a *NotFoundError when no DegreeLevel ID was found.
func (dlq *DegreeLevelQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dlq.Limit(1).IDs(setContextOp(ctx, dlq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{degreelevel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dlq *DegreeLevelQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DegreeLevel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DegreeLevel entity is found.
// Returns a *NotFoundError when no DegreeLevel entities are found.
func (dlq *DegreeLevelQuery) Only(ctx context.Context) (*DegreeLevel, error) {
	nodes, err := dlq.Limit(2).All(setContextOp(ctx, dlq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{degreelevel.Label}
	default:
		return nil, &NotSingularError{degreelevel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dlq *DegreeLevelQuery) OnlyX(ctx context.Context) *DegreeLevel {
	node, err := dlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DegreeLevel ID in the query.
// Returns a *NotSingularError when more than one DegreeLevel ID is found.
// Returns a *NotFoundError when no entities are found.
func (dlq *DegreeLevelQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dlq.Limit(2).IDs(setContextOp(ctx, dlq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{degreelevel.Label}
	default:
		err = &NotSingularError{degreelevel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dlq *DegreeLevelQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DegreeLevels.
func (dlq *DegreeLevelQuery) All(ctx context.Context) ([]*DegreeLevel, error) {
	ctx = setContextOp(ctx, dlq.ctx, "All")
	if err := dlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DegreeLevel, *DegreeLevelQuery]()
	return withInterceptors[[]*DegreeLevel](ctx, dlq, qr, dlq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dlq *DegreeLevelQuery) AllX(ctx context.Context) []*DegreeLevel {
	nodes, err := dlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DegreeLevel IDs.
func (dlq *DegreeLevelQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dlq.ctx.Unique == nil && dlq.path != nil {
		dlq.Unique(true)
	}
	ctx = setContextOp(ctx, dlq.ctx, "IDs")
	if err = dlq.Select(degreelevel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dlq *DegreeLevelQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dlq *DegreeLevelQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dlq.ctx, "Count")
	if err := dlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dlq, querierCount[*DegreeLevelQuery](), dlq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dlq *DegreeLevelQuery) CountX(ctx context.Context) int {
	count, err := dlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dlq *DegreeLevelQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dlq.ctx, "Exist")
	switch _, err := dlq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dlq *DegreeLevelQuery) ExistX(ctx context.Context) bool {
	exist, err := dlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DegreeLevelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dlq *DegreeLevelQuery) Clone() *DegreeLevelQuery {
	if dlq == nil {
		return nil
	}
	return &DegreeLevelQuery{
		config:         dlq.config,
		ctx:            dlq.ctx.Clone(),
		order:          append([]degreelevel.OrderOption{}, dlq.order...),
		inters:         append([]Interceptor{}, dlq.inters...),
		predicates:     append([]predicate.DegreeLevel{}, dlq.predicates...),
		withEducations: dlq.withEducations.Clone(),
		// clone intermediate query.
		sql:  dlq.sql.Clone(),
		path: dlq.path,
	}
}

// WithEducations tells the query-builder to eager-load the nodes that are connected to
// the "educations" edge. The optional arguments are used to configure the query builder of the edge.
func (dlq *DegreeLevelQuery) WithEducations(opts ...func(*EducationQuery)) *DegreeLevelQuery {
	query := (&EducationClient{config: dlq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dlq.withEducations = query
	return dlq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DegreeLevel.Query().
//		GroupBy(degreelevel.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dlq *DegreeLevelQuery) GroupBy(field string, fields ...string) *DegreeLevelGroupBy {
	dlq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DegreeLevelGroupBy{build: dlq}
	grbuild.flds = &dlq.ctx.Fields
	grbuild.label = degreelevel.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.DegreeLevel.Query().
//		Select(degreelevel.FieldName).
//		Scan(ctx, &v)
func (dlq *DegreeLevelQuery) Select(fields ...string) *DegreeLevelSelect {
	dlq.ctx.Fields = append(dlq.ctx.Fields, fields...)
	sbuild := &DegreeLevelSelect{DegreeLevelQuery: dlq}
	sbuild.label = degreelevel.Label
	sbuild.flds, sbuild.scan = &dlq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DegreeLevelSelect configured with the given aggregations.
func (dlq *DegreeLevelQuery) Aggregate(fns ...AggregateFunc) *DegreeLevelSelect {
	return dlq.Select().Aggregate(fns...)
}

func (dlq *DegreeLevelQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dlq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dlq); err != nil {
				return err
			}
		}
	}
	for _, f := range dlq.ctx.Fields {
		if !degreelevel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dlq.path != nil {
		prev, err := dlq.path(ctx)
		if err != nil {
			return err
		}
		dlq.sql = prev
	}
	return nil
}

func (dlq *DegreeLevelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DegreeLevel, error) {
	var (
		nodes       = []*DegreeLevel{}
		_spec       = dlq.querySpec()
		loadedTypes = [1]bool{
			dlq.withEducations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DegreeLevel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DegreeLevel{config: dlq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dlq.withEducations; query != nil {
		if err := dlq.loadEducations(ctx, query, nodes,
			func(n *DegreeLevel) { n.Edges.Educations = []*Education{} },
			func(n *DegreeLevel, e *Education) { n.Edges.Educations = append(n.Edges.Educations, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dlq *DegreeLevelQuery) loadEducations(ctx context.Context, query *EducationQuery, nodes []*DegreeLevel, init func(*DegreeLevel), assign func(*DegreeLevel, *Education)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*DegreeLevel)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(education.FieldDegreeLevelID)
	}
	query.Where(predicate.Education(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(degreelevel.EducationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.DegreeLevelID
		if fk == nil {
			return fmt.Errorf(`foreign-key "degree_level_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "degree_level_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dlq *DegreeLevelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dlq.querySpec()
	_spec.Node.Columns = dlq.ctx.Fields
	if len(dlq.ctx.Fields) > 0 {
		_spec.Unique = dlq.ctx.Unique != nil && *dlq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dlq.driver, _spec)
}

func (dlq *DegreeLevelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(degreelevel.Table, degreelevel.Columns, sqlgraph.NewFieldSpec(degreelevel.FieldID, field.TypeUUID))
	_spec.From = dlq.sql
	if unique := dlq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dlq.path != nil {
		_spec.Unique = true
	}
	if fields := dlq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, degreelevel.FieldID)
		for i := range fields {
			if fields[i] != degreelevel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dlq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dlq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dlq *DegreeLevelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dlq.driver.Dialect())
	t1 := builder.Table(degreelevel.Table)
	columns := dlq.ctx.Fields
	if len(columns) == 0 {
		columns = degreelevel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dlq.sql != nil {
		selector = dlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dlq.ctx.Unique != nil && *dlq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dlq.predicates {
		p(selector)
	}
	for _, p := range dlq.order {
		p(selector)
	}
	if offset := dlq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dlq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DegreeLevelGroupBy is the group-by builder for DegreeLevel entities.
type DegreeLevelGroupBy struct {
	selector
	build *DegreeLevelQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dlgb *DegreeLevelGroupBy) Aggregate(fns ...AggregateFunc) *DegreeLevelGroupBy {
	dlgb.fns = append(dlgb.fns, fns...)
	return dlgb
}

// Scan applies the selector query and scans the result into the given value.
func (dlgb *DegreeLevelGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dlgb.build.ctx, "GroupBy")
	if err := dlgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DegreeLevelQuery, *DegreeLevelGroupBy](ctx, dlgb.build, dlgb, dlgb.build.inters, v)
}

func (dlgb *DegreeLevelGroupBy) sqlScan(ctx context.Context, root *DegreeLevelQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dlgb.fns))
	for _, fn := range dlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dlgb.flds)+len(dlgb.fns))
		for _, f := range *dlgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dlgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dlgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DegreeLevelSelect is the builder for selecting fields of DegreeLevel entities.
type DegreeLevelSelect struct {
	*DegreeLevelQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dls *DegreeLevelSelect) Aggregate(fns ...AggregateFunc) *DegreeLevelSelect {
	dls.fns = append(dls.fns, fns...)
	return dls
}

// Scan applies the selector query and scans the result into the given value.
func (dls *DegreeLevelSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dls.ctx, "Select")
	if err := dls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DegreeLevelQuery, *DegreeLevelSelect](ctx, dls.DegreeLevelQuery, dls, dls.inters, v)
}

func (dls *DegreeLevelSelect) sqlScan(ctx context.Context, root *DegreeLevelQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dls.fns))
	for _, fn := range dls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
