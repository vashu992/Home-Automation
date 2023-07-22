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

func (server *Server) GetSensorReadings(c *gin.Context) (*[]model.SensorReading, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo , model.ServerPackageLavel , model.GetSensorReadings , "reading all sensorReading data", nil )
	sensors, err := server.Pgress.GetSensorReadings()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetSensorReadings, 
			"error while reading sensors data from pgress", err)
		return sensors, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetSensorReadings , 
		"returning all sensorReading data to api and setting response" , sensors )
	c.JSON(http.StatusOK, sensors)
	return sensors, nil

}

func (server *Server) GetSensorReadingsByFilter(c *gin.Context) (*[]model.SensorReading, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetSensorReadings, "reading all sensorReading data", nil)
	condition := server.readQueryParams(c)
	sensors, err := server.Pgress.GetSensorReadingsByFilter(condition)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetSensorReadings,
			"error while reading sensors data from pgress", err)
		return sensors, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetSensorReadings,
		"returning all sensorReading data to api and setting response", sensors)
	c.JSON(http.StatusOK, sensors)
	return sensors, nil

}

func (server *Server) GetSensorReading(c *gin.Context) (*model.SensorReading, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetSensorReading, 
		"reading sensorReading data from pgress", nil)
	sensorReading, err := server.Pgress.GetSensorReading(c.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetSensorReading, 
			"error while reading sensorReading data from pgress", err )
		return sensorReading, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetSensorReading , 
		"returning sensorReading data to api and setting response" , sensorReading)
	c.JSON(http.StatusOK, sensorReading)
	return sensorReading, nil

}
func (server *Server) CreateSensorReading(c *gin.Context) error {

	var sensorReading model.SensorReading
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateSensorReading, 
		"unmarshaling sensorReading data",nil)

	err := c.ShouldBindJSON(&sensorReading)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateSensorReading, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	sensorReading.CreatedAt = time.Now().UTC()
	sensorReading.ID = uuid.New()
	//validation is to be done here
	//DB call
	err = server.Pgress.CreateSensorReading(&sensorReading)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateSensorReading, 
			"error while creating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.CreateSensorReading , 
		"successfully created sensorReading record and setting response" , sensorReading)
	c.JSON(http.StatusCreated, sensorReading)
	return nil

}

func (server *Server) UpdateSensorReading(c *gin.Context) error {

	var sensorReading model.SensorReading
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdateSensorReading, 
		"unmarshaling sensorReading data",nil)
	err := c.ShouldBindJSON(&sensorReading)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateSensorReading, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	sensorReading.UpdatedAt = time.Now().UTC()
	err = server.Pgress.UpdateSensorReading(&sensorReading)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateSensorReading, 
			"error while updating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.UpdateSensorReading , 
		"successfully updated sensorReading record and setting response" , sensorReading)
	c.JSON(http.StatusOK, sensorReading)
	return nil

}

func (server *Server) DeleteSensorReading(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteSensorReading, 
		"reading sensorReading id",nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteSensorReading, 
			"missing sensorReading id", nil )
		return fmt.Errorf("")
	}
	//DB call
	err := server.Pgress.DeleteSensorReading(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteSensorReading, 
			"error while deleting sensorReading record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.DeleteSensorReading , 
		"successfully deleted sensorReading record " , nil)
	return nil

}