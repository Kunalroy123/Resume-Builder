package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Education holds the schema definition for the Education entity.
type Education struct {
	ent.Schema
}

// Fields of the Education.
func (Education) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).Immutable(),
		field.String("institutionName").NotEmpty(),
		field.String("degree").NotEmpty(),
		field.String("fieldOfStudy").NotEmpty(),
		field.Time("startDate"),
		field.Time("endDate").Optional().Nillable(),
		field.String("gpa").Optional(),
		field.String("location").Optional(),
		field.JSON("relevantCoursework", map[string]interface{}{}).Optional(),
		field.Text("description").Optional(),
		field.Int("orderIndex").Default(0),
	}
}

// Edges of the Education.
func (Education) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("educations").Unique().Required(),
	}
}

func (Education) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("institutionName"),
		index.Fields("fieldOfStudy"),
	}
}
