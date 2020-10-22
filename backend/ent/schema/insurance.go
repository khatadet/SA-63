package schema

import (
    
    "github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Insurance holds the schema definition for the Insurance entity.
type Insurance struct {
	ent.Schema
}

// Fields of the INSURANCE.
func (Insurance) Fields() []ent.Field {
	return []ent.Field{


		field.Int("Insurance"),
    }
}

// Edges of the INSURANCE.
func (Insurance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("InsuranceRent", Rent.Type),
	
    }
}
