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

func (server *Server) GetRooms(c *gin.Context) (*[]model.Room, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo , model.ServerPackageLavel , model.GetRooms , "reading all room data", nil )
	rooms, err := server.Pgress.GetRooms()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetRooms, 
			"error while reading rooms data from pgress", err)
		return rooms, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetRooms , 
		"returning all room data to api and setting response" , rooms )
	c.JSON(http.StatusOK, rooms)
	return rooms, nil

}

func (server *Server) GetRoomsByFilter(c *gin.Context) (*[]model.Room, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetRooms, "reading all room data", nil)
	condition := server.readQueryParams(c)
	rooms, err := server.Pgress.GetRoomsByFilter(condition)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetRooms,
			"error while reading rooms data from pgress", err)
		return rooms, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetRooms,
		"returning all room data to api and setting response", rooms)
	c.JSON(http.StatusOK, rooms)
	return rooms, nil

}

func (server *Server) GetRoom(c *gin.Context) (*model.Room, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetRoom, 
		"reading room data from pgress", nil)
	room, err := server.Pgress.GetRoom(c.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetRoom, 
			"error while reading room data from pgress", err )
		return room, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetRooms , 
		"returning room data to api and setting response" , room)
	c.JSON(http.StatusOK, room)
	return room, nil

}
func (server *Server) CreateRoom(c *gin.Context) error {

	var room model.Room
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateRoom, 
		"unmarshaling room data",nil)

	err := c.ShouldBindJSON(&room)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateRoom, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	room.CreatedAt = time.Now().UTC()
	room.ID = uuid.New()
	//validation is to be done here
	//DB call
	err = server.Pgress.CreateRoom(&room)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateRoom, 
			"error while creating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetRooms , 
		"successfully created room record and setting response" , room)
	c.JSON(http.StatusCreated, room)
	return nil

}

func (server *Server) UpdateRoom(c *gin.Context) error {

	var room model.Room
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdateRoom, 
		"unmarshaling room data",nil)
	err := c.ShouldBindJSON(&room)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateRoom, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	room.UpdatedAt = time.Now().UTC()
	err = server.Pgress.UpdateRoom(&room)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateRoom, 
			"error while updating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetRooms , 
		"successfully updated room record and setting response" , room)
	c.JSON(http.StatusOK, room)
	return nil

}

func (server *Server) DeleteRoom(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteRoom, 
		"reading room id",nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteRoom, 
			"missing room id", nil )
		return fmt.Errorf("")
	}
	//DB call
	err := server.Pgress.DeleteRoom(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteRoom, 
			"error while deleting room record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.DeleteRoom , 
		"successfully deleted room record " , nil)
	return nil

}