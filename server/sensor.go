package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
)

func (server *Server) GetSensors(c *gin.Context) (*[]model.Sensor, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo , model.ServerPackageLavel , model.GetSensors , "reading all sensor data", nil )
	sensors, err := server.Pgress.GetSensors()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetSensors, 
			"error while reading sensors data from pgress", err)
		return sensors, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetSensors , 
		"returning all sensor data to api and setting response" , sensors )
	c.JSON(http.StatusOK, sensors)
	return sensors, nil

}

func (server *Server) GetSensorsByFilter(c *gin.Context) (*[]model.Sensor, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetSensors, "reading all sensor data", nil)
	condition := server.readQueryParams(c)
	sensors, err := server.Pgress.GetSensorsByFilter(condition)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetSensors,
			"error while reading sensors data from pgress", err)
		return sensors, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetSensors,
		"returning all sensor data to api and setting response", sensors)
	c.JSON(http.StatusOK, sensors)
	return sensors, nil

}

func (server *Server) GetSensor(c *gin.Context) (*model.Sensor, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetSensor, 
		"reading sensor data from pgress", nil)
	sensor, err := server.Pgress.GetSensor(c.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetSensor, 
			"error while reading sensor data from pgress", err )
		return sensor, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetSensor , 
		"returning sensor data to api and setting response" , sensor)
	c.JSON(http.StatusOK, sensor)
	return sensor, nil

}
func (server *Server) CreateSensor(c *gin.Context) error {

	var sensor model.Sensor
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateSensor, 
		"unmarshaling sensor data",nil)

	err := c.ShouldBindJSON(&sensor)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateSensor, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	sensor.CreatedAt = time.Now().UTC()
	sensor.ID = uuid.New()
	//validation is to be done here
	//DB call
	err = server.Pgress.CreateSensor(&sensor)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateSensor, 
			"error while creating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.CreateSensor , 
		"successfully created sensor record and setting response" , sensor)
	c.JSON(http.StatusCreated, sensor)
	return nil

}

func (server *Server) UpdateSensor(c *gin.Context) error {

	var sensor model.Sensor
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdateSensor, 
		"unmarshaling sensor data",nil)
	err := c.ShouldBindJSON(&sensor)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateSensor, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	sensor.UpdatedAt = time.Now().UTC()
	err = server.Pgress.UpdateSensor(&sensor)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateSensor, 
			"error while updating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.UpdateSensor , 
		"successfully updated sensor record and setting response" , sensor)
	c.JSON(http.StatusOK, sensor)
	return nil

}

func (server *Server) DeleteSensor(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteSensor, 
		"reading sensor id",nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteSensor, 
			"missing sensor id", nil )
		return fmt.Errorf("")
	}
	//DB call
	err := server.Pgress.DeleteSensor(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteSensor, 
			"error while deleting sensor record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.DeleteSensor , 
		"successfully deleted sensor record " , nil)
	return nil

}