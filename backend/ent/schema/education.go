package schema

import "entgo.io/ent"

// Education holds the schema definition for the Education entity.
type Education struct {
	ent.Schema
}

// Fields of the Education.
func (Education) Fields() []ent.Field {
	return nil
}

// Edges of the Education.
func (Education) Edges() []ent.Edge {
	return nil
}
