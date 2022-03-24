package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Artcile holds the schema definition for the Artcile entity.
type Artcile struct {
	ent.Schema
}

// Fields of the Artcile.
func (Artcile) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Default(""),
		field.String("description").Default(""),
	}
}

// Edges of the Artcile.
func (Artcile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("articles").Unique(),
	}
}
