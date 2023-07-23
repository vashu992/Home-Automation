package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-"
	_"github.com/vashu992/Home-Automation/docs"
	"github.com/vashu992/Home-Automation/server"
	"github.com/vashu992/Home-Automation/store/pgress"
)

type APIRoutes struct {
	Server server.ServerOperation
}

func (api APIRoutes) StartApp(router *gin.Engine, server.Server) {

	// SetUp  Server
	fmt.Println("Craeting new server .....")
	api.Server = &server
	api.Server.NewServer(pgress.PgressStore{})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Hadler))

	//user routs
	api.ActuatorRouts(router)
	api.FloorRouts(router)
	api.HouseRouts(router)
	api.OrganizationRouts(router)
	api.PackageRouts(router)
	api.PointRateRouts(router)
	api.RoomRouts(router)
	api.SensorRouts(router)
	api.SensorReadingRouts(router)
	api.UserRouts(router)

}