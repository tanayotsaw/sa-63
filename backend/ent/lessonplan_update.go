// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/USER/app/ent/lessonplan"
	"github.com/USER/app/ent/predicate"
	"github.com/USER/app/ent/section"
	"github.com/USER/app/ent/subject"
	"github.com/USER/app/ent/teacher"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// LessonplanUpdate is the builder for updating Lessonplan entities.
type LessonplanUpdate struct {
	config
	hooks      []Hook
	mutation   *LessonplanMutation
	predicates []predicate.Lessonplan
}

// Where adds a new predicate for the builder.
func (lu *LessonplanUpdate) Where(ps ...predicate.Lessonplan) *LessonplanUpdate {
	lu.predicates = append(lu.predicates, ps...)
	return lu
}

// SetRoom sets the Room field.
func (lu *LessonplanUpdate) SetRoom(s string) *LessonplanUpdate {
	lu.mutation.SetRoom(s)
	return lu
}

// SetGroupIDID sets the Group_id edge to Section by id.
func (lu *LessonplanUpdate) SetGroupIDID(id int) *LessonplanUpdate {
	lu.mutation.SetGroupIDID(id)
	return lu
}

// SetNillableGroupIDID sets the Group_id edge to Section by id if the given value is not nil.
func (lu *LessonplanUpdate) SetNillableGroupIDID(id *int) *LessonplanUpdate {
	if id != nil {
		lu = lu.SetGroupIDID(*id)
	}
	return lu
}

// SetGroupID sets the Group_id edge to Section.
func (lu *LessonplanUpdate) SetGroupID(s *Section) *LessonplanUpdate {
	return lu.SetGroupIDID(s.ID)
}

// SetCourseIDID sets the Course_ID edge to Subject by id.
func (lu *LessonplanUpdate) SetCourseIDID(id int) *LessonplanUpdate {
	lu.mutation.SetCourseIDID(id)
	return lu
}

// SetNillableCourseIDID sets the Course_ID edge to Subject by id if the given value is not nil.
func (lu *LessonplanUpdate) SetNillableCourseIDID(id *int) *LessonplanUpdate {
	if id != nil {
		lu = lu.SetCourseIDID(*id)
	}
	return lu
}

// SetCourseID sets the Course_ID edge to Subject.
func (lu *LessonplanUpdate) SetCourseID(s *Subject) *LessonplanUpdate {
	return lu.SetCourseIDID(s.ID)
}

// SetProfessorIDID sets the Professor_ID edge to Teacher by id.
func (lu *LessonplanUpdate) SetProfessorIDID(id int) *LessonplanUpdate {
	lu.mutation.SetProfessorIDID(id)
	return lu
}

// SetNillableProfessorIDID sets the Professor_ID edge to Teacher by id if the given value is not nil.
func (lu *LessonplanUpdate) SetNillableProfessorIDID(id *int) *LessonplanUpdate {
	if id != nil {
		lu = lu.SetProfessorIDID(*id)
	}
	return lu
}

// SetProfessorID sets the Professor_ID edge to Teacher.
func (lu *LessonplanUpdate) SetProfessorID(t *Teacher) *LessonplanUpdate {
	return lu.SetProfessorIDID(t.ID)
}

// Mutation returns the LessonplanMutation object of the builder.
func (lu *LessonplanUpdate) Mutation() *LessonplanMutation {
	return lu.mutation
}

// ClearGroupID clears the Group_id edge to Section.
func (lu *LessonplanUpdate) ClearGroupID() *LessonplanUpdate {
	lu.mutation.ClearGroupID()
	return lu
}

// ClearCourseID clears the Course_ID edge to Subject.
func (lu *LessonplanUpdate) ClearCourseID() *LessonplanUpdate {
	lu.mutation.ClearCourseID()
	return lu
}

