package pgress

import (
	"github.com/vashu992/Home-Automation/model"

)

type StoreOperation interface {
	// Postgress  config methods
	NewStore()


	// Actuator operation methods 
	GetActuators() (*[]model.Actuator, error)
	GetActuatorByFilter(filter map[string]string) (*[]model.Actuator, error)
	GetActuator(actuatorID string) (*model.Actuator, error)
	CreateActuator(actuator *model.Actuator) error
	UpdateActuator(actuator *model.Actuator) error
	DeleteActuator(actuatorID string) error


	// Floor Operation methods
	GetFloors() (*[]model.Floor,error)
	GetFloorsByFilter(filter map[string]string) (*[]model.Floor, error)
	GetFloor(floorID string)( *model.Floor,error)
	CreateFloor(floor *model.Floor) error
	UpdateFloor(floor *model.Floor) error
	DeleteFloor(floorID string) error


	// House Operation methods 
	GetHouses() (*[]model.House, error)
	GetHousesByFilter(filter map[string]string) (*[]model.House, error)
	GetHouse(houseID string) (*model.House,error)
	CreateHouse(house *model.House) error
	UpdateHouse(house *model.House) error
	DeleteHouse(houseID string) error


	// Organization operation methods 
	GetOraganizations() (*[]model.Organization, error)
	GetOraganizationsByFilter(filter map[string]string) (*[]model.Organization, error)
	GetOraganization(organizationID string) (*model.Organization, error)
	CreateOrganization(organization *model.Organization) error
	UpdateOrganization(organization *model.Organization) error
	DeleteOrganization(organizationID string) error


	// Package operation methods
	GetPackages() (*[]model.Package, error)
	GetPackagesByFilter(filter map[string]string) (*[]model.Package, error)
	GetPackage(pkgID string) (*model.Package, error)
	CreatePackage(pkg *model.Package) error
	UpdatePackage(pkg *model.Package) error
	DeletePackage(pkgID string) error


	// PointRate operation method
	GetPointRates() (*[]model.PointRate, error)
	GetPointRatesByFilter(filter map[string]string) (*[]model.PointRate, error)
	GetPointRate(pointRateID string) (*model.PointRate, error)
	CreatePointRate(pointRate *model.PointRate) error
	UpdatePointRate(pointRate *model.PointRate) error
	DeletePointRate(pointRateID string) error


	// Room Operation methods
	GetRooms() (*[]model.Room, error)
	GetRoomsByFilter(filter map[string]string) (*[]model.Room, error)
	GetRoom(roomID string) (*model.Room, error)
	CreateRoom(room *model.Room) error
	UpdateRoom(room *model.Room) error
	DeleteRoom(roomID string) error

	//Sensor operation methods
	GetSensors() (*[]model.Sensor, error)
	GetSensorsByFilter(filter map[string]string) (*[]model.Sensor, error)
	GetSensor(sensorID string) (*model.Sensor, error)
	CreateSensor(sensor *model.Sensor) error
	UpdateSensor(sensor *model.Sensor) error
	DeleteSensor(sensorID string) error

	//SensorReading operation methods
	GetSensorReadings() (*[]model.SensorReading, error)
	GetSensorReadingsByFilter(filter map[string]string) (*[]model.SensorReading, error)
	GetSensorReading(sensorID string) (*model.SensorReading, error)
	CreateSensorReading(sensor *model.SensorReading) error
	UpdateSensorReading(sensor *model.SensorReading) error
	DeleteSensorReading(sensorID string) error

	//User operation methods
	GetUsers() (*[]model.User, error)
	GetUsersByFilter(filter map[string]string) (*[]model.User, error)
	GetUser(userId string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(userId string) error
	SingIn(userSignIn model.UserSignIn) (*model.User, error)
	SignUp(user *model.User) error


}