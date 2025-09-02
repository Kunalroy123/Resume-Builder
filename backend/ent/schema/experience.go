package schema

import "entgo.io/ent"

// Experience holds the schema definition for the Experience entity.
type Experience struct {
	ent.Schema
}

// Fields of the Experience.
func (Experience) Fields() []ent.Field {
	return nil
}

// Edges of the Experience.
func (Experience) Edges() []ent.Edge {
	return nil
}
