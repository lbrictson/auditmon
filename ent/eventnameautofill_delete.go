// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/auditmon/ent/eventnameautofill"
	"github.com/lbrictson/auditmon/ent/predicate"
)

// EventNameAutofillDelete is the builder for deleting a EventNameAutofill entity.
type EventNameAutofillDelete struct {
	config
	hooks    []Hook
	mutation *EventNameAutofillMutation
}

// Where appends a list predicates to the EventNameAutofillDelete builder.
func (enad *EventNameAutofillDelete) Where(ps ...predicate.EventNameAutofill) *EventNameAutofillDelete {
	enad.mutation.Where(ps...)
	return enad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (enad *EventNameAutofillDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(enad.hooks) == 0 {
		affected, err = enad.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventNameAutofillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			enad.mutation = mutation
			affected, err = enad.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(enad.hooks) - 1; i >= 0; i-- {
			if enad.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = enad.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, enad.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (enad *EventNameAutofillDelete) ExecX(ctx context.Context) int {
	n, err := enad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (enad *EventNameAutofillDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: eventnameautofill.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: eventnameautofill.FieldID,
			},
		},
	}
	if ps := enad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, enad.driver, _spec)
}

// EventNameAutofillDeleteOne is the builder for deleting a single EventNameAutofill entity.
type EventNameAutofillDeleteOne struct {
	enad *EventNameAutofillDelete
}

// Exec executes the deletion query.
func (enado *EventNameAutofillDeleteOne) Exec(ctx context.Context) error {
	n, err := enado.enad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{eventnameautofill.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (enado *EventNameAutofillDeleteOne) ExecX(ctx context.Context) {
	enado.enad.ExecX(ctx)
}
