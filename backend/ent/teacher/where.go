// Code generated by entc, DO NOT EDIT.

package teacher

import (
	"github.com/USER/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TeacherName applies equality check predicate on the "Teacher_Name" field. It's identical to TeacherNameEQ.
func TeacherName(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeacherName), v))
	})
}

// TeacherNameEQ applies the EQ predicate on the "Teacher_Name" field.
func TeacherNameEQ(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeacherName), v))
	})
}

// TeacherNameNEQ applies the NEQ predicate on the "Teacher_Name" field.
func TeacherNameNEQ(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTeacherName), v))
	})
}

// TeacherNameIn applies the In predicate on the "Teacher_Name" field.
func TeacherNameIn(vs ...string) predicate.Teacher {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teacher(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTeacherName), v...))
	})
}

// TeacherNameNotIn applies the NotIn predicate on the "Teacher_Name" field.
func TeacherNameNotIn(vs ...string) predicate.Teacher {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Teacher(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTeacherName), v...))
	})
}

// TeacherNameGT applies the GT predicate on the "Teacher_Name" field.
func TeacherNameGT(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTeacherName), v))
	})
}

// TeacherNameGTE applies the GTE predicate on the "Teacher_Name" field.
func TeacherNameGTE(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTeacherName), v))
	})
}

// TeacherNameLT applies the LT predicate on the "Teacher_Name" field.
func TeacherNameLT(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTeacherName), v))
	})
}

// TeacherNameLTE applies the LTE predicate on the "Teacher_Name" field.
func TeacherNameLTE(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTeacherName), v))
	})
}

// TeacherNameContains applies the Contains predicate on the "Teacher_Name" field.
func TeacherNameContains(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTeacherName), v))
	})
}

// TeacherNameHasPrefix applies the HasPrefix predicate on the "Teacher_Name" field.
func TeacherNameHasPrefix(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTeacherName), v))
	})
}

// TeacherNameHasSuffix applies the HasSuffix predicate on the "Teacher_Name" field.
func TeacherNameHasSuffix(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTeacherName), v))
	})
}

// TeacherNameEqualFold applies the EqualFold predicate on the "Teacher_Name" field.
func TeacherNameEqualFold(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTeacherName), v))
	})
}

// TeacherNameContainsFold applies the ContainsFold predicate on the "Teacher_Name" field.
func TeacherNameContainsFold(v string) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTeacherName), v))
	})
}

// HasTeacherID applies the HasEdge predicate on the "Teacher_ID" edge.
func HasTeacherID() predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeacherIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeacherIDTable, TeacherIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeacherIDWith applies the HasEdge predicate on the "Teacher_ID" edge with a given conditions (other predicates).
func HasTeacherIDWith(preds ...predicate.Subject) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeacherIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeacherIDTable, TeacherIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeachersID applies the HasEdge predicate on the "Teachers_ID" edge.
func HasTeachersID() predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeachersIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeachersIDTable, TeachersIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeachersIDWith applies the HasEdge predicate on the "Teachers_ID" edge with a given conditions (other predicates).
func HasTeachersIDWith(preds ...predicate.Lessonplan) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeachersIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TeachersIDTable, TeachersIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Teacher) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Teacher) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Teacher) predicate.Teacher {
	return predicate.Teacher(func(s *sql.Selector) {
		p(s.Not())
	})
}
