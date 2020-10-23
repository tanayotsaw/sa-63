package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Lessonplan holds the schema definition for the Lessonplan entity.
type Lessonplan struct {
	ent.Schema
}

// Fields of the Lessonplan.
func (Lessonplan) Fields() []ent.Field {
	return []ent.Field{

		field.String("Room").NotEmpty(),
	}

}

// Edges of the Lessonplan.
func (Lessonplan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Group_id", Section.Type).Ref("sec_id").Unique(),
		edge.From("Course_ID", Subject.Type).Ref("Subject_ID").Unique(),
		edge.From("Professor_ID", Teacher.Type).Ref("Teachers_ID").Unique(),
	}

}
