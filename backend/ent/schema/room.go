package schema

import (
    
    "github.com/facebookincubator/ent"
	
	"github.com/facebookincubator/ent/schema/edge"
)
// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{



		



	
    }
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
        	

			edge.From("RoomRoomstatus", RoomStatus.Type).
            Ref("RoomstatusRoom").
			Unique(),
			
			edge.From("RoomRoomtype", RoomType.Type).
            Ref("RoomtypeRoom").
			Unique(),
			
			edge.To("RoomRent", Rent.Type),
    }
}
