// Code generated by entc, DO NOT EDIT.

package procedure

import (
	"time"

	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// SelectionType applies equality check predicate on the "selectionType" field. It's identical to SelectionTypeEQ.
func SelectionType(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSelectionType), v))
	})
}

// ParticipantCount applies equality check predicate on the "participantCount" field. It's identical to ParticipantCountEQ.
func ParticipantCount(v int) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParticipantCount), v))
	})
}

// InternalCriteria applies equality check predicate on the "internalCriteria" field. It's identical to InternalCriteriaEQ.
func InternalCriteria(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInternalCriteria), v))
	})
}

// MturkCriteria applies equality check predicate on the "mturkCriteria" field. It's identical to MturkCriteriaEQ.
func MturkCriteria(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMturkCriteria), v))
	})
}

// Adult applies equality check predicate on the "adult" field. It's identical to AdultEQ.
func Adult(v bool) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdult), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
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
func CreatedAtNotIn(vs ...time.Time) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
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
func CreatedAtGT(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
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
func UpdatedAtNotIn(vs ...time.Time) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
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
func UpdatedAtGT(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
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
func NameGT(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SelectionTypeEQ applies the EQ predicate on the "selectionType" field.
func SelectionTypeEQ(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSelectionType), v))
	})
}

// SelectionTypeNEQ applies the NEQ predicate on the "selectionType" field.
func SelectionTypeNEQ(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSelectionType), v))
	})
}

// SelectionTypeIn applies the In predicate on the "selectionType" field.
func SelectionTypeIn(vs ...string) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSelectionType), v...))
	})
}

// SelectionTypeNotIn applies the NotIn predicate on the "selectionType" field.
func SelectionTypeNotIn(vs ...string) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSelectionType), v...))
	})
}

// SelectionTypeGT applies the GT predicate on the "selectionType" field.
func SelectionTypeGT(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSelectionType), v))
	})
}

// SelectionTypeGTE applies the GTE predicate on the "selectionType" field.
func SelectionTypeGTE(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSelectionType), v))
	})
}

// SelectionTypeLT applies the LT predicate on the "selectionType" field.
func SelectionTypeLT(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSelectionType), v))
	})
}

// SelectionTypeLTE applies the LTE predicate on the "selectionType" field.
func SelectionTypeLTE(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSelectionType), v))
	})
}

// SelectionTypeContains applies the Contains predicate on the "selectionType" field.
func SelectionTypeContains(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSelectionType), v))
	})
}

// SelectionTypeHasPrefix applies the HasPrefix predicate on the "selectionType" field.
func SelectionTypeHasPrefix(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSelectionType), v))
	})
}

// SelectionTypeHasSuffix applies the HasSuffix predicate on the "selectionType" field.
func SelectionTypeHasSuffix(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSelectionType), v))
	})
}

// SelectionTypeEqualFold applies the EqualFold predicate on the "selectionType" field.
func SelectionTypeEqualFold(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSelectionType), v))
	})
}

// SelectionTypeContainsFold applies the ContainsFold predicate on the "selectionType" field.
func SelectionTypeContainsFold(v string) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSelectionType), v))
	})
}

// ParticipantCountEQ applies the EQ predicate on the "participantCount" field.
func ParticipantCountEQ(v int) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParticipantCount), v))
	})
}

// ParticipantCountNEQ applies the NEQ predicate on the "participantCount" field.
func ParticipantCountNEQ(v int) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParticipantCount), v))
	})
}

// ParticipantCountIn applies the In predicate on the "participantCount" field.
func ParticipantCountIn(vs ...int) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldParticipantCount), v...))
	})
}

// ParticipantCountNotIn applies the NotIn predicate on the "participantCount" field.
func ParticipantCountNotIn(vs ...int) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldParticipantCount), v...))
	})
}

// ParticipantCountGT applies the GT predicate on the "participantCount" field.
func ParticipantCountGT(v int) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldParticipantCount), v))
	})
}

