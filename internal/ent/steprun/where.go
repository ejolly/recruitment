// Code generated by entc, DO NOT EDIT.

package steprun

import (
	"time"

	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
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
func IDGT(id string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// StartedAt applies equality check predicate on the "startedAt" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// EndedAt applies equality check predicate on the "endedAt" field. It's identical to EndedAtEQ.
func EndedAt(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndedAt), v))
	})
}

// Index applies equality check predicate on the "index" field. It's identical to IndexEQ.
func Index(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIndex), v))
	})
}

// ParticipantsCount applies equality check predicate on the "participantsCount" field. It's identical to ParticipantsCountEQ.
func ParticipantsCount(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParticipantsCount), v))
	})
}

// HitID applies equality check predicate on the "hitID" field. It's identical to HitIDEQ.
func HitID(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHitID), v))
	})
}

// UrlToken applies equality check predicate on the "urlToken" field. It's identical to UrlTokenEQ.
func UrlToken(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUrlToken), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
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
func StatusNotIn(vs ...Status) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StartedAtEQ applies the EQ predicate on the "startedAt" field.
func StartedAtEQ(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtNEQ applies the NEQ predicate on the "startedAt" field.
func StartedAtNEQ(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartedAt), v))
	})
}

// StartedAtIn applies the In predicate on the "startedAt" field.
func StartedAtIn(vs ...time.Time) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
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
func StartedAtNotIn(vs ...time.Time) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
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
func StartedAtGT(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartedAt), v))
	})
}

// StartedAtGTE applies the GTE predicate on the "startedAt" field.
func StartedAtGTE(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtLT applies the LT predicate on the "startedAt" field.
func StartedAtLT(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartedAt), v))
	})
}

// StartedAtLTE applies the LTE predicate on the "startedAt" field.
func StartedAtLTE(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartedAt), v))
	})
}

// StartedAtIsNil applies the IsNil predicate on the "startedAt" field.
func StartedAtIsNil() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartedAt)))
	})
}

// StartedAtNotNil applies the NotNil predicate on the "startedAt" field.
func StartedAtNotNil() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartedAt)))
	})
}

// EndedAtEQ applies the EQ predicate on the "endedAt" field.
func EndedAtEQ(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndedAt), v))
	})
}

// EndedAtNEQ applies the NEQ predicate on the "endedAt" field.
func EndedAtNEQ(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndedAt), v))
	})
}

// EndedAtIn applies the In predicate on the "endedAt" field.
func EndedAtIn(vs ...time.Time) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
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
func EndedAtNotIn(vs ...time.Time) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
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
func EndedAtGT(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndedAt), v))
	})
}

// EndedAtGTE applies the GTE predicate on the "endedAt" field.
func EndedAtGTE(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndedAt), v))
	})
}

// EndedAtLT applies the LT predicate on the "endedAt" field.
func EndedAtLT(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndedAt), v))
	})
}

// EndedAtLTE applies the LTE predicate on the "endedAt" field.
func EndedAtLTE(v time.Time) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndedAt), v))
	})
}

// EndedAtIsNil applies the IsNil predicate on the "endedAt" field.
func EndedAtIsNil() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEndedAt)))
	})
}

// EndedAtNotNil applies the NotNil predicate on the "endedAt" field.
func EndedAtNotNil() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEndedAt)))
	})
}

// IndexEQ applies the EQ predicate on the "index" field.
func IndexEQ(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIndex), v))
	})
}

// IndexNEQ applies the NEQ predicate on the "index" field.
func IndexNEQ(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIndex), v))
	})
}

// IndexIn applies the In predicate on the "index" field.
func IndexIn(vs ...int) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIndex), v...))
	})
}

// IndexNotIn applies the NotIn predicate on the "index" field.
func IndexNotIn(vs ...int) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIndex), v...))
	})
}

// IndexGT applies the GT predicate on the "index" field.
func IndexGT(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIndex), v))
	})
}

// IndexGTE applies the GTE predicate on the "index" field.
func IndexGTE(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIndex), v))
	})
}

// IndexLT applies the LT predicate on the "index" field.
func IndexLT(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIndex), v))
	})
}

// IndexLTE applies the LTE predicate on the "index" field.
func IndexLTE(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIndex), v))
	})
}

// ParticipantsCountEQ applies the EQ predicate on the "participantsCount" field.
func ParticipantsCountEQ(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParticipantsCount), v))
	})
}

// ParticipantsCountNEQ applies the NEQ predicate on the "participantsCount" field.
func ParticipantsCountNEQ(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParticipantsCount), v))
	})
}

// ParticipantsCountIn applies the In predicate on the "participantsCount" field.
func ParticipantsCountIn(vs ...int) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldParticipantsCount), v...))
	})
}

// ParticipantsCountNotIn applies the NotIn predicate on the "participantsCount" field.
func ParticipantsCountNotIn(vs ...int) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldParticipantsCount), v...))
	})
}

// ParticipantsCountGT applies the GT predicate on the "participantsCount" field.
func ParticipantsCountGT(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldParticipantsCount), v))
	})
}

// ParticipantsCountGTE applies the GTE predicate on the "participantsCount" field.
func ParticipantsCountGTE(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldParticipantsCount), v))
	})
}

// ParticipantsCountLT applies the LT predicate on the "participantsCount" field.
func ParticipantsCountLT(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldParticipantsCount), v))
	})
}

// ParticipantsCountLTE applies the LTE predicate on the "participantsCount" field.
func ParticipantsCountLTE(v int) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldParticipantsCount), v))
	})
}

// HitIDEQ applies the EQ predicate on the "hitID" field.
func HitIDEQ(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHitID), v))
	})
}

