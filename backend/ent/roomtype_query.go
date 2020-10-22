// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/PON/app/ent/predicate"
	"github.com/PON/app/ent/room"
	"github.com/PON/app/ent/roomtype"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoomTypeQuery is the builder for querying RoomType entities.
type RoomTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.RoomType
	// eager-loading edges.
	withRoomtypeRoom *RoomQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rtq *RoomTypeQuery) Where(ps ...predicate.RoomType) *RoomTypeQuery {
	rtq.predicates = append(rtq.predicates, ps...)
	return rtq
}

// Limit adds a limit step to the query.
func (rtq *RoomTypeQuery) Limit(limit int) *RoomTypeQuery {
	rtq.limit = &limit
	return rtq
}

// Offset adds an offset step to the query.
func (rtq *RoomTypeQuery) Offset(offset int) *RoomTypeQuery {
	rtq.offset = &offset
	return rtq
}

// Order adds an order step to the query.
func (rtq *RoomTypeQuery) Order(o ...OrderFunc) *RoomTypeQuery {
	rtq.order = append(rtq.order, o...)
	return rtq
}

// QueryRoomtypeRoom chains the current query on the RoomtypeRoom edge.
func (rtq *RoomTypeQuery) QueryRoomtypeRoom() *RoomQuery {
	query := &RoomQuery{config: rtq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roomtype.Table, roomtype.FieldID, rtq.sqlQuery()),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, roomtype.RoomtypeRoomTable, roomtype.RoomtypeRoomColumn),
		)
		fromU = sqlgraph.SetNeighbors(rtq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RoomType entity in the query. Returns *NotFoundError when no roomtype was found.
func (rtq *RoomTypeQuery) First(ctx context.Context) (*RoomType, error) {
	rts, err := rtq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rts) == 0 {
		return nil, &NotFoundError{roomtype.Label}
	}
	return rts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rtq *RoomTypeQuery) FirstX(ctx context.Context) *RoomType {
	rt, err := rtq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return rt
}

