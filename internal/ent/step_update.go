// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/empiricaly/recruitment/internal/ent/procedure"
	"github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// StepUpdate is the builder for updating Step entities.
type StepUpdate struct {
	config
	hooks      []Hook
	mutation   *StepMutation
	predicates []predicate.Step
}

// Where adds a new predicate for the builder.
func (su *StepUpdate) Where(ps ...predicate.Step) *StepUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetUpdatedAt sets the updatedAt field.
func (su *StepUpdate) SetUpdatedAt(t time.Time) *StepUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetType sets the type field.
func (su *StepUpdate) SetType(s step.Type) *StepUpdate {
	su.mutation.SetType(s)
	return su
}

// SetIndex sets the index field.
func (su *StepUpdate) SetIndex(i int) *StepUpdate {
	su.mutation.ResetIndex()
	su.mutation.SetIndex(i)
	return su
}

// AddIndex adds i to index.
func (su *StepUpdate) AddIndex(i int) *StepUpdate {
	su.mutation.AddIndex(i)
	return su
}

// SetDuration sets the duration field.
func (su *StepUpdate) SetDuration(i int) *StepUpdate {
	su.mutation.ResetDuration()
	su.mutation.SetDuration(i)
	return su
}

// AddDuration adds i to duration.
func (su *StepUpdate) AddDuration(i int) *StepUpdate {
	su.mutation.AddDuration(i)
	return su
}

// SetMsgArgs sets the msgArgs field.
func (su *StepUpdate) SetMsgArgs(b []byte) *StepUpdate {
	su.mutation.SetMsgArgs(b)
	return su
}

// ClearMsgArgs clears the value of msgArgs.
func (su *StepUpdate) ClearMsgArgs() *StepUpdate {
	su.mutation.ClearMsgArgs()
	return su
}

// SetHitArgs sets the hitArgs field.
func (su *StepUpdate) SetHitArgs(b []byte) *StepUpdate {
	su.mutation.SetHitArgs(b)
	return su
}

// ClearHitArgs clears the value of hitArgs.
func (su *StepUpdate) ClearHitArgs() *StepUpdate {
	su.mutation.ClearHitArgs()
	return su
}

// SetFilterArgs sets the filterArgs field.
func (su *StepUpdate) SetFilterArgs(b []byte) *StepUpdate {
	su.mutation.SetFilterArgs(b)
	return su
}

// ClearFilterArgs clears the value of filterArgs.
func (su *StepUpdate) ClearFilterArgs() *StepUpdate {
	su.mutation.ClearFilterArgs()
	return su
}

// SetProcedureID sets the procedure edge to Procedure by id.
func (su *StepUpdate) SetProcedureID(id string) *StepUpdate {
	su.mutation.SetProcedureID(id)
	return su
}

// SetNillableProcedureID sets the procedure edge to Procedure by id if the given value is not nil.
func (su *StepUpdate) SetNillableProcedureID(id *string) *StepUpdate {
	if id != nil {
		su = su.SetProcedureID(*id)
	}
	return su
}

// SetProcedure sets the procedure edge to Procedure.
func (su *StepUpdate) SetProcedure(p *Procedure) *StepUpdate {
	return su.SetProcedureID(p.ID)
}

// Mutation returns the StepMutation object of the builder.
func (su *StepUpdate) Mutation() *StepMutation {
	return su.mutation
}

