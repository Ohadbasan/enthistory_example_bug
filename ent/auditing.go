// Code generated by enthistory, DO NOT EDIT.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	"entdemo/ent/testentityhistory"

	"github.com/flume/enthistory"
)

type Change struct {
	FieldName string
	Old       any
	New       any
}

func NewChange(fieldName string, old, new any) Change {
	return Change{
		FieldName: fieldName,
		Old:       old,
		New:       new,
	}
}

type HistoryDiff[T any] struct {
	Old     *T
	New     *T
	Changes []Change
}

var (
	MismatchedRefError = errors.New("cannot take diff of histories with different Refs")
)

func (teh *TestEntityHistory) changes(new *TestEntityHistory) []Change {
	var changes []Change
	if !reflect.DeepEqual(teh.Status, new.Status) {
		changes = append(changes, NewChange(testentityhistory.FieldStatus, teh.Status, new.Status))
	}
	return changes
}

func (teh *TestEntityHistory) Diff(history *TestEntityHistory) (*HistoryDiff[TestEntityHistory], error) {
	if teh.Ref != history.Ref {
		return nil, MismatchedRefError
	}
	if teh.HistoryTime.UnixMilli() > history.HistoryTime.UnixMilli() || (teh.HistoryTime.UnixMilli() == history.HistoryTime.UnixMilli() && teh.ID > history.ID) {
		return &HistoryDiff[TestEntityHistory]{
			Old:     history,
			New:     teh,
			Changes: history.changes(teh),
		}, nil
	}
	return &HistoryDiff[TestEntityHistory]{
		Old:     teh,
		New:     history,
		Changes: teh.changes(history),
	}, nil
}

func (c Change) String(op enthistory.OpType) string {
	var newstr, oldstr string
	if c.New != nil {
		val, err := json.Marshal(c.New)
		if err != nil {
			newstr = fmt.Sprintf("%#v", c.New)
		} else {
			newstr = string(val)
		}
	}
	if c.Old != nil {
		val, err := json.Marshal(c.Old)
		if err != nil {
			oldstr = fmt.Sprintf("%#v", c.Old)
		} else {
			oldstr = string(val)
		}
	}
	switch op {
	case enthistory.OpTypeInsert:
		return fmt.Sprintf("%s: %s", c.FieldName, newstr)
	case enthistory.OpTypeDelete:
		return fmt.Sprintf("%s: %s", c.FieldName, oldstr)
	default:
		return fmt.Sprintf("%s: %s -> %s", c.FieldName, oldstr, newstr)
	}
}

func (c *Client) Audit(ctx context.Context) ([][]string, error) {
	records := [][]string{
		{"Table", "Ref Id", "History Time", "Operation", "Changes", "Updated By"},
	}
	var rec [][]string
	var err error
	rec, err = auditTestEntityHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, rec...)

	return records, nil
}

type record struct {
	Table       string
	RefId       any
	HistoryTime time.Time
	Operation   enthistory.OpType
	Changes     []Change
	UpdatedBy   *int
}

func (r *record) toRow() []string {
	row := make([]string, 6)

	row[0] = r.Table
	row[1] = fmt.Sprintf("%v", r.RefId)
	row[2] = r.HistoryTime.Format(time.ANSIC)
	row[3] = r.Operation.String()
	for i, change := range r.Changes {
		if i == 0 {
			row[4] = change.String(r.Operation)
			continue
		}
		row[4] = fmt.Sprintf("%s\n%s", row[4], change.String(r.Operation))
	}
	if r.UpdatedBy != nil {
		row[5] = fmt.Sprintf("%v", *r.UpdatedBy)
	}
	return row
}

type testentityhistoryref struct {
	Ref int
}

func auditTestEntityHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []testentityhistoryref
	client := NewTestEntityHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(testentityhistory.ByHistoryTime()).
		Select(testentityhistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(testentityhistory.Ref(currRef.Ref)).
			Order(testentityhistory.ByHistoryTime()).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			r := record{
				Table:       "TestEntityHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				r.Changes = (&TestEntityHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				r.Changes = curr.changes(&TestEntityHistory{})
			default:
				if i == 0 {
					r.Changes = (&TestEntityHistory{}).changes(curr)
				} else {
					r.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, r.toRow())
		}
	}
	return records, nil
}