// ClearProfessorID clears the Professor_ID edge to Teacher.
func (lu *LessonplanUpdate) ClearProfessorID() *LessonplanUpdate {
	lu.mutation.ClearProfessorID()
	return lu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (lu *LessonplanUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := lu.mutation.Room(); ok {
		if err := lessonplan.RoomValidator(v); err != nil {
			return 0, &ValidationError{Name: "Room", err: fmt.Errorf("ent: validator failed for field \"Room\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LessonplanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LessonplanUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LessonplanUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LessonplanUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LessonplanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lessonplan.Table,
			Columns: lessonplan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lessonplan.FieldID,
			},
		},
	}
	if ps := lu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Room(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lessonplan.FieldRoom,
		})
	}
	if lu.mutation.GroupIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.GroupIDTable,
			Columns: []string{lessonplan.GroupIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: section.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.GroupIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.GroupIDTable,
			Columns: []string{lessonplan.GroupIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: section.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.CourseIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.CourseIDTable,
			Columns: []string{lessonplan.CourseIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.CourseIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.CourseIDTable,
			Columns: []string{lessonplan.CourseIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.ProfessorIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.ProfessorIDTable,
			Columns: []string{lessonplan.ProfessorIDColumn},
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
	if nodes := lu.mutation.ProfessorIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.ProfessorIDTable,
			Columns: []string{lessonplan.ProfessorIDColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lessonplan.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LessonplanUpdateOne is the builder for updating a single Lessonplan entity.
type LessonplanUpdateOne struct {
	config
	hooks    []Hook
	mutation *LessonplanMutation
}

// SetRoom sets the Room field.
func (luo *LessonplanUpdateOne) SetRoom(s string) *LessonplanUpdateOne {
	luo.mutation.SetRoom(s)
	return luo
}

// SetGroupIDID sets the Group_id edge to Section by id.
func (luo *LessonplanUpdateOne) SetGroupIDID(id int) *LessonplanUpdateOne {
	luo.mutation.SetGroupIDID(id)
	return luo
}

// SetNillableGroupIDID sets the Group_id edge to Section by id if the given value is not nil.
func (luo *LessonplanUpdateOne) SetNillableGroupIDID(id *int) *LessonplanUpdateOne {
	if id != nil {
		luo = luo.SetGroupIDID(*id)
	}
	return luo
}

// SetGroupID sets the Group_id edge to Section.
func (luo *LessonplanUpdateOne) SetGroupID(s *Section) *LessonplanUpdateOne {
	return luo.SetGroupIDID(s.ID)
}

// SetCourseIDID sets the Course_ID edge to Subject by id.
func (luo *LessonplanUpdateOne) SetCourseIDID(id int) *LessonplanUpdateOne {
	luo.mutation.SetCourseIDID(id)
	return luo
}

// SetNillableCourseIDID sets the Course_ID edge to Subject by id if the given value is not nil.
func (luo *LessonplanUpdateOne) SetNillableCourseIDID(id *int) *LessonplanUpdateOne {
	if id != nil {
		luo = luo.SetCourseIDID(*id)
	}
	return luo
}

// SetCourseID sets the Course_ID edge to Subject.
func (luo *LessonplanUpdateOne) SetCourseID(s *Subject) *LessonplanUpdateOne {
	return luo.SetCourseIDID(s.ID)
}

// SetProfessorIDID sets the Professor_ID edge to Teacher by id.
func (luo *LessonplanUpdateOne) SetProfessorIDID(id int) *LessonplanUpdateOne {
	luo.mutation.SetProfessorIDID(id)
	return luo
}

// SetNillableProfessorIDID sets the Professor_ID edge to Teacher by id if the given value is not nil.
func (luo *LessonplanUpdateOne) SetNillableProfessorIDID(id *int) *LessonplanUpdateOne {
	if id != nil {
		luo = luo.SetProfessorIDID(*id)
	}
	return luo
}

// SetProfessorID sets the Professor_ID edge to Teacher.
func (luo *LessonplanUpdateOne) SetProfessorID(t *Teacher) *LessonplanUpdateOne {
	return luo.SetProfessorIDID(t.ID)
}

// Mutation returns the LessonplanMutation object of the builder.
func (luo *LessonplanUpdateOne) Mutation() *LessonplanMutation {
	return luo.mutation
}

// ClearGroupID clears the Group_id edge to Section.
func (luo *LessonplanUpdateOne) ClearGroupID() *LessonplanUpdateOne {
	luo.mutation.ClearGroupID()
	return luo
}

// ClearCourseID clears the Course_ID edge to Subject.
func (luo *LessonplanUpdateOne) ClearCourseID() *LessonplanUpdateOne {
	luo.mutation.ClearCourseID()
	return luo
}

// ClearProfessorID clears the Professor_ID edge to Teacher.
func (luo *LessonplanUpdateOne) ClearProfessorID() *LessonplanUpdateOne {
	luo.mutation.ClearProfessorID()
	return luo
}

// Save executes the query and returns the updated entity.
func (luo *LessonplanUpdateOne) Save(ctx context.Context) (*Lessonplan, error) {
	if v, ok := luo.mutation.Room(); ok {
		if err := lessonplan.RoomValidator(v); err != nil {
			return nil, &ValidationError{Name: "Room", err: fmt.Errorf("ent: validator failed for field \"Room\": %w", err)}
		}
	}

	var (
		err  error
		node *Lessonplan
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LessonplanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LessonplanUpdateOne) SaveX(ctx context.Context) *Lessonplan {
	l, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return l
}

// Exec executes the query on the entity.
func (luo *LessonplanUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LessonplanUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LessonplanUpdateOne) sqlSave(ctx context.Context) (l *Lessonplan, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   lessonplan.Table,
			Columns: lessonplan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: lessonplan.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Lessonplan.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := luo.mutation.Room(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: lessonplan.FieldRoom,
		})
	}
	if luo.mutation.GroupIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.GroupIDTable,
			Columns: []string{lessonplan.GroupIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: section.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.GroupIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.GroupIDTable,
			Columns: []string{lessonplan.GroupIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: section.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.CourseIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.CourseIDTable,
			Columns: []string{lessonplan.CourseIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.CourseIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.CourseIDTable,
			Columns: []string{lessonplan.CourseIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subject.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.ProfessorIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.ProfessorIDTable,
			Columns: []string{lessonplan.ProfessorIDColumn},
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
	if nodes := luo.mutation.ProfessorIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lessonplan.ProfessorIDTable,
			Columns: []string{lessonplan.ProfessorIDColumn},
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
	l = &Lessonplan{config: luo.config}
	_spec.Assign = l.assignValues
	_spec.ScanValues = l.scanValues()
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{lessonplan.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return l, nil
}
