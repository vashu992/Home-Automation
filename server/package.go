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

func (server *Server) GetPackages(c *gin.Context) (*[]model.Package, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo , model.ServerPackageLavel , model.GetPackages , "reading all pkg data", nil )
	pkgs, err := server.Pgress.GetPackages()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetPackages, 
			"error while reading pkgs data from pgress", err)
		return pkgs, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetPackages , 
		"returning all pkg data to api and setting response" , pkgs )
	c.JSON(http.StatusOK, pkgs)
	return pkgs, nil

}

func (server *Server) GetPackagesByFilter(c *gin.Context) (*[]model.Package, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetPackages, "reading all pkg data", nil)
	condition := server.readQueryParams(c)
	pkgs, err := server.Pgress.GetPackagesByFilter(condition)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetPackages,
			"error while reading pkgs data from pgress", err)
		return pkgs, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetPackages,
		"returning all pkg data to api and setting response", pkgs)
	c.JSON(http.StatusOK, pkgs)
	return pkgs, nil

}

func (server *Server) GetPackage(c *gin.Context) (*model.Package, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetPackage, 
		"reading pkg data from pgress", nil)
	pkg, err := server.Pgress.GetPackage(c.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetPackage, 
			"error while reading pkg data from pgress", err )
		return pkg, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetPackages , 
		"returning pkg data to api and setting response" , pkg)
	c.JSON(http.StatusOK, pkg)
	return pkg, nil

}
func (server *Server) CreatePackage(c *gin.Context) error {

	var pkg model.Package
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreatePackage, 
		"unmarshaling pkg data",nil)

	err := c.ShouldBindJSON(&pkg)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreatePackage, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	pkg.CreatedAt = time.Now().UTC()
	pkg.ID = uuid.New()
	//validation is to be done here
	//DB call
	err = server.Pgress.CreatePackage(&pkg)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreatePackage, 
			"error while creating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetPackages , 
		"successfully created pkg record and setting response" , pkg)
	c.JSON(http.StatusCreated, pkg)
	return nil

}

func (server *Server) UpdatePackage(c *gin.Context) error {

	var pkg model.Package
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdatePackage, 
		"unmarshaling pkg data",nil)
	err := c.ShouldBindJSON(&pkg)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdatePackage, 
			"error while unmarshaling payload", err )
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	pkg.UpdatedAt = time.Now().UTC()
	err = server.Pgress.UpdatePackage(&pkg)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdatePackage, 
			"error while updating record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.GetPackages , 
		"successfully updated pkg record and setting response" , pkg)
	c.JSON(http.StatusOK, pkg)
	return nil

}

func (server *Server) DeletePackage(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeletePackage, 
		"reading pkg id",nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeletePackage, 
			"missing pkg id", nil )
		return fmt.Errorf("")
	}
	//DB call
	err := server.Pgress.DeletePackage(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeletePackage, 
			"error while deleting pkg record from pgress", err )
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel , model.DeletePackage , 
		"successfully deleted pkg record " , nil)
	return nil

}
