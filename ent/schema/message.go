package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created").Default(time.Now),
		field.Time("updated").Default(time.Now),
		field.Time("deleted").Default(time.Now),
		field.String("subject").NotEmpty(),
		field.String("body").NotEmpty(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("messages").
			Unique(),
	}
}
