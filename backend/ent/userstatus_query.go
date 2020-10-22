// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/PON/app/ent/predicate"
	"github.com/PON/app/ent/user"
	"github.com/PON/app/ent/userstatus"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserStatusQuery is the builder for querying UserStatus entities.
type UserStatusQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.UserStatus
	// eager-loading edges.
	withUserstatusUser *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (usq *UserStatusQuery) Where(ps ...predicate.UserStatus) *UserStatusQuery {
	usq.predicates = append(usq.predicates, ps...)
	return usq
}

// Limit adds a limit step to the query.
func (usq *UserStatusQuery) Limit(limit int) *UserStatusQuery {
	usq.limit = &limit
	return usq
}

// Offset adds an offset step to the query.
func (usq *UserStatusQuery) Offset(offset int) *UserStatusQuery {
	usq.offset = &offset
	return usq
}

// Order adds an order step to the query.
func (usq *UserStatusQuery) Order(o ...OrderFunc) *UserStatusQuery {
	usq.order = append(usq.order, o...)
	return usq
}

// QueryUserstatusUser chains the current query on the UserstatusUser edge.
func (usq *UserStatusQuery) QueryUserstatusUser() *UserQuery {
	query := &UserQuery{config: usq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userstatus.Table, userstatus.FieldID, usq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, userstatus.UserstatusUserTable, userstatus.UserstatusUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(usq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserStatus entity in the query. Returns *NotFoundError when no userstatus was found.
func (usq *UserStatusQuery) First(ctx context.Context) (*UserStatus, error) {
	usSlice, err := usq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(usSlice) == 0 {
		return nil, &NotFoundError{userstatus.Label}
	}
	return usSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (usq *UserStatusQuery) FirstX(ctx context.Context) *UserStatus {
	us, err := usq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return us
}

// FirstID returns the first UserStatus id in the query. Returns *NotFoundError when no id was found.
func (usq *UserStatusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = usq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (usq *UserStatusQuery) FirstXID(ctx context.Context) int {
	id, err := usq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only UserStatus entity in the query, returns an error if not exactly one entity was returned.
func (usq *UserStatusQuery) Only(ctx context.Context) (*UserStatus, error) {
	usSlice, err := usq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(usSlice) {
	case 1:
		return usSlice[0], nil
	case 0:
		return nil, &NotFoundError{userstatus.Label}
	default:
		return nil, &NotSingularError{userstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (usq *UserStatusQuery) OnlyX(ctx context.Context) *UserStatus {
	us, err := usq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return us
}

// OnlyID returns the only UserStatus id in the query, returns an error if not exactly one id was returned.
func (usq *UserStatusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = usq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = &NotSingularError{userstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (usq *UserStatusQuery) OnlyIDX(ctx context.Context) int {
	id, err := usq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserStatusSlice.
func (usq *UserStatusQuery) All(ctx context.Context) ([]*UserStatus, error) {
	if err := usq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return usq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (usq *UserStatusQuery) AllX(ctx context.Context) []*UserStatus {
	usSlice, err := usq.All(ctx)
	if err != nil {
		panic(err)
	}
	return usSlice
}

// IDs executes the query and returns a list of UserStatus ids.
func (usq *UserStatusQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := usq.Select(userstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (usq *UserStatusQuery) IDsX(ctx context.Context) []int {
	ids, err := usq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (usq *UserStatusQuery) Count(ctx context.Context) (int, error) {
	if err := usq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return usq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (usq *UserStatusQuery) CountX(ctx context.Context) int {
	count, err := usq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (usq *UserStatusQuery) Exist(ctx context.Context) (bool, error) {
	if err := usq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return usq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (usq *UserStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := usq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (usq *UserStatusQuery) Clone() *UserStatusQuery {
	return &UserStatusQuery{
		config:     usq.config,
		limit:      usq.limit,
		offset:     usq.offset,
		order:      append([]OrderFunc{}, usq.order...),
		unique:     append([]string{}, usq.unique...),
		predicates: append([]predicate.UserStatus{}, usq.predicates...),
		// clone intermediate query.
		sql:  usq.sql.Clone(),
		path: usq.path,
	}
}

//  WithUserstatusUser tells the query-builder to eager-loads the nodes that are connected to
// the "UserstatusUser" edge. The optional arguments used to configure the query builder of the edge.
func (usq *UserStatusQuery) WithUserstatusUser(opts ...func(*UserQuery)) *UserStatusQuery {
	query := &UserQuery{config: usq.config}
	for _, opt := range opts {
		opt(query)
	}
	usq.withUserstatusUser = query
	return usq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserStatus string `json:"UserStatus,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserStatus.Query().
//		GroupBy(userstatus.FieldUserStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (usq *UserStatusQuery) GroupBy(field string, fields ...string) *UserStatusGroupBy {
	group := &UserStatusGroupBy{config: usq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return usq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		UserStatus string `json:"UserStatus,omitempty"`
//	}
//
//	client.UserStatus.Query().
//		Select(userstatus.FieldUserStatus).
//		Scan(ctx, &v)
//
func (usq *UserStatusQuery) Select(field string, fields ...string) *UserStatusSelect {
	selector := &UserStatusSelect{config: usq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return usq.sqlQuery(), nil
	}
	return selector
}

func (usq *UserStatusQuery) prepareQuery(ctx context.Context) error {
	if usq.path != nil {
		prev, err := usq.path(ctx)
		if err != nil {
			return err
		}
		usq.sql = prev
	}
	return nil
}

func (usq *UserStatusQuery) sqlAll(ctx context.Context) ([]*UserStatus, error) {
	var (
		nodes       = []*UserStatus{}
		_spec       = usq.querySpec()
		loadedTypes = [1]bool{
			usq.withUserstatusUser != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &UserStatus{config: usq.config}
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
	if err := sqlgraph.QueryNodes(ctx, usq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := usq.withUserstatusUser; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*UserStatus)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(userstatus.UserstatusUserColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.user_status_userstatus_user
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "user_status_userstatus_user" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_status_userstatus_user" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.UserstatusUser = append(node.Edges.UserstatusUser, n)
		}
	}

	return nodes, nil
}

func (usq *UserStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := usq.querySpec()
	return sqlgraph.CountNodes(ctx, usq.driver, _spec)
}

func (usq *UserStatusQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := usq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (usq *UserStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userstatus.Table,
			Columns: userstatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userstatus.FieldID,
			},
		},
		From:   usq.sql,
		Unique: true,
	}
	if ps := usq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := usq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := usq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := usq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (usq *UserStatusQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(usq.driver.Dialect())
	t1 := builder.Table(userstatus.Table)
	selector := builder.Select(t1.Columns(userstatus.Columns...)...).From(t1)
	if usq.sql != nil {
		selector = usq.sql
		selector.Select(selector.Columns(userstatus.Columns...)...)
	}
	for _, p := range usq.predicates {
		p(selector)
	}
	for _, p := range usq.order {
		p(selector)
	}
	if offset := usq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := usq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserStatusGroupBy is the builder for group-by UserStatus entities.
type UserStatusGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (usgb *UserStatusGroupBy) Aggregate(fns ...AggregateFunc) *UserStatusGroupBy {
	usgb.fns = append(usgb.fns, fns...)
	return usgb
}

// Scan applies the group-by query and scan the result into the given value.
func (usgb *UserStatusGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := usgb.path(ctx)
	if err != nil {
		return err
	}
	usgb.sql = query
	return usgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (usgb *UserStatusGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := usgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserStatusGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(usgb.fields) > 1 {
		return nil, errors.New("ent: UserStatusGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := usgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (usgb *UserStatusGroupBy) StringsX(ctx context.Context) []string {
	v, err := usgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserStatusGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = usgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserStatusGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (usgb *UserStatusGroupBy) StringX(ctx context.Context) string {
	v, err := usgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserStatusGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(usgb.fields) > 1 {
		return nil, errors.New("ent: UserStatusGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := usgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (usgb *UserStatusGroupBy) IntsX(ctx context.Context) []int {
	v, err := usgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserStatusGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = usgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserStatusGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (usgb *UserStatusGroupBy) IntX(ctx context.Context) int {
	v, err := usgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserStatusGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(usgb.fields) > 1 {
		return nil, errors.New("ent: UserStatusGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := usgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (usgb *UserStatusGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := usgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserStatusGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = usgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserStatusGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (usgb *UserStatusGroupBy) Float64X(ctx context.Context) float64 {
	v, err := usgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserStatusGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(usgb.fields) > 1 {
		return nil, errors.New("ent: UserStatusGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := usgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (usgb *UserStatusGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := usgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (usgb *UserStatusGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = usgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserStatusGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (usgb *UserStatusGroupBy) BoolX(ctx context.Context) bool {
	v, err := usgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (usgb *UserStatusGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := usgb.sqlQuery().Query()
	if err := usgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (usgb *UserStatusGroupBy) sqlQuery() *sql.Selector {
	selector := usgb.sql
	columns := make([]string, 0, len(usgb.fields)+len(usgb.fns))
	columns = append(columns, usgb.fields...)
	for _, fn := range usgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(usgb.fields...)
}

// UserStatusSelect is the builder for select fields of UserStatus entities.
type UserStatusSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (uss *UserStatusSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := uss.path(ctx)
	if err != nil {
		return err
	}
	uss.sql = query
	return uss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uss *UserStatusSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (uss *UserStatusSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uss.fields) > 1 {
		return nil, errors.New("ent: UserStatusSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uss *UserStatusSelect) StringsX(ctx context.Context) []string {
	v, err := uss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (uss *UserStatusSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserStatusSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uss *UserStatusSelect) StringX(ctx context.Context) string {
	v, err := uss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (uss *UserStatusSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uss.fields) > 1 {
		return nil, errors.New("ent: UserStatusSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uss *UserStatusSelect) IntsX(ctx context.Context) []int {
	v, err := uss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (uss *UserStatusSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserStatusSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uss *UserStatusSelect) IntX(ctx context.Context) int {
	v, err := uss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (uss *UserStatusSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uss.fields) > 1 {
		return nil, errors.New("ent: UserStatusSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uss *UserStatusSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (uss *UserStatusSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserStatusSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uss *UserStatusSelect) Float64X(ctx context.Context) float64 {
	v, err := uss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (uss *UserStatusSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uss.fields) > 1 {
		return nil, errors.New("ent: UserStatusSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uss *UserStatusSelect) BoolsX(ctx context.Context) []bool {
	v, err := uss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (uss *UserStatusSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userstatus.Label}
	default:
		err = fmt.Errorf("ent: UserStatusSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uss *UserStatusSelect) BoolX(ctx context.Context) bool {
	v, err := uss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uss *UserStatusSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uss.sqlQuery().Query()
	if err := uss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uss *UserStatusSelect) sqlQuery() sql.Querier {
	selector := uss.sql
	selector.Select(selector.Columns(uss.fields...)...)
	return selector
}