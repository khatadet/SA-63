package schema

import (
    
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{




		field.String("UserEmail").Unique(),

        field.String("NAME"),
			

		
		
    }
}

// Edges of the USER.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		
		edge.To("UserRent", Rent.Type),

		edge.From("UserUserstatus", UserStatus.Type).
            Ref("UserstatusUser").
			Unique(),
    }
}
