// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/matheuslc/dty/ent/ingredient"
)

// IngredientCreate is the builder for creating a Ingredient entity.
type IngredientCreate struct {
	config
	name     *string
	unitType *string
}

// SetName sets the name field.
func (ic *IngredientCreate) SetName(s string) *IngredientCreate {
	ic.name = &s
	return ic
}

// SetUnitType sets the unitType field.
func (ic *IngredientCreate) SetUnitType(s string) *IngredientCreate {
	ic.unitType = &s
	return ic
}

// Save creates the Ingredient in the database.
func (ic *IngredientCreate) Save(ctx context.Context) (*Ingredient, error) {
	if ic.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if ic.unitType == nil {
		return nil, errors.New("ent: missing required field \"unitType\"")
	}
	return ic.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IngredientCreate) SaveX(ctx context.Context) *Ingredient {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ic *IngredientCreate) sqlSave(ctx context.Context) (*Ingredient, error) {
	var (
		res sql.Result
		i   = &Ingredient{config: ic.config}
	)
	tx, err := ic.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Dialect(ic.driver.Dialect()).
		Insert(ingredient.Table).
		Default()
	if value := ic.name; value != nil {
		builder.Set(ingredient.FieldName, *value)
		i.Name = *value
	}
	if value := ic.unitType; value != nil {
		builder.Set(ingredient.FieldUnitType, *value)
		i.UnitType = *value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	i.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return i, nil
}
