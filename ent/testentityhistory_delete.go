// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/predicate"
	"entdemo/ent/testentityhistory"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestEntityHistoryDelete is the builder for deleting a TestEntityHistory entity.
type TestEntityHistoryDelete struct {
	config
	hooks    []Hook
	mutation *TestEntityHistoryMutation
}

// Where appends a list predicates to the TestEntityHistoryDelete builder.
func (tehd *TestEntityHistoryDelete) Where(ps ...predicate.TestEntityHistory) *TestEntityHistoryDelete {
	tehd.mutation.Where(ps...)
	return tehd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tehd *TestEntityHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tehd.sqlExec, tehd.mutation, tehd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tehd *TestEntityHistoryDelete) ExecX(ctx context.Context) int {
	n, err := tehd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tehd *TestEntityHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(testentityhistory.Table, sqlgraph.NewFieldSpec(testentityhistory.FieldID, field.TypeInt))
	if ps := tehd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tehd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tehd.mutation.done = true
	return affected, err
}

// TestEntityHistoryDeleteOne is the builder for deleting a single TestEntityHistory entity.
type TestEntityHistoryDeleteOne struct {
	tehd *TestEntityHistoryDelete
}

// Where appends a list predicates to the TestEntityHistoryDelete builder.
func (tehdo *TestEntityHistoryDeleteOne) Where(ps ...predicate.TestEntityHistory) *TestEntityHistoryDeleteOne {
	tehdo.tehd.mutation.Where(ps...)
	return tehdo
}

// Exec executes the deletion query.
func (tehdo *TestEntityHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := tehdo.tehd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{testentityhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tehdo *TestEntityHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := tehdo.Exec(ctx); err != nil {
		panic(err)
	}
}
