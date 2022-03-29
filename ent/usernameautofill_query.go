// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/auditmon/ent/predicate"
	"github.com/lbrictson/auditmon/ent/usernameautofill"
)

// UsernameAutofillQuery is the builder for querying UsernameAutofill entities.
type UsernameAutofillQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UsernameAutofill
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UsernameAutofillQuery builder.
func (uaq *UsernameAutofillQuery) Where(ps ...predicate.UsernameAutofill) *UsernameAutofillQuery {
	uaq.predicates = append(uaq.predicates, ps...)
	return uaq
}

// Limit adds a limit step to the query.
func (uaq *UsernameAutofillQuery) Limit(limit int) *UsernameAutofillQuery {
	uaq.limit = &limit
	return uaq
}

// Offset adds an offset step to the query.
func (uaq *UsernameAutofillQuery) Offset(offset int) *UsernameAutofillQuery {
	uaq.offset = &offset
	return uaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uaq *UsernameAutofillQuery) Unique(unique bool) *UsernameAutofillQuery {
	uaq.unique = &unique
	return uaq
}

// Order adds an order step to the query.
func (uaq *UsernameAutofillQuery) Order(o ...OrderFunc) *UsernameAutofillQuery {
	uaq.order = append(uaq.order, o...)
	return uaq
}

