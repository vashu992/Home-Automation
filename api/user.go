package api

import (
	"github.com/gin-gonic/gin"
)

func (api APIRoutes) UserRouts(router *gin.Engine) {
	// Define routes
	userapi := router.Group("/user")
	{
		userapi.GET("/all", api.AuthMiddlewareComplete(), api.GetUsers)
		userapi.GET("/filter", api.AuthMiddlewareComplete(), api.GetUsersByFilter)
		userapi.GET("/:id", api.AuthMiddlewareComplete(), api.GetUser)
		userapi.POST("/sighup", api.SighUp)
		userapi.POST("/sighin", api.SighIn)
		userapi.POST("/create", api.AuthMiddlewareComplete(), api.CreateUser)
		userapi.PUT("/update/:id", api.AuthMiddlewareComplete(), api.UpdateUser)
		userapi.DELETE("/delete/:id", api.AuthMiddlewareComplete(), api.DeleteUser)

	}

}

// Handler to get all users
// @router /user/all [get]
// @summary Get all users
// @tags user
// @produce json
// @param page query int false "page number (Default: 1)"
// @param limit query int false "page of results per page (Default:10)"
// @success 200 {array} model.User
// @Security ApiKeyAuth
func (api APIRoutes) GetUsers(c *gin.Context) {
	api.Server.GetUsers(c)
}

// Handler to get all users based on filter
// @router /user/filter [get]
// @summary Get all users based on given filters
// @tags users
// @produce json
// @param org_id query string false "org_id"
// @param name query string false "name"
// @param publisher query string false "publisher"
// @param email query string false "email"
// @param password query string false "password"
// @param primary_mobile_number query string false "primary_mobile_number"
// @param secondary_mobile_number query string false "secondary_mobile_number"
// @param landline_number query string false "landline_number"
// @param active_status query string false "active_status"
// @param type query string false "type"
// @param address_type query string false "address_type"
// @param house_no query string false "house_no"
// @param house_name query string false "house_name"
// @param lane_number query string false "lane_number"
// @param lane_name query string false "lane_name"
// @param landmark query string false "landmark"
// @param district query string false "district"
// @param post query string false "post"
// @param city query string false "city"
// @param village query string false "village"
// @param state query string false "state"
// @param nation query string false "nation"
// @param created_by query string false "created_by"
// @param updated_by query string false "updated_by"
// @param start_date query string false "start date"
// @param end_date query string false "end date"
// @param page query int false "Page number (default: 1)"
// @param limit query int false "Number of results per page (default: 10)"
// @success 200 {array} model.User
// @Security ApiKeyAuth
func (api APIRoutes) GetUsersByFilter(c *gin.Context) {
	api.Server.GetUsersByFilter(c)
}

// Handler to SighUp a user
// @router /user/sighup [post]
// @summary SighUp a user
// @tags users
// @accepts json
// @produce json
// @param user body model.User true "User object"
// @Success 200 {string } string "Successful SighUp"
// @failure 400 {object} model.ErrorResponse
func (api APIRoutes) SighUp(c *gin.Context) {
	api.Server.SignUp(c)
}

// Handler to SighIn a user by email and password
// @router /user/sighin [post]
// @summary SighIn user
// @tags users
// @produce json
// @param user body model.UserSignIn true "User object"
// @Success 200 {string} string "Successful SighIn "
// @failure 404 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (api APIRoutes) SighIn(c *gin.Context) {
	api.Server.SignIn(c)
}

// Handler to get a user by ID
// @router /user/{id} [get]
// @summary Get a user by ID
// @tags users
// @produce json
// @param id path string true "User ID "
// @success 200 {object} model.User
// @failure 404 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (api APIRoutes) GetUser(c *gin.Context) {
	api.Server.GetUser(c)
}

// Handler to create a user
// @router /user/create [post]
// @summary create a user
// @tags json
// @produce json
// @param user body model.User true "User object"
// @success 201 {object} model.User
// @failure 400 {object} model.ErrorResponse
// @Security ApiKeyAuth

func (api APIRoutes) CreateUser(c *gin.Context) {
	api.Server.CreateUser(c)
}

// Hadler to update a user 
// @router /user/update/{id} [put]
// @summary Update a user
// @tags users
// @accepts json
// @produce json
// @param id path string true "User ID "
// @param user body model.User true "User object"
// @success 200 {object} model.User
// @failure 400 {object} model.ErrorResponse
// @Security ApiKeyPath
func (api APIRoutes) UpdateUser(c *gin.Context) {
	api.Server.UpdateUser(c)
}

// Handler to delete a user
// @router /user/delete/{id} [delete]
// @summary Delete a user
// @tags users
// @param id path string true "User ID"
// @success 200 {object} model.SuccessResponse
// @failure 404 {object} model.ErrorResponse
// @Security ApiKeyAuth
func (api APIRoutes) DeleteUser(c *gin.Context) {
	api.Server.DeleteUser(c)
}