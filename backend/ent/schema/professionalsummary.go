package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// ProfessionalSummary holds the schema definition for the ProfessionalSummary entity.
type ProfessionalSummary struct {
	ent.Schema
}

// Fields of the ProfessionalSummary.
func (ProfessionalSummary) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).Immutable(),
		field.UUID("resumeId", uuid.UUID{}),
		field.Text("summary").NotEmpty(),
		field.Int("yearsOfExperience").Optional().Nillable(),
		field.JSON("keyStrengths", []string{}).Optional(),
		field.Text("careerObjective").Optional(),
	}
}

// Edges of the ProfessionalSummary.
func (ProfessionalSummary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("professionalSummary").Field("resumeId").Unique().Required(),
	}
}

func (ProfessionalSummary) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("yearsOfExperience"),
	}
}