// HitIDNEQ applies the NEQ predicate on the "hitID" field.
func HitIDNEQ(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHitID), v))
	})
}

// HitIDIn applies the In predicate on the "hitID" field.
func HitIDIn(vs ...string) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHitID), v...))
	})
}

// HitIDNotIn applies the NotIn predicate on the "hitID" field.
func HitIDNotIn(vs ...string) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHitID), v...))
	})
}

// HitIDGT applies the GT predicate on the "hitID" field.
func HitIDGT(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHitID), v))
	})
}

// HitIDGTE applies the GTE predicate on the "hitID" field.
func HitIDGTE(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHitID), v))
	})
}

// HitIDLT applies the LT predicate on the "hitID" field.
func HitIDLT(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHitID), v))
	})
}

// HitIDLTE applies the LTE predicate on the "hitID" field.
func HitIDLTE(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHitID), v))
	})
}

// HitIDContains applies the Contains predicate on the "hitID" field.
func HitIDContains(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHitID), v))
	})
}

// HitIDHasPrefix applies the HasPrefix predicate on the "hitID" field.
func HitIDHasPrefix(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHitID), v))
	})
}

// HitIDHasSuffix applies the HasSuffix predicate on the "hitID" field.
func HitIDHasSuffix(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHitID), v))
	})
}

// HitIDIsNil applies the IsNil predicate on the "hitID" field.
func HitIDIsNil() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHitID)))
	})
}

// HitIDNotNil applies the NotNil predicate on the "hitID" field.
func HitIDNotNil() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHitID)))
	})
}

// HitIDEqualFold applies the EqualFold predicate on the "hitID" field.
func HitIDEqualFold(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHitID), v))
	})
}

// HitIDContainsFold applies the ContainsFold predicate on the "hitID" field.
func HitIDContainsFold(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHitID), v))
	})
}

// UrlTokenEQ applies the EQ predicate on the "urlToken" field.
func UrlTokenEQ(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUrlToken), v))
	})
}

// UrlTokenNEQ applies the NEQ predicate on the "urlToken" field.
func UrlTokenNEQ(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUrlToken), v))
	})
}

// UrlTokenIn applies the In predicate on the "urlToken" field.
func UrlTokenIn(vs ...string) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUrlToken), v...))
	})
}

// UrlTokenNotIn applies the NotIn predicate on the "urlToken" field.
func UrlTokenNotIn(vs ...string) predicate.StepRun {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.StepRun(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUrlToken), v...))
	})
}

// UrlTokenGT applies the GT predicate on the "urlToken" field.
func UrlTokenGT(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUrlToken), v))
	})
}

// UrlTokenGTE applies the GTE predicate on the "urlToken" field.
func UrlTokenGTE(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUrlToken), v))
	})
}

// UrlTokenLT applies the LT predicate on the "urlToken" field.
func UrlTokenLT(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUrlToken), v))
	})
}

// UrlTokenLTE applies the LTE predicate on the "urlToken" field.
func UrlTokenLTE(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUrlToken), v))
	})
}

// UrlTokenContains applies the Contains predicate on the "urlToken" field.
func UrlTokenContains(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUrlToken), v))
	})
}

// UrlTokenHasPrefix applies the HasPrefix predicate on the "urlToken" field.
func UrlTokenHasPrefix(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUrlToken), v))
	})
}

// UrlTokenHasSuffix applies the HasSuffix predicate on the "urlToken" field.
func UrlTokenHasSuffix(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUrlToken), v))
	})
}

// UrlTokenEqualFold applies the EqualFold predicate on the "urlToken" field.
func UrlTokenEqualFold(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUrlToken), v))
	})
}

// UrlTokenContainsFold applies the ContainsFold predicate on the "urlToken" field.
func UrlTokenContainsFold(v string) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUrlToken), v))
	})
}

// HasCreatedParticipants applies the HasEdge predicate on the "createdParticipants" edge.
func HasCreatedParticipants() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatedParticipantsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CreatedParticipantsTable, CreatedParticipantsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatedParticipantsWith applies the HasEdge predicate on the "createdParticipants" edge with a given conditions (other predicates).
func HasCreatedParticipantsWith(preds ...predicate.Participant) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatedParticipantsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CreatedParticipantsTable, CreatedParticipantsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParticipants applies the HasEdge predicate on the "participants" edge.
func HasParticipants() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParticipantsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ParticipantsTable, ParticipantsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParticipantsWith applies the HasEdge predicate on the "participants" edge with a given conditions (other predicates).
func HasParticipantsWith(preds ...predicate.Participant) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParticipantsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ParticipantsTable, ParticipantsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParticipations applies the HasEdge predicate on the "participations" edge.
func HasParticipations() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParticipationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ParticipationsTable, ParticipationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParticipationsWith applies the HasEdge predicate on the "participations" edge with a given conditions (other predicates).
func HasParticipationsWith(preds ...predicate.Participation) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParticipationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ParticipationsTable, ParticipationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStep applies the HasEdge predicate on the "step" edge.
func HasStep() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StepTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, StepTable, StepColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStepWith applies the HasEdge predicate on the "step" edge with a given conditions (other predicates).
func HasStepWith(preds ...predicate.Step) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StepInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, StepTable, StepColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRun applies the HasEdge predicate on the "run" edge.
func HasRun() predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RunTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RunTable, RunColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRunWith applies the HasEdge predicate on the "run" edge with a given conditions (other predicates).
func HasRunWith(preds ...predicate.Run) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RunInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RunTable, RunColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.StepRun) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.StepRun) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
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
func Not(p predicate.StepRun) predicate.StepRun {
	return predicate.StepRun(func(s *sql.Selector) {
		p(s.Not())
	})
}
