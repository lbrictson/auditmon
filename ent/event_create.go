// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lbrictson/auditmon/ent/event"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
}

// SetEventTime sets the "event_time" field.
func (ec *EventCreate) SetEventTime(t time.Time) *EventCreate {
	ec.mutation.SetEventTime(t)
	return ec
}

// SetNillableEventTime sets the "event_time" field if the given value is not nil.
func (ec *EventCreate) SetNillableEventTime(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetEventTime(*t)
	}
	return ec
}

// SetEventName sets the "event_name" field.
func (ec *EventCreate) SetEventName(s string) *EventCreate {
	ec.mutation.SetEventName(s)
	return ec
}

// SetUsername sets the "username" field.
func (ec *EventCreate) SetUsername(s string) *EventCreate {
	ec.mutation.SetUsername(s)
	return ec
}

// SetResource sets the "resource" field.
func (ec *EventCreate) SetResource(s string) *EventCreate {
	ec.mutation.SetResource(s)
	return ec
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (ec *EventCreate) SetNillableResource(s *string) *EventCreate {
	if s != nil {
		ec.SetResource(*s)
	}
	return ec
}

// SetSourceIP sets the "source_ip" field.
func (ec *EventCreate) SetSourceIP(s string) *EventCreate {
	ec.mutation.SetSourceIP(s)
	return ec
}

// SetNillableSourceIP sets the "source_ip" field if the given value is not nil.
func (ec *EventCreate) SetNillableSourceIP(s *string) *EventCreate {
	if s != nil {
		ec.SetSourceIP(*s)
	}
	return ec
}

// SetRequestID sets the "request_id" field.
func (ec *EventCreate) SetRequestID(s string) *EventCreate {
	ec.mutation.SetRequestID(s)
	return ec
}

// SetNillableRequestID sets the "request_id" field if the given value is not nil.
func (ec *EventCreate) SetNillableRequestID(s *string) *EventCreate {
	if s != nil {
		ec.SetRequestID(*s)
	}
	return ec
}

// SetReadOnly sets the "read_only" field.
func (ec *EventCreate) SetReadOnly(b bool) *EventCreate {
	ec.mutation.SetReadOnly(b)
	return ec
}

// SetNillableReadOnly sets the "read_only" field if the given value is not nil.
func (ec *EventCreate) SetNillableReadOnly(b *bool) *EventCreate {
	if b != nil {
		ec.SetReadOnly(*b)
	}
	return ec
}

// SetEventData sets the "event_data" field.
func (ec *EventCreate) SetEventData(m map[string]interface{}) *EventCreate {
	ec.mutation.SetEventData(m)
	return ec
}

// SetID sets the "id" field.
func (ec *EventCreate) SetID(u uuid.UUID) *EventCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EventCreate) SetNillableID(u *uuid.UUID) *EventCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventCreate) defaults() {
	if _, ok := ec.mutation.EventTime(); !ok {
		v := event.DefaultEventTime
		ec.mutation.SetEventTime(v)
	}
	if _, ok := ec.mutation.Resource(); !ok {
		v := event.DefaultResource
		ec.mutation.SetResource(v)
	}
	if _, ok := ec.mutation.SourceIP(); !ok {
		v := event.DefaultSourceIP
		ec.mutation.SetSourceIP(v)
	}
	if _, ok := ec.mutation.RequestID(); !ok {
		v := event.DefaultRequestID
		ec.mutation.SetRequestID(v)
	}
	if _, ok := ec.mutation.ReadOnly(); !ok {
		v := event.DefaultReadOnly
		ec.mutation.SetReadOnly(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := event.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if _, ok := ec.mutation.EventTime(); !ok {
		return &ValidationError{Name: "event_time", err: errors.New(`ent: missing required field "Event.event_time"`)}
	}
	if _, ok := ec.mutation.EventName(); !ok {
		return &ValidationError{Name: "event_name", err: errors.New(`ent: missing required field "Event.event_name"`)}
	}
	if _, ok := ec.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Event.username"`)}
	}
	if _, ok := ec.mutation.Resource(); !ok {
		return &ValidationError{Name: "resource", err: errors.New(`ent: missing required field "Event.resource"`)}
	}
	if _, ok := ec.mutation.SourceIP(); !ok {
		return &ValidationError{Name: "source_ip", err: errors.New(`ent: missing required field "Event.source_ip"`)}
	}
	if _, ok := ec.mutation.RequestID(); !ok {
		return &ValidationError{Name: "request_id", err: errors.New(`ent: missing required field "Event.request_id"`)}
	}
	if _, ok := ec.mutation.ReadOnly(); !ok {
		return &ValidationError{Name: "read_only", err: errors.New(`ent: missing required field "Event.read_only"`)}
	}
	if _, ok := ec.mutation.EventData(); !ok {
		return &ValidationError{Name: "event_data", err: errors.New(`ent: missing required field "Event.event_data"`)}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: event.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		}
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.EventTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldEventTime,
		})
		_node.EventTime = value
	}
	if value, ok := ec.mutation.EventName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldEventName,
		})
		_node.EventName = value
	}
	if value, ok := ec.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := ec.mutation.Resource(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldResource,
		})
		_node.Resource = value
	}
	if value, ok := ec.mutation.SourceIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldSourceIP,
		})
		_node.SourceIP = value
	}
	if value, ok := ec.mutation.RequestID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldRequestID,
		})
		_node.RequestID = value
	}
	if value, ok := ec.mutation.ReadOnly(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: event.FieldReadOnly,
		})
		_node.ReadOnly = value
	}
	if value, ok := ec.mutation.EventData(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: event.FieldEventData,
		})
		_node.EventData = value
	}
	return _node, _spec
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	builders []*EventCreate
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
