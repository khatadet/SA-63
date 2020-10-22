// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/PON/app/ent/room"
	"github.com/PON/app/ent/roomtype"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// RoomTypeCreate is the builder for creating a RoomType entity.
type RoomTypeCreate struct {
	config
	mutation *RoomTypeMutation
	hooks    []Hook
}

// SetRoomType sets the RoomType field.
func (rtc *RoomTypeCreate) SetRoomType(s string) *RoomTypeCreate {
	rtc.mutation.SetRoomType(s)
	return rtc
}

// SetCOST sets the COST field.
func (rtc *RoomTypeCreate) SetCOST(i int) *RoomTypeCreate {
	rtc.mutation.SetCOST(i)
	return rtc
}

// AddRoomtypeRoomIDs adds the RoomtypeRoom edge to Room by ids.
func (rtc *RoomTypeCreate) AddRoomtypeRoomIDs(ids ...int) *RoomTypeCreate {
	rtc.mutation.AddRoomtypeRoomIDs(ids...)
	return rtc
}

// AddRoomtypeRoom adds the RoomtypeRoom edges to Room.
func (rtc *RoomTypeCreate) AddRoomtypeRoom(r ...*Room) *RoomTypeCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rtc.AddRoomtypeRoomIDs(ids...)
}

// Mutation returns the RoomTypeMutation object of the builder.
func (rtc *RoomTypeCreate) Mutation() *RoomTypeMutation {
	return rtc.mutation
}

// Save creates the RoomType in the database.
func (rtc *RoomTypeCreate) Save(ctx context.Context) (*RoomType, error) {
	if _, ok := rtc.mutation.RoomType(); !ok {
		return nil, &ValidationError{Name: "RoomType", err: errors.New("ent: missing required field \"RoomType\"")}
	}
	if _, ok := rtc.mutation.COST(); !ok {
		return nil, &ValidationError{Name: "COST", err: errors.New("ent: missing required field \"COST\"")}
	}
	var (
		err  error
		node *RoomType
	)
	if len(rtc.hooks) == 0 {
		node, err = rtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rtc.mutation = mutation
			node, err = rtc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rtc.hooks) - 1; i >= 0; i-- {
			mut = rtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rtc *RoomTypeCreate) SaveX(ctx context.Context) *RoomType {
	v, err := rtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rtc *RoomTypeCreate) sqlSave(ctx context.Context) (*RoomType, error) {
	rt, _spec := rtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rtc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	rt.ID = int(id)
	return rt, nil
}

func (rtc *RoomTypeCreate) createSpec() (*RoomType, *sqlgraph.CreateSpec) {
	var (
		rt    = &RoomType{config: rtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: roomtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roomtype.FieldID,
			},
		}
	)
	if value, ok := rtc.mutation.RoomType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: roomtype.FieldRoomType,
		})
		rt.RoomType = value
	}
	if value, ok := rtc.mutation.COST(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roomtype.FieldCOST,
		})
		rt.COST = value
	}
	if nodes := rtc.mutation.RoomtypeRoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   roomtype.RoomtypeRoomTable,
			Columns: []string{roomtype.RoomtypeRoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return rt, _spec
}