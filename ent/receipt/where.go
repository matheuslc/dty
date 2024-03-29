// Code generated by entc, DO NOT EDIT.

package receipt

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/matheuslc/dty/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
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
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
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
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
	)
}

// Portions applies equality check predicate on the "portions" field. It's identical to PortionsEQ.
func Portions(v int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldPortions), v))
		},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldName), v))
		},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldName), v...))
		},
	)
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldName), v...))
		},
	)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldName), v))
		},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldName), v))
		},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldName), v))
		},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldName), v))
		},
	)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldName), v))
		},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldName), v))
		},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldName), v))
		},
	)
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldName), v))
		},
	)
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldName), v))
		},
	)
}

// PortionsEQ applies the EQ predicate on the "portions" field.
func PortionsEQ(v int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldPortions), v))
		},
	)
}

// PortionsNEQ applies the NEQ predicate on the "portions" field.
func PortionsNEQ(v int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldPortions), v))
		},
	)
}

// PortionsIn applies the In predicate on the "portions" field.
func PortionsIn(vs ...int) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldPortions), v...))
		},
	)
}

// PortionsNotIn applies the NotIn predicate on the "portions" field.
func PortionsNotIn(vs ...int) predicate.Receipt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Receipt(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldPortions), v...))
		},
	)
}

// PortionsGT applies the GT predicate on the "portions" field.
func PortionsGT(v int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldPortions), v))
		},
	)
}

// PortionsGTE applies the GTE predicate on the "portions" field.
func PortionsGTE(v int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldPortions), v))
		},
	)
}

// PortionsLT applies the LT predicate on the "portions" field.
func PortionsLT(v int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldPortions), v))
		},
	)
}

// PortionsLTE applies the LTE predicate on the "portions" field.
func PortionsLTE(v int) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldPortions), v))
		},
	)
}

// HasIngredients applies the HasEdge predicate on the "ingredients" edge.
func HasIngredients() predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(IngredientsColumn).
						From(sql.Table(IngredientsTable)).
						Where(sql.NotNull(IngredientsColumn)),
				),
			)
		},
	)
}

// HasIngredientsWith applies the HasEdge predicate on the "ingredients" edge with a given conditions (other predicates).
func HasIngredientsWith(preds ...predicate.Ingredient) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(IngredientsColumn).From(sql.Table(IngredientsTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(FieldID), t2))
		},
	)
}

// HasDirections applies the HasEdge predicate on the "directions" edge.
func HasDirections() predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(DirectionsColumn).
						From(sql.Table(DirectionsTable)).
						Where(sql.NotNull(DirectionsColumn)),
				),
			)
		},
	)
}

// HasDirectionsWith applies the HasEdge predicate on the "directions" edge with a given conditions (other predicates).
func HasDirectionsWith(preds ...predicate.Direction) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(DirectionsColumn).From(sql.Table(DirectionsTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(FieldID), t2))
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Receipt) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			for _, p := range predicates {
				p(s)
			}
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Receipt) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			for i, p := range predicates {
				if i > 0 {
					s.Or()
				}
				p(s)
			}
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Receipt) predicate.Receipt {
	return predicate.Receipt(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
