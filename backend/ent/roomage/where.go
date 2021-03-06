// Code generated by entc, DO NOT EDIT.

package roomage

import (
	"github.com/PON/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
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
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
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
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// RoomAge applies equality check predicate on the "RoomAge" field. It's identical to RoomAgeEQ.
func RoomAge(v int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomAge), v))
	})
}

// Text applies equality check predicate on the "Text" field. It's identical to TextEQ.
func Text(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// RoomAgeEQ applies the EQ predicate on the "RoomAge" field.
func RoomAgeEQ(v int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomAge), v))
	})
}

// RoomAgeNEQ applies the NEQ predicate on the "RoomAge" field.
func RoomAgeNEQ(v int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoomAge), v))
	})
}

// RoomAgeIn applies the In predicate on the "RoomAge" field.
func RoomAgeIn(vs ...int) predicate.Roomage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoomAge), v...))
	})
}

// RoomAgeNotIn applies the NotIn predicate on the "RoomAge" field.
func RoomAgeNotIn(vs ...int) predicate.Roomage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoomAge), v...))
	})
}

// RoomAgeGT applies the GT predicate on the "RoomAge" field.
func RoomAgeGT(v int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoomAge), v))
	})
}

// RoomAgeGTE applies the GTE predicate on the "RoomAge" field.
func RoomAgeGTE(v int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoomAge), v))
	})
}

// RoomAgeLT applies the LT predicate on the "RoomAge" field.
func RoomAgeLT(v int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoomAge), v))
	})
}

// RoomAgeLTE applies the LTE predicate on the "RoomAge" field.
func RoomAgeLTE(v int) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoomAge), v))
	})
}

// TextEQ applies the EQ predicate on the "Text" field.
func TextEQ(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldText), v))
	})
}

// TextNEQ applies the NEQ predicate on the "Text" field.
func TextNEQ(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldText), v))
	})
}

// TextIn applies the In predicate on the "Text" field.
func TextIn(vs ...string) predicate.Roomage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldText), v...))
	})
}

// TextNotIn applies the NotIn predicate on the "Text" field.
func TextNotIn(vs ...string) predicate.Roomage {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomage(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldText), v...))
	})
}

// TextGT applies the GT predicate on the "Text" field.
func TextGT(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldText), v))
	})
}

// TextGTE applies the GTE predicate on the "Text" field.
func TextGTE(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldText), v))
	})
}

// TextLT applies the LT predicate on the "Text" field.
func TextLT(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldText), v))
	})
}

// TextLTE applies the LTE predicate on the "Text" field.
func TextLTE(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldText), v))
	})
}

// TextContains applies the Contains predicate on the "Text" field.
func TextContains(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldText), v))
	})
}

// TextHasPrefix applies the HasPrefix predicate on the "Text" field.
func TextHasPrefix(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldText), v))
	})
}

// TextHasSuffix applies the HasSuffix predicate on the "Text" field.
func TextHasSuffix(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldText), v))
	})
}

// TextEqualFold applies the EqualFold predicate on the "Text" field.
func TextEqualFold(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldText), v))
	})
}

// TextContainsFold applies the ContainsFold predicate on the "Text" field.
func TextContainsFold(v string) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldText), v))
	})
}

// HasRoomageRent applies the HasEdge predicate on the "RoomageRent" edge.
func HasRoomageRent() predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomageRentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomageRentTable, RoomageRentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomageRentWith applies the HasEdge predicate on the "RoomageRent" edge with a given conditions (other predicates).
func HasRoomageRentWith(preds ...predicate.Rent) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomageRentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomageRentTable, RoomageRentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Roomage) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Roomage) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Roomage) predicate.Roomage {
	return predicate.Roomage(func(s *sql.Selector) {
		p(s.Not())
	})
}
