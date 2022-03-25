// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/auditmon/ent/event"
	"github.com/lbrictson/auditmon/ent/predicate"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetEventTime sets the "event_time" field.
func (eu *EventUpdate) SetEventTime(t time.Time) *EventUpdate {
	eu.mutation.SetEventTime(t)
	return eu
}

// SetNillableEventTime sets the "event_time" field if the given value is not nil.
func (eu *EventUpdate) SetNillableEventTime(t *time.Time) *EventUpdate {
	if t != nil {
		eu.SetEventTime(*t)
	}
	return eu
}

// SetEventName sets the "event_name" field.
func (eu *EventUpdate) SetEventName(s string) *EventUpdate {
	eu.mutation.SetEventName(s)
	return eu
}

// SetUsername sets the "username" field.
func (eu *EventUpdate) SetUsername(s string) *EventUpdate {
	eu.mutation.SetUsername(s)
	return eu
}

// SetResource sets the "resource" field.
func (eu *EventUpdate) SetResource(s string) *EventUpdate {
	eu.mutation.SetResource(s)
	return eu
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (eu *EventUpdate) SetNillableResource(s *string) *EventUpdate {
	if s != nil {
		eu.SetResource(*s)
	}
	return eu
}

// SetSourceIP sets the "source_ip" field.
func (eu *EventUpdate) SetSourceIP(s string) *EventUpdate {
	eu.mutation.SetSourceIP(s)
	return eu
}

// SetNillableSourceIP sets the "source_ip" field if the given value is not nil.
func (eu *EventUpdate) SetNillableSourceIP(s *string) *EventUpdate {
	if s != nil {
		eu.SetSourceIP(*s)
	}
	return eu
}

// SetRequestID sets the "request_id" field.
func (eu *EventUpdate) SetRequestID(s string) *EventUpdate {
	eu.mutation.SetRequestID(s)
	return eu
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (eu *EventUpdate) SetNillableRequestID(s *string) *EventUpdate {
	if s != nil {
		eu.SetRequestID(*s)
	}
	return eu
}

// SetReadOnly sets the "read_only" field.
func (eu *EventUpdate) SetReadOnly(b bool) *EventUpdate {
	eu.mutation.SetReadOnly(b)
	return eu
}

// SetNillableReadOnly sets the "read_only" field if the given value is not nil.
func (eu *EventUpdate) SetNillableReadOnly(b *bool) *EventUpdate {
	if b != nil {
		eu.SetReadOnly(*b)
	}
	return eu
}

// SetEventData sets the "event_data" field.
func (eu *EventUpdate) SetEventData(m map[string]interface{}) *EventUpdate {
	eu.mutation.SetEventData(m)
	return eu
}

// SetEventSource sets the "event_source" field.
func (eu *EventUpdate) SetEventSource(s string) *EventUpdate {
	eu.mutation.SetEventSource(s)
	return eu
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.EventTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldEventTime,
		})
	}
	if value, ok := eu.mutation.EventName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventName,
		})
	}
	if value, ok := eu.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldUsername,
		})
	}
	if value, ok := eu.mutation.Resource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldResource,
		})
	}
	if value, ok := eu.mutation.SourceIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSourceIP,
		})
	}
	if value, ok := eu.mutation.RequestID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldRequestID,
		})
	}
	if value, ok := eu.mutation.ReadOnly(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: event.FieldReadOnly,
		})
	}
	if value, ok := eu.mutation.EventData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: event.FieldEventData,
		})
	}
	if value, ok := eu.mutation.EventSource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventSource,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventMutation
}

// SetEventTime sets the "event_time" field.
func (euo *EventUpdateOne) SetEventTime(t time.Time) *EventUpdateOne {
	euo.mutation.SetEventTime(t)
	return euo
}

// SetNillableEventTime sets the "event_time" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableEventTime(t *time.Time) *EventUpdateOne {
	if t != nil {
		euo.SetEventTime(*t)
	}
	return euo
}

// SetEventName sets the "event_name" field.
func (euo *EventUpdateOne) SetEventName(s string) *EventUpdateOne {
	euo.mutation.SetEventName(s)
	return euo
}

// SetUsername sets the "username" field.
func (euo *EventUpdateOne) SetUsername(s string) *EventUpdateOne {
	euo.mutation.SetUsername(s)
	return euo
}

// SetResource sets the "resource" field.
func (euo *EventUpdateOne) SetResource(s string) *EventUpdateOne {
	euo.mutation.SetResource(s)
	return euo
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableResource(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetResource(*s)
	}
	return euo
}

// SetSourceIP sets the "source_ip" field.
func (euo *EventUpdateOne) SetSourceIP(s string) *EventUpdateOne {
	euo.mutation.SetSourceIP(s)
	return euo
}

// SetNillableSourceIP sets the "source_ip" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableSourceIP(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetSourceIP(*s)
	}
	return euo
}

// SetRequestID sets the "request_id" field.
func (euo *EventUpdateOne) SetRequestID(s string) *EventUpdateOne {
	euo.mutation.SetRequestID(s)
	return euo
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableRequestID(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetRequestID(*s)
	}
	return euo
}

// SetReadOnly sets the "read_only" field.
func (euo *EventUpdateOne) SetReadOnly(b bool) *EventUpdateOne {
	euo.mutation.SetReadOnly(b)
	return euo
}

// SetNillableReadOnly sets the "read_only" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableReadOnly(b *bool) *EventUpdateOne {
	if b != nil {
		euo.SetReadOnly(*b)
	}
	return euo
}

// SetEventData sets the "event_data" field.
func (euo *EventUpdateOne) SetEventData(m map[string]interface{}) *EventUpdateOne {
	euo.mutation.SetEventData(m)
	return euo
}

// SetEventSource sets the "event_source" field.
func (euo *EventUpdateOne) SetEventSource(s string) *EventUpdateOne {
	euo.mutation.SetEventSource(s)
	return euo
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Event.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.EventTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldEventTime,
		})
	}
	if value, ok := euo.mutation.EventName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventName,
		})
	}
	if value, ok := euo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldUsername,
		})
	}
	if value, ok := euo.mutation.Resource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldResource,
		})
	}
	if value, ok := euo.mutation.SourceIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSourceIP,
		})
	}
	if value, ok := euo.mutation.RequestID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldRequestID,
		})
	}
	if value, ok := euo.mutation.ReadOnly(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: event.FieldReadOnly,
		})
	}
	if value, ok := euo.mutation.EventData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: event.FieldEventData,
		})
	}
	if value, ok := euo.mutation.EventSource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventSource,
		})
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
