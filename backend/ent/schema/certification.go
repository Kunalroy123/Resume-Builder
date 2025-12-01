package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Certification holds the schema definition for the Certification entity.
type Certification struct {
	ent.Schema
}

// Fields of the Certification.
func (Certification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).Immutable(),
		field.String("name").Optional(),
		field.String("issuingOrganization").Optional(),
		field.Time("issueDate").Optional().Nillable(),
		field.Time("expiryDate").Optional().Nillable(),
		field.String("credentialId").Optional(),
		field.String("credentialUrl").Optional(),
		field.Text("description").Optional(),
		field.Int("orderIndex").Default(0),
	}
}

// Edges of the Certification.
func (Certification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("certifications").Unique().Required(),
	}
}

func (Certification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("name"),
		index.Fields("credentialId"),
		index.Fields("issueDate"),
		index.Fields("expiryDate"),
	}
}
