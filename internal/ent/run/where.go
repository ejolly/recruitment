// Code generated by entc, DO NOT EDIT.

package run

import (
	"time"

	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// StartAt applies equality check predicate on the "startAt" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartedAt applies equality check predicate on the "startedAt" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// EndedAt applies equality check predicate on the "endedAt" field. It's identical to EndedAtEQ.
func EndedAt(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndedAt), v))
	})
}

// Error applies equality check predicate on the "error" field. It's identical to ErrorEQ.
func Error(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldError), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StartAtEQ applies the EQ predicate on the "startAt" field.
func StartAtEQ(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartAtNEQ applies the NEQ predicate on the "startAt" field.
func StartAtNEQ(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAt), v))
	})
}

// StartAtIn applies the In predicate on the "startAt" field.
func StartAtIn(vs ...time.Time) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartAt), v...))
	})
}

// StartAtNotIn applies the NotIn predicate on the "startAt" field.
func StartAtNotIn(vs ...time.Time) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartAt), v...))
	})
}

// StartAtGT applies the GT predicate on the "startAt" field.
func StartAtGT(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAt), v))
	})
}

// StartAtGTE applies the GTE predicate on the "startAt" field.
func StartAtGTE(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAt), v))
	})
}

// StartAtLT applies the LT predicate on the "startAt" field.
func StartAtLT(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAt), v))
	})
}

// StartAtLTE applies the LTE predicate on the "startAt" field.
func StartAtLTE(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAt), v))
	})
}

// StartAtIsNil applies the IsNil predicate on the "startAt" field.
func StartAtIsNil() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartAt)))
	})
}

// StartAtNotNil applies the NotNil predicate on the "startAt" field.
func StartAtNotNil() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartAt)))
	})
}

// StartedAtEQ applies the EQ predicate on the "startedAt" field.
func StartedAtEQ(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtNEQ applies the NEQ predicate on the "startedAt" field.
func StartedAtNEQ(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtIn applies the In predicate on the "startedAt" field.
func StartedAtIn(vs ...time.Time) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartedAt), v...))
	})
}

// StartedAtNotIn applies the NotIn predicate on the "startedAt" field.
func StartedAtNotIn(vs ...time.Time) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartedAt), v...))
	})
}

// StartedAtGT applies the GT predicate on the "startedAt" field.
func StartedAtGT(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartedAt), v))
	})
}

// StartedAtGTE applies the GTE predicate on the "startedAt" field.
func StartedAtGTE(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtLT applies the LT predicate on the "startedAt" field.
func StartedAtLT(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartedAt), v))
	})
}

// StartedAtLTE applies the LTE predicate on the "startedAt" field.
func StartedAtLTE(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtIsNil applies the IsNil predicate on the "startedAt" field.
func StartedAtIsNil() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartedAt)))
	})
}

// StartedAtNotNil applies the NotNil predicate on the "startedAt" field.
func StartedAtNotNil() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartedAt)))
	})
}

// EndedAtEQ applies the EQ predicate on the "endedAt" field.
func EndedAtEQ(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndedAt), v))
	})
}

// EndedAtNEQ applies the NEQ predicate on the "endedAt" field.
func EndedAtNEQ(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndedAt), v))
	})
}

// EndedAtIn applies the In predicate on the "endedAt" field.
func EndedAtIn(vs ...time.Time) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndedAt), v...))
	})
}

// EndedAtNotIn applies the NotIn predicate on the "endedAt" field.
func EndedAtNotIn(vs ...time.Time) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndedAt), v...))
	})
}

// EndedAtGT applies the GT predicate on the "endedAt" field.
func EndedAtGT(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndedAt), v))
	})
}

// EndedAtGTE applies the GTE predicate on the "endedAt" field.
func EndedAtGTE(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndedAt), v))
	})
}

// EndedAtLT applies the LT predicate on the "endedAt" field.
func EndedAtLT(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndedAt), v))
	})
}

// EndedAtLTE applies the LTE predicate on the "endedAt" field.
func EndedAtLTE(v time.Time) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndedAt), v))
	})
}

// EndedAtIsNil applies the IsNil predicate on the "endedAt" field.
func EndedAtIsNil() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndedAt)))
	})
}

// EndedAtNotNil applies the NotNil predicate on the "endedAt" field.
func EndedAtNotNil() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndedAt)))
	})
}

// ErrorEQ applies the EQ predicate on the "error" field.
func ErrorEQ(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldError), v))
	})
}

// ErrorNEQ applies the NEQ predicate on the "error" field.
func ErrorNEQ(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldError), v))
	})
}

// ErrorIn applies the In predicate on the "error" field.
func ErrorIn(vs ...string) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldError), v...))
	})
}

// ErrorNotIn applies the NotIn predicate on the "error" field.
func ErrorNotIn(vs ...string) predicate.Run {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Run(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldError), v...))
	})
}

// ErrorGT applies the GT predicate on the "error" field.
func ErrorGT(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldError), v))
	})
}

// ErrorGTE applies the GTE predicate on the "error" field.
func ErrorGTE(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldError), v))
	})
}

// ErrorLT applies the LT predicate on the "error" field.
func ErrorLT(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldError), v))
	})
}

// ErrorLTE applies the LTE predicate on the "error" field.
func ErrorLTE(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldError), v))
	})
}

// ErrorContains applies the Contains predicate on the "error" field.
func ErrorContains(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldError), v))
	})
}

// ErrorHasPrefix applies the HasPrefix predicate on the "error" field.
func ErrorHasPrefix(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldError), v))
	})
}

// ErrorHasSuffix applies the HasSuffix predicate on the "error" field.
func ErrorHasSuffix(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldError), v))
	})
}

// ErrorIsNil applies the IsNil predicate on the "error" field.
func ErrorIsNil() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldError)))
	})
}

// ErrorNotNil applies the NotNil predicate on the "error" field.
func ErrorNotNil() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldError)))
	})
}

// ErrorEqualFold applies the EqualFold predicate on the "error" field.
func ErrorEqualFold(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldError), v))
	})
}

// ErrorContainsFold applies the ContainsFold predicate on the "error" field.
func ErrorContainsFold(v string) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldError), v))
	})
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProcedure applies the HasEdge predicate on the "procedure" edge.
func HasProcedure() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProcedureTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ProcedureTable, ProcedureColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProcedureWith applies the HasEdge predicate on the "procedure" edge with a given conditions (other predicates).
func HasProcedureWith(preds ...predicate.Procedure) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProcedureInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ProcedureTable, ProcedureColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSteps applies the HasEdge predicate on the "steps" edge.
func HasSteps() predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StepsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StepsTable, StepsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStepsWith applies the HasEdge predicate on the "steps" edge with a given conditions (other predicates).
func HasStepsWith(preds ...predicate.StepRun) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StepsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, StepsTable, StepsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Run) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Run) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
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
func Not(p predicate.Run) predicate.Run {
	return predicate.Run(func(s *sql.Selector) {
		p(s.Not())
	})
}