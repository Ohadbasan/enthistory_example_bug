// Code generated by enthistory, DO NOT EDIT.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent"

	"github.com/flume/enthistory"
)

var (
	idNotFoundError = errors.New("could not get id from mutation")
)

func EntOpToHistoryOp(op ent.Op) enthistory.OpType {
	switch op {
	case ent.OpDelete, ent.OpDeleteOne:
		return enthistory.OpTypeDelete
	case ent.OpUpdate, ent.OpUpdateOne:
		return enthistory.OpTypeUpdate
	default:
		return enthistory.OpTypeInsert
	}
}

func rollback(tx *Tx, err error) error {
	if tx != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}
	return err
}

func (m *TestEntityMutation) CreateHistoryFromCreate(ctx context.Context) error {
	client := m.Client()
	tx, err := m.Tx()
	if err != nil {
		tx = nil
	}

	updatedBy, _ := ctx.Value("userId").(int)

	id, ok := m.ID()
	if !ok {
		return rollback(tx, idNotFoundError)
	}

	create := client.TestEntityHistory.Create()
	if tx != nil {
		create = tx.TestEntityHistory.Create()
	}

	create = create.
		SetOperation(EntOpToHistoryOp(m.Op())).
		SetHistoryTime(time.Now()).
		SetRef(id)
	if updatedBy != 0 {
		create = create.SetUpdatedBy(updatedBy)
	}

	if status, exists := m.Status(); exists {
		create = create.SetStatus(status)
	}

	_, err = create.Save(ctx)
	if err != nil {
		rollback(tx, err)
	}
	return nil
}

func (m *TestEntityMutation) CreateHistoryFromUpdate(ctx context.Context) error {
	client := m.Client()
	tx, err := m.Tx()
	if err != nil {
		tx = nil
	}

	updatedBy, _ := ctx.Value("userId").(int)

	ids, err := m.IDs(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("getting ids: %w", err))
	}

	for _, id := range ids {
		testentity, err := client.TestEntity.Get(ctx, id)
		if err != nil {
			return rollback(tx, err)
		}

		create := client.TestEntityHistory.Create()
		if tx != nil {
			create = tx.TestEntityHistory.Create()
		}

		create = create.
			SetOperation(EntOpToHistoryOp(m.Op())).
			SetHistoryTime(time.Now()).
			SetRef(id)
		if updatedBy != 0 {
			create = create.SetUpdatedBy(updatedBy)
		}

		if status, exists := m.Status(); exists {
			create = create.SetStatus(status)
		} else {
			create = create.SetStatus(testentity.Status)
		}

		_, err = create.Save(ctx)
		if err != nil {
			rollback(tx, err)
		}
	}

	return nil
}

func (m *TestEntityMutation) CreateHistoryFromDelete(ctx context.Context) error {
	client := m.Client()
	tx, err := m.Tx()
	if err != nil {
		tx = nil
	}

	updatedBy, _ := ctx.Value("userId").(int)

	ids, err := m.IDs(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("getting ids: %w", err))
	}

	for _, id := range ids {
		testentity, err := client.TestEntity.Get(ctx, id)
		if err != nil {
			return rollback(tx, err)
		}

		create := client.TestEntityHistory.Create()
		if tx != nil {
			create = tx.TestEntityHistory.Create()
		}
		if updatedBy != 0 {
			create = create.SetUpdatedBy(updatedBy)
		}

		_, err = create.
			SetOperation(EntOpToHistoryOp(m.Op())).
			SetHistoryTime(time.Now()).
			SetRef(id).
			SetStatus(testentity.Status).
			Save(ctx)
		if err != nil {
			rollback(tx, err)
		}
	}

	return nil
}
