// Code generated by entc, DO NOT EDIT.

package roomtype

import (
	"github.com/PON/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
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
func IDGT(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// RoomType applies equality check predicate on the "RoomType" field. It's identical to RoomTypeEQ.
func RoomType(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomType), v))
	})
}

// COST applies equality check predicate on the "COST" field. It's identical to COSTEQ.
func COST(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCOST), v))
	})
}

// RoomTypeEQ applies the EQ predicate on the "RoomType" field.
func RoomTypeEQ(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomType), v))
	})
}

// RoomTypeNEQ applies the NEQ predicate on the "RoomType" field.
func RoomTypeNEQ(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoomType), v))
	})
}

// RoomTypeIn applies the In predicate on the "RoomType" field.
func RoomTypeIn(vs ...string) predicate.RoomType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RoomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoomType), v...))
	})
}

// RoomTypeNotIn applies the NotIn predicate on the "RoomType" field.
func RoomTypeNotIn(vs ...string) predicate.RoomType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RoomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoomType), v...))
	})
}

// RoomTypeGT applies the GT predicate on the "RoomType" field.
func RoomTypeGT(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoomType), v))
	})
}

// RoomTypeGTE applies the GTE predicate on the "RoomType" field.
func RoomTypeGTE(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoomType), v))
	})
}

// RoomTypeLT applies the LT predicate on the "RoomType" field.
func RoomTypeLT(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoomType), v))
	})
}

// RoomTypeLTE applies the LTE predicate on the "RoomType" field.
func RoomTypeLTE(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoomType), v))
	})
}

// RoomTypeContains applies the Contains predicate on the "RoomType" field.
func RoomTypeContains(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoomType), v))
	})
}

// RoomTypeHasPrefix applies the HasPrefix predicate on the "RoomType" field.
func RoomTypeHasPrefix(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoomType), v))
	})
}

// RoomTypeHasSuffix applies the HasSuffix predicate on the "RoomType" field.
func RoomTypeHasSuffix(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoomType), v))
	})
}

// RoomTypeEqualFold applies the EqualFold predicate on the "RoomType" field.
func RoomTypeEqualFold(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoomType), v))
	})
}

// RoomTypeContainsFold applies the ContainsFold predicate on the "RoomType" field.
func RoomTypeContainsFold(v string) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoomType), v))
	})
}

// COSTEQ applies the EQ predicate on the "COST" field.
func COSTEQ(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCOST), v))
	})
}

// COSTNEQ applies the NEQ predicate on the "COST" field.
func COSTNEQ(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCOST), v))
	})
}

// COSTIn applies the In predicate on the "COST" field.
func COSTIn(vs ...int) predicate.RoomType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RoomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCOST), v...))
	})
}

// COSTNotIn applies the NotIn predicate on the "COST" field.
func COSTNotIn(vs ...int) predicate.RoomType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RoomType(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCOST), v...))
	})
}

// COSTGT applies the GT predicate on the "COST" field.
func COSTGT(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCOST), v))
	})
}

// COSTGTE applies the GTE predicate on the "COST" field.
func COSTGTE(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCOST), v))
	})
}

// COSTLT applies the LT predicate on the "COST" field.
func COSTLT(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCOST), v))
	})
}

// COSTLTE applies the LTE predicate on the "COST" field.
func COSTLTE(v int) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCOST), v))
	})
}

// HasRoomtypeRoom applies the HasEdge predicate on the "RoomtypeRoom" edge.
func HasRoomtypeRoom() predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomtypeRoomTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomtypeRoomTable, RoomtypeRoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomtypeRoomWith applies the HasEdge predicate on the "RoomtypeRoom" edge with a given conditions (other predicates).
func HasRoomtypeRoomWith(preds ...predicate.Room) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomtypeRoomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomtypeRoomTable, RoomtypeRoomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.RoomType) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.RoomType) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
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
func Not(p predicate.RoomType) predicate.RoomType {
	return predicate.RoomType(func(s *sql.Selector) {
		p(s.Not())
	})
}