// ClearProcedure clears the procedure edge to Procedure.
func (su *StepUpdate) ClearProcedure() *StepUpdate {
	su.mutation.ClearProcedure()
	return su
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *StepUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := step.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
	if v, ok := su.mutation.GetType(); ok {
		if err := step.TypeValidator(v); err != nil {
			return 0, &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StepMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StepUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StepUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StepUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StepUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   step.Table,
			Columns: step.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: step.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: step.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: step.FieldType,
		})
	}
	if value, ok := su.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: step.FieldIndex,
		})
	}
	if value, ok := su.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: step.FieldIndex,
		})
	}
	if value, ok := su.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: step.FieldDuration,
		})
	}
	if value, ok := su.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: step.FieldDuration,
		})
	}
	if value, ok := su.mutation.MsgArgs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: step.FieldMsgArgs,
		})
	}
	if su.mutation.MsgArgsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: step.FieldMsgArgs,
		})
	}
	if value, ok := su.mutation.HitArgs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: step.FieldHitArgs,
		})
	}
	if su.mutation.HitArgsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: step.FieldHitArgs,
		})
	}
	if value, ok := su.mutation.FilterArgs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: step.FieldFilterArgs,
		})
	}
	if su.mutation.FilterArgsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: step.FieldFilterArgs,
		})
	}
	if su.mutation.ProcedureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   step.ProcedureTable,
			Columns: []string{step.ProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: procedure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ProcedureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   step.ProcedureTable,
			Columns: []string{step.ProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: procedure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{step.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StepUpdateOne is the builder for updating a single Step entity.
type StepUpdateOne struct {
	config
	hooks    []Hook
	mutation *StepMutation
}

// SetUpdatedAt sets the updatedAt field.
func (suo *StepUpdateOne) SetUpdatedAt(t time.Time) *StepUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetType sets the type field.
func (suo *StepUpdateOne) SetType(s step.Type) *StepUpdateOne {
	suo.mutation.SetType(s)
	return suo
}

// SetIndex sets the index field.
func (suo *StepUpdateOne) SetIndex(i int) *StepUpdateOne {
	suo.mutation.ResetIndex()
	suo.mutation.SetIndex(i)
	return suo
}

// AddIndex adds i to index.
func (suo *StepUpdateOne) AddIndex(i int) *StepUpdateOne {
	suo.mutation.AddIndex(i)
	return suo
}

// SetDuration sets the duration field.
func (suo *StepUpdateOne) SetDuration(i int) *StepUpdateOne {
	suo.mutation.ResetDuration()
	suo.mutation.SetDuration(i)
	return suo
}

// AddDuration adds i to duration.
func (suo *StepUpdateOne) AddDuration(i int) *StepUpdateOne {
	suo.mutation.AddDuration(i)
	return suo
}

// SetMsgArgs sets the msgArgs field.
func (suo *StepUpdateOne) SetMsgArgs(b []byte) *StepUpdateOne {
	suo.mutation.SetMsgArgs(b)
	return suo
}

// ClearMsgArgs clears the value of msgArgs.
func (suo *StepUpdateOne) ClearMsgArgs() *StepUpdateOne {
	suo.mutation.ClearMsgArgs()
	return suo
}

// SetHitArgs sets the hitArgs field.
func (suo *StepUpdateOne) SetHitArgs(b []byte) *StepUpdateOne {
	suo.mutation.SetHitArgs(b)
	return suo
}

// ClearHitArgs clears the value of hitArgs.
func (suo *StepUpdateOne) ClearHitArgs() *StepUpdateOne {
	suo.mutation.ClearHitArgs()
	return suo
}

// SetFilterArgs sets the filterArgs field.
func (suo *StepUpdateOne) SetFilterArgs(b []byte) *StepUpdateOne {
	suo.mutation.SetFilterArgs(b)
	return suo
}

// ClearFilterArgs clears the value of filterArgs.
func (suo *StepUpdateOne) ClearFilterArgs() *StepUpdateOne {
	suo.mutation.ClearFilterArgs()
	return suo
}

// SetProcedureID sets the procedure edge to Procedure by id.
func (suo *StepUpdateOne) SetProcedureID(id string) *StepUpdateOne {
	suo.mutation.SetProcedureID(id)
	return suo
}

// SetNillableProcedureID sets the procedure edge to Procedure by id if the given value is not nil.
func (suo *StepUpdateOne) SetNillableProcedureID(id *string) *StepUpdateOne {
	if id != nil {
		suo = suo.SetProcedureID(*id)
	}
	return suo
}

// SetProcedure sets the procedure edge to Procedure.
func (suo *StepUpdateOne) SetProcedure(p *Procedure) *StepUpdateOne {
	return suo.SetProcedureID(p.ID)
}

// Mutation returns the StepMutation object of the builder.
func (suo *StepUpdateOne) Mutation() *StepMutation {
	return suo.mutation
}

// ClearProcedure clears the procedure edge to Procedure.
func (suo *StepUpdateOne) ClearProcedure() *StepUpdateOne {
	suo.mutation.ClearProcedure()
	return suo
}

// Save executes the query and returns the updated entity.
func (suo *StepUpdateOne) Save(ctx context.Context) (*Step, error) {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := step.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
	if v, ok := suo.mutation.GetType(); ok {
		if err := step.TypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}

	var (
		err  error
		node *Step
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StepMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StepUpdateOne) SaveX(ctx context.Context) *Step {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *StepUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StepUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StepUpdateOne) sqlSave(ctx context.Context) (s *Step, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   step.Table,
			Columns: step.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: step.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Step.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: step.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: step.FieldType,
		})
	}
	if value, ok := suo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: step.FieldIndex,
		})
	}
	if value, ok := suo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: step.FieldIndex,
		})
	}
	if value, ok := suo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: step.FieldDuration,
		})
	}
	if value, ok := suo.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: step.FieldDuration,
		})
	}
	if value, ok := suo.mutation.MsgArgs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: step.FieldMsgArgs,
		})
	}
	if suo.mutation.MsgArgsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: step.FieldMsgArgs,
		})
	}
	if value, ok := suo.mutation.HitArgs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: step.FieldHitArgs,
		})
	}
	if suo.mutation.HitArgsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: step.FieldHitArgs,
		})
	}
	if value, ok := suo.mutation.FilterArgs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: step.FieldFilterArgs,
		})
	}
	if suo.mutation.FilterArgsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: step.FieldFilterArgs,
		})
	}
	if suo.mutation.ProcedureCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   step.ProcedureTable,
			Columns: []string{step.ProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: procedure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ProcedureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   step.ProcedureTable,
			Columns: []string{step.ProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: procedure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Step{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{step.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}