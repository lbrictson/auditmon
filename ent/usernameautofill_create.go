// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/auditmon/ent/usernameautofill"
)

// UsernameAutofillCreate is the builder for creating a UsernameAutofill entity.
type UsernameAutofillCreate struct {
	config
	mutation *UsernameAutofillMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (uac *UsernameAutofillCreate) SetUsername(s string) *UsernameAutofillCreate {
	uac.mutation.SetUsername(s)
	return uac
}

// Mutation returns the UsernameAutofillMutation object of the builder.
func (uac *UsernameAutofillCreate) Mutation() *UsernameAutofillMutation {
	return uac.mutation
}

// Save creates the UsernameAutofill in the database.
func (uac *UsernameAutofillCreate) Save(ctx context.Context) (*UsernameAutofill, error) {
	var (
		err  error
		node *UsernameAutofill
	)
	if len(uac.hooks) == 0 {
		if err = uac.check(); err != nil {
			return nil, err
		}
		node, err = uac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsernameAutofillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uac.check(); err != nil {
				return nil, err
			}
			uac.mutation = mutation
			if node, err = uac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uac.hooks) - 1; i >= 0; i-- {
			if uac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uac *UsernameAutofillCreate) SaveX(ctx context.Context) *UsernameAutofill {
	v, err := uac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uac *UsernameAutofillCreate) Exec(ctx context.Context) error {
	_, err := uac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uac *UsernameAutofillCreate) ExecX(ctx context.Context) {
	if err := uac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uac *UsernameAutofillCreate) check() error {
	if _, ok := uac.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "UsernameAutofill.username"`)}
	}
	return nil
}

func (uac *UsernameAutofillCreate) sqlSave(ctx context.Context) (*UsernameAutofill, error) {
	_node, _spec := uac.createSpec()
	if err := sqlgraph.CreateNode(ctx, uac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uac *UsernameAutofillCreate) createSpec() (*UsernameAutofill, *sqlgraph.CreateSpec) {
	var (
		_node = &UsernameAutofill{config: uac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usernameautofill.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usernameautofill.FieldID,
			},
		}
	)
	if value, ok := uac.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usernameautofill.FieldUsername,
		})
		_node.Username = value
	}
	return _node, _spec
}

// UsernameAutofillCreateBulk is the builder for creating many UsernameAutofill entities in bulk.
type UsernameAutofillCreateBulk struct {
	config
	builders []*UsernameAutofillCreate
}

// Save creates the UsernameAutofill entities in the database.
func (uacb *UsernameAutofillCreateBulk) Save(ctx context.Context) ([]*UsernameAutofill, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uacb.builders))
	nodes := make([]*UsernameAutofill, len(uacb.builders))
	mutators := make([]Mutator, len(uacb.builders))
	for i := range uacb.builders {
		func(i int, root context.Context) {
			builder := uacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsernameAutofillMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uacb *UsernameAutofillCreateBulk) SaveX(ctx context.Context) []*UsernameAutofill {
	v, err := uacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uacb *UsernameAutofillCreateBulk) Exec(ctx context.Context) error {
	_, err := uacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uacb *UsernameAutofillCreateBulk) ExecX(ctx context.Context) {
	if err := uacb.Exec(ctx); err != nil {
		panic(err)
	}
}
