package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Direction holds the schema definition for the Direction entity.
type Direction struct {
	ent.Schema
}

// Fields of the Direction.
func (Direction) Fields() []ent.Field {
	return []ent.Field{
		field.String("description"),
		field.String("name"),
	}
}
