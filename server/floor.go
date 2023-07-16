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

func (server *Server) GetFloors(c *gin.Context) (*[]model.Floor, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo , model.ServerPackageLavel , model.GetFloors , "reading all floor data", nil )
	floors, err := server.Pgress.GetFloors()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetFloors, 
			"error while reading floors data from pgress", err)
		return floors, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetFloors , 
		"returning all floor data to api and setting response" , floors )
	c.JSON(http.StatusOK, floors)
	return floors, nil

}

func (server *Server) GetFloorsByFilter(c *gin.Context) (*[]model.Floor, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetFloors, "reading all floor data", nil)
	condition := server.readQueryParams(c)
	floors, err := server.Pgress.GetFloorsByFilter(condition)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetFloors,
			"error while reading floors data from pgress", err)
		return floors, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetFloors,
		"returning all floor data to api and setting response", floors)
	c.JSON(http.StatusOK, floors)
	return floors, nil

}

func (server *Server) GetFloor(c *gin.Context) (*model.Floor, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetFloor, 
		"reading floor data from pgress", nil)
	floor, err := server.Pgress.GetFloor(c.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetFloor, 
			"error while reading floor data from pgress", err )
		return floor, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetFloor , 
		"returning floor data to api and setting response" , floor)
	c.JSON(http.StatusOK, floor)
	return floor, nil

}
func (server *Server) CreateFloor(c *gin.Context) error {

	var floor model.Floor
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateFloor, 
		"unmarshaling floor data",nil)

	err := c.ShouldBindJSON(&floor)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateFloor, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	floor.CreatedAt = time.Now().UTC()
	floor.ID = uuid.New()
	//validation is to be done here
	//DB call
	err = server.Pgress.CreateFloor(&floor)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateFloor, 
			"error while creating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetFloors , 
		"successfully created floor record and setting response" , floor)
	c.JSON(http.StatusCreated, floor)
	return nil

}

func (server *Server) UpdateFloor(c *gin.Context) error {

	var floor model.Floor
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdateFloor, 
		"unmarshaling floor data",nil)
	err := c.ShouldBindJSON(&floor)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateFloor, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	floor.UpdatedAt = time.Now().UTC()
	err = server.Pgress.UpdateFloor(&floor)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateFloor, 
			"error while updating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetFloors , 
		"successfully updated floor record and setting response" , floor)
	c.JSON(http.StatusOK, floor)
	return nil

}

func (server *Server) DeleteFloor(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteFloor, 
		"reading floor id",nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteFloor, 
			"missing floor id", nil )
		return fmt.Errorf("")
	}
	//DB call
	err := server.Pgress.DeleteFloor(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteFloor, 
			"error while deleting floor record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.DeleteFloor , 
		"successfully deleted floor record " , nil)
	return nil

}