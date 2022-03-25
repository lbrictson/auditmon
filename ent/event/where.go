// Code generated by entc, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/lbrictson/auditmon/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EventTime applies equality check predicate on the "event_time" field. It's identical to EventTimeEQ.
func EventTime(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventTime), v))
	})
}

// EventName applies equality check predicate on the "event_name" field. It's identical to EventNameEQ.
func EventName(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventName), v))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// Resource applies equality check predicate on the "resource" field. It's identical to ResourceEQ.
func Resource(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResource), v))
	})
}

// SourceIP applies equality check predicate on the "source_ip" field. It's identical to SourceIPEQ.
func SourceIP(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSourceIP), v))
	})
}

// RequestID applies equality check predicate on the "request_id" field. It's identical to RequestIDEQ.
func RequestID(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRequestID), v))
	})
}

// ReadOnly applies equality check predicate on the "read_only" field. It's identical to ReadOnlyEQ.
func ReadOnly(v bool) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadOnly), v))
	})
}

// EventTimeEQ applies the EQ predicate on the "event_time" field.
func EventTimeEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventTime), v))
	})
}

// EventTimeNEQ applies the NEQ predicate on the "event_time" field.
func EventTimeNEQ(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventTime), v))
	})
}

// EventTimeIn applies the In predicate on the "event_time" field.
func EventTimeIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventTime), v...))
	})
}

// EventTimeNotIn applies the NotIn predicate on the "event_time" field.
func EventTimeNotIn(vs ...time.Time) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventTime), v...))
	})
}

// EventTimeGT applies the GT predicate on the "event_time" field.
func EventTimeGT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventTime), v))
	})
}

// EventTimeGTE applies the GTE predicate on the "event_time" field.
func EventTimeGTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventTime), v))
	})
}

// EventTimeLT applies the LT predicate on the "event_time" field.
func EventTimeLT(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventTime), v))
	})
}

// EventTimeLTE applies the LTE predicate on the "event_time" field.
func EventTimeLTE(v time.Time) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventTime), v))
	})
}

// EventNameEQ applies the EQ predicate on the "event_name" field.
func EventNameEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventName), v))
	})
}

// EventNameNEQ applies the NEQ predicate on the "event_name" field.
func EventNameNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventName), v))
	})
}

// EventNameIn applies the In predicate on the "event_name" field.
func EventNameIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventName), v...))
	})
}

// EventNameNotIn applies the NotIn predicate on the "event_name" field.
func EventNameNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventName), v...))
	})
}

// EventNameGT applies the GT predicate on the "event_name" field.
func EventNameGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventName), v))
	})
}

// EventNameGTE applies the GTE predicate on the "event_name" field.
func EventNameGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventName), v))
	})
}

// EventNameLT applies the LT predicate on the "event_name" field.
func EventNameLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventName), v))
	})
}

// EventNameLTE applies the LTE predicate on the "event_name" field.
func EventNameLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventName), v))
	})
}

// EventNameContains applies the Contains predicate on the "event_name" field.
func EventNameContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventName), v))
	})
}

// EventNameHasPrefix applies the HasPrefix predicate on the "event_name" field.
func EventNameHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventName), v))
	})
}

// EventNameHasSuffix applies the HasSuffix predicate on the "event_name" field.
func EventNameHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventName), v))
	})
}

// EventNameEqualFold applies the EqualFold predicate on the "event_name" field.
func EventNameEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventName), v))
	})
}

// EventNameContainsFold applies the ContainsFold predicate on the "event_name" field.
func EventNameContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventName), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// ResourceEQ applies the EQ predicate on the "resource" field.
func ResourceEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResource), v))
	})
}

// ResourceNEQ applies the NEQ predicate on the "resource" field.
func ResourceNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResource), v))
	})
}

// ResourceIn applies the In predicate on the "resource" field.
func ResourceIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldResource), v...))
	})
}

// ResourceNotIn applies the NotIn predicate on the "resource" field.
func ResourceNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldResource), v...))
	})
}

// ResourceGT applies the GT predicate on the "resource" field.
func ResourceGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResource), v))
	})
}

// ResourceGTE applies the GTE predicate on the "resource" field.
func ResourceGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResource), v))
	})
}

// ResourceLT applies the LT predicate on the "resource" field.
func ResourceLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResource), v))
	})
}

// ResourceLTE applies the LTE predicate on the "resource" field.
func ResourceLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResource), v))
	})
}

