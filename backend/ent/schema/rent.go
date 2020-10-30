package schema

import (
	"github.com/facebookincubator/ent"
	//"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Rent holds the schema definition for the Rent entity.
type Rent struct {
	ent.Schema
}

// Fields of the Rent.
func (Rent) Fields() []ent.Field {
	return []ent.Field{

		field.String("RentAge").NotEmpty(),

		field.String("CidUser").NotEmpty(),
		
		field.String("RentDate").NotEmpty(),
	}
}

// Edges of the Rent.
func (Rent) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("RentRoom", Room.Type).
			Ref("RoomRent").
			Unique(),

		edge.From("RentUser", User.Type).
			Ref("UserRent").
			Unique(),

		edge.From("RentInsurance", Insurance.Type).
			Ref("InsuranceRent").
			Unique(),

		edge.From("RentRoomage", Roomage.Type).
			Ref("RoomageRent").
			Unique(),
	}
}
