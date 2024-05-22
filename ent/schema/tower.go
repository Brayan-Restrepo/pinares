package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tower holds the schema definition for the Tower entity.
type Tower struct {
	ent.Schema
}

// Fields of the Tower.
func (Tower) Fields() []ent.Field {

	return []ent.Field{
		field.Int("number_of_floors").
			Positive(),
		field.Int("number_of_apartments_per_floor").
			Positive(),
		field.String("name").
			Default("unknown"),
		field.Time("registered_at"),
	}
}

// Edges of the Tower.
func (Tower) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("apartments", Apartment.Type),
	}
}
