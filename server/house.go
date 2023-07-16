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

func (server *Server) GetHouses(c *gin.Context) (*[]model.House, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo , model.ServerPackageLavel , model.GetHouses , "reading all house data", nil )
	houses, err := server.Pgress.GetHouses()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetHouses, 
			"error while reading houses data from pgress", err)
		return houses, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetHouses , 
		"returning all house data to api and setting response" , houses )
	c.JSON(http.StatusOK, houses)
	return houses, nil

}

func (server *Server) GetHousesByFilter(c *gin.Context) (*[]model.House, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetHouses, "reading all house data", nil)
	condition := server.readQueryParams(c)
	houses, err := server.Pgress.GetHousesByFilter(condition)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetHouses,
			"error while reading houses data from pgress", err)
		return houses, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetHouses,
		"returning all house data to api and setting response", houses)
	c.JSON(http.StatusOK, houses)
	return houses, nil

}

func (server *Server) GetHouse(c *gin.Context) (*model.House, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetHouse, 
		"reading house data from pgress", nil)
	house, err := server.Pgress.GetHouse(c.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetHouse, 
			"error while reading house data from pgress", err )
		return house, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetHouses , 
		"returning house data to api and setting response" , house)
	c.JSON(http.StatusOK, house)
	return house, nil

}
func (server *Server) CreateHouse(c *gin.Context) error {

	var house model.House
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateHouse, 
		"unmarshaling house data",nil)

	err := c.ShouldBindJSON(&house)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateHouse, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	house.CreatedAt = time.Now().UTC()
	house.ID = uuid.New()
	//validation is to be done here
	//DB call
	err = server.Pgress.CreateHouse(&house)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateHouse, 
			"error while creating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetHouses , 
		"successfully created house record and setting response" , house)
	c.JSON(http.StatusCreated, house)
	return nil

}

func (server *Server) UpdateHouse(c *gin.Context) error {

	var house model.House
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdateHouse, 
		"unmarshaling house data",nil)
	err := c.ShouldBindJSON(&house)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateHouse, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	house.UpdatedAt = time.Now().UTC()
	err = server.Pgress.UpdateHouse(&house)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateHouse, 
			"error while updating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetHouses , 
		"successfully updated house record and setting response" , house)
	c.JSON(http.StatusOK, house)
	return nil

}

func (server *Server) DeleteHouse(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteHouse, 
		"reading house id",nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteHouse, 
			"missing house id", nil )
		return fmt.Errorf("")
	}
	//DB call
	err := server.Pgress.DeleteHouse(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteHouse, 
			"error while deleting house record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.DeleteHouse , 
		"successfully deleted house record " , nil)
	return nil

}