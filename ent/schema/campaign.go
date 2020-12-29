package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Campaign holds the schema definition for the Campaign entity.
type Campaign struct {
	ent.Schema
}

// Fields of the Campaign.
func (Campaign) Fields() []ent.Field {
	return []ent.Field{
		idField("CA"),
		field.String("name"),
	}
}

// Edges of the Campaign.
func (Campaign) Edges() []ent.Edge {
	return nil
}
