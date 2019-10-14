package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Ingredient holds the schema definition for the Ingredient entity.
type Ingredient struct {
	ent.Schema
}

// Fields of the Ingredient.
func (Ingredient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("unitType"),
	}
}