// ResourceContains applies the Contains predicate on the "resource" field.
func ResourceContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResource), v))
	})
}

// ResourceHasPrefix applies the HasPrefix predicate on the "resource" field.
func ResourceHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResource), v))
	})
}

// ResourceHasSuffix applies the HasSuffix predicate on the "resource" field.
func ResourceHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResource), v))
	})
}

// ResourceEqualFold applies the EqualFold predicate on the "resource" field.
func ResourceEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResource), v))
	})
}

// ResourceContainsFold applies the ContainsFold predicate on the "resource" field.
func ResourceContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResource), v))
	})
}

// SourceIPEQ applies the EQ predicate on the "source_ip" field.
func SourceIPEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSourceIP), v))
	})
}

// SourceIPNEQ applies the NEQ predicate on the "source_ip" field.
func SourceIPNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSourceIP), v))
	})
}

// SourceIPIn applies the In predicate on the "source_ip" field.
func SourceIPIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSourceIP), v...))
	})
}

// SourceIPNotIn applies the NotIn predicate on the "source_ip" field.
func SourceIPNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSourceIP), v...))
	})
}

// SourceIPGT applies the GT predicate on the "source_ip" field.
func SourceIPGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSourceIP), v))
	})
}

// SourceIPGTE applies the GTE predicate on the "source_ip" field.
func SourceIPGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSourceIP), v))
	})
}

// SourceIPLT applies the LT predicate on the "source_ip" field.
func SourceIPLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSourceIP), v))
	})
}

// SourceIPLTE applies the LTE predicate on the "source_ip" field.
func SourceIPLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSourceIP), v))
	})
}

// SourceIPContains applies the Contains predicate on the "source_ip" field.
func SourceIPContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSourceIP), v))
	})
}

// SourceIPHasPrefix applies the HasPrefix predicate on the "source_ip" field.
func SourceIPHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSourceIP), v))
	})
}

// SourceIPHasSuffix applies the HasSuffix predicate on the "source_ip" field.
func SourceIPHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSourceIP), v))
	})
}

// SourceIPEqualFold applies the EqualFold predicate on the "source_ip" field.
func SourceIPEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSourceIP), v))
	})
}

// SourceIPContainsFold applies the ContainsFold predicate on the "source_ip" field.
func SourceIPContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSourceIP), v))
	})
}

// RequestIDEQ applies the EQ predicate on the "request_id" field.
func RequestIDEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRequestID), v))
	})
}

// RequestIDNEQ applies the NEQ predicate on the "request_id" field.
func RequestIDNEQ(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRequestID), v))
	})
}

// RequestIDIn applies the In predicate on the "request_id" field.
func RequestIDIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRequestID), v...))
	})
}

// RequestIDNotIn applies the NotIn predicate on the "request_id" field.
func RequestIDNotIn(vs ...string) predicate.Event {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Event(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRequestID), v...))
	})
}

// RequestIDGT applies the GT predicate on the "request_id" field.
func RequestIDGT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRequestID), v))
	})
}

// RequestIDGTE applies the GTE predicate on the "request_id" field.
func RequestIDGTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRequestID), v))
	})
}

// RequestIDLT applies the LT predicate on the "request_id" field.
func RequestIDLT(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRequestID), v))
	})
}

// RequestIDLTE applies the LTE predicate on the "request_id" field.
func RequestIDLTE(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRequestID), v))
	})
}

// RequestIDContains applies the Contains predicate on the "request_id" field.
func RequestIDContains(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRequestID), v))
	})
}

// RequestIDHasPrefix applies the HasPrefix predicate on the "request_id" field.
func RequestIDHasPrefix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRequestID), v))
	})
}

// RequestIDHasSuffix applies the HasSuffix predicate on the "request_id" field.
func RequestIDHasSuffix(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRequestID), v))
	})
}

// RequestIDEqualFold applies the EqualFold predicate on the "request_id" field.
func RequestIDEqualFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRequestID), v))
	})
}

// RequestIDContainsFold applies the ContainsFold predicate on the "request_id" field.
func RequestIDContainsFold(v string) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRequestID), v))
	})
}

// ReadOnlyEQ applies the EQ predicate on the "read_only" field.
func ReadOnlyEQ(v bool) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadOnly), v))
	})
}

// ReadOnlyNEQ applies the NEQ predicate on the "read_only" field.
func ReadOnlyNEQ(v bool) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReadOnly), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		p(s.Not())
	})
}
