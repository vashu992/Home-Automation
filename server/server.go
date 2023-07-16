package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/store/pgress"
	"github.com/vashu992/Home-Automation/util"
	// "github.com/sirupsen/logrus"
)

type Server struct {
	Pgress pgress.StoreOperation
}

func (s *Server) NewServer(store pgress.PgressStore) {

	s.Pgress = &store
	util.SetLoger()
	util.Logger.Info("Logger Setup Done at server level")
	fmt.Println("Creating new Store .....")

	s.Pgress.NewStore()
}

type ServerOperation interface {
	//Postgress config methods
	NewServer(store pgress.PgressStore)

	// Actuator operation methods
	GetActuators(c *gin.Context) (*[]model.Actuator, error)
	GetActuatorsByFilter(c *gin.Context) (*[]model.Actuator, error)
	GetActuator(c *gin.Context) (*model.Actuator, error)
	CreateActuator(c *gin.Context) error
	UpdateActuator(c *gin.Context) error
	DeleteActuator(c *gin.Context) error

	//Floor operation methods
	GetFloors(c *gin.Context) (*[]model.Floor, error)
	GetFloorsByFilter(c *gin.Context) (*[]model.Floor, error)
	GetFloor(c *gin.Context) (*model.Floor, error)
	CreateFloor(c *gin.Context) error
	UpdateFloor(c *gin.Context) error
	DeleteFloor(c *gin.Context) error

	//House operation methods
	GetHouses(c *gin.Context) (*[]model.House, error)
	GetHousesByFilter(c *gin.Context) (*[]model.House, error)
	GetHouse(c *gin.Context) (*model.House, error)
	CreateHouse(c *gin.Context) error
	UpdateHouse(c *gin.Context) error
	DeleteHouse(c *gin.Context) error

	//Organization operation methods
	GetOrganizations(c *gin.Context) (*[]model.Organization, error)
	GetOrganizationsByFilter(c *gin.Context) (*[]model.Organization, error)
	GetOrganization(c *gin.Context) (*model.Organization, error)
	CreateOrganization(c *gin.Context) error
	UpdateOrganization(c *gin.Context) error
	DeleteOrganization(c *gin.Context) error

	//Package operation methods
	GetPackages(c *gin.Context) (*[]model.Package, error)
	GetPackagesByFilter(c *gin.Context) (*[]model.Package, error)
	GetPackage(c *gin.Context) (*model.Package, error)
	CreatePackage(c *gin.Context) error
	UpdatePackage(c *gin.Context) error
	DeletePackage(c *gin.Context) error

	//PointRate operation method
	GetPointRates(c *gin.Context) (*[]model.PointRate, error)
	GetPointRatesByFilter(c *gin.Context) (*[]model.PointRate, error)
	GetPointRate(c *gin.Context) (*model.PointRate, error)
	CreatePointRate(c *gin.Context) error
	UpdatePointRate(c *gin.Context) error
	DeletePointRate(c *gin.Context) error

	//Room operation methods
	GetRooms(c *gin.Context) (*[]model.Room, error)
	GetRoomsByFilter(c *gin.Context) (*[]model.Room, error)
	GetRoom(c *gin.Context) (*model.Room, error)
	CreateRoom(c *gin.Context) error
	UpdateRoom(c *gin.Context) error
	DeleteRoom(c *gin.Context) error

	//Sensor operation methods
	GetSensors(c *gin.Context) (*[]model.Sensor, error)
	GetSensorsByFilter(c *gin.Context) (*[]model.Sensor, error)
	GetSensor(c *gin.Context) (*model.Sensor, error)
	CreateSensor(c *gin.Context) error
	UpdateSensor(c *gin.Context) error
	DeleteSensor(c *gin.Context) error

	//SensorReading operation methods
	GetSensorReadings(c *gin.Context) (*[]model.SensorReading, error)
	GetSensorReadingsByFilter(c *gin.Context) (*[]model.SensorReading, error)
	GetSensorReading(c *gin.Context) (*model.SensorReading, error)
	CreateSensorReading(c *gin.Context) error
	UpdateSensorReading(c *gin.Context) error
	DeleteSensorReading(c *gin.Context) error

	//User operation methods
	GetUsers(c *gin.Context) (*[]model.User, error)
	GetUser(c *gin.Context) (*model.User, error)
	GetUsersByFilter(c *gin.Context) (*[]model.User, error)
	CreateUser(c *gin.Context) error
	UpdateUser(c *gin.Context) error
	DeleteUser(c *gin.Context) error
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)

	//middleware
	AuthMiddleware() gin.HandlerFunc
}