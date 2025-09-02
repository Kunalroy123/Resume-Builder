package schema

import "entgo.io/ent"

// ProfessionalSummary holds the schema definition for the ProfessionalSummary entity.
type ProfessionalSummary struct {
	ent.Schema
}

// Fields of the ProfessionalSummary.
func (ProfessionalSummary) Fields() []ent.Field {
	return nil
}

// Edges of the ProfessionalSummary.
func (ProfessionalSummary) Edges() []ent.Edge {
	return nil
}
