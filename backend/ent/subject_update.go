// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/USER/app/ent/lessonplan"
	"github.com/USER/app/ent/predicate"
	"github.com/USER/app/ent/subject"
	"github.com/USER/app/ent/teacher"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// SubjectUpdate is the builder for updating Subject entities.
type SubjectUpdate struct {
	config
	hooks      []Hook
	mutation   *SubjectMutation
	predicates []predicate.Subject
}

// Where adds a new predicate for the builder.
func (su *SubjectUpdate) Where(ps ...predicate.Subject) *SubjectUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetSubjectName sets the Subject_Name field.
func (su *SubjectUpdate) SetSubjectName(s string) *SubjectUpdate {
	su.mutation.SetSubjectName(s)
	return su
}

// SetCredit sets the Credit field.
func (su *SubjectUpdate) SetCredit(i int) *SubjectUpdate {
	su.mutation.ResetCredit()
	su.mutation.SetCredit(i)
	return su
}

// AddCredit adds i to Credit.
func (su *SubjectUpdate) AddCredit(i int) *SubjectUpdate {
	su.mutation.AddCredit(i)
	return su
}

// SetPrice sets the Price field.
func (su *SubjectUpdate) SetPrice(f float64) *SubjectUpdate {
	su.mutation.ResetPrice()
	su.mutation.SetPrice(f)
	return su
}

// AddPrice adds f to Price.
func (su *SubjectUpdate) AddPrice(f float64) *SubjectUpdate {
	su.mutation.AddPrice(f)
	return su
}

// SetProfessorIDID sets the Professor_ID edge to Teacher by id.
func (su *SubjectUpdate) SetProfessorIDID(id int) *SubjectUpdate {
	su.mutation.SetProfessorIDID(id)
	return su
}

// SetNillableProfessorIDID sets the Professor_ID edge to Teacher by id if the given value is not nil.
func (su *SubjectUpdate) SetNillableProfessorIDID(id *int) *SubjectUpdate {
	if id != nil {
		su = su.SetProfessorIDID(*id)
	}
	return su
}

// SetProfessorID sets the Professor_ID edge to Teacher.
func (su *SubjectUpdate) SetProfessorID(t *Teacher) *SubjectUpdate {
	return su.SetProfessorIDID(t.ID)
}

// AddSubjectIDIDs adds the Subject_ID edge to Lessonplan by ids.
func (su *SubjectUpdate) AddSubjectIDIDs(ids ...int) *SubjectUpdate {
	su.mutation.AddSubjectIDIDs(ids...)
	return su
}

// AddSubjectID adds the Subject_ID edges to Lessonplan.
func (su *SubjectUpdate) AddSubjectID(l ...*Lessonplan) *SubjectUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return su.AddSubjectIDIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (su *SubjectUpdate) Mutation() *SubjectMutation {
	return su.mutation
}

// ClearProfessorID clears the Professor_ID edge to Teacher.
func (su *SubjectUpdate) ClearProfessorID() *SubjectUpdate {
	su.mutation.ClearProfessorID()
	return su
}

// RemoveSubjectIDIDs removes the Subject_ID edge to Lessonplan by ids.
func (su *SubjectUpdate) RemoveSubjectIDIDs(ids ...int) *SubjectUpdate {
	su.mutation.RemoveSubjectIDIDs(ids...)
	return su
}

