// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/ent/predicate"
	"github.com/empiricaly/recruitment/ent/project"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks      []Hook
	mutation   *ProjectMutation
	predicates []predicate.Project
}

// Where adds a new predicate for the builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetUpdatedAt sets the updatedAt field.
func (pu *ProjectUpdate) SetUpdatedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetStartAt sets the startAt field.
func (pu *ProjectUpdate) SetStartAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetStartAt(t)
	return pu
}

// SetStartedAt sets the startedAt field.
func (pu *ProjectUpdate) SetStartedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetStartedAt(t)
	return pu
}

// SetEndedAt sets the endedAt field.
func (pu *ProjectUpdate) SetEndedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetEndedAt(t)
	return pu
}

// SetError sets the error field.
func (pu *ProjectUpdate) SetError(s string) *ProjectUpdate {
	pu.mutation.SetError(s)
	return pu
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := project.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   project.Table,
			Columns: project.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: project.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldStartAt,
		})
	}
	if value, ok := pu.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldStartedAt,
		})
	}
	if value, ok := pu.mutation.EndedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldEndedAt,
		})
	}
	if value, ok := pu.mutation.Error(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldError,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// SetUpdatedAt sets the updatedAt field.
func (puo *ProjectUpdateOne) SetUpdatedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetStartAt sets the startAt field.
func (puo *ProjectUpdateOne) SetStartAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetStartAt(t)
	return puo
}

// SetStartedAt sets the startedAt field.
func (puo *ProjectUpdateOne) SetStartedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetStartedAt(t)
	return puo
}

// SetEndedAt sets the endedAt field.
func (puo *ProjectUpdateOne) SetEndedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetEndedAt(t)
	return puo
}

// SetError sets the error field.
func (puo *ProjectUpdateOne) SetError(s string) *ProjectUpdateOne {
	puo.mutation.SetError(s)
	return puo
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// Save executes the query and returns the updated entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := project.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	var (
		err  error
		node *Project
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (pr *Project, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   project.Table,
			Columns: project.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: project.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Project.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldStartAt,
		})
	}
	if value, ok := puo.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldStartedAt,
		})
	}
	if value, ok := puo.mutation.EndedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: project.FieldEndedAt,
		})
	}
	if value, ok := puo.mutation.Error(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldError,
		})
	}
	pr = &Project{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
