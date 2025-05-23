// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kratos-admin/app/admin/service/internal/data/ent/adminloginlog"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdminLoginLogQuery is the builder for querying AdminLoginLog entities.
type AdminLoginLogQuery struct {
	config
	ctx        *QueryContext
	order      []adminloginlog.OrderOption
	inters     []Interceptor
	predicates []predicate.AdminLoginLog
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdminLoginLogQuery builder.
func (allq *AdminLoginLogQuery) Where(ps ...predicate.AdminLoginLog) *AdminLoginLogQuery {
	allq.predicates = append(allq.predicates, ps...)
	return allq
}

// Limit the number of records to be returned by this query.
func (allq *AdminLoginLogQuery) Limit(limit int) *AdminLoginLogQuery {
	allq.ctx.Limit = &limit
	return allq
}

// Offset to start from.
func (allq *AdminLoginLogQuery) Offset(offset int) *AdminLoginLogQuery {
	allq.ctx.Offset = &offset
	return allq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (allq *AdminLoginLogQuery) Unique(unique bool) *AdminLoginLogQuery {
	allq.ctx.Unique = &unique
	return allq
}

// Order specifies how the records should be ordered.
func (allq *AdminLoginLogQuery) Order(o ...adminloginlog.OrderOption) *AdminLoginLogQuery {
	allq.order = append(allq.order, o...)
	return allq
}

// First returns the first AdminLoginLog entity from the query.
// Returns a *NotFoundError when no AdminLoginLog was found.
func (allq *AdminLoginLogQuery) First(ctx context.Context) (*AdminLoginLog, error) {
	nodes, err := allq.Limit(1).All(setContextOp(ctx, allq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{adminloginlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (allq *AdminLoginLogQuery) FirstX(ctx context.Context) *AdminLoginLog {
	node, err := allq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdminLoginLog ID from the query.
// Returns a *NotFoundError when no AdminLoginLog ID was found.
func (allq *AdminLoginLogQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = allq.Limit(1).IDs(setContextOp(ctx, allq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{adminloginlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (allq *AdminLoginLogQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := allq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdminLoginLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AdminLoginLog entity is found.
// Returns a *NotFoundError when no AdminLoginLog entities are found.
func (allq *AdminLoginLogQuery) Only(ctx context.Context) (*AdminLoginLog, error) {
	nodes, err := allq.Limit(2).All(setContextOp(ctx, allq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{adminloginlog.Label}
	default:
		return nil, &NotSingularError{adminloginlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (allq *AdminLoginLogQuery) OnlyX(ctx context.Context) *AdminLoginLog {
	node, err := allq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdminLoginLog ID in the query.
// Returns a *NotSingularError when more than one AdminLoginLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (allq *AdminLoginLogQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = allq.Limit(2).IDs(setContextOp(ctx, allq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{adminloginlog.Label}
	default:
		err = &NotSingularError{adminloginlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (allq *AdminLoginLogQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := allq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdminLoginLogs.
func (allq *AdminLoginLogQuery) All(ctx context.Context) ([]*AdminLoginLog, error) {
	ctx = setContextOp(ctx, allq.ctx, ent.OpQueryAll)
	if err := allq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AdminLoginLog, *AdminLoginLogQuery]()
	return withInterceptors[[]*AdminLoginLog](ctx, allq, qr, allq.inters)
}

// AllX is like All, but panics if an error occurs.
func (allq *AdminLoginLogQuery) AllX(ctx context.Context) []*AdminLoginLog {
	nodes, err := allq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdminLoginLog IDs.
func (allq *AdminLoginLogQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if allq.ctx.Unique == nil && allq.path != nil {
		allq.Unique(true)
	}
	ctx = setContextOp(ctx, allq.ctx, ent.OpQueryIDs)
	if err = allq.Select(adminloginlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (allq *AdminLoginLogQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := allq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (allq *AdminLoginLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, allq.ctx, ent.OpQueryCount)
	if err := allq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, allq, querierCount[*AdminLoginLogQuery](), allq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (allq *AdminLoginLogQuery) CountX(ctx context.Context) int {
	count, err := allq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (allq *AdminLoginLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, allq.ctx, ent.OpQueryExist)
	switch _, err := allq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (allq *AdminLoginLogQuery) ExistX(ctx context.Context) bool {
	exist, err := allq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdminLoginLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (allq *AdminLoginLogQuery) Clone() *AdminLoginLogQuery {
	if allq == nil {
		return nil
	}
	return &AdminLoginLogQuery{
		config:     allq.config,
		ctx:        allq.ctx.Clone(),
		order:      append([]adminloginlog.OrderOption{}, allq.order...),
		inters:     append([]Interceptor{}, allq.inters...),
		predicates: append([]predicate.AdminLoginLog{}, allq.predicates...),
		// clone intermediate query.
		sql:       allq.sql.Clone(),
		path:      allq.path,
		modifiers: append([]func(*sql.Selector){}, allq.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdminLoginLog.Query().
//		GroupBy(adminloginlog.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (allq *AdminLoginLogQuery) GroupBy(field string, fields ...string) *AdminLoginLogGroupBy {
	allq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AdminLoginLogGroupBy{build: allq}
	grbuild.flds = &allq.ctx.Fields
	grbuild.label = adminloginlog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.AdminLoginLog.Query().
//		Select(adminloginlog.FieldCreateTime).
//		Scan(ctx, &v)
func (allq *AdminLoginLogQuery) Select(fields ...string) *AdminLoginLogSelect {
	allq.ctx.Fields = append(allq.ctx.Fields, fields...)
	sbuild := &AdminLoginLogSelect{AdminLoginLogQuery: allq}
	sbuild.label = adminloginlog.Label
	sbuild.flds, sbuild.scan = &allq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AdminLoginLogSelect configured with the given aggregations.
func (allq *AdminLoginLogQuery) Aggregate(fns ...AggregateFunc) *AdminLoginLogSelect {
	return allq.Select().Aggregate(fns...)
}

func (allq *AdminLoginLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range allq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, allq); err != nil {
				return err
			}
		}
	}
	for _, f := range allq.ctx.Fields {
		if !adminloginlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if allq.path != nil {
		prev, err := allq.path(ctx)
		if err != nil {
			return err
		}
		allq.sql = prev
	}
	return nil
}

func (allq *AdminLoginLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AdminLoginLog, error) {
	var (
		nodes = []*AdminLoginLog{}
		_spec = allq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AdminLoginLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AdminLoginLog{config: allq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(allq.modifiers) > 0 {
		_spec.Modifiers = allq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, allq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (allq *AdminLoginLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := allq.querySpec()
	if len(allq.modifiers) > 0 {
		_spec.Modifiers = allq.modifiers
	}
	_spec.Node.Columns = allq.ctx.Fields
	if len(allq.ctx.Fields) > 0 {
		_spec.Unique = allq.ctx.Unique != nil && *allq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, allq.driver, _spec)
}

func (allq *AdminLoginLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(adminloginlog.Table, adminloginlog.Columns, sqlgraph.NewFieldSpec(adminloginlog.FieldID, field.TypeUint32))
	_spec.From = allq.sql
	if unique := allq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if allq.path != nil {
		_spec.Unique = true
	}
	if fields := allq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adminloginlog.FieldID)
		for i := range fields {
			if fields[i] != adminloginlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := allq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := allq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := allq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := allq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (allq *AdminLoginLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(allq.driver.Dialect())
	t1 := builder.Table(adminloginlog.Table)
	columns := allq.ctx.Fields
	if len(columns) == 0 {
		columns = adminloginlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if allq.sql != nil {
		selector = allq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if allq.ctx.Unique != nil && *allq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range allq.modifiers {
		m(selector)
	}
	for _, p := range allq.predicates {
		p(selector)
	}
	for _, p := range allq.order {
		p(selector)
	}
	if offset := allq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := allq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (allq *AdminLoginLogQuery) ForUpdate(opts ...sql.LockOption) *AdminLoginLogQuery {
	if allq.driver.Dialect() == dialect.Postgres {
		allq.Unique(false)
	}
	allq.modifiers = append(allq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return allq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (allq *AdminLoginLogQuery) ForShare(opts ...sql.LockOption) *AdminLoginLogQuery {
	if allq.driver.Dialect() == dialect.Postgres {
		allq.Unique(false)
	}
	allq.modifiers = append(allq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return allq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (allq *AdminLoginLogQuery) Modify(modifiers ...func(s *sql.Selector)) *AdminLoginLogSelect {
	allq.modifiers = append(allq.modifiers, modifiers...)
	return allq.Select()
}

// AdminLoginLogGroupBy is the group-by builder for AdminLoginLog entities.
type AdminLoginLogGroupBy struct {
	selector
	build *AdminLoginLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (allgb *AdminLoginLogGroupBy) Aggregate(fns ...AggregateFunc) *AdminLoginLogGroupBy {
	allgb.fns = append(allgb.fns, fns...)
	return allgb
}

// Scan applies the selector query and scans the result into the given value.
func (allgb *AdminLoginLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, allgb.build.ctx, ent.OpQueryGroupBy)
	if err := allgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminLoginLogQuery, *AdminLoginLogGroupBy](ctx, allgb.build, allgb, allgb.build.inters, v)
}

func (allgb *AdminLoginLogGroupBy) sqlScan(ctx context.Context, root *AdminLoginLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(allgb.fns))
	for _, fn := range allgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*allgb.flds)+len(allgb.fns))
		for _, f := range *allgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*allgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := allgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AdminLoginLogSelect is the builder for selecting fields of AdminLoginLog entities.
type AdminLoginLogSelect struct {
	*AdminLoginLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (alls *AdminLoginLogSelect) Aggregate(fns ...AggregateFunc) *AdminLoginLogSelect {
	alls.fns = append(alls.fns, fns...)
	return alls
}

// Scan applies the selector query and scans the result into the given value.
func (alls *AdminLoginLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, alls.ctx, ent.OpQuerySelect)
	if err := alls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdminLoginLogQuery, *AdminLoginLogSelect](ctx, alls.AdminLoginLogQuery, alls, alls.inters, v)
}

func (alls *AdminLoginLogSelect) sqlScan(ctx context.Context, root *AdminLoginLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(alls.fns))
	for _, fn := range alls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*alls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := alls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (alls *AdminLoginLogSelect) Modify(modifiers ...func(s *sql.Selector)) *AdminLoginLogSelect {
	alls.modifiers = append(alls.modifiers, modifiers...)
	return alls
}
