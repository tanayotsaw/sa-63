// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/USER/app/ent/predicate"
	"github.com/USER/app/ent/subject"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// SubjectDelete is the builder for deleting a Subject entity.
type SubjectDelete struct {
	config
	hooks      []Hook
	mutation   *SubjectMutation
	predicates []predicate.Subject
}

// Where adds a new predicate to the delete builder.
func (sd *SubjectDelete) Where(ps ...predicate.Subject) *SubjectDelete {
	sd.predicates = append(sd.predicates, ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SubjectDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sd.hooks) == 0 {
		affected, err = sd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubjectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sd.mutation = mutation
			affected, err = sd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sd.hooks) - 1; i >= 0; i-- {
			mut = sd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SubjectDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SubjectDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: subject.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subject.FieldID,
			},
		},
	}
	if ps := sd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
}

// SubjectDeleteOne is the builder for deleting a single Subject entity.
type SubjectDeleteOne struct {
	sd *SubjectDelete
}

// Exec executes the deletion query.
func (sdo *SubjectDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{subject.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SubjectDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}
