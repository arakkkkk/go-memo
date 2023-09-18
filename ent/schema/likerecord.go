package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LikeRecord holds the schema definition for the LikeRecord entity.
type LikeRecord struct {
	ent.Schema
}

// Fields of the LikeRecord.
func (LikeRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.UUID("memo_id", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the LikeRecord.
func (LikeRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("memos", Memo.Type),
		edge.To("users", User.Type),
	}
}
