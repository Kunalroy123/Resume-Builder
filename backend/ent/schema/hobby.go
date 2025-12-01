package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Hobby holds the schema definition for the Hobby entity.
type Hobby struct {
	ent.Schema
}

// Fields of the Hobby.
func (Hobby) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).Immutable(),
		field.String("name").Optional(),
		field.Text("description").Optional(),
		field.String("skillLevel").Optional(),
		field.Int("yearsInvolved").Optional().Nillable(),
		field.Text("achievements").Optional(),
		field.Int("orderIndex").Default(0),
	}
}

// Edges of the Hobby.
func (Hobby) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resume", Resume.Type).Ref("hobbies").Unique().Required(),
	}
}

func (Hobby) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
