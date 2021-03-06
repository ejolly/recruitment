// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/admin"
	"github.com/empiricaly/recruitment/internal/ent/participant"
	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/ent/template"
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

// SetUpdatedAt sets the updated_at field.
func (pu *ProjectUpdate) SetUpdatedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetProjectID sets the projectID field.
func (pu *ProjectUpdate) SetProjectID(s string) *ProjectUpdate {
	pu.mutation.SetProjectID(s)
	return pu
}

// SetName sets the name field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// AddRunIDs adds the runs edge to Run by ids.
func (pu *ProjectUpdate) AddRunIDs(ids ...string) *ProjectUpdate {
	pu.mutation.AddRunIDs(ids...)
	return pu
}

// AddRuns adds the runs edges to Run.
func (pu *ProjectUpdate) AddRuns(r ...*Run) *ProjectUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRunIDs(ids...)
}

// AddTemplateIDs adds the templates edge to Template by ids.
func (pu *ProjectUpdate) AddTemplateIDs(ids ...string) *ProjectUpdate {
	pu.mutation.AddTemplateIDs(ids...)
	return pu
}

// AddTemplates adds the templates edges to Template.
func (pu *ProjectUpdate) AddTemplates(t ...*Template) *ProjectUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTemplateIDs(ids...)
}

// AddParticipantIDs adds the participants edge to Participant by ids.
func (pu *ProjectUpdate) AddParticipantIDs(ids ...string) *ProjectUpdate {
	pu.mutation.AddParticipantIDs(ids...)
	return pu
}

// AddParticipants adds the participants edges to Participant.
func (pu *ProjectUpdate) AddParticipants(p ...*Participant) *ProjectUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddParticipantIDs(ids...)
}

// SetOwnerID sets the owner edge to Admin by id.
func (pu *ProjectUpdate) SetOwnerID(id string) *ProjectUpdate {
	pu.mutation.SetOwnerID(id)
	return pu
}

// SetNillableOwnerID sets the owner edge to Admin by id if the given value is not nil.
func (pu *ProjectUpdate) SetNillableOwnerID(id *string) *ProjectUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the owner edge to Admin.
func (pu *ProjectUpdate) SetOwner(a *Admin) *ProjectUpdate {
	return pu.SetOwnerID(a.ID)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearRuns clears all "runs" edges to type Run.
func (pu *ProjectUpdate) ClearRuns() *ProjectUpdate {
	pu.mutation.ClearRuns()
	return pu
}

// RemoveRunIDs removes the runs edge to Run by ids.
func (pu *ProjectUpdate) RemoveRunIDs(ids ...string) *ProjectUpdate {
	pu.mutation.RemoveRunIDs(ids...)
	return pu
}

// RemoveRuns removes runs edges to Run.
func (pu *ProjectUpdate) RemoveRuns(r ...*Run) *ProjectUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRunIDs(ids...)
}

// ClearTemplates clears all "templates" edges to type Template.
func (pu *ProjectUpdate) ClearTemplates() *ProjectUpdate {
	pu.mutation.ClearTemplates()
	return pu
}

// RemoveTemplateIDs removes the templates edge to Template by ids.
func (pu *ProjectUpdate) RemoveTemplateIDs(ids ...string) *ProjectUpdate {
	pu.mutation.RemoveTemplateIDs(ids...)
	return pu
}

// RemoveTemplates removes templates edges to Template.
func (pu *ProjectUpdate) RemoveTemplates(t ...*Template) *ProjectUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTemplateIDs(ids...)
}

// ClearParticipants clears all "participants" edges to type Participant.
func (pu *ProjectUpdate) ClearParticipants() *ProjectUpdate {
	pu.mutation.ClearParticipants()
	return pu
}

// RemoveParticipantIDs removes the participants edge to Participant by ids.
func (pu *ProjectUpdate) RemoveParticipantIDs(ids ...string) *ProjectUpdate {
	pu.mutation.RemoveParticipantIDs(ids...)
	return pu
}

// RemoveParticipants removes participants edges to Participant.
func (pu *ProjectUpdate) RemoveParticipants(p ...*Participant) *ProjectUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveParticipantIDs(ids...)
}

