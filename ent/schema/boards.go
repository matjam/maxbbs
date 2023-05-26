package schema

import "entgo.io/ent"

// Boards holds the schema definition for the Boards entity.
type Boards struct {
	ent.Schema
}

// Fields of the Boards.
func (Boards) Fields() []ent.Field {
	return nil
}

// Edges of the Boards.
func (Boards) Edges() []ent.Edge {
	return nil
}
