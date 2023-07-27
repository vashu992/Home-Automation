package api

import (
	"github.com/gin-gonic/gin"
)

func (api APIRoutes) SensorReadingRouts(router *gin.Engine) {
	// Define routes
	sensorReadingapi := router.Group("/sensorreading")
	{
		sensorReadingapi.GET("/all", api.AuthMiddlewareComplete(), api.GetSensorReadings)
		sensorReadingapi.GET("/filter", api.AuthMiddlewareComplete(), api.GetSensorReadingsByFilter)
		sensorReadingapi.GET("/:id", api.AuthMiddlewareComplete(), api.GetSensorReading)
		sensorReadingapi.POST("/create", api.AuthMiddlewareComplete(), api.CreateSensorReading)
		sensorReadingapi.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateSensorReading)
		sensorReadingapi.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteSensorReading)
	}

}

// Handler to get all sensorReadings
// @router /sensorreading/all [get]
// @summary Get all sensorReadings
// @tags sensorReadings
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.SensorReading
func (api APIRoutes) GetSensorReadings(c *gin.Context) {
	api.Server.GetSensorReadings(c)
}

// Handler to get all sensorReadings based on filter
// @router /sensorreading/filter [get]
// @summary Get all sensorReadings based on given filters
// @tags sensorReadings
// @produce json
// @param id query string false "id"
// @param sensor_id query string false "sensor_id"
// @param reading query string false "reading"
// @param created_at query string false "created_at"
// @param updated_at query string false "updated_at"
// @param start_date query string false "start_date"
// @param end_date query string false "end_date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.SensorReading
// @Security ApiKeyAuth
func (api APIRoutes) GetSensorReadingsByFilter(c *gin.Context) {
	api.Server.GetSensorReadingsByFilter(c)
}

// Handler to get a sensor by ID
// @router /sensorreading/{id} [get]
// @summary Get a sensor by ID
// @tags sensorReadings
// @produce json
// @param id path string true "SensorReading ID"
// @success 200 {object} model.SensorReading
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) GetSensorReading(c *gin.Context) {
	api.Server.GetSensorReading(c)
}

// Handler to create a sensor
// @router /sensorreading/create [post]
// @summary Create a sensor
// @tags sensorReadings
// @accept json
// @produce json
// @param sensor body model.SensorReading true "SensorReading object"
// @success 201 {object} model.SensorReading
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) CreateSensorReading(c *gin.Context) {
	api.Server.CreateSensorReading(c)
}

// Handler to update a sensor
// @router /sensorreading/update/{id} [put]
// @summary Update a sensor
// @tags sensorReadings
// @accept json
// @produce json
// @param id path string true "SensorReading ID"
// @param sensor body model.SensorReading true "SensorReading object"
// @success 200 {object} model.SensorReading
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) UpdateSensorReading(c *gin.Context) {
	api.Server.UpdateSensorReading(c)
}

// Handler to delete a sensor
// @router  /sensorreading/delete/{id} [delete]
// @summary Delete a sensor
// @tags sensorReadings
// @param id path string true "SensorReading ID"
// @success 200 {object} model.SuccessResponse
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) DeleteSensorReading(c *gin.Context) {
	api.Server.DeleteSensorReading(c)
}