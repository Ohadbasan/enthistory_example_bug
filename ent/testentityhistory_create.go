// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/testentityhistory"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/flume/enthistory"
)

// TestEntityHistoryCreate is the builder for creating a TestEntityHistory entity.
type TestEntityHistoryCreate struct {
	config
	mutation *TestEntityHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (tehc *TestEntityHistoryCreate) SetHistoryTime(t time.Time) *TestEntityHistoryCreate {
	tehc.mutation.SetHistoryTime(t)
	return tehc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (tehc *TestEntityHistoryCreate) SetNillableHistoryTime(t *time.Time) *TestEntityHistoryCreate {
	if t != nil {
		tehc.SetHistoryTime(*t)
	}
	return tehc
}

// SetOperation sets the "operation" field.
func (tehc *TestEntityHistoryCreate) SetOperation(et enthistory.OpType) *TestEntityHistoryCreate {
	tehc.mutation.SetOperation(et)
	return tehc
}

// SetRef sets the "ref" field.
func (tehc *TestEntityHistoryCreate) SetRef(i int) *TestEntityHistoryCreate {
	tehc.mutation.SetRef(i)
	return tehc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (tehc *TestEntityHistoryCreate) SetNillableRef(i *int) *TestEntityHistoryCreate {
	if i != nil {
		tehc.SetRef(*i)
	}
	return tehc
}

// SetUpdatedBy sets the "updated_by" field.
func (tehc *TestEntityHistoryCreate) SetUpdatedBy(i int) *TestEntityHistoryCreate {
	tehc.mutation.SetUpdatedBy(i)
	return tehc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tehc *TestEntityHistoryCreate) SetNillableUpdatedBy(i *int) *TestEntityHistoryCreate {
	if i != nil {
		tehc.SetUpdatedBy(*i)
	}
	return tehc
}

// SetStatus sets the "status" field.
func (tehc *TestEntityHistoryCreate) SetStatus(t testentityhistory.Status) *TestEntityHistoryCreate {
	tehc.mutation.SetStatus(t)
	return tehc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tehc *TestEntityHistoryCreate) SetNillableStatus(t *testentityhistory.Status) *TestEntityHistoryCreate {
	if t != nil {
		tehc.SetStatus(*t)
	}
	return tehc
}

// Mutation returns the TestEntityHistoryMutation object of the builder.
func (tehc *TestEntityHistoryCreate) Mutation() *TestEntityHistoryMutation {
	return tehc.mutation
}

// Save creates the TestEntityHistory in the database.
func (tehc *TestEntityHistoryCreate) Save(ctx context.Context) (*TestEntityHistory, error) {
	tehc.defaults()
	return withHooks(ctx, tehc.sqlSave, tehc.mutation, tehc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tehc *TestEntityHistoryCreate) SaveX(ctx context.Context) *TestEntityHistory {
	v, err := tehc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tehc *TestEntityHistoryCreate) Exec(ctx context.Context) error {
	_, err := tehc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tehc *TestEntityHistoryCreate) ExecX(ctx context.Context) {
	if err := tehc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tehc *TestEntityHistoryCreate) defaults() {
	if _, ok := tehc.mutation.HistoryTime(); !ok {
		v := testentityhistory.DefaultHistoryTime()
		tehc.mutation.SetHistoryTime(v)
	}
	if _, ok := tehc.mutation.Status(); !ok {
		v := testentityhistory.DefaultStatus
		tehc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tehc *TestEntityHistoryCreate) check() error {
	if _, ok := tehc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`ent: missing required field "TestEntityHistory.history_time"`)}
	}
	if _, ok := tehc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "TestEntityHistory.operation"`)}
	}
	if v, ok := tehc.mutation.Operation(); ok {
		if err := testentityhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "TestEntityHistory.operation": %w`, err)}
		}
	}
	if _, ok := tehc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "TestEntityHistory.status"`)}
	}
	if v, ok := tehc.mutation.Status(); ok {
		if err := testentityhistory.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "TestEntityHistory.status": %w`, err)}
		}
	}
	return nil
}

func (tehc *TestEntityHistoryCreate) sqlSave(ctx context.Context) (*TestEntityHistory, error) {
	if err := tehc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tehc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tehc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tehc.mutation.id = &_node.ID
	tehc.mutation.done = true
	return _node, nil
}

func (tehc *TestEntityHistoryCreate) createSpec() (*TestEntityHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &TestEntityHistory{config: tehc.config}
		_spec = sqlgraph.NewCreateSpec(testentityhistory.Table, sqlgraph.NewFieldSpec(testentityhistory.FieldID, field.TypeInt))
	)
	if value, ok := tehc.mutation.HistoryTime(); ok {
		_spec.SetField(testentityhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := tehc.mutation.Operation(); ok {
		_spec.SetField(testentityhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := tehc.mutation.Ref(); ok {
		_spec.SetField(testentityhistory.FieldRef, field.TypeInt, value)
		_node.Ref = value
	}
	if value, ok := tehc.mutation.UpdatedBy(); ok {
		_spec.SetField(testentityhistory.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = &value
	}
	if value, ok := tehc.mutation.Status(); ok {
		_spec.SetField(testentityhistory.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	return _node, _spec
}

// TestEntityHistoryCreateBulk is the builder for creating many TestEntityHistory entities in bulk.
type TestEntityHistoryCreateBulk struct {
	config
	err      error
	builders []*TestEntityHistoryCreate
}

// Save creates the TestEntityHistory entities in the database.
func (tehcb *TestEntityHistoryCreateBulk) Save(ctx context.Context) ([]*TestEntityHistory, error) {
	if tehcb.err != nil {
		return nil, tehcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tehcb.builders))
	nodes := make([]*TestEntityHistory, len(tehcb.builders))
	mutators := make([]Mutator, len(tehcb.builders))
	for i := range tehcb.builders {
		func(i int, root context.Context) {
			builder := tehcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestEntityHistoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tehcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tehcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tehcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tehcb *TestEntityHistoryCreateBulk) SaveX(ctx context.Context) []*TestEntityHistory {
	v, err := tehcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tehcb *TestEntityHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := tehcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tehcb *TestEntityHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := tehcb.Exec(ctx); err != nil {
		panic(err)
	}
}