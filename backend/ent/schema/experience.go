package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Experience holds the schema definition for the Experience entity.
type Experience struct {
	ent.Schema
}

// Fields of the Experience.
func (Experience) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).Immutable(),
		field.String("companyName").NotEmpty(),
		field.String("position").NotEmpty(),
		field.Time("startDate"),
		field.Time("endDate").Optional().Nillable(),
		field.Bool("isCurrent").Default(false),
		field.Text("description").NotEmpty(),
		field.String("location").Optional(),
		field.JSON("acheivements", map[string]interface{}{}).Optional(),
		field.JSON("technologiesUsed", map[string]interface{}{}).Optional(),
		field.Int("orderIndex").Default(0),
	}
}

// Edges of the Experience.
func (Experience) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("experiences").Unique().Required(),
	}
}

func (Experience) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("companyName"),
	}
}
