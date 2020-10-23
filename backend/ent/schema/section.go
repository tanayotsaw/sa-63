package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Section holds the schema definition for the Section entity.
type Section struct {
	ent.Schema
}

// Fields of the Section.
func (Section) Fields() []ent.Field {
	return []ent.Field{

		field.Int("Group").Positive(),
		field.Int("recieve").Positive(),
		field.Int("Seat_left").Positive(),
		field.String("Room").NotEmpty(),
		field.String("Date_Time").NotEmpty(),
	}
}

// Edges of the Section.
func (Section) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sec_id", Lessonplan.Type),
		
	}

}
