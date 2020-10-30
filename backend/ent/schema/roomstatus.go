package schema

import (
    
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// RoomStatus holds the schema definition for the RoomStatus entity.
type RoomStatus struct {
	ent.Schema
}

// Fields of the RoomStatus.
func (RoomStatus) Fields() []ent.Field {
	return []ent.Field{

		field.String("RoomStatus").Unique(),
	
    }
}

// Edges of the RoomStatus.
func (RoomStatus) Edges() []ent.Edge {
	return []ent.Edge{
		
		edge.To("RoomstatusRoom", Room.Type),
    }
}
