// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/USER/app/ent/section"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Section is the model entity for the Section schema.
type Section struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Group holds the value of the "Group" field.
	Group int `json:"Group,omitempty"`
	// Recieve holds the value of the "recieve" field.
	Recieve int `json:"recieve,omitempty"`
	// SeatLeft holds the value of the "Seat_left" field.
	SeatLeft int `json:"Seat_left,omitempty"`
	// Room holds the value of the "Room" field.
	Room string `json:"Room,omitempty"`
	// DateTime holds the value of the "Date_Time" field.
	DateTime string `json:"Date_Time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SectionQuery when eager-loading is set.
	Edges SectionEdges `json:"edges"`
}

// SectionEdges holds the relations/edges for other nodes in the graph.
type SectionEdges struct {
	// SecID holds the value of the sec_id edge.
	SecID []*Lessonplan
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SecIDOrErr returns the SecID value or an error if the edge
// was not loaded in eager-loading.
func (e SectionEdges) SecIDOrErr() ([]*Lessonplan, error) {
	if e.loadedTypes[0] {
		return e.SecID, nil
	}
	return nil, &NotLoadedError{edge: "sec_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Section) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // Group
		&sql.NullInt64{},  // recieve
		&sql.NullInt64{},  // Seat_left
		&sql.NullString{}, // Room
		&sql.NullString{}, // Date_Time
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Section fields.
func (s *Section) assignValues(values ...interface{}) error {
	if m, n := len(values), len(section.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Group", values[0])
	} else if value.Valid {
		s.Group = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field recieve", values[1])
	} else if value.Valid {
		s.Recieve = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Seat_left", values[2])
	} else if value.Valid {
		s.SeatLeft = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Room", values[3])
	} else if value.Valid {
		s.Room = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Date_Time", values[4])
	} else if value.Valid {
		s.DateTime = value.String
	}
	return nil
}

// QuerySecID queries the sec_id edge of the Section.
func (s *Section) QuerySecID() *LessonplanQuery {
	return (&SectionClient{config: s.config}).QuerySecID(s)
}

// Update returns a builder for updating this Section.
// Note that, you need to call Section.Unwrap() before calling this method, if this Section
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Section) Update() *SectionUpdateOne {
	return (&SectionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Section) Unwrap() *Section {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Section is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Section) String() string {
	var builder strings.Builder
	builder.WriteString("Section(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", Group=")
	builder.WriteString(fmt.Sprintf("%v", s.Group))
	builder.WriteString(", recieve=")
	builder.WriteString(fmt.Sprintf("%v", s.Recieve))
	builder.WriteString(", Seat_left=")
	builder.WriteString(fmt.Sprintf("%v", s.SeatLeft))
	builder.WriteString(", Room=")
	builder.WriteString(s.Room)
	builder.WriteString(", Date_Time=")
	builder.WriteString(s.DateTime)
	builder.WriteByte(')')
	return builder.String()
}

// Sections is a parsable slice of Section.
type Sections []*Section

func (s Sections) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