// RemoveSubjectID removes Subject_ID edges to Lessonplan.
func (su *SubjectUpdate) RemoveSubjectID(l ...*Lessonplan) *SubjectUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return su.RemoveSubjectIDIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SubjectUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := su.mutation.Credit(); ok {
		if err := subject.CreditValidator(v); err != nil {
			return 0, &ValidationError{Name: "Credit", err: fmt.Errorf("ent: validator failed for field \"Credit\": %w", err)}
		}
	}
	if v, ok := su.mutation.Price(); ok {
		if err := subject.PriceValidator(v); err != nil {
			return 0, &ValidationError{Name: "Price", err: fmt.Errorf("ent: validator failed for field \"Price\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubjectUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubjectUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubjectUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SubjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subject.Table,
			Columns: subject.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subject.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.SubjectName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subject.FieldSubjectName,
		})
	}
	if value, ok := su.mutation.Credit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subject.FieldCredit,
		})
	}
	if value, ok := su.mutation.AddedCredit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subject.FieldCredit,
		})
	}
	if value, ok := su.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: subject.FieldPrice,
		})
	}
	if value, ok := su.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: subject.FieldPrice,
		})
	}
	if su.mutation.ProfessorIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.ProfessorIDTable,
			Columns: []string{subject.ProfessorIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ProfessorIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.ProfessorIDTable,
			Columns: []string{subject.ProfessorIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := su.mutation.RemovedSubjectIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.SubjectIDTable,
			Columns: []string{subject.SubjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lessonplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SubjectIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.SubjectIDTable,
			Columns: []string{subject.SubjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lessonplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subject.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SubjectUpdateOne is the builder for updating a single Subject entity.
type SubjectUpdateOne struct {
	config
	hooks    []Hook
	mutation *SubjectMutation
}

// SetSubjectName sets the Subject_Name field.
func (suo *SubjectUpdateOne) SetSubjectName(s string) *SubjectUpdateOne {
	suo.mutation.SetSubjectName(s)
	return suo
}

// SetCredit sets the Credit field.
func (suo *SubjectUpdateOne) SetCredit(i int) *SubjectUpdateOne {
	suo.mutation.ResetCredit()
	suo.mutation.SetCredit(i)
	return suo
}

// AddCredit adds i to Credit.
func (suo *SubjectUpdateOne) AddCredit(i int) *SubjectUpdateOne {
	suo.mutation.AddCredit(i)
	return suo
}

// SetPrice sets the Price field.
func (suo *SubjectUpdateOne) SetPrice(f float64) *SubjectUpdateOne {
	suo.mutation.ResetPrice()
	suo.mutation.SetPrice(f)
	return suo
}

// AddPrice adds f to Price.
func (suo *SubjectUpdateOne) AddPrice(f float64) *SubjectUpdateOne {
	suo.mutation.AddPrice(f)
	return suo
}

// SetProfessorIDID sets the Professor_ID edge to Teacher by id.
func (suo *SubjectUpdateOne) SetProfessorIDID(id int) *SubjectUpdateOne {
	suo.mutation.SetProfessorIDID(id)
	return suo
}

// SetNillableProfessorIDID sets the Professor_ID edge to Teacher by id if the given value is not nil.
func (suo *SubjectUpdateOne) SetNillableProfessorIDID(id *int) *SubjectUpdateOne {
	if id != nil {
		suo = suo.SetProfessorIDID(*id)
	}
	return suo
}

// SetProfessorID sets the Professor_ID edge to Teacher.
func (suo *SubjectUpdateOne) SetProfessorID(t *Teacher) *SubjectUpdateOne {
	return suo.SetProfessorIDID(t.ID)
}

// AddSubjectIDIDs adds the Subject_ID edge to Lessonplan by ids.
func (suo *SubjectUpdateOne) AddSubjectIDIDs(ids ...int) *SubjectUpdateOne {
	suo.mutation.AddSubjectIDIDs(ids...)
	return suo
}

// AddSubjectID adds the Subject_ID edges to Lessonplan.
func (suo *SubjectUpdateOne) AddSubjectID(l ...*Lessonplan) *SubjectUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return suo.AddSubjectIDIDs(ids...)
}

// Mutation returns the SubjectMutation object of the builder.
func (suo *SubjectUpdateOne) Mutation() *SubjectMutation {
	return suo.mutation
}

// ClearProfessorID clears the Professor_ID edge to Teacher.
func (suo *SubjectUpdateOne) ClearProfessorID() *SubjectUpdateOne {
	suo.mutation.ClearProfessorID()
	return suo
}

// RemoveSubjectIDIDs removes the Subject_ID edge to Lessonplan by ids.
func (suo *SubjectUpdateOne) RemoveSubjectIDIDs(ids ...int) *SubjectUpdateOne {
	suo.mutation.RemoveSubjectIDIDs(ids...)
	return suo
}

// RemoveSubjectID removes Subject_ID edges to Lessonplan.
func (suo *SubjectUpdateOne) RemoveSubjectID(l ...*Lessonplan) *SubjectUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return suo.RemoveSubjectIDIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SubjectUpdateOne) Save(ctx context.Context) (*Subject, error) {
	if v, ok := suo.mutation.Credit(); ok {
		if err := subject.CreditValidator(v); err != nil {
			return nil, &ValidationError{Name: "Credit", err: fmt.Errorf("ent: validator failed for field \"Credit\": %w", err)}
		}
	}
	if v, ok := suo.mutation.Price(); ok {
		if err := subject.PriceValidator(v); err != nil {
			return nil, &ValidationError{Name: "Price", err: fmt.Errorf("ent: validator failed for field \"Price\": %w", err)}
		}
	}

	var (
		err  error
		node *Subject
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubjectUpdateOne) SaveX(ctx context.Context) *Subject {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SubjectUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubjectUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SubjectUpdateOne) sqlSave(ctx context.Context) (s *Subject, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subject.Table,
			Columns: subject.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subject.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Subject.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.SubjectName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subject.FieldSubjectName,
		})
	}
	if value, ok := suo.mutation.Credit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subject.FieldCredit,
		})
	}
	if value, ok := suo.mutation.AddedCredit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: subject.FieldCredit,
		})
	}
	if value, ok := suo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: subject.FieldPrice,
		})
	}
	if value, ok := suo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: subject.FieldPrice,
		})
	}
	if suo.mutation.ProfessorIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.ProfessorIDTable,
			Columns: []string{subject.ProfessorIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ProfessorIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subject.ProfessorIDTable,
			Columns: []string{subject.ProfessorIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: teacher.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := suo.mutation.RemovedSubjectIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.SubjectIDTable,
			Columns: []string{subject.SubjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lessonplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SubjectIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subject.SubjectIDTable,
			Columns: []string{subject.SubjectIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: lessonplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Subject{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subject.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
