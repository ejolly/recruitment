// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/admin"
	"github.com/empiricaly/recruitment/internal/ent/procedure"
	"github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ProcedureCreate is the builder for creating a Procedure entity.
type ProcedureCreate struct {
	config
	mutation *ProcedureMutation
	hooks    []Hook
}

// SetCreatedAt sets the createdAt field.
func (pc *ProcedureCreate) SetCreatedAt(t time.Time) *ProcedureCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (pc *ProcedureCreate) SetNillableCreatedAt(t *time.Time) *ProcedureCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the updatedAt field.
func (pc *ProcedureCreate) SetUpdatedAt(t time.Time) *ProcedureCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (pc *ProcedureCreate) SetNillableUpdatedAt(t *time.Time) *ProcedureCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetName sets the name field.
func (pc *ProcedureCreate) SetName(s string) *ProcedureCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetSelectionType sets the selectionType field.
func (pc *ProcedureCreate) SetSelectionType(s string) *ProcedureCreate {
	pc.mutation.SetSelectionType(s)
	return pc
}

// SetParticipantCount sets the participantCount field.
func (pc *ProcedureCreate) SetParticipantCount(i int) *ProcedureCreate {
	pc.mutation.SetParticipantCount(i)
	return pc
}

// SetNillableParticipantCount sets the participantCount field if the given value is not nil.
func (pc *ProcedureCreate) SetNillableParticipantCount(i *int) *ProcedureCreate {
	if i != nil {
		pc.SetParticipantCount(*i)
	}
	return pc
}

// SetInternalCriteria sets the internalCriteria field.
func (pc *ProcedureCreate) SetInternalCriteria(b []byte) *ProcedureCreate {
	pc.mutation.SetInternalCriteria(b)
	return pc
}

// SetMturkCriteria sets the mturkCriteria field.
func (pc *ProcedureCreate) SetMturkCriteria(b []byte) *ProcedureCreate {
	pc.mutation.SetMturkCriteria(b)
	return pc
}

// SetAdult sets the adult field.
func (pc *ProcedureCreate) SetAdult(b bool) *ProcedureCreate {
	pc.mutation.SetAdult(b)
	return pc
}

// SetNillableAdult sets the adult field if the given value is not nil.
func (pc *ProcedureCreate) SetNillableAdult(b *bool) *ProcedureCreate {
	if b != nil {
		pc.SetAdult(*b)
	}
	return pc
}

// SetID sets the id field.
func (pc *ProcedureCreate) SetID(s string) *ProcedureCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetProjectID sets the project edge to Project by id.
func (pc *ProcedureCreate) SetProjectID(id string) *ProcedureCreate {
	pc.mutation.SetProjectID(id)
	return pc
}

// SetNillableProjectID sets the project edge to Project by id if the given value is not nil.
func (pc *ProcedureCreate) SetNillableProjectID(id *string) *ProcedureCreate {
	if id != nil {
		pc = pc.SetProjectID(*id)
	}
	return pc
}

// SetProject sets the project edge to Project.
func (pc *ProcedureCreate) SetProject(p *Project) *ProcedureCreate {
	return pc.SetProjectID(p.ID)
}

// SetCreatorID sets the creator edge to Admin by id.
func (pc *ProcedureCreate) SetCreatorID(id string) *ProcedureCreate {
	pc.mutation.SetCreatorID(id)
	return pc
}

// SetNillableCreatorID sets the creator edge to Admin by id if the given value is not nil.
func (pc *ProcedureCreate) SetNillableCreatorID(id *string) *ProcedureCreate {
	if id != nil {
		pc = pc.SetCreatorID(*id)
	}
	return pc
}

// SetCreator sets the creator edge to Admin.
func (pc *ProcedureCreate) SetCreator(a *Admin) *ProcedureCreate {
	return pc.SetCreatorID(a.ID)
}

// SetRunID sets the run edge to Run by id.
func (pc *ProcedureCreate) SetRunID(id string) *ProcedureCreate {
	pc.mutation.SetRunID(id)
	return pc
}

