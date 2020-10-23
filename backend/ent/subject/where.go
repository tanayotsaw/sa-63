// Code generated by entc, DO NOT EDIT.

package subject

import (
	"github.com/USER/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SubjectName applies equality check predicate on the "Subject_Name" field. It's identical to SubjectNameEQ.
func SubjectName(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubjectName), v))
	})
}

// Credit applies equality check predicate on the "Credit" field. It's identical to CreditEQ.
func Credit(v int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCredit), v))
	})
}

// Price applies equality check predicate on the "Price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// SubjectNameEQ applies the EQ predicate on the "Subject_Name" field.
func SubjectNameEQ(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubjectName), v))
	})
}

// SubjectNameNEQ applies the NEQ predicate on the "Subject_Name" field.
func SubjectNameNEQ(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSubjectName), v))
	})
}

// SubjectNameIn applies the In predicate on the "Subject_Name" field.
func SubjectNameIn(vs ...string) predicate.Subject {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subject(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSubjectName), v...))
	})
}

// SubjectNameNotIn applies the NotIn predicate on the "Subject_Name" field.
func SubjectNameNotIn(vs ...string) predicate.Subject {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subject(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSubjectName), v...))
	})
}

// SubjectNameGT applies the GT predicate on the "Subject_Name" field.
func SubjectNameGT(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSubjectName), v))
	})
}

// SubjectNameGTE applies the GTE predicate on the "Subject_Name" field.
func SubjectNameGTE(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSubjectName), v))
	})
}

// SubjectNameLT applies the LT predicate on the "Subject_Name" field.
func SubjectNameLT(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSubjectName), v))
	})
}

// SubjectNameLTE applies the LTE predicate on the "Subject_Name" field.
func SubjectNameLTE(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSubjectName), v))
	})
}

// SubjectNameContains applies the Contains predicate on the "Subject_Name" field.
func SubjectNameContains(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSubjectName), v))
	})
}

// SubjectNameHasPrefix applies the HasPrefix predicate on the "Subject_Name" field.
func SubjectNameHasPrefix(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSubjectName), v))
	})
}

// SubjectNameHasSuffix applies the HasSuffix predicate on the "Subject_Name" field.
func SubjectNameHasSuffix(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSubjectName), v))
	})
}

// SubjectNameEqualFold applies the EqualFold predicate on the "Subject_Name" field.
func SubjectNameEqualFold(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSubjectName), v))
	})
}

// SubjectNameContainsFold applies the ContainsFold predicate on the "Subject_Name" field.
func SubjectNameContainsFold(v string) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSubjectName), v))
	})
}

// CreditEQ applies the EQ predicate on the "Credit" field.
func CreditEQ(v int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCredit), v))
	})
}

// CreditNEQ applies the NEQ predicate on the "Credit" field.
func CreditNEQ(v int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCredit), v))
	})
}

// CreditIn applies the In predicate on the "Credit" field.
func CreditIn(vs ...int) predicate.Subject {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subject(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCredit), v...))
	})
}

// CreditNotIn applies the NotIn predicate on the "Credit" field.
func CreditNotIn(vs ...int) predicate.Subject {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subject(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCredit), v...))
	})
}

// CreditGT applies the GT predicate on the "Credit" field.
func CreditGT(v int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCredit), v))
	})
}

// CreditGTE applies the GTE predicate on the "Credit" field.
func CreditGTE(v int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCredit), v))
	})
}

// CreditLT applies the LT predicate on the "Credit" field.
func CreditLT(v int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCredit), v))
	})
}

// CreditLTE applies the LTE predicate on the "Credit" field.
func CreditLTE(v int) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCredit), v))
	})
}

// PriceEQ applies the EQ predicate on the "Price" field.
func PriceEQ(v float64) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "Price" field.
func PriceNEQ(v float64) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "Price" field.
func PriceIn(vs ...float64) predicate.Subject {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subject(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "Price" field.
func PriceNotIn(vs ...float64) predicate.Subject {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Subject(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "Price" field.
func PriceGT(v float64) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "Price" field.
func PriceGTE(v float64) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "Price" field.
func PriceLT(v float64) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "Price" field.
func PriceLTE(v float64) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// HasProfessorID applies the HasEdge predicate on the "Professor_ID" edge.
func HasProfessorID() predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfessorIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProfessorIDTable, ProfessorIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfessorIDWith applies the HasEdge predicate on the "Professor_ID" edge with a given conditions (other predicates).
func HasProfessorIDWith(preds ...predicate.Teacher) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProfessorIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProfessorIDTable, ProfessorIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubjectID applies the HasEdge predicate on the "Subject_ID" edge.
func HasSubjectID() predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubjectIDTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubjectIDTable, SubjectIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectIDWith applies the HasEdge predicate on the "Subject_ID" edge with a given conditions (other predicates).
func HasSubjectIDWith(preds ...predicate.Lessonplan) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubjectIDInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubjectIDTable, SubjectIDColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Subject) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Subject) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
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
func Not(p predicate.Subject) predicate.Subject {
	return predicate.Subject(func(s *sql.Selector) {
		p(s.Not())
	})
}
