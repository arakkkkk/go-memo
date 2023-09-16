package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Memo holds the schema definition for the Memo entity.
type Memo struct {
	ent.Schema
}

// Fields of the Memo.
func (Memo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Int("user_id").
			Optional(),
		field.String("description").
			Default(""),
		field.Time("updated_at").
			Default(time.Now),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Memo.
func (Memo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
