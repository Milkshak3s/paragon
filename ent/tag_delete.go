// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/predicate"
	"github.com/kcarretto/paragon/ent/tag"
)

// TagDelete is the builder for deleting a Tag entity.
type TagDelete struct {
	config
	predicates []predicate.Tag
}

// Where adds a new predicate to the delete builder.
func (td *TagDelete) Where(ps ...predicate.Tag) *TagDelete {
	td.predicates = append(td.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TagDelete) Exec(ctx context.Context) (int, error) {
	return td.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TagDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TagDelete) sqlExec(ctx context.Context) (int, error) {
	spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tag.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tag.FieldID,
			},
		},
	}
	if ps := td.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, td.driver, spec)
}

// TagDeleteOne is the builder for deleting a single Tag entity.
type TagDeleteOne struct {
	td *TagDelete
}

// Exec executes the deletion query.
func (tdo *TagDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{tag.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TagDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}