// ParticipantCountGTE applies the GTE predicate on the "participantCount" field.
func ParticipantCountGTE(v int) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldParticipantCount), v))
	})
}

// ParticipantCountLT applies the LT predicate on the "participantCount" field.
func ParticipantCountLT(v int) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldParticipantCount), v))
	})
}

// ParticipantCountLTE applies the LTE predicate on the "participantCount" field.
func ParticipantCountLTE(v int) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldParticipantCount), v))
	})
}

// InternalCriteriaEQ applies the EQ predicate on the "internalCriteria" field.
func InternalCriteriaEQ(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInternalCriteria), v))
	})
}

// InternalCriteriaNEQ applies the NEQ predicate on the "internalCriteria" field.
func InternalCriteriaNEQ(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInternalCriteria), v))
	})
}

// InternalCriteriaIn applies the In predicate on the "internalCriteria" field.
func InternalCriteriaIn(vs ...[]byte) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInternalCriteria), v...))
	})
}

// InternalCriteriaNotIn applies the NotIn predicate on the "internalCriteria" field.
func InternalCriteriaNotIn(vs ...[]byte) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInternalCriteria), v...))
	})
}

// InternalCriteriaGT applies the GT predicate on the "internalCriteria" field.
func InternalCriteriaGT(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInternalCriteria), v))
	})
}

// InternalCriteriaGTE applies the GTE predicate on the "internalCriteria" field.
func InternalCriteriaGTE(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInternalCriteria), v))
	})
}

// InternalCriteriaLT applies the LT predicate on the "internalCriteria" field.
func InternalCriteriaLT(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInternalCriteria), v))
	})
}

// InternalCriteriaLTE applies the LTE predicate on the "internalCriteria" field.
func InternalCriteriaLTE(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInternalCriteria), v))
	})
}

// MturkCriteriaEQ applies the EQ predicate on the "mturkCriteria" field.
func MturkCriteriaEQ(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMturkCriteria), v))
	})
}

// MturkCriteriaNEQ applies the NEQ predicate on the "mturkCriteria" field.
func MturkCriteriaNEQ(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMturkCriteria), v))
	})
}

// MturkCriteriaIn applies the In predicate on the "mturkCriteria" field.
func MturkCriteriaIn(vs ...[]byte) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMturkCriteria), v...))
	})
}

// MturkCriteriaNotIn applies the NotIn predicate on the "mturkCriteria" field.
func MturkCriteriaNotIn(vs ...[]byte) predicate.Procedure {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Procedure(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMturkCriteria), v...))
	})
}

// MturkCriteriaGT applies the GT predicate on the "mturkCriteria" field.
func MturkCriteriaGT(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMturkCriteria), v))
	})
}

// MturkCriteriaGTE applies the GTE predicate on the "mturkCriteria" field.
func MturkCriteriaGTE(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMturkCriteria), v))
	})
}

// MturkCriteriaLT applies the LT predicate on the "mturkCriteria" field.
func MturkCriteriaLT(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMturkCriteria), v))
	})
}

// MturkCriteriaLTE applies the LTE predicate on the "mturkCriteria" field.
func MturkCriteriaLTE(v []byte) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMturkCriteria), v))
	})
}

// AdultEQ applies the EQ predicate on the "adult" field.
func AdultEQ(v bool) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdult), v))
	})
}

// AdultNEQ applies the NEQ predicate on the "adult" field.
func AdultNEQ(v bool) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdult), v))
	})
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProjectTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
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

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.Admin) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRun applies the HasEdge predicate on the "run" edge.
func HasRun() predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RunTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, RunTable, RunColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRunWith applies the HasEdge predicate on the "run" edge with a given conditions (other predicates).
func HasRunWith(preds ...predicate.Run) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RunInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, RunTable, RunColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Procedure) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Procedure) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
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
func Not(p predicate.Procedure) predicate.Procedure {
	return predicate.Procedure(func(s *sql.Selector) {
		p(s.Not())
	})
}
