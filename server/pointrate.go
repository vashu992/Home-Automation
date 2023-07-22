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

func (server *Server) GetPointRates(c *gin.Context) (*[]model.PointRate, error) {

	// validation is to be done here
	// DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetPointRates, "reading all pointer data ", nil)
	pointers, err := server.Pgress.GetPointRates()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetPointRates,
			"error while reading pointers data from pgress ", err)
		return pointers, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetPointRates,
		"returning all pointer data to api and setting response ", pointers)
	c.JSON(http.StatusOK, pointers)
	return pointers, nil

}

func (server *Server) GetPointRatesByFilter(c *gin.Context) (*[]model.PointRate, error) {

	// validation is  to be done here
	// DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetPointRates, "reading all pointRate data ", nil)
	condition := server.readQueryParams(c)
	pointRates, err := server.Pgress.GetPointRatesByFilter(condition)
	if err != nil {
		util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetPointRates,
			"error while reading pointRates data from pgress ", err)
		return pointRates, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetPointRates,
		"returning all pointRate data to api and setting response ", pointRates)
	c.JSON(http.StatusOK, pointRates)
	return pointRates, nil

}

func (server *Server) GetPointRate(c *gin.Context) (*model.PointRate, error) {

	// validation is to be done here
	// DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetPointRate,
		"reading pointer data from pgress", nil)
	pointer, err := server.Pgress.GetPointRate(c.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetPointRate,
			"error while reading pointer data from pgress ", err)
		return pointer, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetPointRates,
		"returning pointer data to api and setting  response ", pointer)
	c.JSON(http.StatusOK, pointer)
	return pointer, nil

}

func (server *Server) CreatePointRate(c *gin.Context) error {

	var pointer model.PointRate
	// Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreatePointRate,
		"unmarshaling pointer data ", nil)

	err := c.ShouldBindJSON(&pointer)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreatePointRate,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	pointer.ID = uuid.New()
	pointer.CreatedAt = time.Now().UTC()
	// validation is to be done here
	// DB call
	err = server.Pgress.CreatePointRate(&pointer)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreatePointRate,
			"error while creating record from pgress ", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreatePointRate,
		"successfully created pointer record and setting response ", pointer)
	c.JSON(http.StatusCreated, pointer)
	return nil
}

func (server *Server) UpdatePointRate(c *gin.Context) error {

	var pointer model.PointRate
	// Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdatePointRate,
		"unmarshaling pointer data ", nil)
	err := c.ShouldBindJSON(&pointer)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdatePointRate,
			"error while unmarshaling payload ", err)
		return fmt.Errorf("")
	}
	// validation is to be done here
	// DB call
	pointer.UpdatedAt = time.Now().UTC()
	err = server.Pgress.UpdatePointRate(&pointer)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdatePointRate,
			"error while updating record from pgress ", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdatePointRate,
		"successfully updated pointer record and setting response ", pointer)
	c.JSON(http.StatusOK, pointer)
	return nil

}

func (server *Server) DeletePointRate(c *gin.Context) error {

	// validation is to be done here 
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeletePointRate,
	"reading pointer id ", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeletePointRate,
		"missing pointer id ", nil)
		return fmt.Errorf("")
	}
	// DB call
	err := server.Pgress.DeletePointRate(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeletePointRate,
		"error while deleting pointer record from pgress ", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeletePointRate,
	"successfully deleted pointer record ", err)
	return nil
}
