package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Achievement holds the schema definition for the Achievement entity.
type Achievement struct {
	ent.Schema
}

// Fields of the Achievement.
func (Achievement) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).Immutable(),
		field.Text("discription").NotEmpty(),
		field.Time("dateAchieved").Optional().Nillable(),
		field.String("issuingOrganization").Optional(),
		field.Enum("achievementType").Values("AWARD", "RECOGNITION", "PUBLICATION", "PATENT", "PROMOTION", "PERFORMANCE", "LEADERSHIP", "OTHER").Default("AWARD"),
		field.String("achievementUrl").Optional(),
		field.String("impactMetrics").Optional(),
		field.Int("orderIndex").Default(0),
	}
}

// Edges of the Achievement.
func (Achievement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("achievements").Unique().Required(),
	}
}

func (Achievement) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("achievementType"),
	}
}
