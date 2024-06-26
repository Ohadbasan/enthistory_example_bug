// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/testentity"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// TestEntity is the model entity for the TestEntity schema.
type TestEntity struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status       testentity.Status `json:"status,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestEntity) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case testentity.FieldID:
			values[i] = new(sql.NullInt64)
		case testentity.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestEntity fields.
func (te *TestEntity) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testentity.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			te.ID = int(value.Int64)
		case testentity.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				te.Status = testentity.Status(value.String)
			}
		default:
			te.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TestEntity.
// This includes values selected through modifiers, order, etc.
func (te *TestEntity) Value(name string) (ent.Value, error) {
	return te.selectValues.Get(name)
}

// Update returns a builder for updating this TestEntity.
// Note that you need to call TestEntity.Unwrap() before calling this method if this TestEntity
// was returned from a transaction, and the transaction was committed or rolled back.
func (te *TestEntity) Update() *TestEntityUpdateOne {
	return NewTestEntityClient(te.config).UpdateOne(te)
}

// Unwrap unwraps the TestEntity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (te *TestEntity) Unwrap() *TestEntity {
	_tx, ok := te.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestEntity is not a transactional entity")
	}
	te.config.driver = _tx.drv
	return te
}

// String implements the fmt.Stringer.
func (te *TestEntity) String() string {
	var builder strings.Builder
	builder.WriteString("TestEntity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", te.ID))
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", te.Status))
	builder.WriteByte(')')
	return builder.String()
}

// TestEntities is a parsable slice of TestEntity.
type TestEntities []*TestEntity
