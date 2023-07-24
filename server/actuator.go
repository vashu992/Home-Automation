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

func (server *Server) GetActuators(c *gin.Context) (*[]model.Actuator, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetActuators, "reading all actuator data", nil)
	actuators, err := server.Pgress.GetActuators()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetActuators,
			"error while reading actuators data from pgress", err)
		return actuators, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetActuators,
		"returning all actuator data to api and setting response", actuators)
	c.JSON(http.StatusOK, actuators)
	return actuators, nil

}

func (server *Server) GetActuatorsByFilter(c *gin.Context) (*[]model.Actuator, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetActuators, "reading all actuator data", nil)
	condition := server.readQueryParams(c)
	actuators, err := server.Pgress.GetActuatorByFilter(condition)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetActuators,
			"error while reading actuators data from pgress", err)
		return actuators, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetActuators,
		"returning all actuator data to api and setting response", actuators)
	c.JSON(http.StatusOK, actuators)
	return actuators, nil

}

func (server *Server) GetActuator(c *gin.Context) (*model.Actuator, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetActuator,
		"reading actuator data from pgress", nil)
	actuator, err := server.Pgress.GetActuator(c.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetActuator,
			"error while reading actuator data from pgress", err)
		return actuator, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetActuators,
		"returning actuator data to api and setting response", actuator)
	c.JSON(http.StatusOK, actuator)
	return actuator, nil

}
func (server *Server) CreateActuator(c *gin.Context) error {

	var actuator model.Actuator
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateActuator,
		"unmarshaling actuator data", nil)

	err := c.ShouldBindJSON(&actuator)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateActuator,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	actuator.CreatedAt = time.Now().UTC()
	actuator.Id = uuid.New()
	//validation is to be done here
	//DB call
	err = server.Pgress.CreateActuator(&actuator)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateActuator,
			"error while creating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetActuators,
		"successfully created actuator record and setting response", actuator)
	c.JSON(http.StatusCreated, actuator)
	return nil

}

func (server *Server) UpdateActuator(c *gin.Context) error {

	var actuator model.Actuator
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdateActuator,
		"unmarshaling actuator data", nil)
	err := c.ShouldBindJSON(&actuator)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateActuator,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	actuator.UpdatedAt = time.Now().UTC()
	err = server.Pgress.UpdateActuator(&actuator)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateActuator,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetActuators,
		"successfully updated actuator record and setting response", actuator)
	c.JSON(http.StatusOK, actuator)
	return nil

}

func (server *Server) DeleteActuator(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteActuator,
		"reading actuator id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteActuator,
			"missing actuator id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.Pgress.DeleteActuator(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteActuator,
			"error while deleting actuator record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteActuator,
		"successfully deleted actuator record ", nil)
	return nil

}
