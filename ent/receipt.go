// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Receipt is the model entity for the Receipt schema.
type Receipt struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Portions holds the value of the "portions" field.
	Portions int `json:"portions,omitempty"`
}

// FromRows scans the sql response data into Receipt.
func (r *Receipt) FromRows(rows *sql.Rows) error {
	var vr struct {
		ID       int
		Name     sql.NullString
		Portions sql.NullInt64
	}
	// the order here should be the same as in the `receipt.Columns`.
	if err := rows.Scan(
		&vr.ID,
		&vr.Name,
		&vr.Portions,
	); err != nil {
		return err
	}
	r.ID = vr.ID
	r.Name = vr.Name.String
	r.Portions = int(vr.Portions.Int64)
	return nil
}

// QueryIngredients queries the ingredients edge of the Receipt.
func (r *Receipt) QueryIngredients() *IngredientQuery {
	return (&ReceiptClient{r.config}).QueryIngredients(r)
}

// QueryDirections queries the directions edge of the Receipt.
func (r *Receipt) QueryDirections() *DirectionQuery {
	return (&ReceiptClient{r.config}).QueryDirections(r)
}

// Update returns a builder for updating this Receipt.
// Note that, you need to call Receipt.Unwrap() before calling this method, if this Receipt
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Receipt) Update() *ReceiptUpdateOne {
	return (&ReceiptClient{r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Receipt) Unwrap() *Receipt {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Receipt is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Receipt) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Receipt(")
	buf.WriteString(fmt.Sprintf("id=%v", r.ID))
	buf.WriteString(fmt.Sprintf(", name=%v", r.Name))
	buf.WriteString(fmt.Sprintf(", portions=%v", r.Portions))
	buf.WriteString(")")
	return buf.String()
}

// Receipts is a parsable slice of Receipt.
type Receipts []*Receipt

// FromRows scans the sql response data into Receipts.
func (r *Receipts) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vr := &Receipt{}
		if err := vr.FromRows(rows); err != nil {
			return err
		}
		*r = append(*r, vr)
	}
	return nil
}

func (r Receipts) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
