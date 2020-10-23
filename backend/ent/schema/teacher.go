package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Teacher holds the schema definition for the Teacher entity.
type Teacher struct {
	ent.Schema
}

// Fields of the Teacher.
func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.String("Teacher_Name").Unique(),
	}
}

// Edges of the Teacher.
func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Teacher_ID", Subject.Type),
		edge.To("Teachers_ID", Lessonplan.Type),
	}
}
