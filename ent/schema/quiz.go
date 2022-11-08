package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Quiz holds the schema definition for the Quiz entity.
type Quiz struct {
	ent.Schema
}

// Fields of the Quiz.
func (Quiz) Fields() []ent.Field {
	return []ent.Field{
		field.String("postUrl").StorageKey("post_url"),
		field.String("question").StorageKey("question"),
		field.String("answer").StorageKey("answer"),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Quiz.
func (Quiz) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("quiz").Unique(),
	}
}
