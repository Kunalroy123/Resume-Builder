package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Skill holds the schema definition for the Skill entity.
type Skill struct {
	ent.Schema
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).Immutable(),
		field.String("name").NotEmpty(),
		field.String("category").NotEmpty(),
		field.Enum("skillType").Values("TECHNICAL", "SOFT").Default("TECHNICAL"),
		field.Enum("proficiencyLevel").Values("BEGINNER", "INTERMEDIATE", "ADVANCED", "EXPERT").Default("INTERMEDIATE"),
		field.Int("yearsOfExperience").Optional().Nillable(),
		field.Int("orderIndex").Default(0),
	}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("skills").Unique().Required(),
	}
}

func (Skill) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("category"),
		index.Fields("yearsOfExperience"),
		index.Fields("skillType"),
		index.Fields("proficiencyLevel"),
	}
}
