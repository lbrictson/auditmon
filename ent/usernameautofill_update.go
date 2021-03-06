// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/auditmon/ent/predicate"
	"github.com/lbrictson/auditmon/ent/usernameautofill"
)

// UsernameAutofillUpdate is the builder for updating UsernameAutofill entities.
type UsernameAutofillUpdate struct {
	config
	hooks    []Hook
	mutation *UsernameAutofillMutation
}

// Where appends a list predicates to the UsernameAutofillUpdate builder.
func (uau *UsernameAutofillUpdate) Where(ps ...predicate.UsernameAutofill) *UsernameAutofillUpdate {
	uau.mutation.Where(ps...)
	return uau
}

// SetUsername sets the "username" field.
func (uau *UsernameAutofillUpdate) SetUsername(s string) *UsernameAutofillUpdate {
	uau.mutation.SetUsername(s)
	return uau
}

// Mutation returns the UsernameAutofillMutation object of the builder.
func (uau *UsernameAutofillUpdate) Mutation() *UsernameAutofillMutation {
	return uau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uau *UsernameAutofillUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uau.hooks) == 0 {
		affected, err = uau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsernameAutofillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uau.mutation = mutation
			affected, err = uau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uau.hooks) - 1; i >= 0; i-- {
			if uau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uau *UsernameAutofillUpdate) SaveX(ctx context.Context) int {
	affected, err := uau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uau *UsernameAutofillUpdate) Exec(ctx context.Context) error {
	_, err := uau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uau *UsernameAutofillUpdate) ExecX(ctx context.Context) {
	if err := uau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uau *UsernameAutofillUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usernameautofill.Table,
			Columns: usernameautofill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernameautofill.FieldID,
			},
		},
	}
	if ps := uau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uau.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usernameautofill.FieldUsername,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usernameautofill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UsernameAutofillUpdateOne is the builder for updating a single UsernameAutofill entity.
type UsernameAutofillUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsernameAutofillMutation
}

// SetUsername sets the "username" field.
func (uauo *UsernameAutofillUpdateOne) SetUsername(s string) *UsernameAutofillUpdateOne {
	uauo.mutation.SetUsername(s)
	return uauo
}

// Mutation returns the UsernameAutofillMutation object of the builder.
func (uauo *UsernameAutofillUpdateOne) Mutation() *UsernameAutofillMutation {
	return uauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uauo *UsernameAutofillUpdateOne) Select(field string, fields ...string) *UsernameAutofillUpdateOne {
	uauo.fields = append([]string{field}, fields...)
	return uauo
}

// Save executes the query and returns the updated UsernameAutofill entity.
func (uauo *UsernameAutofillUpdateOne) Save(ctx context.Context) (*UsernameAutofill, error) {
	var (
		err  error
		node *UsernameAutofill
	)
	if len(uauo.hooks) == 0 {
		node, err = uauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsernameAutofillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uauo.mutation = mutation
			node, err = uauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uauo.hooks) - 1; i >= 0; i-- {
			if uauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uauo *UsernameAutofillUpdateOne) SaveX(ctx context.Context) *UsernameAutofill {
	node, err := uauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uauo *UsernameAutofillUpdateOne) Exec(ctx context.Context) error {
	_, err := uauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uauo *UsernameAutofillUpdateOne) ExecX(ctx context.Context) {
	if err := uauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uauo *UsernameAutofillUpdateOne) sqlSave(ctx context.Context) (_node *UsernameAutofill, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usernameautofill.Table,
			Columns: usernameautofill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernameautofill.FieldID,
			},
		},
	}
	id, ok := uauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UsernameAutofill.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usernameautofill.FieldID)
		for _, f := range fields {
			if !usernameautofill.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usernameautofill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uauo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usernameautofill.FieldUsername,
		})
	}
	_node = &UsernameAutofill{config: uauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usernameautofill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
