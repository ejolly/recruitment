// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/empiricaly/recruitment/internal/ent"
)

// The AdminFunc type is an adapter to allow the use of ordinary
// function as Admin mutator.
type AdminFunc func(context.Context, *ent.AdminMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AdminFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AdminMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AdminMutation", m)
	}
	return f(ctx, mv)
}

// The ParticipantFunc type is an adapter to allow the use of ordinary
// function as Participant mutator.
type ParticipantFunc func(context.Context, *ent.ParticipantMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ParticipantFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ParticipantMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ParticipantMutation", m)
	}
	return f(ctx, mv)
}

// The ParticipationFunc type is an adapter to allow the use of ordinary
// function as Participation mutator.
type ParticipationFunc func(context.Context, *ent.ParticipationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ParticipationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ParticipationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ParticipationMutation", m)
	}
	return f(ctx, mv)
}

// The ProjectFunc type is an adapter to allow the use of ordinary
// function as Project mutator.
type ProjectFunc func(context.Context, *ent.ProjectMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProjectFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProjectMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProjectMutation", m)
	}
	return f(ctx, mv)
}

// The ProviderIDFunc type is an adapter to allow the use of ordinary
// function as ProviderID mutator.
type ProviderIDFunc func(context.Context, *ent.ProviderIDMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProviderIDFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProviderIDMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProviderIDMutation", m)
	}
	return f(ctx, mv)
}

// The RunFunc type is an adapter to allow the use of ordinary
// function as Run mutator.
type RunFunc func(context.Context, *ent.RunMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RunFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RunMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RunMutation", m)
	}
	return f(ctx, mv)
}

// The StepFunc type is an adapter to allow the use of ordinary
// function as Step mutator.
type StepFunc func(context.Context, *ent.StepMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StepFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StepMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StepMutation", m)
	}
	return f(ctx, mv)
}

// The StepRunFunc type is an adapter to allow the use of ordinary
// function as StepRun mutator.
type StepRunFunc func(context.Context, *ent.StepRunMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StepRunFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StepRunMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StepRunMutation", m)
	}
	return f(ctx, mv)
}

// The TemplateFunc type is an adapter to allow the use of ordinary
// function as Template mutator.
type TemplateFunc func(context.Context, *ent.TemplateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TemplateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TemplateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TemplateMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	Hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(_ context.Context, m ent.Mutation) (ent.Value, error) {
			return nil, fmt.Errorf("%s operation is not allowed", m.Op())
		})
	}
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
