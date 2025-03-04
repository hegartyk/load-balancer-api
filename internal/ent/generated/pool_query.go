// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/origin"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/pool"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/port"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// PoolQuery is the builder for querying Pool entities.
type PoolQuery struct {
	config
	ctx              *QueryContext
	order            []pool.OrderOption
	inters           []Interceptor
	predicates       []predicate.Pool
	withPorts        *PortQuery
	withOrigins      *OriginQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*Pool) error
	withNamedPorts   map[string]*PortQuery
	withNamedOrigins map[string]*OriginQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PoolQuery builder.
func (pq *PoolQuery) Where(ps ...predicate.Pool) *PoolQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PoolQuery) Limit(limit int) *PoolQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PoolQuery) Offset(offset int) *PoolQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PoolQuery) Unique(unique bool) *PoolQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PoolQuery) Order(o ...pool.OrderOption) *PoolQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryPorts chains the current query on the "ports" edge.
func (pq *PoolQuery) QueryPorts() *PortQuery {
	query := (&PortClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pool.Table, pool.FieldID, selector),
			sqlgraph.To(port.Table, port.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, pool.PortsTable, pool.PortsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrigins chains the current query on the "origins" edge.
func (pq *PoolQuery) QueryOrigins() *OriginQuery {
	query := (&OriginClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pool.Table, pool.FieldID, selector),
			sqlgraph.To(origin.Table, origin.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, pool.OriginsTable, pool.OriginsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Pool entity from the query.
// Returns a *NotFoundError when no Pool was found.
func (pq *PoolQuery) First(ctx context.Context) (*Pool, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pool.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PoolQuery) FirstX(ctx context.Context) *Pool {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pool ID from the query.
// Returns a *NotFoundError when no Pool ID was found.
func (pq *PoolQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pool.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PoolQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pool entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pool entity is found.
// Returns a *NotFoundError when no Pool entities are found.
func (pq *PoolQuery) Only(ctx context.Context) (*Pool, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pool.Label}
	default:
		return nil, &NotSingularError{pool.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PoolQuery) OnlyX(ctx context.Context) *Pool {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pool ID in the query.
// Returns a *NotSingularError when more than one Pool ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PoolQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pool.Label}
	default:
		err = &NotSingularError{pool.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PoolQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pools.
func (pq *PoolQuery) All(ctx context.Context) ([]*Pool, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Pool, *PoolQuery]()
	return withInterceptors[[]*Pool](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PoolQuery) AllX(ctx context.Context) []*Pool {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pool IDs.
func (pq *PoolQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(pool.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PoolQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PoolQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PoolQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PoolQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PoolQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PoolQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PoolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PoolQuery) Clone() *PoolQuery {
	if pq == nil {
		return nil
	}
	return &PoolQuery{
		config:      pq.config,
		ctx:         pq.ctx.Clone(),
		order:       append([]pool.OrderOption{}, pq.order...),
		inters:      append([]Interceptor{}, pq.inters...),
		predicates:  append([]predicate.Pool{}, pq.predicates...),
		withPorts:   pq.withPorts.Clone(),
		withOrigins: pq.withOrigins.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithPorts tells the query-builder to eager-load the nodes that are connected to
// the "ports" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PoolQuery) WithPorts(opts ...func(*PortQuery)) *PoolQuery {
	query := (&PortClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPorts = query
	return pq
}

// WithOrigins tells the query-builder to eager-load the nodes that are connected to
// the "origins" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PoolQuery) WithOrigins(opts ...func(*OriginQuery)) *PoolQuery {
	query := (&OriginClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withOrigins = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Pool.Query().
//		GroupBy(pool.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (pq *PoolQuery) GroupBy(field string, fields ...string) *PoolGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PoolGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = pool.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Pool.Query().
//		Select(pool.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *PoolQuery) Select(fields ...string) *PoolSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PoolSelect{PoolQuery: pq}
	sbuild.label = pool.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PoolSelect configured with the given aggregations.
func (pq *PoolQuery) Aggregate(fns ...AggregateFunc) *PoolSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PoolQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !pool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PoolQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Pool, error) {
	var (
		nodes       = []*Pool{}
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withPorts != nil,
			pq.withOrigins != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Pool).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Pool{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withPorts; query != nil {
		if err := pq.loadPorts(ctx, query, nodes,
			func(n *Pool) { n.Edges.Ports = []*Port{} },
			func(n *Pool, e *Port) { n.Edges.Ports = append(n.Edges.Ports, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withOrigins; query != nil {
		if err := pq.loadOrigins(ctx, query, nodes,
			func(n *Pool) { n.Edges.Origins = []*Origin{} },
			func(n *Pool, e *Origin) { n.Edges.Origins = append(n.Edges.Origins, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedPorts {
		if err := pq.loadPorts(ctx, query, nodes,
			func(n *Pool) { n.appendNamedPorts(name) },
			func(n *Pool, e *Port) { n.appendNamedPorts(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedOrigins {
		if err := pq.loadOrigins(ctx, query, nodes,
			func(n *Pool) { n.appendNamedOrigins(name) },
			func(n *Pool, e *Origin) { n.appendNamedOrigins(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PoolQuery) loadPorts(ctx context.Context, query *PortQuery, nodes []*Pool, init func(*Pool), assign func(*Pool, *Port)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[gidx.PrefixedID]*Pool)
	nids := make(map[gidx.PrefixedID]map[*Pool]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(pool.PortsTable)
		s.Join(joinT).On(s.C(port.FieldID), joinT.C(pool.PortsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(pool.PortsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(pool.PortsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(gidx.PrefixedID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*gidx.PrefixedID)
				inValue := *values[1].(*gidx.PrefixedID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Pool]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Port](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "ports" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (pq *PoolQuery) loadOrigins(ctx context.Context, query *OriginQuery, nodes []*Pool, init func(*Pool), assign func(*Pool, *Origin)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID]*Pool)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(origin.FieldPoolID)
	}
	query.Where(predicate.Origin(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(pool.OriginsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PoolID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "pool_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PoolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PoolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pool.Table, pool.Columns, sqlgraph.NewFieldSpec(pool.FieldID, field.TypeString))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pool.FieldID)
		for i := range fields {
			if fields[i] != pool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PoolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pool.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = pool.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedPorts tells the query-builder to eager-load the nodes that are connected to the "ports"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *PoolQuery) WithNamedPorts(name string, opts ...func(*PortQuery)) *PoolQuery {
	query := (&PortClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedPorts == nil {
		pq.withNamedPorts = make(map[string]*PortQuery)
	}
	pq.withNamedPorts[name] = query
	return pq
}

// WithNamedOrigins tells the query-builder to eager-load the nodes that are connected to the "origins"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *PoolQuery) WithNamedOrigins(name string, opts ...func(*OriginQuery)) *PoolQuery {
	query := (&OriginClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedOrigins == nil {
		pq.withNamedOrigins = make(map[string]*OriginQuery)
	}
	pq.withNamedOrigins[name] = query
	return pq
}

// PoolGroupBy is the group-by builder for Pool entities.
type PoolGroupBy struct {
	selector
	build *PoolQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PoolGroupBy) Aggregate(fns ...AggregateFunc) *PoolGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PoolGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PoolQuery, *PoolGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PoolGroupBy) sqlScan(ctx context.Context, root *PoolQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PoolSelect is the builder for selecting fields of Pool entities.
type PoolSelect struct {
	*PoolQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PoolSelect) Aggregate(fns ...AggregateFunc) *PoolSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PoolSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PoolQuery, *PoolSelect](ctx, ps.PoolQuery, ps, ps.inters, v)
}

func (ps *PoolSelect) sqlScan(ctx context.Context, root *PoolQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
