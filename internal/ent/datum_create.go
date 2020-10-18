// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/datum"
	"github.com/empiricaly/recruitment/internal/ent/participant"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// DatumCreate is the builder for creating a Datum entity.
type DatumCreate struct {
	config
	mutation *DatumMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (dc *DatumCreate) SetCreatedAt(t time.Time) *DatumCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (dc *DatumCreate) SetNillableCreatedAt(t *time.Time) *DatumCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the updated_at field.
func (dc *DatumCreate) SetUpdatedAt(t time.Time) *DatumCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (dc *DatumCreate) SetNillableUpdatedAt(t *time.Time) *DatumCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetKey sets the key field.
func (dc *DatumCreate) SetKey(s string) *DatumCreate {
	dc.mutation.SetKey(s)
	return dc
}

// SetVal sets the val field.
func (dc *DatumCreate) SetVal(b []byte) *DatumCreate {
	dc.mutation.SetVal(b)
	return dc
}

// SetIndex sets the index field.
func (dc *DatumCreate) SetIndex(i int) *DatumCreate {
	dc.mutation.SetIndex(i)
	return dc
}

// SetNillableIndex sets the index field if the given value is not nil.
func (dc *DatumCreate) SetNillableIndex(i *int) *DatumCreate {
	if i != nil {
		dc.SetIndex(*i)
	}
	return dc
}

// SetCurrent sets the current field.
func (dc *DatumCreate) SetCurrent(b bool) *DatumCreate {
	dc.mutation.SetCurrent(b)
	return dc
}

// SetNillableCurrent sets the current field if the given value is not nil.
func (dc *DatumCreate) SetNillableCurrent(b *bool) *DatumCreate {
	if b != nil {
		dc.SetCurrent(*b)
	}
	return dc
}

// SetVersion sets the version field.
func (dc *DatumCreate) SetVersion(i int) *DatumCreate {
	dc.mutation.SetVersion(i)
	return dc
}

// SetNillableVersion sets the version field if the given value is not nil.
func (dc *DatumCreate) SetNillableVersion(i *int) *DatumCreate {
	if i != nil {
		dc.SetVersion(*i)
	}
	return dc
}

// SetDeletedAt sets the deletedAt field.
func (dc *DatumCreate) SetDeletedAt(t time.Time) *DatumCreate {
	dc.mutation.SetDeletedAt(t)
	return dc
}

// SetNillableDeletedAt sets the deletedAt field if the given value is not nil.
func (dc *DatumCreate) SetNillableDeletedAt(t *time.Time) *DatumCreate {
	if t != nil {
		dc.SetDeletedAt(*t)
	}
	return dc
}

// SetID sets the id field.
func (dc *DatumCreate) SetID(s string) *DatumCreate {
	dc.mutation.SetID(s)
	return dc
}

// SetParticipantID sets the participant edge to Participant by id.
func (dc *DatumCreate) SetParticipantID(id string) *DatumCreate {
	dc.mutation.SetParticipantID(id)
	return dc
}

// SetParticipant sets the participant edge to Participant.
func (dc *DatumCreate) SetParticipant(p *Participant) *DatumCreate {
	return dc.SetParticipantID(p.ID)
}

// Mutation returns the DatumMutation object of the builder.
func (dc *DatumCreate) Mutation() *DatumMutation {
	return dc.mutation
}

// Save creates the Datum in the database.
func (dc *DatumCreate) Save(ctx context.Context) (*Datum, error) {
	var (
		err  error
		node *Datum
	)
	dc.defaults()
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DatumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DatumCreate) SaveX(ctx context.Context) *Datum {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (dc *DatumCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := datum.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := datum.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.Index(); !ok {
		v := datum.DefaultIndex
		dc.mutation.SetIndex(v)
	}
	if _, ok := dc.mutation.Current(); !ok {
		v := datum.DefaultCurrent
		dc.mutation.SetCurrent(v)
	}
	if _, ok := dc.mutation.Version(); !ok {
		v := datum.DefaultVersion
		dc.mutation.SetVersion(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DatumCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := dc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New("ent: missing required field \"key\"")}
	}
	if _, ok := dc.mutation.Val(); !ok {
		return &ValidationError{Name: "val", err: errors.New("ent: missing required field \"val\"")}
	}
	if _, ok := dc.mutation.Index(); !ok {
		return &ValidationError{Name: "index", err: errors.New("ent: missing required field \"index\"")}
	}
	if _, ok := dc.mutation.Current(); !ok {
		return &ValidationError{Name: "current", err: errors.New("ent: missing required field \"current\"")}
	}
	if _, ok := dc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New("ent: missing required field \"version\"")}
	}
	if v, ok := dc.mutation.ID(); ok {
		if err := datum.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	if _, ok := dc.mutation.ParticipantID(); !ok {
		return &ValidationError{Name: "participant", err: errors.New("ent: missing required edge \"participant\"")}
	}
	return nil
}

func (dc *DatumCreate) sqlSave(ctx context.Context) (*Datum, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (dc *DatumCreate) createSpec() (*Datum, *sqlgraph.CreateSpec) {
	var (
		_node = &Datum{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: datum.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: datum.FieldID,
			},
		}
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: datum.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: datum.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.Key(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datum.FieldKey,
		})
		_node.Key = value
	}
	if value, ok := dc.mutation.Val(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: datum.FieldVal,
		})
		_node.Val = value
	}
	if value, ok := dc.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: datum.FieldIndex,
		})
		_node.Index = value
	}
	if value, ok := dc.mutation.Current(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: datum.FieldCurrent,
		})
		_node.Current = value
	}
	if value, ok := dc.mutation.Version(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: datum.FieldVersion,
		})
		_node.Version = value
	}
	if value, ok := dc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: datum.FieldDeletedAt,
		})
		_node.DeletedAt = &value
	}
	if nodes := dc.mutation.ParticipantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   datum.ParticipantTable,
			Columns: []string{datum.ParticipantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DatumCreateBulk is the builder for creating a bulk of Datum entities.
type DatumCreateBulk struct {
	config
	builders []*DatumCreate
}

// Save creates the Datum entities in the database.
func (dcb *DatumCreateBulk) Save(ctx context.Context) ([]*Datum, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Datum, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DatumMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (dcb *DatumCreateBulk) SaveX(ctx context.Context) []*Datum {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
