// Code generated by entc, DO NOT EDIT.

package roomstatus

import (
	"github.com/PON/app/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
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
func IDGT(id int) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// RoomStatus applies equality check predicate on the "RoomStatus" field. It's identical to RoomStatusEQ.
func RoomStatus(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusEQ applies the EQ predicate on the "RoomStatus" field.
func RoomStatusEQ(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusNEQ applies the NEQ predicate on the "RoomStatus" field.
func RoomStatusNEQ(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusIn applies the In predicate on the "RoomStatus" field.
func RoomStatusIn(vs ...string) predicate.RoomStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RoomStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRoomStatus), v...))
	})
}

// RoomStatusNotIn applies the NotIn predicate on the "RoomStatus" field.
func RoomStatusNotIn(vs ...string) predicate.RoomStatus {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RoomStatus(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRoomStatus), v...))
	})
}

// RoomStatusGT applies the GT predicate on the "RoomStatus" field.
func RoomStatusGT(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusGTE applies the GTE predicate on the "RoomStatus" field.
func RoomStatusGTE(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusLT applies the LT predicate on the "RoomStatus" field.
func RoomStatusLT(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusLTE applies the LTE predicate on the "RoomStatus" field.
func RoomStatusLTE(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusContains applies the Contains predicate on the "RoomStatus" field.
func RoomStatusContains(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusHasPrefix applies the HasPrefix predicate on the "RoomStatus" field.
func RoomStatusHasPrefix(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusHasSuffix applies the HasSuffix predicate on the "RoomStatus" field.
func RoomStatusHasSuffix(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusEqualFold applies the EqualFold predicate on the "RoomStatus" field.
func RoomStatusEqualFold(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRoomStatus), v))
	})
}

// RoomStatusContainsFold applies the ContainsFold predicate on the "RoomStatus" field.
func RoomStatusContainsFold(v string) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRoomStatus), v))
	})
}

// HasRoomstatusRoom applies the HasEdge predicate on the "RoomstatusRoom" edge.
func HasRoomstatusRoom() predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomstatusRoomTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomstatusRoomTable, RoomstatusRoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomstatusRoomWith applies the HasEdge predicate on the "RoomstatusRoom" edge with a given conditions (other predicates).
func HasRoomstatusRoomWith(preds ...predicate.Room) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RoomstatusRoomInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoomstatusRoomTable, RoomstatusRoomColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.RoomStatus) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.RoomStatus) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
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
func Not(p predicate.RoomStatus) predicate.RoomStatus {
	return predicate.RoomStatus(func(s *sql.Selector) {
		p(s.Not())
	})
}
