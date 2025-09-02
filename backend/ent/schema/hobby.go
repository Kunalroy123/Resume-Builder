package schema

import "entgo.io/ent"

// Hobby holds the schema definition for the Hobby entity.
type Hobby struct {
	ent.Schema
}

// Fields of the Hobby.
func (Hobby) Fields() []ent.Field {
	return nil
}

// Edges of the Hobby.
func (Hobby) Edges() []ent.Edge {
	return nil
}
