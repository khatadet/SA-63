package controllers

import (
   "context"
   
   "strconv"
   "github.com/PON/app/ent"
   "github.com/PON/app/ent/roomstatus"
   "github.com/gin-gonic/gin"
)

// RoomStatusController defines the struct for the RoomStatus controller
type RoomStatusController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateRoomStatus handles POST requests for adding RoomStatus entities
// @Summary Create RoomStatus
// @Description Create RoomStatus
// @ID create-RoomStatus
// @Accept   json
// @Produce  json
// @Param RoomStatus body ent.RoomStatus true "RoomStatus entity"
// @Success 200 {object} ent.RoomStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /RoomStatuss [post]
func (ctl *RoomStatusController) CreateRoomStatus(c *gin.Context) {
   obj := ent.RoomStatus{}
   if err := c.ShouldBind(&obj); err != nil {
       c.JSON(400, gin.H{
           "error": "roomstatus binding failed",
       })
       return
   }

   u, err := ctl.client.RoomStatus.
       Create().
       SetRoomStatus(obj.RoomStatus).

       Save(context.Background())
   if err != nil {
       c.JSON(400, gin.H{
           "error": "saving failed",
       })
       return
   }

   c.JSON(200, u)
}
// GetRoomStatus handles GET requests to retrieve a roomstatus entity
// @Summary Get a RoomStatus entity by ID
// @Description get RoomStatus by ID
// @ID get-roomstatus
// @Produce  json
// @Param id path int true "RoomStatus ID"
// @Success 200 {object} ent.RoomStatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstatuss/{id} [get]
func (ctl *RoomStatusController) GetRoomStatus(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }

   u, err := ctl.client.RoomStatus.
       Query().
       Where(roomstatus.IDEQ(int(id))).
       Only(context.Background())
   if err != nil {
       c.JSON(404, gin.H{
           "error": err.Error(),
       })
       return
   }

   c.JSON(200, u)
}

// ListRoomStatus handles request to get a list of roomstatus entities
// @Summary List roomstatus entities
// @Description list roomstatus entities
// @ID list-roomstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.RoomStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstatuss [get]
func (ctl *RoomStatusController) ListRoomStatus(c *gin.Context) {
   limitQuery := c.Query("limit")
   limit := 10
   if limitQuery != "" {
       limit64, err := strconv.ParseInt(limitQuery, 10, 64)
       if err == nil {limit = int(limit64)}
   }

   offsetQuery := c.Query("offset")
   offset := 0
   if offsetQuery != "" {
       offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
       if err == nil {offset = int(offset64)}
   }

   roomstatuss, err := ctl.client.RoomStatus.
       Query().
       Limit(limit).
       Offset(offset).
       All(context.Background())
   	if err != nil {
       c.JSON(400, gin.H{"error": err.Error(),})
       return
   }

   c.JSON(200, roomstatuss)
}


// UpdateRoomStatus handles PUT requests to update a roomstatus entity
// @Summary Update a roomstatus entity by ID
// @Description update roomstatus by ID
// @ID update-roomstatus
// @Accept   json
// @Produce  json
// @Param id path int true "RoomStatus ID"
// @Param roomstatus body ent.RoomStatus true "RoomStatus entity"
// @Success 200 {object} ent.RoomStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomstatuss/{id} [put]
func (ctl *RoomStatusController) UpdateRoomStatus(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }

   obj := ent.RoomStatus{}
   if err := c.ShouldBind(&obj); err != nil {
       c.JSON(400, gin.H{
           "error": "roomstatus binding failed",
       })
       return
   }
   obj.ID = int(id)
   u, err := ctl.client.RoomStatus.
       UpdateOne(&obj).
       Save(context.Background())
   if err != nil {
       c.JSON(400, gin.H{"error": "update failed",})
       return
   }

   c.JSON(200, u)
}

// NewRoomStatusController creates and registers handles for the roomstatus controller
func NewRoomStatusController(router gin.IRouter, client *ent.Client) *RoomStatusController {
   uc := &RoomStatusController{
       client: client,
       router: router,
   }
   uc.register()
   return uc
}

// InitRoomStatusController registers routes to the main engine
func (ctl *RoomStatusController) register() {
    roomstatuss := ctl.router.Group("/roomstatuss")

    roomstatuss.GET("", ctl.ListRoomStatus)

   // CRUD
   roomstatuss.POST("", ctl.CreateRoomStatus)
   roomstatuss.GET(":id", ctl.GetRoomStatus)
   roomstatuss.PUT(":id", ctl.UpdateRoomStatus)
   //roomstatuss.DELETE(":id", ctl.DeleteRoomStatus)
}


