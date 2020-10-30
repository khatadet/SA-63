package controllers

import (
	"context"

	"strconv"

	"github.com/PON/app/ent"
	"github.com/PON/app/ent/room"
	"github.com/PON/app/ent/roomstatus"
	"github.com/PON/app/ent/roomtype"

	"github.com/gin-gonic/gin"
)

// RoomController defines the struct for the Room controller
type RoomController struct {
	client *ent.Client
	router gin.IRouter
}

// Room defines the struct for the Room
type Room struct {
	roomtype   int
	roomstatus int
	RoomName   string
}

// CreateRoom handles POST requests for adding Room entities
// @Summary Create Room
// @Description Create Room
// @ID create-Room
// @Accept   json
// @Produce  json
// @Param Room body ent.Room true "Room entity"
// @Success 200 {object} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Rooms [post]
func (ctl *RoomController) CreateRoom(c *gin.Context) {
	obj := Room{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "room binding failed",
		})
		return
	}

	RoomType, err := ctl.client.RoomType.
		Query().
		Where(roomtype.IDEQ(int(obj.roomtype))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "RoomType not found",
		})
		return
	}

	RoomStatus, err := ctl.client.RoomStatus.
		Query().
		Where(roomstatus.IDEQ(int(obj.roomstatus))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "RoomStatus not found",
		})
		return
	}

	u, err := ctl.client.Room.
		Create().
		SetRoomRoomstatus(RoomStatus).
		SetRoomRoomtype(RoomType).
		SetRoomName(obj.RoomName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetRoom handles GET requests to retrieve a room entity
// @Summary Get a Room entity by ID
// @Description get Room by ID
// @ID get-room
// @Produce  json
// @Param id path int true "Room ID"
// @Success 200 {object} ent.Room
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms/{id} [get]
func (ctl *RoomController) GetRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListRoom handles request to get a list of room entities
// @Summary List room entities
// @Description list room entities
// @ID list-room
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms [get]
func (ctl *RoomController) ListRoom(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	rooms, err := ctl.client.Room.
		Query().
		WithRoomRent().
		WithRoomRoomstatus().
		WithRoomRoomtype().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, rooms)
}

// UpdateRoom handles PUT requests to update a room entity
// @Summary Update a room entity by ID
// @Description update room by ID
// @ID update-room
// @Accept   json
// @Produce  json
// @Param id path int true "Room ID"
// @Param room body ent.Room true "Room entity"
// @Success 200 {object} ent.Room
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rooms/{id} [put]
func (ctl *RoomController) UpdateRoom(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Room{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "room binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Room.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewRoomController creates and registers handles for the room controller
func NewRoomController(router gin.IRouter, client *ent.Client) *RoomController {
	uc := &RoomController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitRoomController registers routes to the main engine
func (ctl *RoomController) register() {
	rooms := ctl.router.Group("/rooms")

	rooms.GET("", ctl.ListRoom)

	// CRUD
	rooms.POST("", ctl.CreateRoom)
	rooms.GET(":id", ctl.GetRoom)
	rooms.PUT(":id", ctl.UpdateRoom)
	//rooms.DELETE(":id", ctl.DeleteRoom)
}
