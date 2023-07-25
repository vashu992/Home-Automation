package api

import "github.com/gin-gonic/gin"

func (api APIRoutes) HouseRouts(router *gin.Engine) {
	// Define routes
	houseapi := router.Group("/house")
	{
		houseapi.GET("/all ", api.GetAllHouses)
		houseapi.GET("/filter", api.AuthMiddlewareComplete(), api.GetHousesByFilter)
		houseapi.GET("/:id", api.GetHouse)
		houseapi.POST("/create", api.CreateHouse)
		houseapi.PUT("/update/:id", api.UpdateHouse)
		houseapi.DELETE("/delete/:id", api.DeleteHouse)
	}
}

// Handler to get all houses
// @router /house/all [get]
// @summary Get all houses
// @tags houses
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default:10)"
// @success 200 {array} model.House
func (api APIRoutes) GetAllHouses(c *gin.Context) {
	api.Server.GetHouses(c)
}

// Handler to get all houses based on filter
// @router /house/filter [get]
// @summary Get all houses based on given filters
// @tags houses
// @produce json
// @param id query string false "id"
// @param name query string false "name"
// @param description query string false "description"
// @param area query string false "area"
// @param occupancy query string false "occupancy"
// @param number_of_floors query string false "number_of_floors"
// @param created_at query string false "created_at"
// @param created_by query string false "created_by"
// @param updated_at query string false "updated_at"
// @param updated_by query string false "updated_by"
// @param latitude query string false "latitude"
// @param longitude query string false "longitude"
// @param start_date query string false "start_date"
// @param end_date query string false "end_date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.House
// @Security ApiKeyAuth

func (api APIRoutes) GetHousesByFilter(c *gin.Context) {
	api.Server.GetHousesByFilter(c)
}

// Handler to get a house by ID
// @router /house/{id} [get]
// @summary Get a house by ID
// @tags houses
// @produce json
// @param id path string true "House ID"
// @success 200 {object} model.House
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) GetHouse(c *gin.Context) {
	api.Server.GetHouse(c)
}
// Handler to create a house
// @router /house/create [post]
// @summary Create a house
// @tags houses
// @accept json
// @produce json
// @param house body model.House true "House object"
// @success 201 {object} model.House
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) CreateHouse(c *gin.Context) {
	api.Server.CreateHouse(c)
}

// Handler to update a house
// @router /house/update/{id} [put]
// @summary Update a house
// @tags houses
// @accept json
// @produce json
// @param id path string true "House ID"
// @param house body model.House true "House object"
// @success 200 {object} model.House
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) UpdateHouse(c *gin.Context) {
	api.Server.UpdateHouse(c)
}

// Handler to delete a house
// @router  /house/delete/{id} [delete]
// @summary Delete a house
// @tags houses
// @param id path string true "House ID"
// @success 200 {object} model.SuccessResponse
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) DeleteHouse(c *gin.Context) {
	api.Server.DeleteHouse(c)
}