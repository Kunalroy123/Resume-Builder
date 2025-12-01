package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).Immutable(),
		field.String("title").Optional(),
		field.Text("description").Optional(),
		field.Time("startDate").Optional().Nillable(),
		field.Time("endDate").Optional().Nillable(),
		field.String("projectUrl").Optional(),
		field.String("githubUrl").Optional(),
		field.String("demoUrl").Optional(),
		field.JSON("technologiesUsed", map[string]interface{}{}).Optional(),
		field.JSON("keyFeatures", map[string]interface{}{}).Optional(),
		field.String("role").Optional(),
		field.Int("teamSize").Optional().Nillable(),
		field.Int("orderIndex").Default(0),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("projects").Unique().Required(),
	}
}

func (Project) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("projectUrl"),
	}
}
