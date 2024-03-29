// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Direction is the model entity for the Direction schema.
type Direction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// FromRows scans the sql response data into Direction.
func (d *Direction) FromRows(rows *sql.Rows) error {
	var vd struct {
		ID          int
		Description sql.NullString
		Name        sql.NullString
	}
	// the order here should be the same as in the `direction.Columns`.
	if err := rows.Scan(
		&vd.ID,
		&vd.Description,
		&vd.Name,
	); err != nil {
		return err
	}
	d.ID = vd.ID
	d.Description = vd.Description.String
	d.Name = vd.Name.String
	return nil
}

// Update returns a builder for updating this Direction.
// Note that, you need to call Direction.Unwrap() before calling this method, if this Direction
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Direction) Update() *DirectionUpdateOne {
	return (&DirectionClient{d.config}).UpdateOne(d)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (d *Direction) Unwrap() *Direction {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Direction is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Direction) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Direction(")
	buf.WriteString(fmt.Sprintf("id=%v", d.ID))
	buf.WriteString(fmt.Sprintf(", description=%v", d.Description))
	buf.WriteString(fmt.Sprintf(", name=%v", d.Name))
	buf.WriteString(")")
	return buf.String()
}

// Directions is a parsable slice of Direction.
type Directions []*Direction

// FromRows scans the sql response data into Directions.
func (d *Directions) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vd := &Direction{}
		if err := vd.FromRows(rows); err != nil {
			return err
		}
		*d = append(*d, vd)
	}
	return nil
}

func (d Directions) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
