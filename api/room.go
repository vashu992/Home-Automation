package api

import (
	"github.com/gin-gonic/gin"
)

func (api APIRoutes) RoomRouts(router *gin.Engine) {
	// Define routes
	roomapi := router.Group("/room")
	{
		roomapi.GET("/all", api.GetAllRooms)
		roomapi.GET("/filter", api.AuthMiddlewareComplete(), api.GetRoomsByFilter)
		roomapi.GET("/:id", api.GetRoom)
		roomapi.POST("/create", api.CreateRoom)
		roomapi.PUT("/update/:id", api.UpdateRoom)
		roomapi.DELETE("/delete/:id", api.DeleteRoom)
	}

}

// Handler to get all rooms
// @router /room/all [get]
// @summary Get all rooms
// @tags rooms
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Room
func (api APIRoutes) GetAllRooms(c *gin.Context) {
	api.Server.GetRooms(c)
}

// Handler to get all rooms based on filter
// @router /room/filter [get]
// @summary Get all rooms based on given filters
// @tags rooms
// @produce json
// @param id query string false "id"
// @param house_id query string false "house_id"
// @param floor_id query string false "floor_id"
// @param name query string false "name"
// @param location query string false "location"
// @param area query string false "area"
// @param occupancy query string false "occupancy"
// @param heating_enabled query boolean false "heating_enabled"
// @param heatin_type query string false "heatin_type"
// @param cooling_enabled query boolean false "cooling_enabled"
// @param cooling_type query string false "cooling_type"
// @param ventilation_enabled query boolean false "ventilation_enabled"
// @param ventilation_type query string false "ventilation_type"
// @param temperature query string false "temperature"
// @param humidity query string false "humidity"
// @param light_level query string false "light_level"
// @param co2_level query string false "co2_level"
// @param created_at query string false "created_at"
// @param created_by query string false "created_by"
// @param updated_at query string false "updated_at"
// @param updated_by query string false "updated_by"
// @param start_date query string false "start_date"
// @param end_date query string false "end_date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Room
// @Security ApiKeyAuth
func (api APIRoutes) GetRoomsByFilter(c *gin.Context) {
	api.Server.GetRoomsByFilter(c)
}

// Handler to get a room by ID
// @router /room/{id} [get]
// @summary Get a room by ID
// @tags rooms
// @produce json
// @param id path string true "Room ID"
// @success 200 {object} model.Room
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) GetRoom(c *gin.Context) {
	api.Server.GetRoom(c)
}

// Handler to create a room
// @router /room/create [post]
// @summary Create a room
// @tags rooms
// @accept json
// @produce json
// @param room body model.Room true "Room object"
// @success 201 {object} model.Room
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) CreateRoom(c *gin.Context) {
	api.Server.CreateRoom(c)
}

// Handler to update a room
// @router /room/update/{id} [put]
// @summary Update a room
// @tags rooms
// @accept json
// @produce json
// @param id path string true "Room ID"
// @param room body model.Room true "Room object"
// @success 200 {object} model.Room
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) UpdateRoom(c *gin.Context) {
	api.Server.UpdateRoom(c)
}

// Handler to delete a room
// @router  /room/delete/{id} [delete]
// @summary Delete a room
// @tags rooms
// @param id path string true "Room ID"
// @success 200 {object} model.SuccessResponse
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) DeleteRoom(c *gin.Context) {
	api.Server.DeleteRoom(c)
}