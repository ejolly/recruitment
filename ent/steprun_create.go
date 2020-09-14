// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/ent/steprun"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// StepRunCreate is the builder for creating a StepRun entity.
type StepRunCreate struct {
	config
	mutation *StepRunMutation
	hooks    []Hook
}

// SetStartAt sets the startAt field.
func (src *StepRunCreate) SetStartAt(t time.Time) *StepRunCreate {
	src.mutation.SetStartAt(t)
	return src
}

// SetEndedAt sets the endedAt field.
func (src *StepRunCreate) SetEndedAt(t time.Time) *StepRunCreate {
	src.mutation.SetEndedAt(t)
	return src
}

// SetParticipantsCount sets the participantsCount field.
func (src *StepRunCreate) SetParticipantsCount(i int) *StepRunCreate {
	src.mutation.SetParticipantsCount(i)
	return src
}

// Mutation returns the StepRunMutation object of the builder.
func (src *StepRunCreate) Mutation() *StepRunMutation {
	return src.mutation
}

// Save creates the StepRun in the database.
func (src *StepRunCreate) Save(ctx context.Context) (*StepRun, error) {
	if err := src.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *StepRun
	)
	if len(src.hooks) == 0 {
		node, err = src.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StepRunMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			src.mutation = mutation
			node, err = src.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(src.hooks) - 1; i >= 0; i-- {
			mut = src.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, src.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (src *StepRunCreate) SaveX(ctx context.Context) *StepRun {
	v, err := src.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (src *StepRunCreate) preSave() error {
	if _, ok := src.mutation.StartAt(); !ok {
		return &ValidationError{Name: "startAt", err: errors.New("ent: missing required field \"startAt\"")}
	}
	if _, ok := src.mutation.EndedAt(); !ok {
		return &ValidationError{Name: "endedAt", err: errors.New("ent: missing required field \"endedAt\"")}
	}
	if _, ok := src.mutation.ParticipantsCount(); !ok {
		return &ValidationError{Name: "participantsCount", err: errors.New("ent: missing required field \"participantsCount\"")}
	}
	return nil
}

func (src *StepRunCreate) sqlSave(ctx context.Context) (*StepRun, error) {
	sr, _spec := src.createSpec()
	if err := sqlgraph.CreateNode(ctx, src.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	sr.ID = int(id)
	return sr, nil
}

func (src *StepRunCreate) createSpec() (*StepRun, *sqlgraph.CreateSpec) {
	var (
		sr    = &StepRun{config: src.config}
		_spec = &sqlgraph.CreateSpec{
			Table: steprun.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: steprun.FieldID,
			},
		}
	)
	if value, ok := src.mutation.StartAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: steprun.FieldStartAt,
		})
		sr.StartAt = value
	}
	if value, ok := src.mutation.EndedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: steprun.FieldEndedAt,
		})
		sr.EndedAt = value
	}
	if value, ok := src.mutation.ParticipantsCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: steprun.FieldParticipantsCount,
		})
		sr.ParticipantsCount = value
	}
	return sr, _spec
}

// StepRunCreateBulk is the builder for creating a bulk of StepRun entities.
type StepRunCreateBulk struct {
	config
	builders []*StepRunCreate
}

// Save creates the StepRun entities in the database.
func (srcb *StepRunCreateBulk) Save(ctx context.Context) ([]*StepRun, error) {
	specs := make([]*sqlgraph.CreateSpec, len(srcb.builders))
	nodes := make([]*StepRun, len(srcb.builders))
	mutators := make([]Mutator, len(srcb.builders))
	for i := range srcb.builders {
		func(i int, root context.Context) {
			builder := srcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*StepRunMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, srcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, srcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, srcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (srcb *StepRunCreateBulk) SaveX(ctx context.Context) []*StepRun {
	v, err := srcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
