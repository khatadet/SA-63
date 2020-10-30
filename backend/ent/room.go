// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/PON/app/ent/room"
	"github.com/PON/app/ent/roomstatus"
	"github.com/PON/app/ent/roomtype"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Room is the model entity for the Room schema.
type Room struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RoomName holds the value of the "RoomName" field.
	RoomName string `json:"RoomName,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomQuery when eager-loading is set.
	Edges                       RoomEdges `json:"edges"`
	room_status_roomstatus_room *int
	room_type_roomtype_room     *int
}

// RoomEdges holds the relations/edges for other nodes in the graph.
type RoomEdges struct {
	// RoomRoomstatus holds the value of the RoomRoomstatus edge.
	RoomRoomstatus *RoomStatus
	// RoomRoomtype holds the value of the RoomRoomtype edge.
	RoomRoomtype *RoomType
	// RoomRent holds the value of the RoomRent edge.
	RoomRent []*Rent
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// RoomRoomstatusOrErr returns the RoomRoomstatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomEdges) RoomRoomstatusOrErr() (*RoomStatus, error) {
	if e.loadedTypes[0] {
		if e.RoomRoomstatus == nil {
			// The edge RoomRoomstatus was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: roomstatus.Label}
		}
		return e.RoomRoomstatus, nil
	}
	return nil, &NotLoadedError{edge: "RoomRoomstatus"}
}

// RoomRoomtypeOrErr returns the RoomRoomtype value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomEdges) RoomRoomtypeOrErr() (*RoomType, error) {
	if e.loadedTypes[1] {
		if e.RoomRoomtype == nil {
			// The edge RoomRoomtype was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: roomtype.Label}
		}
		return e.RoomRoomtype, nil
	}
	return nil, &NotLoadedError{edge: "RoomRoomtype"}
}

// RoomRentOrErr returns the RoomRent value or an error if the edge
// was not loaded in eager-loading.
func (e RoomEdges) RoomRentOrErr() ([]*Rent, error) {
	if e.loadedTypes[2] {
		return e.RoomRent, nil
	}
	return nil, &NotLoadedError{edge: "RoomRent"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Room) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // RoomName
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Room) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // room_status_roomstatus_room
		&sql.NullInt64{}, // room_type_roomtype_room
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Room fields.
func (r *Room) assignValues(values ...interface{}) error {
	if m, n := len(values), len(room.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field RoomName", values[0])
	} else if value.Valid {
		r.RoomName = value.String
	}
	values = values[1:]
	if len(values) == len(room.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field room_status_roomstatus_room", value)
		} else if value.Valid {
			r.room_status_roomstatus_room = new(int)
			*r.room_status_roomstatus_room = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field room_type_roomtype_room", value)
		} else if value.Valid {
			r.room_type_roomtype_room = new(int)
			*r.room_type_roomtype_room = int(value.Int64)
		}
	}
	return nil
}

// QueryRoomRoomstatus queries the RoomRoomstatus edge of the Room.
func (r *Room) QueryRoomRoomstatus() *RoomStatusQuery {
	return (&RoomClient{config: r.config}).QueryRoomRoomstatus(r)
}

// QueryRoomRoomtype queries the RoomRoomtype edge of the Room.
func (r *Room) QueryRoomRoomtype() *RoomTypeQuery {
	return (&RoomClient{config: r.config}).QueryRoomRoomtype(r)
}

// QueryRoomRent queries the RoomRent edge of the Room.
func (r *Room) QueryRoomRent() *RentQuery {
	return (&RoomClient{config: r.config}).QueryRoomRent(r)
}

// Update returns a builder for updating this Room.
// Note that, you need to call Room.Unwrap() before calling this method, if this Room
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Room) Update() *RoomUpdateOne {
	return (&RoomClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Room) Unwrap() *Room {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Room is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Room) String() string {
	var builder strings.Builder
	builder.WriteString("Room(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", RoomName=")
	builder.WriteString(r.RoomName)
	builder.WriteByte(')')
	return builder.String()
}

// Rooms is a parsable slice of Room.
type Rooms []*Room

func (r Rooms) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
