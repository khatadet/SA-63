package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/PON/app/ent"
	"github.com/PON/app/ent/rent"
	"github.com/gin-gonic/gin"

	"github.com/PON/app/ent/insurance"
	"github.com/PON/app/ent/room"
	"github.com/PON/app/ent/roomage"
	"github.com/PON/app/ent/user"
)

// RentController defines the struct for the Rent controller
type RentController struct {
	client *ent.Client
	router gin.IRouter
}

// Rent defines the struct for the Rent
type Rent struct {
	RentAge   string
	CidUser   string
	RentDate  string
	Insurance int
	Room      int
	Roomage   int
	User      int
}

// CreateRent handles POST requests for adding Rent entities
// @Summary Create Rent
// @Description Create Rent
// @ID create-Rent
// @Accept   json
// @Produce  json
// @Param Rent body Rent true "Rent entity"
// @Success 200 {object} Rent
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rents [post]
func (ctl *RentController) CreateRent(c *gin.Context) {
	obj := Rent{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "rent binding failed",
		})
		return
	}

	Room, err := ctl.client.Room.
		Query().
		Where(room.IDEQ(int(obj.Room))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Room not found",
		})
		return
	}

	User, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Room not found",
		})
		return
	}

	Insurance, err := ctl.client.Insurance.
		Query().
		Where(insurance.IDEQ(int(obj.Insurance))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Insurance not found",
		})
		return
	}

	Roomage, err := ctl.client.Roomage.
		Query().
		Where(roomage.IDEQ(int(obj.Roomage))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Insurance not found",
		})
		return
	}

	u, err := ctl.client.Rent.
		Create().
		SetRentRoom(Room).
		SetRentUser(User).
		SetRentInsurance(Insurance).
		SetRentRoomage(Roomage).
		SetRentAge(obj.RentAge).
		SetCidUser(obj.CidUser).
		SetRentDate(obj.RentDate).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetRent handles GET requests to retrieve a rent entity
// @Summary Get a Rent entity by ID
// @Description get Rent by ID
// @ID get-rent
// @Produce  json
// @Param id path int true "Rent ID"
// @Success 200 {object} ent.Rent
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rents/{id} [get]
func (ctl *RentController) GetRent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Rent.
		Query().
		Where(rent.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// DeleteRent handles DELETE requests to delete a rent entity
// @Summary Delete a rent entity by ID
// @Description get rent by ID
// @ID delete-rent
// @Produce  json
// @Param id path int true "Rent ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rents/{id} [delete]
func (ctl *RentController) DeleteRent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Rent.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// ListRent handles request to get a list of rent entities
// @Summary List rent entities
// @Description list rent entities
// @ID list-rent
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Rent
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /rents [get]
func (ctl *RentController) ListRent(c *gin.Context) {
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

	rents, err := ctl.client.Rent.
		Query().
		WithRentInsurance().
		WithRentRoom().
		WithRentRoomage().
		WithRentUser().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, rents)
}

// NewRentController creates and registers handles for the rent controller
func NewRentController(router gin.IRouter, client *ent.Client) *RentController {
	uc := &RentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitRentController registers routes to the main engine
func (ctl *RentController) register() {
	rents := ctl.router.Group("/rents")

	rents.GET("", ctl.ListRent)

	// CRUD
	rents.POST("", ctl.CreateRent)
	rents.GET(":id", ctl.GetRent)
	//rents.PUT(":id", ctl.UpdateRent)
	rents.DELETE(":id", ctl.DeleteRent)
}
