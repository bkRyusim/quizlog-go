// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bkRyusim/quizlog-go/ent/predicate"
	"github.com/bkRyusim/quizlog-go/ent/quiz"
	"github.com/bkRyusim/quizlog-go/ent/user"
)

// QuizQuery is the builder for querying Quiz entities.
type QuizQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Quiz
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the QuizQuery builder.
func (qq *QuizQuery) Where(ps ...predicate.Quiz) *QuizQuery {
	qq.predicates = append(qq.predicates, ps...)
	return qq
}

// Limit adds a limit step to the query.
func (qq *QuizQuery) Limit(limit int) *QuizQuery {
	qq.limit = &limit
	return qq
}

// Offset adds an offset step to the query.
func (qq *QuizQuery) Offset(offset int) *QuizQuery {
	qq.offset = &offset
	return qq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (qq *QuizQuery) Unique(unique bool) *QuizQuery {
	qq.unique = &unique
	return qq
}

// Order adds an order step to the query.
func (qq *QuizQuery) Order(o ...OrderFunc) *QuizQuery {
	qq.order = append(qq.order, o...)
	return qq
}

// QueryUser chains the current query on the "user" edge.
func (qq *QuizQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: qq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := qq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(quiz.Table, quiz.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, quiz.UserTable, quiz.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(qq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Quiz entity from the query.
// Returns a *NotFoundError when no Quiz was found.
func (qq *QuizQuery) First(ctx context.Context) (*Quiz, error) {
	nodes, err := qq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{quiz.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (qq *QuizQuery) FirstX(ctx context.Context) *Quiz {
	node, err := qq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Quiz ID from the query.
// Returns a *NotFoundError when no Quiz ID was found.
func (qq *QuizQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{quiz.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (qq *QuizQuery) FirstIDX(ctx context.Context) int {
	id, err := qq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Quiz entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Quiz entity is found.
// Returns a *NotFoundError when no Quiz entities are found.
func (qq *QuizQuery) Only(ctx context.Context) (*Quiz, error) {
	nodes, err := qq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{quiz.Label}
	default:
		return nil, &NotSingularError{quiz.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (qq *QuizQuery) OnlyX(ctx context.Context) *Quiz {
	node, err := qq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Quiz ID in the query.
// Returns a *NotSingularError when more than one Quiz ID is found.
// Returns a *NotFoundError when no entities are found.
func (qq *QuizQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = qq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{quiz.Label}
	default:
		err = &NotSingularError{quiz.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (qq *QuizQuery) OnlyIDX(ctx context.Context) int {
	id, err := qq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Quizs.
func (qq *QuizQuery) All(ctx context.Context) ([]*Quiz, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return qq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (qq *QuizQuery) AllX(ctx context.Context) []*Quiz {
	nodes, err := qq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Quiz IDs.
func (qq *QuizQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := qq.Select(quiz.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (qq *QuizQuery) IDsX(ctx context.Context) []int {
	ids, err := qq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (qq *QuizQuery) Count(ctx context.Context) (int, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return qq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (qq *QuizQuery) CountX(ctx context.Context) int {
	count, err := qq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (qq *QuizQuery) Exist(ctx context.Context) (bool, error) {
	if err := qq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return qq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (qq *QuizQuery) ExistX(ctx context.Context) bool {
	exist, err := qq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the QuizQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (qq *QuizQuery) Clone() *QuizQuery {
	if qq == nil {
		return nil
	}
	return &QuizQuery{
		config:     qq.config,
		limit:      qq.limit,
		offset:     qq.offset,
		order:      append([]OrderFunc{}, qq.order...),
		predicates: append([]predicate.Quiz{}, qq.predicates...),
		withUser:   qq.withUser.Clone(),
		// clone intermediate query.
		sql:    qq.sql.Clone(),
		path:   qq.path,
		unique: qq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (qq *QuizQuery) WithUser(opts ...func(*UserQuery)) *QuizQuery {
	query := &UserQuery{config: qq.config}
	for _, opt := range opts {
		opt(query)
	}
	qq.withUser = query
	return qq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PostUrl string `json:"postUrl,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Quiz.Query().
//		GroupBy(quiz.FieldPostUrl).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (qq *QuizQuery) GroupBy(field string, fields ...string) *QuizGroupBy {
	grbuild := &QuizGroupBy{config: qq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := qq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return qq.sqlQuery(ctx), nil
	}
	grbuild.label = quiz.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PostUrl string `json:"postUrl,omitempty"`
//	}
//
//	client.Quiz.Query().
//		Select(quiz.FieldPostUrl).
//		Scan(ctx, &v)
func (qq *QuizQuery) Select(fields ...string) *QuizSelect {
	qq.fields = append(qq.fields, fields...)
	selbuild := &QuizSelect{QuizQuery: qq}
	selbuild.label = quiz.Label
	selbuild.flds, selbuild.scan = &qq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a QuizSelect configured with the given aggregations.
func (qq *QuizQuery) Aggregate(fns ...AggregateFunc) *QuizSelect {
	return qq.Select().Aggregate(fns...)
}

func (qq *QuizQuery) prepareQuery(ctx context.Context) error {
	for _, f := range qq.fields {
		if !quiz.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if qq.path != nil {
		prev, err := qq.path(ctx)
		if err != nil {
			return err
		}
		qq.sql = prev
	}
	return nil
}

func (qq *QuizQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Quiz, error) {
	var (
		nodes       = []*Quiz{}
		withFKs     = qq.withFKs
		_spec       = qq.querySpec()
		loadedTypes = [1]bool{
			qq.withUser != nil,
		}
	)
	if qq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, quiz.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Quiz).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Quiz{config: qq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, qq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := qq.withUser; query != nil {
		if err := qq.loadUser(ctx, query, nodes, nil,
			func(n *Quiz, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (qq *QuizQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Quiz, init func(*Quiz), assign func(*Quiz, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Quiz)
	for i := range nodes {
		if nodes[i].user_quiz == nil {
			continue
		}
		fk := *nodes[i].user_quiz
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_quiz" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (qq *QuizQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := qq.querySpec()
	_spec.Node.Columns = qq.fields
	if len(qq.fields) > 0 {
		_spec.Unique = qq.unique != nil && *qq.unique
	}
	return sqlgraph.CountNodes(ctx, qq.driver, _spec)
}

func (qq *QuizQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := qq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (qq *QuizQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   quiz.Table,
			Columns: quiz.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: quiz.FieldID,
			},
		},
		From:   qq.sql,
		Unique: true,
	}
	if unique := qq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := qq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, quiz.FieldID)
		for i := range fields {
			if fields[i] != quiz.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := qq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := qq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := qq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := qq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (qq *QuizQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(qq.driver.Dialect())
	t1 := builder.Table(quiz.Table)
	columns := qq.fields
	if len(columns) == 0 {
		columns = quiz.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if qq.sql != nil {
		selector = qq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if qq.unique != nil && *qq.unique {
		selector.Distinct()
	}
	for _, p := range qq.predicates {
		p(selector)
	}
	for _, p := range qq.order {
		p(selector)
	}
	if offset := qq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := qq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// QuizGroupBy is the group-by builder for Quiz entities.
type QuizGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (qgb *QuizGroupBy) Aggregate(fns ...AggregateFunc) *QuizGroupBy {
	qgb.fns = append(qgb.fns, fns...)
	return qgb
}

// Scan applies the group-by query and scans the result into the given value.
func (qgb *QuizGroupBy) Scan(ctx context.Context, v any) error {
	query, err := qgb.path(ctx)
	if err != nil {
		return err
	}
	qgb.sql = query
	return qgb.sqlScan(ctx, v)
}

func (qgb *QuizGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range qgb.fields {
		if !quiz.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := qgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := qgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (qgb *QuizGroupBy) sqlQuery() *sql.Selector {
	selector := qgb.sql.Select()
	aggregation := make([]string, 0, len(qgb.fns))
	for _, fn := range qgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(qgb.fields)+len(qgb.fns))
		for _, f := range qgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(qgb.fields...)...)
}

// QuizSelect is the builder for selecting fields of Quiz entities.
type QuizSelect struct {
	*QuizQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (qs *QuizSelect) Aggregate(fns ...AggregateFunc) *QuizSelect {
	qs.fns = append(qs.fns, fns...)
	return qs
}

// Scan applies the selector query and scans the result into the given value.
func (qs *QuizSelect) Scan(ctx context.Context, v any) error {
	if err := qs.prepareQuery(ctx); err != nil {
		return err
	}
	qs.sql = qs.QuizQuery.sqlQuery(ctx)
	return qs.sqlScan(ctx, v)
}

func (qs *QuizSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(qs.fns))
	for _, fn := range qs.fns {
		aggregation = append(aggregation, fn(qs.sql))
	}
	switch n := len(*qs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		qs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		qs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := qs.sql.Query()
	if err := qs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