// FirstID returns the first RoomType id in the query. Returns *NotFoundError when no id was found.
func (rtq *RoomTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{roomtype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (rtq *RoomTypeQuery) FirstXID(ctx context.Context) int {
	id, err := rtq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only RoomType entity in the query, returns an error if not exactly one entity was returned.
func (rtq *RoomTypeQuery) Only(ctx context.Context) (*RoomType, error) {
	rts, err := rtq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(rts) {
	case 1:
		return rts[0], nil
	case 0:
		return nil, &NotFoundError{roomtype.Label}
	default:
		return nil, &NotSingularError{roomtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rtq *RoomTypeQuery) OnlyX(ctx context.Context) *RoomType {
	rt, err := rtq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return rt
}

// OnlyID returns the only RoomType id in the query, returns an error if not exactly one id was returned.
func (rtq *RoomTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rtq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{roomtype.Label}
	default:
		err = &NotSingularError{roomtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rtq *RoomTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := rtq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RoomTypes.
func (rtq *RoomTypeQuery) All(ctx context.Context) ([]*RoomType, error) {
	if err := rtq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rtq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rtq *RoomTypeQuery) AllX(ctx context.Context) []*RoomType {
	rts, err := rtq.All(ctx)
	if err != nil {
		panic(err)
	}
	return rts
}

// IDs executes the query and returns a list of RoomType ids.
func (rtq *RoomTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rtq.Select(roomtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rtq *RoomTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := rtq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rtq *RoomTypeQuery) Count(ctx context.Context) (int, error) {
	if err := rtq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rtq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rtq *RoomTypeQuery) CountX(ctx context.Context) int {
	count, err := rtq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rtq *RoomTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := rtq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rtq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rtq *RoomTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := rtq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rtq *RoomTypeQuery) Clone() *RoomTypeQuery {
	return &RoomTypeQuery{
		config:     rtq.config,
		limit:      rtq.limit,
		offset:     rtq.offset,
		order:      append([]OrderFunc{}, rtq.order...),
		unique:     append([]string{}, rtq.unique...),
		predicates: append([]predicate.RoomType{}, rtq.predicates...),
		// clone intermediate query.
		sql:  rtq.sql.Clone(),
		path: rtq.path,
	}
}

//  WithRoomtypeRoom tells the query-builder to eager-loads the nodes that are connected to
// the "RoomtypeRoom" edge. The optional arguments used to configure the query builder of the edge.
func (rtq *RoomTypeQuery) WithRoomtypeRoom(opts ...func(*RoomQuery)) *RoomTypeQuery {
	query := &RoomQuery{config: rtq.config}
	for _, opt := range opts {
		opt(query)
	}
	rtq.withRoomtypeRoom = query
	return rtq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RoomType string `json:"RoomType,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RoomType.Query().
//		GroupBy(roomtype.FieldRoomType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rtq *RoomTypeQuery) GroupBy(field string, fields ...string) *RoomTypeGroupBy {
	group := &RoomTypeGroupBy{config: rtq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rtq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		RoomType string `json:"RoomType,omitempty"`
//	}
//
//	client.RoomType.Query().
//		Select(roomtype.FieldRoomType).
//		Scan(ctx, &v)
//
func (rtq *RoomTypeQuery) Select(field string, fields ...string) *RoomTypeSelect {
	selector := &RoomTypeSelect{config: rtq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rtq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rtq.sqlQuery(), nil
	}
	return selector
}

func (rtq *RoomTypeQuery) prepareQuery(ctx context.Context) error {
	if rtq.path != nil {
		prev, err := rtq.path(ctx)
		if err != nil {
			return err
		}
		rtq.sql = prev
	}
	return nil
}

func (rtq *RoomTypeQuery) sqlAll(ctx context.Context) ([]*RoomType, error) {
	var (
		nodes       = []*RoomType{}
		_spec       = rtq.querySpec()
		loadedTypes = [1]bool{
			rtq.withRoomtypeRoom != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &RoomType{config: rtq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, rtq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rtq.withRoomtypeRoom; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*RoomType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Room(func(s *sql.Selector) {
			s.Where(sql.InValues(roomtype.RoomtypeRoomColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.room_type_roomtype_room
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "room_type_roomtype_room" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "room_type_roomtype_room" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.RoomtypeRoom = append(node.Edges.RoomtypeRoom, n)
		}
	}

	return nodes, nil
}

func (rtq *RoomTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rtq.querySpec()
	return sqlgraph.CountNodes(ctx, rtq.driver, _spec)
}

func (rtq *RoomTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rtq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rtq *RoomTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roomtype.Table,
			Columns: roomtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomtype.FieldID,
			},
		},
		From:   rtq.sql,
		Unique: true,
	}
	if ps := rtq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rtq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rtq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rtq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rtq *RoomTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rtq.driver.Dialect())
	t1 := builder.Table(roomtype.Table)
	selector := builder.Select(t1.Columns(roomtype.Columns...)...).From(t1)
	if rtq.sql != nil {
		selector = rtq.sql
		selector.Select(selector.Columns(roomtype.Columns...)...)
	}
	for _, p := range rtq.predicates {
		p(selector)
	}
	for _, p := range rtq.order {
		p(selector)
	}
	if offset := rtq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rtq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoomTypeGroupBy is the builder for group-by RoomType entities.
type RoomTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rtgb *RoomTypeGroupBy) Aggregate(fns ...AggregateFunc) *RoomTypeGroupBy {
	rtgb.fns = append(rtgb.fns, fns...)
	return rtgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rtgb *RoomTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rtgb.path(ctx)
	if err != nil {
		return err
	}
	rtgb.sql = query
	return rtgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rtgb *RoomTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rtgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rtgb *RoomTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rtgb.fields) > 1 {
		return nil, errors.New("ent: RoomTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rtgb *RoomTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := rtgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rtgb *RoomTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rtgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomtype.Label}
	default:
		err = fmt.Errorf("ent: RoomTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rtgb *RoomTypeGroupBy) StringX(ctx context.Context) string {
	v, err := rtgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rtgb *RoomTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rtgb.fields) > 1 {
		return nil, errors.New("ent: RoomTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rtgb *RoomTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := rtgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rtgb *RoomTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rtgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomtype.Label}
	default:
		err = fmt.Errorf("ent: RoomTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rtgb *RoomTypeGroupBy) IntX(ctx context.Context) int {
	v, err := rtgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rtgb *RoomTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rtgb.fields) > 1 {
		return nil, errors.New("ent: RoomTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rtgb *RoomTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rtgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rtgb *RoomTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rtgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomtype.Label}
	default:
		err = fmt.Errorf("ent: RoomTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rtgb *RoomTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rtgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rtgb *RoomTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rtgb.fields) > 1 {
		return nil, errors.New("ent: RoomTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rtgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rtgb *RoomTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rtgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rtgb *RoomTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rtgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomtype.Label}
	default:
		err = fmt.Errorf("ent: RoomTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rtgb *RoomTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := rtgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rtgb *RoomTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rtgb.sqlQuery().Query()
	if err := rtgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rtgb *RoomTypeGroupBy) sqlQuery() *sql.Selector {
	selector := rtgb.sql
	columns := make([]string, 0, len(rtgb.fields)+len(rtgb.fns))
	columns = append(columns, rtgb.fields...)
	for _, fn := range rtgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rtgb.fields...)
}

// RoomTypeSelect is the builder for select fields of RoomType entities.
type RoomTypeSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rts *RoomTypeSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rts.path(ctx)
	if err != nil {
		return err
	}
	rts.sql = query
	return rts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rts *RoomTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rts *RoomTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rts.fields) > 1 {
		return nil, errors.New("ent: RoomTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rts *RoomTypeSelect) StringsX(ctx context.Context) []string {
	v, err := rts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rts *RoomTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomtype.Label}
	default:
		err = fmt.Errorf("ent: RoomTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rts *RoomTypeSelect) StringX(ctx context.Context) string {
	v, err := rts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rts *RoomTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rts.fields) > 1 {
		return nil, errors.New("ent: RoomTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rts *RoomTypeSelect) IntsX(ctx context.Context) []int {
	v, err := rts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rts *RoomTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomtype.Label}
	default:
		err = fmt.Errorf("ent: RoomTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rts *RoomTypeSelect) IntX(ctx context.Context) int {
	v, err := rts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rts *RoomTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rts.fields) > 1 {
		return nil, errors.New("ent: RoomTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rts *RoomTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rts *RoomTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomtype.Label}
	default:
		err = fmt.Errorf("ent: RoomTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rts *RoomTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := rts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rts *RoomTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rts.fields) > 1 {
		return nil, errors.New("ent: RoomTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rts *RoomTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := rts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rts *RoomTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{roomtype.Label}
	default:
		err = fmt.Errorf("ent: RoomTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rts *RoomTypeSelect) BoolX(ctx context.Context) bool {
	v, err := rts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rts *RoomTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rts.sqlQuery().Query()
	if err := rts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rts *RoomTypeSelect) sqlQuery() sql.Querier {
	selector := rts.sql
	selector.Select(selector.Columns(rts.fields...)...)
	return selector
}
