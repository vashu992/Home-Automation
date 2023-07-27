package api

import (
	"github.com/gin-gonic/gin"
)

func (api APIRoutes) ActuatorRouts(router *gin.Engine) {
	// Define routes
	actuatorapi := router.Group("/actuator")
	{
		actuatorapi.GET("/all", api.GetAllActuators)
		actuatorapi.GET("/filter", api.AuthMiddlewareComplete(), api.GetUsersByFilter)
		actuatorapi.GET("/:id", api.GetActuator)
		actuatorapi.POST("/create", api.CreateActuator)
		actuatorapi.PUT("/update/:id", api.UpdateActuator)
		actuatorapi.DELETE("/delete/:id", api.DeleteActuator)
	}
}

// Handler to get all actuators
// @router /actuator/all [get]
// @summary Get all actuators
// @tags actuators
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (defaults: 10)"
// @success 200 {array} model.Actuator
func (api APIRoutes) GetAllActuators(c *gin.Context) {
	api.Server.GetActuators(c)
}

// Handler to get all actuators based on filter
// @router /actuator/filter [get]
// @summary Get all actuators based on given filters
// @tags actuators
// @produce json
// @param id query string false "id"
// @param rooms_id query string false "rooms_id"
// @param house_id query string false "house_id"
// @param floor_id query string false "floor_id"
// @param name query string false "name"
// @param description query string false "description"
// @param area query string false "area"
// @param watt_consumption query string false "watt_consumption"
// @param occupancy query string false "occupancy"
// @param heating_enabled query boolean false "heating_enabled"
// @param heating_type query string false "heating_type"
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
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Actuator
// @Security ApiKeyAuth
func (api APIRoutes) GetActuatorsByFilter(c *gin.Context) {
	api.Server.GetActuatorsByFilter(c)
}

// Handler to get a actuator by ID
// @router /actuator/{id} [get]
// @summary Get a actuator by ID
// @tags actuators 
// @produce json 
// @param id path string true "Actuator ID"
// @success 200 {object} model.Actuator
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) GetActuator(c *gin.Context) {
	api.Server.GetActuator(c)
}

// Handler to create a actuator
// @router /actuator/create [post]
// @summary Create a actuator
// @tags actuators
// @accept json
// @produce json
// @param actuator body model.Actuator true "Actuator object"
// @success 201 {object} model.Actuator
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) CreateActuator(c *gin.Context) {
	api.Server.CreateActuator(c)
}

// Handler to update a actuator
// @router /actuator/update/{id} [put]
// @summary Update a actuator
// @tags actuators
// @accept json
// @produce json
// @param id path string true "Actuator ID"
// @param actuator body model.Actuator true "Actuator object"
// @success 200 {object} model.Actuator
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) UpdateActuator(c *gin.Context) {
	api.Server.UpdateActuator(c)
}

// Handler to delete a actuator
// @router  /actuator/delete/{id} [delete]
// @summary Delete a actuator
// @tags actuators
// @param id path string true "Actuator ID"
// @success 200 {object} model.SuccessResponse
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) DeleteActuator(c *gin.Context) {
	api.Server.DeleteActuator(c)
}