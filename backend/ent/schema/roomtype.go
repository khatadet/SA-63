package schema

import (
    
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// RoomType holds the schema definition for the RoomType entity.
type RoomType struct {
	ent.Schema
}

// Fields of the ROOMTYPE.
func (RoomType) Fields() []ent.Field {
	return []ent.Field{


		field.String("RoomType").Unique(),
		field.Int("COST").Unique(),
    }
}

// Edges of the ROOMTYPE.
func (RoomType) Edges() []ent.Edge {
	return []ent.Edge{
		
		edge.To("RoomtypeRoom", Room.Type),
    }
}
