// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/testentityhistory"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/flume/enthistory"
)

// TestEntityHistory is the model entity for the TestEntityHistory schema.
type TestEntityHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref int `json:"ref,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy *int `json:"updated_by,omitempty"`
	// Status holds the value of the "status" field.
	Status       testentityhistory.Status `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestEntityHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case testentityhistory.FieldID, testentityhistory.FieldRef, testentityhistory.FieldUpdatedBy:
			values[i] = new(sql.NullInt64)
		case testentityhistory.FieldOperation, testentityhistory.FieldStatus:
			values[i] = new(sql.NullString)
		case testentityhistory.FieldHistoryTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestEntityHistory fields.
func (teh *TestEntityHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testentityhistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			teh.ID = int(value.Int64)
		case testentityhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				teh.HistoryTime = value.Time
			}
		case testentityhistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				teh.Operation = enthistory.OpType(value.String)
			}
		case testentityhistory.FieldRef:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				teh.Ref = int(value.Int64)
			}
		case testentityhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				teh.UpdatedBy = new(int)
				*teh.UpdatedBy = int(value.Int64)
			}
		case testentityhistory.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				teh.Status = testentityhistory.Status(value.String)
			}
		default:
			teh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TestEntityHistory.
// This includes values selected through modifiers, order, etc.
func (teh *TestEntityHistory) Value(name string) (ent.Value, error) {
	return teh.selectValues.Get(name)
}

// Update returns a builder for updating this TestEntityHistory.
// Note that you need to call TestEntityHistory.Unwrap() before calling this method if this TestEntityHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (teh *TestEntityHistory) Update() *TestEntityHistoryUpdateOne {
	return NewTestEntityHistoryClient(teh.config).UpdateOne(teh)
}

// Unwrap unwraps the TestEntityHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (teh *TestEntityHistory) Unwrap() *TestEntityHistory {
	_tx, ok := teh.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestEntityHistory is not a transactional entity")
	}
	teh.config.driver = _tx.drv
	return teh
}

// String implements the fmt.Stringer.
func (teh *TestEntityHistory) String() string {
	var builder strings.Builder
	builder.WriteString("TestEntityHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", teh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(teh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", teh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", teh.Ref))
	builder.WriteString(", ")
	if v := teh.UpdatedBy; v != nil {
		builder.WriteString("updated_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", teh.Status))
	builder.WriteByte(')')
	return builder.String()
}

// TestEntityHistories is a parsable slice of TestEntityHistory.
type TestEntityHistories []*TestEntityHistory