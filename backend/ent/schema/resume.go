package schema

import "entgo.io/ent"

// Resume holds the schema definition for the Resume entity.
type Resume struct {
	ent.Schema
}

// Fields of the Resume.
func (Resume) Fields() []ent.Field {
	return nil
}

// Edges of the Resume.
func (Resume) Edges() []ent.Edge {
	return nil
}
