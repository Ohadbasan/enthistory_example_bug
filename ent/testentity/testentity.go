// Code generated by ent, DO NOT EDIT.

package testentity

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the testentity type in the database.
	Label = "test_entity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the testentity in the database.
	Table = "test_entities"
)

// Columns holds all SQL columns for testentity fields.
var Columns = []string{
	FieldID,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Status defines the type for the "status" enum field.
type Status string

// StatusStarted is the default value of the Status enum.
const DefaultStatus = StatusStarted

// Status values.
const (
	StatusStarted Status = "started"
	StatusSuccess Status = "success"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusStarted, StatusSuccess:
		return nil
	default:
		return fmt.Errorf("testentity: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the TestEntity queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}
