package pgress

import (
	"fmt"

	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/gorm"
)

func (store PgressStore) GetOraganizations() (*[]model.Organization, error) {
	var Organizations []model.Organization
	util.Log(model.LogLevelInfo,model.PgressPackageLevel, model.GetOrganizations, "reading all organization data ", nil)
	resp := store.DB.Find(&Organizations)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetOrganizations,
		"error while reading all organizations data ", resp.Error)
		return &Organizations, fmt.Errorf("error while fetching organizations record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetOrganizations,
	"returning all organization data ", Organizations)
	return &Organizations, nil
}

func (store PgressStore) GetOraganizationsByFilter(filter map[string]string) (*[]model.Organization, error) {
	var organization []model.Organization
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetOrganizationByFilter, 
	"reading organization data from DB based on filter ", filter )
	var query *gorm.DB
	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == filter[model.StartDate]  || key == filter[model.EndDate] {
			continue
		}
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetOrganizationByFilter, 
		"filters key ", key+" value = "+value)
		query = store.DB.Where(fmt.Sprintf("%s = ?", key ), value)
	}
	store.setLimitAndPage(filter, query)
	store.setDateRangeFilter(filter, query)
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetOrganizationByFilter,
	"generated query", query)
	err := query.Find(&organization).Error
	if err != nil {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetOrganizationByFilter,
		"error while reading organization data ", err )
		return &organization, fmt.Errorf("error while fetching organization record from DB for given id, err = %v", err)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetOrganizationByFilter,
	"returning organization data ", organization)
	return &organization, nil
}

func (store PgressStore) GetOraganization(organizationID string) (*model.Organization, error) {
	var organization model.Organization
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetOrganization, "reading organization data ", nil)
	resp := store.DB.Find(&organization, `"id" = '`+organizationID+`'`)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetOrganization,
		"error while reading organization data ",resp.Error)
		return &organization, fmt.Errorf(" error while fetching organization record from DB for given id, err = %v",resp.Error)

	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetOrganization,
	"returning organization data",organization)
	return &organization, nil
}

func (store PgressStore) CreateOrganization(organization *model.Organization )error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateOrganization, "creating organization data ",*organization)
	resp:= store.DB.Create(organization)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.CreateOrganization,
		"error while creating organization data ", resp.Error)
		return fmt.Errorf("error while creating organization record, err = %v ",resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateOrganization,
	"successfully created organization ",nil)
	return nil
}

func(store PgressStore) UpdateOrganization(Organization *model.Organization) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateOrganization, "updating organization data ",*Organization)
	resp:= store.DB.Save(Organization)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.UpdateOrganization, 
		"error while updating organization data ", resp.Error)
		return fmt.Errorf("error while updating organization record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateOrganization,
	"successfully updated organization ", nil)
	return nil
}

// DeleteOrganization is used to delete record by given organizationID
func (store PgressStore) DeleteOrganization(organizationID string) error {

	var organization model.Organization
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteOrganization, "deleting organization data", organizationID)
	if err := store.DB.First(&organization, `"id" = '`+organizationID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.DeleteOrganization,
			"error while deleting organization data", err)
		return fmt.Errorf("organization not found for given id, ID = %v", organizationID)
	}
	resp := store.DB.Delete(organization)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting organization record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteOrganization,
		"successfully deleted organization", nil)
	return nil
}
