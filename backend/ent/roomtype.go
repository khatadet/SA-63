// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/PON/app/ent/roomtype"
	"github.com/facebookincubator/ent/dialect/sql"
)

// RoomType is the model entity for the RoomType schema.
type RoomType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RoomType holds the value of the "RoomType" field.
	RoomType string `json:"RoomType,omitempty"`
	// COST holds the value of the "COST" field.
	COST int `json:"COST,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomTypeQuery when eager-loading is set.
	Edges RoomTypeEdges `json:"edges"`
}

// RoomTypeEdges holds the relations/edges for other nodes in the graph.
type RoomTypeEdges struct {
	// RoomtypeRoom holds the value of the RoomtypeRoom edge.
	RoomtypeRoom []*Room
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RoomtypeRoomOrErr returns the RoomtypeRoom value or an error if the edge
// was not loaded in eager-loading.
func (e RoomTypeEdges) RoomtypeRoomOrErr() ([]*Room, error) {
	if e.loadedTypes[0] {
		return e.RoomtypeRoom, nil
	}
	return nil, &NotLoadedError{edge: "RoomtypeRoom"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoomType) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // RoomType
		&sql.NullInt64{},  // COST
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoomType fields.
func (rt *RoomType) assignValues(values ...interface{}) error {
	if m, n := len(values), len(roomtype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	rt.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field RoomType", values[0])
	} else if value.Valid {
		rt.RoomType = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field COST", values[1])
	} else if value.Valid {
		rt.COST = int(value.Int64)
	}
	return nil
}

// QueryRoomtypeRoom queries the RoomtypeRoom edge of the RoomType.
func (rt *RoomType) QueryRoomtypeRoom() *RoomQuery {
	return (&RoomTypeClient{config: rt.config}).QueryRoomtypeRoom(rt)
}

// Update returns a builder for updating this RoomType.
// Note that, you need to call RoomType.Unwrap() before calling this method, if this RoomType
// was returned from a transaction, and the transaction was committed or rolled back.
func (rt *RoomType) Update() *RoomTypeUpdateOne {
	return (&RoomTypeClient{config: rt.config}).UpdateOne(rt)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (rt *RoomType) Unwrap() *RoomType {
	tx, ok := rt.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoomType is not a transactional entity")
	}
	rt.config.driver = tx.drv
	return rt
}

// String implements the fmt.Stringer.
func (rt *RoomType) String() string {
	var builder strings.Builder
	builder.WriteString("RoomType(")
	builder.WriteString(fmt.Sprintf("id=%v", rt.ID))
	builder.WriteString(", RoomType=")
	builder.WriteString(rt.RoomType)
	builder.WriteString(", COST=")
	builder.WriteString(fmt.Sprintf("%v", rt.COST))
	builder.WriteByte(')')
	return builder.String()
}

// RoomTypes is a parsable slice of RoomType.
type RoomTypes []*RoomType

func (rt RoomTypes) config(cfg config) {
	for _i := range rt {
		rt[_i].config = cfg
	}
}
