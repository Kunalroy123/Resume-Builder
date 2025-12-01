package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Resume holds the schema definition for the Resume entity.
type Resume struct {
	ent.Schema
}

// Fields of the Resume.
func (Resume) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).StorageKey("id").Immutable(),
		field.UUID("userId", uuid.UUID{}),
		field.String("title").NotEmpty(),
		field.UUID("templateId", uuid.UUID{}).Optional().Nillable(),
		field.Bool("isAiGenerated").Default(false),
		field.Bool("isPublic").Default(false),
		field.JSON("content", map[string]interface{}{}).Optional(),
		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Resume.
func (Resume) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("resumes").Field("userId").Unique().Required(),
		edge.From("template", Template.Type).Ref("resumes").Field("templateId").Unique(),
		edge.To("headerContanctInfo", HeaderContactInfo.Type).Unique(),
		edge.To("professionalSummary", ProfessionalSummary.Type).Unique(),
		edge.To("experiences", Experience.Type),
		edge.To("educations", Education.Type),
		edge.To("skills", Skill.Type),
		edge.To("projects", Project.Type),
		edge.To("certifications", Certification.Type),
		edge.To("achievements", Achievement.Type),
		edge.To("hobbies", Hobby.Type),
	}
}

func (Resume) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("userId").Unique(),
		index.Fields("templateId"),
		index.Fields("isAiGenerated"),
		index.Fields("isPublic"),
		index.Fields("createdAt"),
		index.Fields("userId", "isPublic"),
	}
}
