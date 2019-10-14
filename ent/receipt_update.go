// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/matheuslc/dty/ent/direction"
	"github.com/matheuslc/dty/ent/ingredient"
	"github.com/matheuslc/dty/ent/predicate"
	"github.com/matheuslc/dty/ent/receipt"
)

// ReceiptUpdate is the builder for updating Receipt entities.
type ReceiptUpdate struct {
	config
	name               *string
	portions           *int
	addportions        *int
	ingredients        map[int]struct{}
	directions         map[int]struct{}
	removedIngredients map[int]struct{}
	removedDirections  map[int]struct{}
	predicates         []predicate.Receipt
}

// Where adds a new predicate for the builder.
func (ru *ReceiptUpdate) Where(ps ...predicate.Receipt) *ReceiptUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetName sets the name field.
func (ru *ReceiptUpdate) SetName(s string) *ReceiptUpdate {
	ru.name = &s
	return ru
}

// SetPortions sets the portions field.
func (ru *ReceiptUpdate) SetPortions(i int) *ReceiptUpdate {
	ru.portions = &i
	ru.addportions = nil
	return ru
}

// AddPortions adds i to portions.
func (ru *ReceiptUpdate) AddPortions(i int) *ReceiptUpdate {
	if ru.addportions == nil {
		ru.addportions = &i
	} else {
		*ru.addportions += i
	}
	return ru
}

// AddIngredientIDs adds the ingredients edge to Ingredient by ids.
func (ru *ReceiptUpdate) AddIngredientIDs(ids ...int) *ReceiptUpdate {
	if ru.ingredients == nil {
		ru.ingredients = make(map[int]struct{})
	}
	for i := range ids {
		ru.ingredients[ids[i]] = struct{}{}
	}
	return ru
}

// AddIngredients adds the ingredients edges to Ingredient.
func (ru *ReceiptUpdate) AddIngredients(i ...*Ingredient) *ReceiptUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ru.AddIngredientIDs(ids...)
}

// AddDirectionIDs adds the directions edge to Direction by ids.
func (ru *ReceiptUpdate) AddDirectionIDs(ids ...int) *ReceiptUpdate {
	if ru.directions == nil {
		ru.directions = make(map[int]struct{})
	}
	for i := range ids {
		ru.directions[ids[i]] = struct{}{}
	}
	return ru
}

// AddDirections adds the directions edges to Direction.
func (ru *ReceiptUpdate) AddDirections(d ...*Direction) *ReceiptUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.AddDirectionIDs(ids...)
}

// RemoveIngredientIDs removes the ingredients edge to Ingredient by ids.
func (ru *ReceiptUpdate) RemoveIngredientIDs(ids ...int) *ReceiptUpdate {
	if ru.removedIngredients == nil {
		ru.removedIngredients = make(map[int]struct{})
	}
	for i := range ids {
		ru.removedIngredients[ids[i]] = struct{}{}
	}
	return ru
}

// RemoveIngredients removes ingredients edges to Ingredient.
func (ru *ReceiptUpdate) RemoveIngredients(i ...*Ingredient) *ReceiptUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ru.RemoveIngredientIDs(ids...)
}

// RemoveDirectionIDs removes the directions edge to Direction by ids.
func (ru *ReceiptUpdate) RemoveDirectionIDs(ids ...int) *ReceiptUpdate {
	if ru.removedDirections == nil {
		ru.removedDirections = make(map[int]struct{})
	}
	for i := range ids {
		ru.removedDirections[ids[i]] = struct{}{}
	}
	return ru
}