// First returns the first UsernameAutofill entity from the query.
// Returns a *NotFoundError when no UsernameAutofill was found.
func (uaq *UsernameAutofillQuery) First(ctx context.Context) (*UsernameAutofill, error) {
	nodes, err := uaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usernameautofill.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uaq *UsernameAutofillQuery) FirstX(ctx context.Context) *UsernameAutofill {
	node, err := uaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UsernameAutofill ID from the query.
// Returns a *NotFoundError when no UsernameAutofill ID was found.
func (uaq *UsernameAutofillQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usernameautofill.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uaq *UsernameAutofillQuery) FirstIDX(ctx context.Context) int {
	id, err := uaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UsernameAutofill entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UsernameAutofill entity is found.
// Returns a *NotFoundError when no UsernameAutofill entities are found.
func (uaq *UsernameAutofillQuery) Only(ctx context.Context) (*UsernameAutofill, error) {
	nodes, err := uaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usernameautofill.Label}
	default:
		return nil, &NotSingularError{usernameautofill.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uaq *UsernameAutofillQuery) OnlyX(ctx context.Context) *UsernameAutofill {
	node, err := uaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UsernameAutofill ID in the query.
// Returns a *NotSingularError when more than one UsernameAutofill ID is found.
// Returns a *NotFoundError when no entities are found.
func (uaq *UsernameAutofillQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usernameautofill.Label}
	default:
		err = &NotSingularError{usernameautofill.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uaq *UsernameAutofillQuery) OnlyIDX(ctx context.Context) int {
	id, err := uaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UsernameAutofills.
func (uaq *UsernameAutofillQuery) All(ctx context.Context) ([]*UsernameAutofill, error) {
	if err := uaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uaq *UsernameAutofillQuery) AllX(ctx context.Context) []*UsernameAutofill {
	nodes, err := uaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UsernameAutofill IDs.
func (uaq *UsernameAutofillQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uaq.Select(usernameautofill.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uaq *UsernameAutofillQuery) IDsX(ctx context.Context) []int {
	ids, err := uaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uaq *UsernameAutofillQuery) Count(ctx context.Context) (int, error) {
	if err := uaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uaq *UsernameAutofillQuery) CountX(ctx context.Context) int {
	count, err := uaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uaq *UsernameAutofillQuery) Exist(ctx context.Context) (bool, error) {
	if err := uaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uaq *UsernameAutofillQuery) ExistX(ctx context.Context) bool {
	exist, err := uaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UsernameAutofillQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uaq *UsernameAutofillQuery) Clone() *UsernameAutofillQuery {
	if uaq == nil {
		return nil
	}
	return &UsernameAutofillQuery{
		config:     uaq.config,
		limit:      uaq.limit,
		offset:     uaq.offset,
		order:      append([]OrderFunc{}, uaq.order...),
		predicates: append([]predicate.UsernameAutofill{}, uaq.predicates...),
		// clone intermediate query.
		sql:    uaq.sql.Clone(),
		path:   uaq.path,
		unique: uaq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UsernameAutofill.Query().
//		GroupBy(usernameautofill.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uaq *UsernameAutofillQuery) GroupBy(field string, fields ...string) *UsernameAutofillGroupBy {
	group := &UsernameAutofillGroupBy{config: uaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uaq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.UsernameAutofill.Query().
//		Select(usernameautofill.FieldUsername).
//		Scan(ctx, &v)
//
func (uaq *UsernameAutofillQuery) Select(fields ...string) *UsernameAutofillSelect {
	uaq.fields = append(uaq.fields, fields...)
	return &UsernameAutofillSelect{UsernameAutofillQuery: uaq}
}

func (uaq *UsernameAutofillQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uaq.fields {
		if !usernameautofill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uaq.path != nil {
		prev, err := uaq.path(ctx)
		if err != nil {
			return err
		}
		uaq.sql = prev
	}
	return nil
}

func (uaq *UsernameAutofillQuery) sqlAll(ctx context.Context) ([]*UsernameAutofill, error) {
	var (
		nodes = []*UsernameAutofill{}
		_spec = uaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UsernameAutofill{config: uaq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, uaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (uaq *UsernameAutofillQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uaq.querySpec()
	_spec.Node.Columns = uaq.fields
	if len(uaq.fields) > 0 {
		_spec.Unique = uaq.unique != nil && *uaq.unique
	}
	return sqlgraph.CountNodes(ctx, uaq.driver, _spec)
}

func (uaq *UsernameAutofillQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uaq *UsernameAutofillQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usernameautofill.Table,
			Columns: usernameautofill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernameautofill.FieldID,
			},
		},
		From:   uaq.sql,
		Unique: true,
	}
	if unique := uaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usernameautofill.FieldID)
		for i := range fields {
			if fields[i] != usernameautofill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uaq *UsernameAutofillQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uaq.driver.Dialect())
	t1 := builder.Table(usernameautofill.Table)
	columns := uaq.fields
	if len(columns) == 0 {
		columns = usernameautofill.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uaq.sql != nil {
		selector = uaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uaq.unique != nil && *uaq.unique {
		selector.Distinct()
	}
	for _, p := range uaq.predicates {
		p(selector)
	}
	for _, p := range uaq.order {
		p(selector)
	}
	if offset := uaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UsernameAutofillGroupBy is the group-by builder for UsernameAutofill entities.
type UsernameAutofillGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uagb *UsernameAutofillGroupBy) Aggregate(fns ...AggregateFunc) *UsernameAutofillGroupBy {
	uagb.fns = append(uagb.fns, fns...)
	return uagb
}

// Scan applies the group-by query and scans the result into the given value.
func (uagb *UsernameAutofillGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uagb.path(ctx)
	if err != nil {
		return err
	}
	uagb.sql = query
	return uagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uagb *UsernameAutofillGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (uagb *UsernameAutofillGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uagb.fields) > 1 {
		return nil, errors.New("ent: UsernameAutofillGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uagb *UsernameAutofillGroupBy) StringsX(ctx context.Context) []string {
	v, err := uagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uagb *UsernameAutofillGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usernameautofill.Label}
	default:
		err = fmt.Errorf("ent: UsernameAutofillGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uagb *UsernameAutofillGroupBy) StringX(ctx context.Context) string {
	v, err := uagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (uagb *UsernameAutofillGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uagb.fields) > 1 {
		return nil, errors.New("ent: UsernameAutofillGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uagb *UsernameAutofillGroupBy) IntsX(ctx context.Context) []int {
	v, err := uagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uagb *UsernameAutofillGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usernameautofill.Label}
	default:
		err = fmt.Errorf("ent: UsernameAutofillGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uagb *UsernameAutofillGroupBy) IntX(ctx context.Context) int {
	v, err := uagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (uagb *UsernameAutofillGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uagb.fields) > 1 {
		return nil, errors.New("ent: UsernameAutofillGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uagb *UsernameAutofillGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uagb *UsernameAutofillGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usernameautofill.Label}
	default:
		err = fmt.Errorf("ent: UsernameAutofillGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uagb *UsernameAutofillGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (uagb *UsernameAutofillGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uagb.fields) > 1 {
		return nil, errors.New("ent: UsernameAutofillGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uagb *UsernameAutofillGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uagb *UsernameAutofillGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usernameautofill.Label}
	default:
		err = fmt.Errorf("ent: UsernameAutofillGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uagb *UsernameAutofillGroupBy) BoolX(ctx context.Context) bool {
	v, err := uagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uagb *UsernameAutofillGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uagb.fields {
		if !usernameautofill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uagb *UsernameAutofillGroupBy) sqlQuery() *sql.Selector {
	selector := uagb.sql.Select()
	aggregation := make([]string, 0, len(uagb.fns))
	for _, fn := range uagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uagb.fields)+len(uagb.fns))
		for _, f := range uagb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uagb.fields...)...)
}

// UsernameAutofillSelect is the builder for selecting fields of UsernameAutofill entities.
type UsernameAutofillSelect struct {
	*UsernameAutofillQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uas *UsernameAutofillSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uas.prepareQuery(ctx); err != nil {
		return err
	}
	uas.sql = uas.UsernameAutofillQuery.sqlQuery(ctx)
	return uas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uas *UsernameAutofillSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (uas *UsernameAutofillSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uas.fields) > 1 {
		return nil, errors.New("ent: UsernameAutofillSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uas *UsernameAutofillSelect) StringsX(ctx context.Context) []string {
	v, err := uas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (uas *UsernameAutofillSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usernameautofill.Label}
	default:
		err = fmt.Errorf("ent: UsernameAutofillSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uas *UsernameAutofillSelect) StringX(ctx context.Context) string {
	v, err := uas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (uas *UsernameAutofillSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uas.fields) > 1 {
		return nil, errors.New("ent: UsernameAutofillSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uas *UsernameAutofillSelect) IntsX(ctx context.Context) []int {
	v, err := uas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (uas *UsernameAutofillSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usernameautofill.Label}
	default:
		err = fmt.Errorf("ent: UsernameAutofillSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uas *UsernameAutofillSelect) IntX(ctx context.Context) int {
	v, err := uas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (uas *UsernameAutofillSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uas.fields) > 1 {
		return nil, errors.New("ent: UsernameAutofillSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uas *UsernameAutofillSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (uas *UsernameAutofillSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usernameautofill.Label}
	default:
		err = fmt.Errorf("ent: UsernameAutofillSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uas *UsernameAutofillSelect) Float64X(ctx context.Context) float64 {
	v, err := uas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (uas *UsernameAutofillSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uas.fields) > 1 {
		return nil, errors.New("ent: UsernameAutofillSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uas *UsernameAutofillSelect) BoolsX(ctx context.Context) []bool {
	v, err := uas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (uas *UsernameAutofillSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{usernameautofill.Label}
	default:
		err = fmt.Errorf("ent: UsernameAutofillSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uas *UsernameAutofillSelect) BoolX(ctx context.Context) bool {
	v, err := uas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uas *UsernameAutofillSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uas.sql.Query()
	if err := uas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