// ClearOwner clears the "owner" edge to type Admin.
func (pu *ProjectUpdate) ClearOwner() *ProjectUpdate {
	pu.mutation.ClearOwner()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
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

// defaults sets the default values of the builder before save.
func (pu *ProjectUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := project.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
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
	if value, ok := pu.mutation.ProjectID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldProjectID,
		})
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
	}
	if pu.mutation.RunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RunsTable,
			Columns: []string{project.RunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: run.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRunsIDs(); len(nodes) > 0 && !pu.mutation.RunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RunsTable,
			Columns: []string{project.RunsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RunsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RunsTable,
			Columns: []string{project.RunsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.TemplatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TemplatesTable,
			Columns: []string{project.TemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: template.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTemplatesIDs(); len(nodes) > 0 && !pu.mutation.TemplatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TemplatesTable,
			Columns: []string{project.TemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: template.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TemplatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TemplatesTable,
			Columns: []string{project.TemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: template.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.ParticipantsTable,
			Columns: project.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedParticipantsIDs(); len(nodes) > 0 && !pu.mutation.ParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.ParticipantsTable,
			Columns: project.ParticipantsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.ParticipantsTable,
			Columns: project.ParticipantsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.OwnerTable,
			Columns: []string{project.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: admin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.OwnerTable,
			Columns: []string{project.OwnerColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
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

// SetUpdatedAt sets the updated_at field.
func (puo *ProjectUpdateOne) SetUpdatedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetProjectID sets the projectID field.
func (puo *ProjectUpdateOne) SetProjectID(s string) *ProjectUpdateOne {
	puo.mutation.SetProjectID(s)
	return puo
}

// SetName sets the name field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// AddRunIDs adds the runs edge to Run by ids.
func (puo *ProjectUpdateOne) AddRunIDs(ids ...string) *ProjectUpdateOne {
	puo.mutation.AddRunIDs(ids...)
	return puo
}

// AddRuns adds the runs edges to Run.
func (puo *ProjectUpdateOne) AddRuns(r ...*Run) *ProjectUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRunIDs(ids...)
}

// AddTemplateIDs adds the templates edge to Template by ids.
func (puo *ProjectUpdateOne) AddTemplateIDs(ids ...string) *ProjectUpdateOne {
	puo.mutation.AddTemplateIDs(ids...)
	return puo
}

// AddTemplates adds the templates edges to Template.
func (puo *ProjectUpdateOne) AddTemplates(t ...*Template) *ProjectUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTemplateIDs(ids...)
}

// AddParticipantIDs adds the participants edge to Participant by ids.
func (puo *ProjectUpdateOne) AddParticipantIDs(ids ...string) *ProjectUpdateOne {
	puo.mutation.AddParticipantIDs(ids...)
	return puo
}

// AddParticipants adds the participants edges to Participant.
func (puo *ProjectUpdateOne) AddParticipants(p ...*Participant) *ProjectUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddParticipantIDs(ids...)
}

// SetOwnerID sets the owner edge to Admin by id.
func (puo *ProjectUpdateOne) SetOwnerID(id string) *ProjectUpdateOne {
	puo.mutation.SetOwnerID(id)
	return puo
}

// SetNillableOwnerID sets the owner edge to Admin by id if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableOwnerID(id *string) *ProjectUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the owner edge to Admin.
func (puo *ProjectUpdateOne) SetOwner(a *Admin) *ProjectUpdateOne {
	return puo.SetOwnerID(a.ID)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearRuns clears all "runs" edges to type Run.
func (puo *ProjectUpdateOne) ClearRuns() *ProjectUpdateOne {
	puo.mutation.ClearRuns()
	return puo
}

// RemoveRunIDs removes the runs edge to Run by ids.
func (puo *ProjectUpdateOne) RemoveRunIDs(ids ...string) *ProjectUpdateOne {
	puo.mutation.RemoveRunIDs(ids...)
	return puo
}

// RemoveRuns removes runs edges to Run.
func (puo *ProjectUpdateOne) RemoveRuns(r ...*Run) *ProjectUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRunIDs(ids...)
}

// ClearTemplates clears all "templates" edges to type Template.
func (puo *ProjectUpdateOne) ClearTemplates() *ProjectUpdateOne {
	puo.mutation.ClearTemplates()
	return puo
}

// RemoveTemplateIDs removes the templates edge to Template by ids.
func (puo *ProjectUpdateOne) RemoveTemplateIDs(ids ...string) *ProjectUpdateOne {
	puo.mutation.RemoveTemplateIDs(ids...)
	return puo
}

// RemoveTemplates removes templates edges to Template.
func (puo *ProjectUpdateOne) RemoveTemplates(t ...*Template) *ProjectUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTemplateIDs(ids...)
}

// ClearParticipants clears all "participants" edges to type Participant.
func (puo *ProjectUpdateOne) ClearParticipants() *ProjectUpdateOne {
	puo.mutation.ClearParticipants()
	return puo
}

// RemoveParticipantIDs removes the participants edge to Participant by ids.
func (puo *ProjectUpdateOne) RemoveParticipantIDs(ids ...string) *ProjectUpdateOne {
	puo.mutation.RemoveParticipantIDs(ids...)
	return puo
}

// RemoveParticipants removes participants edges to Participant.
func (puo *ProjectUpdateOne) RemoveParticipants(p ...*Participant) *ProjectUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveParticipantIDs(ids...)
}

// ClearOwner clears the "owner" edge to type Admin.
func (puo *ProjectUpdateOne) ClearOwner() *ProjectUpdateOne {
	puo.mutation.ClearOwner()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	var (
		err  error
		node *Project
	)
	puo.defaults()
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
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
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

// defaults sets the default values of the builder before save.
func (puo *ProjectUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := project.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
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
	if value, ok := puo.mutation.ProjectID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldProjectID,
		})
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: project.FieldName,
		})
	}
	if puo.mutation.RunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RunsTable,
			Columns: []string{project.RunsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: run.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRunsIDs(); len(nodes) > 0 && !puo.mutation.RunsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RunsTable,
			Columns: []string{project.RunsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RunsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.RunsTable,
			Columns: []string{project.RunsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.TemplatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TemplatesTable,
			Columns: []string{project.TemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: template.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTemplatesIDs(); len(nodes) > 0 && !puo.mutation.TemplatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TemplatesTable,
			Columns: []string{project.TemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: template.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TemplatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.TemplatesTable,
			Columns: []string{project.TemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: template.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.ParticipantsTable,
			Columns: project.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedParticipantsIDs(); len(nodes) > 0 && !puo.mutation.ParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.ParticipantsTable,
			Columns: project.ParticipantsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   project.ParticipantsTable,
			Columns: project.ParticipantsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.OwnerTable,
			Columns: []string{project.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: admin.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   project.OwnerTable,
			Columns: []string{project.OwnerColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
