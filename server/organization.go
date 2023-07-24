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

func (server *Server) GetOrganizations(c *gin.Context) (*[]model.Organization, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetOrganizations, "reading all organization data", nil)
	organizations, err := server.Pgress.GetOraganizations()
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetOrganizations,
			"error while reading organizations data from pgress", err)
		return organizations, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetOrganizations,
		"returning all organization data to api and setting response", organizations)
	c.JSON(http.StatusOK, organizations)
	return organizations, nil

}

func (server *Server) GetOrganizationsByFilter(c *gin.Context) (*[]model.Organization, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetOrganizations, "reading all organization data", nil)
	condition := server.readQueryParams(c)
	organizations, err := server.Pgress.GetOraganizationsByFilter(condition)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetOrganizations,
			"error while reading organizations data from pgress", err)
		return organizations, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetOrganizations,
		"returning all organization data to api and setting response", organizations)
	c.JSON(http.StatusOK, organizations)
	return organizations, nil

}

func (server *Server) GetOrganization(c *gin.Context) (*model.Organization, error) {

	//validation is to be done here
	//DB call
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetOrganization,
		"reading organization data from pgress", nil)
	organization, err := server.Pgress.GetOraganization(c.Param("id"))
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.GetOrganization,
			"error while reading organization data from pgress", err)
		return organization, fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.GetOrganizations,
		"returning organization data to api and setting response", organization)
	c.JSON(http.StatusOK, organization)
	return organization, nil

}
func (server *Server) CreateOrganization(c *gin.Context) error {

	var organization model.Organization
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateOrganization,
		"unmarshaling organization data", nil)

	err := c.ShouldBindJSON(&organization)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateOrganization,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	organization.ID = uuid.New()
	organization.CreatedAt = time.Now().UTC()
	//validation is to be done here
	//DB call
	err = server.Pgress.CreateOrganization(&organization)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.CreateOrganization,
			"error while creating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.CreateOrganization,
		"successfully created organization record and setting response", organization)
	c.JSON(http.StatusCreated, organization)
	return nil

}

func (server *Server) UpdateOrganization(c *gin.Context) error {

	var organization model.Organization
	//Unmarshal
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdateOrganization,
		"unmarshaling organization data", nil)
	err := c.ShouldBindJSON(&organization)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateOrganization,
			"error while unmarshaling payload", err)
		return fmt.Errorf("")
	}
	//validation is to be done here
	//DB call
	organization.UpdatedAt = time.Now().UTC()
	err = server.Pgress.UpdateOrganization(&organization)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.UpdateOrganization,
			"error while updating record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.UpdateOrganization,
		"successfully updated organization record and setting response", organization)
	c.JSON(http.StatusOK, organization)
	return nil

}

func (server *Server) DeleteOrganization(c *gin.Context) error {

	//validation is to be done here
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteOrganization,
		"reading organization id", nil)
	id := c.Param("id")
	if id == "" {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteOrganization,
			"missing organization id", nil)
		return fmt.Errorf("")
	}
	//DB call
	err := server.Pgress.DeleteOrganization(id)
	if err != nil {
		util.Log(model.LogLevelError, model.ServerPackageLavel, model.DeleteOrganization,
			"error while deleting organization record from pgress", err)
		return fmt.Errorf("")
	}
	util.Log(model.LogLevelInfo, model.ServerPackageLavel, model.DeleteOrganization,
		"successfully deleted organization record ", nil)
	return nil

}