// SetNillableRunID sets the run edge to Run by id if the given value is not nil.
func (pc *ProcedureCreate) SetNillableRunID(id *string) *ProcedureCreate {
	if id != nil {
		pc = pc.SetRunID(*id)
	}
	return pc
}

// SetRun sets the run edge to Run.
func (pc *ProcedureCreate) SetRun(r *Run) *ProcedureCreate {
	return pc.SetRunID(r.ID)
}

// Mutation returns the ProcedureMutation object of the builder.
func (pc *ProcedureCreate) Mutation() *ProcedureMutation {
	return pc.mutation
}

// Save creates the Procedure in the database.
func (pc *ProcedureCreate) Save(ctx context.Context) (*Procedure, error) {
	if err := pc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Procedure
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProcedureMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProcedureCreate) SaveX(ctx context.Context) *Procedure {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *ProcedureCreate) preSave() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := procedure.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := procedure.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := procedure.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := pc.mutation.SelectionType(); !ok {
		return &ValidationError{Name: "selectionType", err: errors.New("ent: missing required field \"selectionType\"")}
	}
	if _, ok := pc.mutation.ParticipantCount(); !ok {
		v := procedure.DefaultParticipantCount
		pc.mutation.SetParticipantCount(v)
	}
	if v, ok := pc.mutation.ParticipantCount(); ok {
		if err := procedure.ParticipantCountValidator(v); err != nil {
			return &ValidationError{Name: "participantCount", err: fmt.Errorf("ent: validator failed for field \"participantCount\": %w", err)}
		}
	}
	if _, ok := pc.mutation.InternalCriteria(); !ok {
		return &ValidationError{Name: "internalCriteria", err: errors.New("ent: missing required field \"internalCriteria\"")}
	}
	if _, ok := pc.mutation.MturkCriteria(); !ok {
		return &ValidationError{Name: "mturkCriteria", err: errors.New("ent: missing required field \"mturkCriteria\"")}
	}
	if _, ok := pc.mutation.Adult(); !ok {
		v := procedure.DefaultAdult
		pc.mutation.SetAdult(v)
	}
	return nil
}

func (pc *ProcedureCreate) sqlSave(ctx context.Context) (*Procedure, error) {
	pr, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}

func (pc *ProcedureCreate) createSpec() (*Procedure, *sqlgraph.CreateSpec) {
	var (
		pr    = &Procedure{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: procedure.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: procedure.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		pr.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: procedure.FieldCreatedAt,
		})
		pr.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: procedure.FieldUpdatedAt,
		})
		pr.UpdatedAt = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: procedure.FieldName,
		})
		pr.Name = value
	}
	if value, ok := pc.mutation.SelectionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: procedure.FieldSelectionType,
		})
		pr.SelectionType = value
	}
	if value, ok := pc.mutation.ParticipantCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: procedure.FieldParticipantCount,
		})
		pr.ParticipantCount = value
	}
	if value, ok := pc.mutation.InternalCriteria(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: procedure.FieldInternalCriteria,
		})
		pr.InternalCriteria = value
	}
	if value, ok := pc.mutation.MturkCriteria(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: procedure.FieldMturkCriteria,
		})
		pr.MturkCriteria = value
	}
	if value, ok := pc.mutation.Adult(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: procedure.FieldAdult,
		})
		pr.Adult = value
	}
	if nodes := pc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   procedure.ProjectTable,
			Columns: []string{procedure.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   procedure.CreatorTable,
			Columns: []string{procedure.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.RunIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   procedure.RunTable,
			Columns: []string{procedure.RunColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: run.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pr, _spec
}

// ProcedureCreateBulk is the builder for creating a bulk of Procedure entities.
type ProcedureCreateBulk struct {
	config
	builders []*ProcedureCreate
}

// Save creates the Procedure entities in the database.
func (pcb *ProcedureCreateBulk) Save(ctx context.Context) ([]*Procedure, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Procedure, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*ProcedureMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pcb *ProcedureCreateBulk) SaveX(ctx context.Context) []*Procedure {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
