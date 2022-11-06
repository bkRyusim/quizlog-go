package schema

import "entgo.io/ent"

// Quiz holds the schema definition for the Quiz entity.
type Quiz struct {
	ent.Schema
}

// Fields of the Quiz.
func (Quiz) Fields() []ent.Field {
	return nil
}

// Edges of the Quiz.
func (Quiz) Edges() []ent.Edge {
	return nil
}
