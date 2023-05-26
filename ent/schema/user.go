package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

type UserOptions struct {
	Superuser bool `json:"superuser"`
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created").Default(time.Now),
		field.Time("updated").Default(time.Now),
		field.String("username").NotEmpty().Unique(),
		field.String("password").Default(""),
		field.String("name").Default(""),
		field.String("email").NotEmpty().Unique(),
		field.Bool("active").Default(true),
		field.JSON("options", &UserOptions{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
	}
}
