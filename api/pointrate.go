package api

import (
	"github.com/gin-gonic/gin"
)

func (api APIRoutes) PointRateRouts(router *gin.Engine) {
	// Define routes
	pointRateapi := router.Group("/pointRate")
	{
		pointRateapi.GET("/all", api.GetAllPointRates)
		pointRateapi.GET("/filter", api.AuthMiddlewareComplete(), api.GetPointRatesByFilter)
		pointRateapi.GET("/:id", api.GetPointRate)
		pointRateapi.POST("/create", api.CreatePointRate)
		pointRateapi.PUT("/update/:id", api.UpdatePointRate)
		pointRateapi.DELETE("/delete/:id", api.DeletePointRate)
	}

}

// Handler to get all pointRates
// @router /pointRate/all [get]
// @summary Get all pointRates
// @tags pointRates
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.PointRate
func (api APIRoutes) GetAllPointRates(c *gin.Context) {
	api.Server.GetPointRates(c)
}

// Handler to get all pointRates based on filter
// @router /pointRate/filter [get]
// @summary Get all pointRates based on given filters
// @tags pointRates
// @produce json
// @param id query string false "id"
// @param name query string false "name"
// @param description query string false "description"
// @param house query integer false "house"
// @param floor query integer false "floor"
// @param room query integer false "room"
// @param user query integer false "user"
// @param sensor query integer false "sensor"
// @param actuator query integer false "actuator"
// @param created_by query string false "created_by"
// @param created_at query string false "created_at"
// @param updated_at query string false "updated_at"
// @param updated_by query string false "updated_by"
// @param start_date query string false "start_date"
// @param end_date query string false "end_date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.PointRate
// @Security ApiKeyAuth
func (api APIRoutes) GetPointRatesByFilter(c *gin.Context) {
	api.Server.GetPointRatesByFilter(c)
}

// Handler to get a pointRate by ID
// @router /pointRate/{id} [get]
// @summary Get a pointRate by ID
// @tags pointRates
// @produce json
// @param id path string true "PointRate ID"
// @success 200 {object} model.PointRate
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) GetPointRate(c *gin.Context) {
	api.Server.GetPointRate(c)
}

// Handler to create a pointRate
// @router /pointRate/create [post]
// @summary Create a pointRate
// @tags pointRates
// @accept json
// @produce json
// @param pointRate body model.PointRate true "PointRate object"
// @success 201 {object} model.PointRate
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) CreatePointRate(c *gin.Context) {
	api.Server.CreatePointRate(c)
}

// Handler to update a pointRate
// @router /pointRate/update/{id} [put]
// @summary Update a pointRate
// @tags pointRates
// @accept json
// @produce json
// @param id path string true "PointRate ID"
// @param pointRate body model.PointRate true "PointRate object"
// @success 200 {object} model.PointRate
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) UpdatePointRate(c *gin.Context) {
	api.Server.UpdatePointRate(c)
}

// Handler to delete a pointRate
// @router  /pointRate/delete/{id} [delete]
// @summary Delete a pointRate
// @tags pointRates
// @param id path string true "PointRate ID"
// @success 200 {object} model.SuccessResponse
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) DeletePointRate(c *gin.Context) {
	api.Server.DeletePointRate(c)
}