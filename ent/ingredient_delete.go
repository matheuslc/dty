// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/matheuslc/dty/ent/ingredient"
	"github.com/matheuslc/dty/ent/predicate"
)

// IngredientDelete is the builder for deleting a Ingredient entity.
type IngredientDelete struct {
	config
	predicates []predicate.Ingredient
}

// Where adds a new predicate to the delete builder.
func (id *IngredientDelete) Where(ps ...predicate.Ingredient) *IngredientDelete {
	id.predicates = append(id.predicates, ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *IngredientDelete) Exec(ctx context.Context) (int, error) {
	return id.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *IngredientDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *IngredientDelete) sqlExec(ctx context.Context) (int, error) {
	var res sql.Result
	selector := sql.Select().From(sql.Table(ingredient.Table))
	for _, p := range id.predicates {
		p(selector)
	}
	query, args := sql.Delete(ingredient.Table).FromSelect(selector).Query()
	if err := id.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// IngredientDeleteOne is the builder for deleting a single Ingredient entity.
type IngredientDeleteOne struct {
	id *IngredientDelete
}

// Exec executes the deletion query.
func (ido *IngredientDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{ingredient.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *IngredientDeleteOne) ExecX(ctx context.Context) {
	ido.id.ExecX(ctx)
}