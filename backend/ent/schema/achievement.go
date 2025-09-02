package schema

import "entgo.io/ent"

// Achievement holds the schema definition for the Achievement entity.
type Achievement struct {
	ent.Schema
}

// Fields of the Achievement.
func (Achievement) Fields() []ent.Field {
	return nil
}

// Edges of the Achievement.
func (Achievement) Edges() []ent.Edge {
	return nil
}
