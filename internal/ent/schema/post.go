package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Text("content").NotEmpty(),
		field.Time("created_at").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).Default(time.Now()),
		field.Time("updated_at").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).Default(time.Now()),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("posts").
			Unique(),
	}
}
