package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// HeaderContactInfo holds the schema definition for the HeaderContactInfo entity.
type HeaderContactInfo struct {
	ent.Schema
}

// Fields of the HeaderContactInfo.
func (HeaderContactInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).Immutable(),
		field.String("fullname").NotEmpty(),
		field.String("professionalTitle").Optional(),
		field.String("address").Optional(),
		field.String("phone").Optional(),
		field.String("email").NotEmpty(),
		field.String("city").Optional(),
		field.String("state").Optional(),
		field.String("zipCode").Optional(),
		field.String("country").Optional(),
		field.String("linkedinUrl").Optional(),
		field.String("githubUrl").Optional(),
		field.String("portfolioUrl").Optional(),
	}
}

// Edges of the HeaderContactInfo.
func (HeaderContactInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("headerContanctInfo").Unique().Required(),
	}
}

func (HeaderContactInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("phone"),
		index.Fields("email"),
	}
}
