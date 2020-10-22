package controllers

import (
	"context"

	"strconv"

	"github.com/PON/app/ent"
	"github.com/PON/app/ent/roomage"
	"github.com/gin-gonic/gin"
)

// RoomageController defines the struct for the Roomage controller
type RoomageController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateRoomage handles POST requests for adding Roomage entities
// @Summary Create Roomage
// @Description Create Roomage
// @ID create-Roomage
// @Accept   json
// @Produce  json
// @Param Roomage body ent.Roomage true "Roomage entity"
// @Success 200 {object} ent.Roomage
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Roomages [post]
func (ctl *RoomageController) CreateRoomage(c *gin.Context) {
	obj := ent.Roomage{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "roomage binding failed",
		})
		return
	}

	u, err := ctl.client.Roomage.
		Create().
		SetRoomAge(obj.RoomAge).
		SetText(obj.Text).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetRoomage handles GET requests to retrieve a roomage entity
// @Summary Get a Roomage entity by ID
// @Description get Roomage by ID
// @ID get-roomage
// @Produce  json
// @Param id path int true "Roomage ID"
// @Success 200 {object} ent.Roomage
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomages/{id} [get]
func (ctl *RoomageController) GetRoomage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Roomage.
		Query().
		Where(roomage.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListRoomage handles request to get a list of roomage entities
// @Summary List roomage entities
// @Description list roomage entities
// @ID list-roomage
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Roomage
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomages [get]
func (ctl *RoomageController) ListRoomage(c *gin.Context) {
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

	roomages, err := ctl.client.Roomage.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, roomages)
}

// UpdateRoomage handles PUT requests to update a roomage entity
// @Summary Update a roomage entity by ID
// @Description update roomage by ID
// @ID update-roomage
// @Accept   json
// @Produce  json
// @Param id path int true "Roomage ID"
// @Param roomage body ent.Roomage true "Roomage entity"
// @Success 200 {object} ent.Roomage
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomages/{id} [put]
func (ctl *RoomageController) UpdateRoomage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Roomage{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "roomage binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Roomage.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewRoomageController creates and registers handles for the roomage controller
func NewRoomageController(router gin.IRouter, client *ent.Client) *RoomageController {
	uc := &RoomageController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitRoomageController registers routes to the main engine
func (ctl *RoomageController) register() {
	roomages := ctl.router.Group("/roomages")

	roomages.GET("", ctl.ListRoomage)

	// CRUD
	roomages.POST("", ctl.CreateRoomage)
	roomages.GET(":id", ctl.GetRoomage)
	roomages.PUT(":id", ctl.UpdateRoomage)
	//roomages.DELETE(":id", ctl.DeleteRoomage)
}
