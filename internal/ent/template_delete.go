// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TemplateDelete is the builder for deleting a Template entity.
type TemplateDelete struct {
	config
	hooks      []Hook
	mutation   *TemplateMutation
	predicates []predicate.Template
}

// Where adds a new predicate to the delete builder.
func (td *TemplateDelete) Where(ps ...predicate.Template) *TemplateDelete {
	td.predicates = append(td.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TemplateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(td.hooks) == 0 {
		affected, err = td.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			td.mutation = mutation
			affected, err = td.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(td.hooks) - 1; i >= 0; i-- {
			mut = td.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, td.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TemplateDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TemplateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: template.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: template.FieldID,
			},
		},
	}
	if ps := td.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, td.driver, _spec)
}

// TemplateDeleteOne is the builder for deleting a single Template entity.
type TemplateDeleteOne struct {
	td *TemplateDelete
}

// Exec executes the deletion query.
func (tdo *TemplateDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{template.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TemplateDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}