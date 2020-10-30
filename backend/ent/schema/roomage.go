package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Roomage holds the schema definition for the Roomage entity.
type Roomage struct {
	ent.Schema
}

// Fields of the Roomage.
func (Roomage) Fields() []ent.Field {
	return []ent.Field{

		field.Int("RoomAge").Unique(),
		field.String("Text").Unique(),
	}
}

// Edges of the ROOMAGE.
func (Roomage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("RoomageRent", Rent.Type),
	}
}
