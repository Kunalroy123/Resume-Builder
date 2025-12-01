package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Template holds the schema definition for the Template entity.
type Template struct {
	ent.Schema
}

// Fields of the Template.
func (Template) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New).Immutable(),
		field.String("name").NotEmpty(),
		field.Text("description").Optional(),
		field.Text("htmlTemplate").NotEmpty(),
		field.Text("cssStyles").NotEmpty(),
		field.Bool("isActive").Default(true),
		field.String("previewImage").Optional(),
	}
}

// Edges of the Template.
func (Template) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("resumes", Resume.Type),
	}
}

func (Template) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("isActive"),
		index.Fields("name"),
	}
}
