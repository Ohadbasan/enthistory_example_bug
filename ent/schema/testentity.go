package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TestEntity holds the schema definition for the TestEntity entity.
type TestEntity struct {
	ent.Schema
}

// Fields of the TestEntity.
func (TestEntity) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").Values("started", "success").Default("started"),
	}
}

// Edges of the TestEntity.
func (TestEntity) Edges() []ent.Edge {
	return nil
}
