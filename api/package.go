package api

import (
	"github.com/gin-gonic/gin"
)

func (api APIRoutes) PackageRouts(router *gin.Engine) {
	// Define routes
	pkgapi := router.Group("/pkg")
	{
		pkgapi.GET("/all", api.GetAllPackages)
		pkgapi.GET("/filter", api.AuthMiddlewareComplete(), api.GetPackagesByFilter)
		pkgapi.GET("/:id", api.GetPackage)
		pkgapi.POST("/create", api.CreatePackage)
		pkgapi.PUT("/update/:id", api.UpdatePackage)
		pkgapi.DELETE("/delete/:id", api.DeletePackage)
	}

}

// Handler to get all pkgs
// @router /pkg/all [get]
// @summary Get all pkgs
// @tags pkgs
// @produce json
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Package
func (api APIRoutes) GetAllPackages(c *gin.Context) {
	api.Server.GetPackages(c)
}

// Handler to get all pkgs based on filter
// @router /pkg/filter [get]
// @summary Get all pkgs based on given filters
// @tags pkgs
// @produce json
// @param id query string false "id"
// @param point_rate_id query string false "point_rate_id"
// @param name query string false "name"
// @param description query string false "description"
// @param price query integer false "price"
// @param Duration query string false "Duration"
// @param limits query string false "limits"
// @param number_of_rooms query integer false "number_of_rooms"
// @param number_of_house query integer false "number_of_house"
// @param number_of_floor query integer false "number_of_floor"
// @param number_of_user query integer false "number_of_user"
// @param number_of_actuator query integer false "number_of_actuator"
// @param number_of_sensor query integer false "number_of_sensor"
// @param points query integer false "points"
// @param benifits query string false "benifits"
// @param avilibility query string false "avilibility"
// @param fratures query string false "fratures"
// @param status query boolean false "status"
// @param created_at query string false "created_at"
// @param created_by query string false "created_by"
// @param updated_at query string false "updated_at"
// @param updated_by query string false "updated_by"
// @param start_date query string false "start_date"
// @param end_date query string false "end_date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.Package
// @Security ApiKeyAuth
func (api APIRoutes) GetPackagesByFilter(c *gin.Context) {
	api.Server.GetPackagesByFilter(c)
}

// Handler to get a pkg by ID
// @router /pkg/{id} [get]
// @summary Get a pkg by ID
// @tags pkgs
// @produce json
// @param id path string true "Package ID"
// @success 200 {object} model.Package
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) GetPackage(c *gin.Context) {
	api.Server.GetPackage(c)
}

// Handler to create a pkg
// @router /pkg/create [post]
// @summary Create a pkg
// @tags pkgs
// @accept json
// @produce json
// @param pkg body model.Package true "Package object"
// @success 201 {object} model.Package
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) CreatePackage(c *gin.Context) {
	api.Server.CreatePackage(c)
}

// Handler to update a pkg
// @router /pkg/update/{id} [put]
// @summary Update a pkg
// @tags pkgs
// @accept json
// @produce json
// @param id path string true "Package ID"
// @param pkg body model.Package true "Package object"
// @success 200 {object} model.Package
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) UpdatePackage(c *gin.Context) {
	api.Server.UpdatePackage(c)
}

// Handler to delete a pkg
// @router  /pkg/delete/{id} [delete]
// @summary Delete a pkg
// @tags pkgs
// @param id path string true "Package ID"
// @success 200 {object} model.SuccessResponse
// @failure 404 {object} model.ErrorResponse
func (api APIRoutes) DeletePackage(c *gin.Context) {
	api.Server.DeletePackage(c)
}