package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Apartment holds the schema definition for the Apartment entity.
type Apartment struct {
	ent.Schema
}

// Fields of the Apartment.
func (Apartment) Fields() []ent.Field {
	return []ent.Field{
		field.Int16("area"),
		field.Int16("floor"),
		field.Int16("number"),
	}
}

// Edges of the Apartment.
func (Apartment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).Unique(),
	}
}