// RemoveDirections removes directions edges to Direction.
func (ru *ReceiptUpdate) RemoveDirections(d ...*Direction) *ReceiptUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ru.RemoveDirectionIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *ReceiptUpdate) Save(ctx context.Context) (int, error) {
	return ru.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReceiptUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReceiptUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReceiptUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ReceiptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(receipt.FieldID).From(sql.Table(receipt.Table))
	for _, p := range ru.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = ru.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := ru.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		builder = sql.Update(receipt.Table).Where(sql.InInts(receipt.FieldID, ids...))
	)
	if value := ru.name; value != nil {
		builder.Set(receipt.FieldName, *value)
	}
	if value := ru.portions; value != nil {
		builder.Set(receipt.FieldPortions, *value)
	}
	if value := ru.addportions; value != nil {
		builder.Add(receipt.FieldPortions, *value)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(ru.removedIngredients) > 0 {
		eids := make([]int, len(ru.removedIngredients))
		for eid := range ru.removedIngredients {
			eids = append(eids, eid)
		}
		query, args := sql.Update(receipt.IngredientsTable).
			SetNull(receipt.IngredientsColumn).
			Where(sql.InInts(receipt.IngredientsColumn, ids...)).
			Where(sql.InInts(ingredient.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(ru.ingredients) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range ru.ingredients {
				p.Or().EQ(ingredient.FieldID, eid)
			}
			query, args := sql.Update(receipt.IngredientsTable).
				Set(receipt.IngredientsColumn, id).
				Where(sql.And(p, sql.IsNull(receipt.IngredientsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(ru.ingredients) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"ingredients\" %v already connected to a different \"Receipt\"", keys(ru.ingredients))})
			}
		}
	}
	if len(ru.removedDirections) > 0 {
		eids := make([]int, len(ru.removedDirections))
		for eid := range ru.removedDirections {
			eids = append(eids, eid)
		}
		query, args := sql.Update(receipt.DirectionsTable).
			SetNull(receipt.DirectionsColumn).
			Where(sql.InInts(receipt.DirectionsColumn, ids...)).
			Where(sql.InInts(direction.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(ru.directions) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range ru.directions {
				p.Or().EQ(direction.FieldID, eid)
			}
			query, args := sql.Update(receipt.DirectionsTable).
				Set(receipt.DirectionsColumn, id).
				Where(sql.And(p, sql.IsNull(receipt.DirectionsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(ru.directions) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"directions\" %v already connected to a different \"Receipt\"", keys(ru.directions))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// ReceiptUpdateOne is the builder for updating a single Receipt entity.
type ReceiptUpdateOne struct {
	config
	id                 int
	name               *string
	portions           *int
	addportions        *int
	ingredients        map[int]struct{}
	directions         map[int]struct{}
	removedIngredients map[int]struct{}
	removedDirections  map[int]struct{}
}

// SetName sets the name field.
func (ruo *ReceiptUpdateOne) SetName(s string) *ReceiptUpdateOne {
	ruo.name = &s
	return ruo
}

// SetPortions sets the portions field.
func (ruo *ReceiptUpdateOne) SetPortions(i int) *ReceiptUpdateOne {
	ruo.portions = &i
	ruo.addportions = nil
	return ruo
}

// AddPortions adds i to portions.
func (ruo *ReceiptUpdateOne) AddPortions(i int) *ReceiptUpdateOne {
	if ruo.addportions == nil {
		ruo.addportions = &i
	} else {
		*ruo.addportions += i
	}
	return ruo
}

// AddIngredientIDs adds the ingredients edge to Ingredient by ids.
func (ruo *ReceiptUpdateOne) AddIngredientIDs(ids ...int) *ReceiptUpdateOne {
	if ruo.ingredients == nil {
		ruo.ingredients = make(map[int]struct{})
	}
	for i := range ids {
		ruo.ingredients[ids[i]] = struct{}{}
	}
	return ruo
}

// AddIngredients adds the ingredients edges to Ingredient.
func (ruo *ReceiptUpdateOne) AddIngredients(i ...*Ingredient) *ReceiptUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ruo.AddIngredientIDs(ids...)
}

// AddDirectionIDs adds the directions edge to Direction by ids.
func (ruo *ReceiptUpdateOne) AddDirectionIDs(ids ...int) *ReceiptUpdateOne {
	if ruo.directions == nil {
		ruo.directions = make(map[int]struct{})
	}
	for i := range ids {
		ruo.directions[ids[i]] = struct{}{}
	}
	return ruo
}

// AddDirections adds the directions edges to Direction.
func (ruo *ReceiptUpdateOne) AddDirections(d ...*Direction) *ReceiptUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.AddDirectionIDs(ids...)
}

// RemoveIngredientIDs removes the ingredients edge to Ingredient by ids.
func (ruo *ReceiptUpdateOne) RemoveIngredientIDs(ids ...int) *ReceiptUpdateOne {
	if ruo.removedIngredients == nil {
		ruo.removedIngredients = make(map[int]struct{})
	}
	for i := range ids {
		ruo.removedIngredients[ids[i]] = struct{}{}
	}
	return ruo
}

// RemoveIngredients removes ingredients edges to Ingredient.
func (ruo *ReceiptUpdateOne) RemoveIngredients(i ...*Ingredient) *ReceiptUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ruo.RemoveIngredientIDs(ids...)
}

// RemoveDirectionIDs removes the directions edge to Direction by ids.
func (ruo *ReceiptUpdateOne) RemoveDirectionIDs(ids ...int) *ReceiptUpdateOne {
	if ruo.removedDirections == nil {
		ruo.removedDirections = make(map[int]struct{})
	}
	for i := range ids {
		ruo.removedDirections[ids[i]] = struct{}{}
	}
	return ruo
}

// RemoveDirections removes directions edges to Direction.
func (ruo *ReceiptUpdateOne) RemoveDirections(d ...*Direction) *ReceiptUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return ruo.RemoveDirectionIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ruo *ReceiptUpdateOne) Save(ctx context.Context) (*Receipt, error) {
	return ruo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReceiptUpdateOne) SaveX(ctx context.Context) *Receipt {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *ReceiptUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReceiptUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ReceiptUpdateOne) sqlSave(ctx context.Context) (r *Receipt, err error) {
	selector := sql.Select(receipt.Columns...).From(sql.Table(receipt.Table))
	receipt.ID(ruo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = ruo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		r = &Receipt{config: ruo.config}
		if err := r.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Receipt: %v", err)
		}
		id = r.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Receipt with id: %v", ruo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Receipt with the same id: %v", ruo.id)
	}

	tx, err := ruo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		builder = sql.Update(receipt.Table).Where(sql.InInts(receipt.FieldID, ids...))
	)
	if value := ruo.name; value != nil {
		builder.Set(receipt.FieldName, *value)
		r.Name = *value
	}
	if value := ruo.portions; value != nil {
		builder.Set(receipt.FieldPortions, *value)
		r.Portions = *value
	}
	if value := ruo.addportions; value != nil {
		builder.Add(receipt.FieldPortions, *value)
		r.Portions += *value
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(ruo.removedIngredients) > 0 {
		eids := make([]int, len(ruo.removedIngredients))
		for eid := range ruo.removedIngredients {
			eids = append(eids, eid)
		}
		query, args := sql.Update(receipt.IngredientsTable).
			SetNull(receipt.IngredientsColumn).
			Where(sql.InInts(receipt.IngredientsColumn, ids...)).
			Where(sql.InInts(ingredient.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(ruo.ingredients) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range ruo.ingredients {
				p.Or().EQ(ingredient.FieldID, eid)
			}
			query, args := sql.Update(receipt.IngredientsTable).
				Set(receipt.IngredientsColumn, id).
				Where(sql.And(p, sql.IsNull(receipt.IngredientsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(ruo.ingredients) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"ingredients\" %v already connected to a different \"Receipt\"", keys(ruo.ingredients))})
			}
		}
	}
	if len(ruo.removedDirections) > 0 {
		eids := make([]int, len(ruo.removedDirections))
		for eid := range ruo.removedDirections {
			eids = append(eids, eid)
		}
		query, args := sql.Update(receipt.DirectionsTable).
			SetNull(receipt.DirectionsColumn).
			Where(sql.InInts(receipt.DirectionsColumn, ids...)).
			Where(sql.InInts(direction.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(ruo.directions) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range ruo.directions {
				p.Or().EQ(direction.FieldID, eid)
			}
			query, args := sql.Update(receipt.DirectionsTable).
				Set(receipt.DirectionsColumn, id).
				Where(sql.And(p, sql.IsNull(receipt.DirectionsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(ruo.directions) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"directions\" %v already connected to a different \"Receipt\"", keys(ruo.directions))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return r, nil
}