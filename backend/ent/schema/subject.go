package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Subject holds the schema definition for the Subject entity.
type Subject struct {
	ent.Schema
}

// Fields of the Subject.
func (Subject) Fields() []ent.Field {
	return []ent.Field{

		field.String("Subject_Name").Unique(),
		
		field.Int("Credit").Positive(),

		field.Float("Price").Positive(),
	}
}

// Edges of the Subject.
func (Subject) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Professor_ID", Teacher.Type).Ref("Teacher_ID").Unique(),
		edge.To("Subject_ID", Lessonplan.Type),
	}

}
