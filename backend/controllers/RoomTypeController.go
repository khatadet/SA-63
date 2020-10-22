package controllers

import (
   "context"
   
   "strconv"
   "github.com/PON/app/ent"
   "github.com/PON/app/ent/roomtype"
   "github.com/gin-gonic/gin"
)

// RoomTypeController defines the struct for the RoomType controller
type RoomTypeController struct {
   client *ent.Client
   router gin.IRouter
}

// CreateRoomType handles POST requests for adding RoomType entities
// @Summary Create RoomType
// @Description Create RoomType
// @ID create-RoomType
// @Accept   json
// @Produce  json
// @Param RoomType body ent.RoomType true "RoomType entity"
// @Success 200 {object} ent.RoomType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /RoomTypes [post]
func (ctl *RoomTypeController) CreateRoomType(c *gin.Context) {
   obj := ent.RoomType{}
   if err := c.ShouldBind(&obj); err != nil {
       c.JSON(400, gin.H{
           "error": "roomtype binding failed",
       })
       return
   }

   u, err := ctl.client.RoomType.
       Create().
       
       SetRoomType(obj.RoomType).

       SetCOST(obj.COST).
       Save(context.Background())
   if err != nil {
       c.JSON(400, gin.H{
           "error": "saving failed",
       })
       return
   }

   c.JSON(200, u)
}
// GetRoomType handles GET requests to retrieve a roomtype entity
// @Summary Get a RoomType entity by ID
// @Description get RoomType by ID
// @ID get-roomtype
// @Produce  json
// @Param id path int true "RoomType ID"
// @Success 200 {object} ent.RoomType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomtypes/{id} [get]
func (ctl *RoomTypeController) GetRoomType(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }

   u, err := ctl.client.RoomType.
       Query().
       Where(roomtype.IDEQ(int(id))).
       Only(context.Background())
   if err != nil {
       c.JSON(404, gin.H{
           "error": err.Error(),
       })
       return
   }

   c.JSON(200, u)
}

// ListRoomType handles request to get a list of roomtype entities
// @Summary List roomtype entities
// @Description list roomtype entities
// @ID list-roomtype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.RoomType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomtypes [get]
func (ctl *RoomTypeController) ListRoomType(c *gin.Context) {
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

   roomtypes, err := ctl.client.RoomType.
       Query().
       Limit(limit).
       Offset(offset).
       All(context.Background())
   	if err != nil {
       c.JSON(400, gin.H{"error": err.Error(),})
       return
   }

   c.JSON(200, roomtypes)
}



// UpdateRoomType handles PUT requests to update a roomtype entity
// @Summary Update a roomtype entity by ID
// @Description update roomtype by ID
// @ID update-roomtype
// @Accept   json
// @Produce  json
// @Param id path int true "RoomType ID"
// @Param roomtype body ent.RoomType true "RoomType entity"
// @Success 200 {object} ent.RoomType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roomtypes/{id} [put]
func (ctl *RoomTypeController) UpdateRoomType(c *gin.Context) {
   id, err := strconv.ParseInt(c.Param("id"), 10, 64)
   if err != nil {
       c.JSON(400, gin.H{
           "error": err.Error(),
       })
       return
   }

   obj := ent.RoomType{}
   if err := c.ShouldBind(&obj); err != nil {
       c.JSON(400, gin.H{
           "error": "roomtype binding failed",
       })
       return
   }
   obj.ID = int(id)
   u, err := ctl.client.RoomType.
       UpdateOne(&obj).
       Save(context.Background())
   if err != nil {
       c.JSON(400, gin.H{"error": "update failed",})
       return
   }

   c.JSON(200, u)
}

// NewRoomTypeController creates and registers handles for the roomtype controller
func NewRoomTypeController(router gin.IRouter, client *ent.Client) *RoomTypeController {
   uc := &RoomTypeController{
       client: client,
       router: router,
   }
   uc.register()
   return uc
}

// InitRoomTypeController registers routes to the main engine
func (ctl *RoomTypeController) register() {
    roomtypes := ctl.router.Group("/roomtypes")

    roomtypes.GET("", ctl.ListRoomType)

   // CRUD
   roomtypes.POST("", ctl.CreateRoomType)
   roomtypes.GET(":id", ctl.GetRoomType)
   roomtypes.PUT(":id", ctl.UpdateRoomType)
   //roomtypes.DELETE(":id", ctl.DeleteRoomType)
}


