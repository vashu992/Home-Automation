package api

import "github.com/gin-gonic/gin"

func (api APIRoutes) FloorRouts(router *gin.Engine) {
	// Define routes
	floorapi := router.Group("/floor")
	{
		floorapi.GET("/all", api.GetAllFloors)
		floorapi.GET("/filter", api.AuthMiddlewareComplete(), api.GetFloorsByFilter)
		floorapi.GET("/:id", api.CreateFloor)
		floorapi.POST("/create", api.CreateFloor)
		floorapi.PUT("/update/:id", api.UpdateFloor)
		floorapi.DELETE("/delete/:id", api.DeleteFloor)
	}
}

// Handler to get all Floors
// @router /floor/all [get]
// @summary Get all floors
// @tags floors
// @produce json
// @param page query int false "Page number (default:1)"
// @param limit query int false "Number of results per page (default:10)"
// @success 200 {array} model.Floor
func (api APIRoutes) GetAllFloors(c *gin.Context) {
	api.Server.GetFloors(c)
}

// Handler to get all floor based on filter
// @router /floor/filter [get]
// @summary Get all floors based on filters
// @tags floors
// @produce json
// @param id query string false "id"
// @param house_id query string false "house_id"
// @param name query string false "name"
// @param description query string false "description"
// @param area query string false "area"
// @param occupancy query string false "occupancy"
// @param heating_enabled query string false "heating_enabled"
// @param heating_type query string false "heating_type"
// @param cooling_enabled query string false "cooling_enabled"
// @param cooling_type query string false "cooling_type"
// @param ventilation_enabled query string false "ventilation_enabled"
// @param ventilation_type query string false "ventilation_type"
// @param temperature query string false "temperature"
// @param humidity query string false "humidity"
// @param light_level query string false "light_level"
// @param co2_level query string false "co2_false"
// @param created_at query string false "created_at"
// @param created_by query string false "created_by"
// @param updated_at query string false "updated_at"
// @param updated_by query string false "updated_by"
// @param start_date query string false "start_date"
// @param end_date query string false "end_date"
// @param page query int false "Page number (default:1)"
// @param limit query int false "Number of results per page (default:10)"
// @success 200 {array} model.Floor
// @Security ApiKeyAuth
func (api APIRoutes) GetFloorsByFilter(c *gin.Context) {
	api.Server.GetFloorsByFilter(c)
}

// Handler to get a floor by ID
// @router /floor/{id} [get]
// @summary Get a floor by ID
// @tags floors
// @produce json
// @param id path string true "Floor ID"
// @success 200 {object} model.Floor
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) GetFloor(c *gin.Context) {
	api.Server.GetFloor(c)
}

// Handler to create a floor
// @router /floor/create [post]
// @summary Create a floor
// @tags floors
// @accept json
// @produce json
// @param floor body model.Floor true "Floor object"
// @success 201 {object} model.Floor
// @failure 400 {object} model.ErrorResponse

func (api APIRoutes) CreateFloor(c *gin.Context) {
	api.Server.CreateFloor(c)
}

// Handler to update a floor
// @router /floor/update/{id} [put]
// @summary Update a floor
// @tags floors
// @accept json
// @produce json
// @param id path string true "Floor ID"
// @param floor body model.Floor true "Floor object"
// @success 200 {object} model.Floor
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) UpdateFloor(c *gin.Context) {
	api.Server.UpdateFloor(c)
}

// Handler to delete a floor
// @router  /floor/delete/{id} [delete]
// @summary Delete a floor
// @tags floors
// @param id path string true "Floor ID"
// @success 200 {object} model.SuccessResponse
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) DeleteFloor(c *gin.Context) {
	api.Server.DeleteFloor(c)
}
