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
	"github.com/marianozunino/jobby/ent/category"
	"github.com/marianozunino/jobby/ent/joboffer"
	"github.com/marianozunino/jobby/ent/joboffercategory"
	"github.com/marianozunino/jobby/ent/predicate"
)

// JobOfferCategoryQuery is the builder for querying JobOfferCategory entities.
type JobOfferCategoryQuery struct {
	config
	ctx          *QueryContext
	order        []joboffercategory.OrderOption
	inters       []Interceptor
	predicates   []predicate.JobOfferCategory
	withCategory *CategoryQuery
	withJobOffer *JobOfferQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobOfferCategoryQuery builder.
func (jocq *JobOfferCategoryQuery) Where(ps ...predicate.JobOfferCategory) *JobOfferCategoryQuery {
	jocq.predicates = append(jocq.predicates, ps...)
	return jocq
}

// Limit the number of records to be returned by this query.
func (jocq *JobOfferCategoryQuery) Limit(limit int) *JobOfferCategoryQuery {
	jocq.ctx.Limit = &limit
	return jocq
}

// Offset to start from.
func (jocq *JobOfferCategoryQuery) Offset(offset int) *JobOfferCategoryQuery {
	jocq.ctx.Offset = &offset
	return jocq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jocq *JobOfferCategoryQuery) Unique(unique bool) *JobOfferCategoryQuery {
	jocq.ctx.Unique = &unique
	return jocq
}

// Order specifies how the records should be ordered.
func (jocq *JobOfferCategoryQuery) Order(o ...joboffercategory.OrderOption) *JobOfferCategoryQuery {
	jocq.order = append(jocq.order, o...)
	return jocq
}

// QueryCategory chains the current query on the "category" edge.
func (jocq *JobOfferCategoryQuery) QueryCategory() *CategoryQuery {
	query := (&CategoryClient{config: jocq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jocq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(joboffercategory.Table, joboffercategory.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, joboffercategory.CategoryTable, joboffercategory.CategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(jocq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryJobOffer chains the current query on the "job_offer" edge.
func (jocq *JobOfferCategoryQuery) QueryJobOffer() *JobOfferQuery {
	query := (&JobOfferClient{config: jocq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jocq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(joboffercategory.Table, joboffercategory.FieldID, selector),
			sqlgraph.To(joboffer.Table, joboffer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, joboffercategory.JobOfferTable, joboffercategory.JobOfferColumn),
		)
		fromU = sqlgraph.SetNeighbors(jocq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first JobOfferCategory entity from the query.
// Returns a *NotFoundError when no JobOfferCategory was found.
func (jocq *JobOfferCategoryQuery) First(ctx context.Context) (*JobOfferCategory, error) {
	nodes, err := jocq.Limit(1).All(setContextOp(ctx, jocq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{joboffercategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jocq *JobOfferCategoryQuery) FirstX(ctx context.Context) *JobOfferCategory {
	node, err := jocq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JobOfferCategory ID from the query.
// Returns a *NotFoundError when no JobOfferCategory ID was found.
func (jocq *JobOfferCategoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = jocq.Limit(1).IDs(setContextOp(ctx, jocq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{joboffercategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jocq *JobOfferCategoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := jocq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single JobOfferCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one JobOfferCategory entity is found.
// Returns a *NotFoundError when no JobOfferCategory entities are found.
func (jocq *JobOfferCategoryQuery) Only(ctx context.Context) (*JobOfferCategory, error) {
	nodes, err := jocq.Limit(2).All(setContextOp(ctx, jocq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{joboffercategory.Label}
	default:
		return nil, &NotSingularError{joboffercategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jocq *JobOfferCategoryQuery) OnlyX(ctx context.Context) *JobOfferCategory {
	node, err := jocq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only JobOfferCategory ID in the query.
// Returns a *NotSingularError when more than one JobOfferCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (jocq *JobOfferCategoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = jocq.Limit(2).IDs(setContextOp(ctx, jocq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{joboffercategory.Label}
	default:
		err = &NotSingularError{joboffercategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jocq *JobOfferCategoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := jocq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JobOfferCategories.
func (jocq *JobOfferCategoryQuery) All(ctx context.Context) ([]*JobOfferCategory, error) {
	ctx = setContextOp(ctx, jocq.ctx, "All")
	if err := jocq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*JobOfferCategory, *JobOfferCategoryQuery]()
	return withInterceptors[[]*JobOfferCategory](ctx, jocq, qr, jocq.inters)
}

// AllX is like All, but panics if an error occurs.
func (jocq *JobOfferCategoryQuery) AllX(ctx context.Context) []*JobOfferCategory {
	nodes, err := jocq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JobOfferCategory IDs.
func (jocq *JobOfferCategoryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if jocq.ctx.Unique == nil && jocq.path != nil {
		jocq.Unique(true)
	}
	ctx = setContextOp(ctx, jocq.ctx, "IDs")
	if err = jocq.Select(joboffercategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jocq *JobOfferCategoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := jocq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jocq *JobOfferCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, jocq.ctx, "Count")
	if err := jocq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, jocq, querierCount[*JobOfferCategoryQuery](), jocq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (jocq *JobOfferCategoryQuery) CountX(ctx context.Context) int {
	count, err := jocq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jocq *JobOfferCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, jocq.ctx, "Exist")
	switch _, err := jocq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (jocq *JobOfferCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := jocq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobOfferCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jocq *JobOfferCategoryQuery) Clone() *JobOfferCategoryQuery {
	if jocq == nil {
		return nil
	}
	return &JobOfferCategoryQuery{
		config:       jocq.config,
		ctx:          jocq.ctx.Clone(),
		order:        append([]joboffercategory.OrderOption{}, jocq.order...),
		inters:       append([]Interceptor{}, jocq.inters...),
		predicates:   append([]predicate.JobOfferCategory{}, jocq.predicates...),
		withCategory: jocq.withCategory.Clone(),
		withJobOffer: jocq.withJobOffer.Clone(),
		// clone intermediate query.
		sql:  jocq.sql.Clone(),
		path: jocq.path,
	}
}

// WithCategory tells the query-builder to eager-load the nodes that are connected to
// the "category" edge. The optional arguments are used to configure the query builder of the edge.
func (jocq *JobOfferCategoryQuery) WithCategory(opts ...func(*CategoryQuery)) *JobOfferCategoryQuery {
	query := (&CategoryClient{config: jocq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	jocq.withCategory = query
	return jocq
}

// WithJobOffer tells the query-builder to eager-load the nodes that are connected to
// the "job_offer" edge. The optional arguments are used to configure the query builder of the edge.
func (jocq *JobOfferCategoryQuery) WithJobOffer(opts ...func(*JobOfferQuery)) *JobOfferCategoryQuery {
	query := (&JobOfferClient{config: jocq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	jocq.withJobOffer = query
	return jocq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		JobOfferID uuid.UUID `json:"job_offer_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.JobOfferCategory.Query().
//		GroupBy(joboffercategory.FieldJobOfferID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (jocq *JobOfferCategoryQuery) GroupBy(field string, fields ...string) *JobOfferCategoryGroupBy {
	jocq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &JobOfferCategoryGroupBy{build: jocq}
	grbuild.flds = &jocq.ctx.Fields
	grbuild.label = joboffercategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		JobOfferID uuid.UUID `json:"job_offer_id,omitempty"`
//	}
//
//	client.JobOfferCategory.Query().
//		Select(joboffercategory.FieldJobOfferID).
//		Scan(ctx, &v)
func (jocq *JobOfferCategoryQuery) Select(fields ...string) *JobOfferCategorySelect {
	jocq.ctx.Fields = append(jocq.ctx.Fields, fields...)
	sbuild := &JobOfferCategorySelect{JobOfferCategoryQuery: jocq}
	sbuild.label = joboffercategory.Label
	sbuild.flds, sbuild.scan = &jocq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a JobOfferCategorySelect configured with the given aggregations.
func (jocq *JobOfferCategoryQuery) Aggregate(fns ...AggregateFunc) *JobOfferCategorySelect {
	return jocq.Select().Aggregate(fns...)
}

func (jocq *JobOfferCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range jocq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, jocq); err != nil {
				return err
			}
		}
	}
	for _, f := range jocq.ctx.Fields {
		if !joboffercategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jocq.path != nil {
		prev, err := jocq.path(ctx)
		if err != nil {
			return err
		}
		jocq.sql = prev
	}
	return nil
}

func (jocq *JobOfferCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*JobOfferCategory, error) {
	var (
		nodes       = []*JobOfferCategory{}
		_spec       = jocq.querySpec()
		loadedTypes = [2]bool{
			jocq.withCategory != nil,
			jocq.withJobOffer != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*JobOfferCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &JobOfferCategory{config: jocq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, jocq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := jocq.withCategory; query != nil {
		if err := jocq.loadCategory(ctx, query, nodes, nil,
			func(n *JobOfferCategory, e *Category) { n.Edges.Category = e }); err != nil {
			return nil, err
		}
	}
	if query := jocq.withJobOffer; query != nil {
		if err := jocq.loadJobOffer(ctx, query, nodes, nil,
			func(n *JobOfferCategory, e *JobOffer) { n.Edges.JobOffer = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (jocq *JobOfferCategoryQuery) loadCategory(ctx context.Context, query *CategoryQuery, nodes []*JobOfferCategory, init func(*JobOfferCategory), assign func(*JobOfferCategory, *Category)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*JobOfferCategory)
	for i := range nodes {
		fk := nodes[i].CategoryID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(category.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "category_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (jocq *JobOfferCategoryQuery) loadJobOffer(ctx context.Context, query *JobOfferQuery, nodes []*JobOfferCategory, init func(*JobOfferCategory), assign func(*JobOfferCategory, *JobOffer)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*JobOfferCategory)
	for i := range nodes {
		fk := nodes[i].JobOfferID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(joboffer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "job_offer_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (jocq *JobOfferCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jocq.querySpec()
	_spec.Node.Columns = jocq.ctx.Fields
	if len(jocq.ctx.Fields) > 0 {
		_spec.Unique = jocq.ctx.Unique != nil && *jocq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, jocq.driver, _spec)
}

func (jocq *JobOfferCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(joboffercategory.Table, joboffercategory.Columns, sqlgraph.NewFieldSpec(joboffercategory.FieldID, field.TypeUUID))
	_spec.From = jocq.sql
	if unique := jocq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if jocq.path != nil {
		_spec.Unique = true
	}
	if fields := jocq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, joboffercategory.FieldID)
		for i := range fields {
			if fields[i] != joboffercategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if jocq.withCategory != nil {
			_spec.Node.AddColumnOnce(joboffercategory.FieldCategoryID)
		}
		if jocq.withJobOffer != nil {
			_spec.Node.AddColumnOnce(joboffercategory.FieldJobOfferID)
		}
	}
	if ps := jocq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jocq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jocq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jocq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jocq *JobOfferCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jocq.driver.Dialect())
	t1 := builder.Table(joboffercategory.Table)
	columns := jocq.ctx.Fields
	if len(columns) == 0 {
		columns = joboffercategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jocq.sql != nil {
		selector = jocq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if jocq.ctx.Unique != nil && *jocq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range jocq.predicates {
		p(selector)
	}
	for _, p := range jocq.order {
		p(selector)
	}
	if offset := jocq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jocq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JobOfferCategoryGroupBy is the group-by builder for JobOfferCategory entities.
type JobOfferCategoryGroupBy struct {
	selector
	build *JobOfferCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jocgb *JobOfferCategoryGroupBy) Aggregate(fns ...AggregateFunc) *JobOfferCategoryGroupBy {
	jocgb.fns = append(jocgb.fns, fns...)
	return jocgb
}

// Scan applies the selector query and scans the result into the given value.
func (jocgb *JobOfferCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jocgb.build.ctx, "GroupBy")
	if err := jocgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobOfferCategoryQuery, *JobOfferCategoryGroupBy](ctx, jocgb.build, jocgb, jocgb.build.inters, v)
}

func (jocgb *JobOfferCategoryGroupBy) sqlScan(ctx context.Context, root *JobOfferCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(jocgb.fns))
	for _, fn := range jocgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*jocgb.flds)+len(jocgb.fns))
		for _, f := range *jocgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*jocgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jocgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// JobOfferCategorySelect is the builder for selecting fields of JobOfferCategory entities.
type JobOfferCategorySelect struct {
	*JobOfferCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (jocs *JobOfferCategorySelect) Aggregate(fns ...AggregateFunc) *JobOfferCategorySelect {
	jocs.fns = append(jocs.fns, fns...)
	return jocs
}

// Scan applies the selector query and scans the result into the given value.
func (jocs *JobOfferCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jocs.ctx, "Select")
	if err := jocs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobOfferCategoryQuery, *JobOfferCategorySelect](ctx, jocs.JobOfferCategoryQuery, jocs, jocs.inters, v)
}

func (jocs *JobOfferCategorySelect) sqlScan(ctx context.Context, root *JobOfferCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(jocs.fns))
	for _, fn := range jocs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*jocs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jocs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
