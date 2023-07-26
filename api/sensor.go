package api

import (
	"github.com/gin-gonic/gin"
)

func (api APIRoutes) SensorRouts(router *gin.Engine) {
	// Define routes
	sensorapi := router.Group("/sensor")
	{
		sensorapi.GET("/all", api.AuthMiddlewareComplete(), api.GetSensors)
		sensorapi.GET("/filter", api.AuthMiddlewareComplete(), api.GetSensorsByFilter)
		sensorapi.GET("/:id", api.AuthMiddlewareComplete(), api.GetSensor)
		sensorapi.POST("/create", api.AuthMiddlewareComplete(), api.CreateSensor)
		sensorapi.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateSensor)
		sensorapi.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteSensor)
	}

}

// Handler to get all sensors
// @router /sensor/all [get]
// @summary Get all sensors
// @tags sensors
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Sensor
func (api APIRoutes) GetSensors(c *gin.Context) {
	api.Server.GetSensors(c)
}

// Handler to get all sensors based on filter
// @router /sensor/filter [get]
// @summary Get all sensors based on given filters
// @tags sensors
// @produce json
// @param id query string false "id"
// @param rooms_id query string false "rooms_id"
// @param house_id query string false "house_id"
// @param floor_id query string false "floor_id"
// @param name query string false "name"
// @param start_time query string false "start_time"
// @param end_time query string false "end_time"
// @param status query string false "status"
// @param updated_at query string false "updated_at"
// @param updated_by query string false "updated_by"
// @param created_at query string false "created_at"
// @param created_by query string false "created_by"
// @param reading query string false "reading"
// @param unit query string false "unit"
// @param refresh_rate query string false "refresh_rate"
// @param min_reading query string false "min_reading"
// @param max_reading query string false "max_reading"
// @param start_date query string false "start_date"
// @param end_date query string false "end_date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Sensor
// @Security ApiKeyAuth
func (api APIRoutes) GetSensorsByFilter(c *gin.Context) {
	api.Server.GetSensorsByFilter(c)
}

// Handler to get a sensor by ID
// @router /sensor/{id} [get]
// @summary Get a sensor by ID
// @tags sensors
// @produce json
// @param id path string true "Sensor ID"
// @success 200 {object} model.Sensor
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) GetSensor(c *gin.Context) {
	api.Server.GetSensor(c)
}

// Handler to create a sensor
// @router /sensor/create [post]
// @summary Create a sensor
// @tags sensors
// @accept json
// @produce json
// @param sensor body model.Sensor true "Sensor object"
// @success 201 {object} model.Sensor
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) CreateSensor(c *gin.Context) {
	api.Server.CreateSensor(c)
}

// Handler to update a sensor
// @router /sensor/update/{id} [put]
// @summary Update a sensor
// @tags sensors
// @accept json
// @produce json
// @param id path string true "Sensor ID"
// @param sensor body model.Sensor true "Sensor object"
// @success 200 {object} model.Sensor
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) UpdateSensor(c *gin.Context) {
	api.Server.UpdateSensor(c)
}

// Handler to delete a sensor
// @router  /sensor/delete/{id} [delete]
// @summary Delete a sensor
// @tags sensors
// @param id path string true "Sensor ID"
// @success 200 {object} model.SuccessResponse
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) DeleteSensor(c *gin.Context) {
	api.Server.DeleteSensor(c)
}