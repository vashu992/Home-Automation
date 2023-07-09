package pgress

import (
	"fmt"


	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/gorm"

)

func (store PgressStore) GetHouses() (*[]model.House, error) {
	var houses []model.House
	util.Log(model.LogLevelInfo,model.PgressPackageLevel, model.GetHouses, "reading all house data ", nil)
	resp := store.DB.Find(&houses)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetHouses, 
		"error while reading all houses data ", resp.Error)
	}
	util.Log(model.LogLevelInfo,model.PgressPackageLevel, model.GetHouses,
	"returning all house data ", houses)
	return &houses, nil

}

func (store PgressStore) GetHousesByFilter(filter map[string]string) (*[]model.House, error) {
	var house []model.House
	util.Log(model.LogLevelInfo, model.PgressPackageLevel,model.GetHouseByFilter,
	"reading house data from DB baesd on filter  ", filter)
	var query *gorm.DB
	for key, value := range filter {
		if key == model.PageNumber  || key == model.DataPerPage || key == filter[model.StartDate] || key == filter[model.EndDate] {
			continue
		}
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetHouseByFilter,
		"filters key ", key+ "value = "+value)
		query = store.DB.Where(fmt.Sprintf("%s = ?",key), value)
	}
	store.setLimitAndPage(filter, query)
	store.setDateRangeFilter(filter, query)
	util.Log(model.LogLevelInfo,model.PgressPackageLevel,model.GetHouseByFilter,
	"generated query ", query)
	err := query.Find(&house).Error
	if err != nil { 
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetHouseByFilter,
		"error while reading house data ",err )
		return &house, fmt.Errorf(" error while fetching house record from DB for given id, err = %v", err)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetHouseByFilter,
	"returning house data", house)
	return &house, nil 

}
func (store PgressStore) GetHouse(houseID string) (*model.House, error) {
	var house model.House
	util.Log(model.LogLevelInfo, model.PgressPackageLevel,model.GetHouse, "reading house data ", nil)
	resp := store.DB.Find(&house, `"id" = '`+houseID+`'`)
	if resp.Error != nil {
		util.Log(model.LogLevelInfo,model.PgressPackageLevel, model.GetHouse,
		"error while reading house data ", resp.Error)
		return &house, fmt.Errorf("error while fetching house record from DB for given id, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetHouse,
	"returning house data ", house)
	return &house, nil
}

func (store PgressStore) CreateHouse(house *model.House) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateHouse, "creating house data", *house)
	resp := store.DB.Create(house)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.CreateHouse,
			"error while creating house data", resp.Error)
		return fmt.Errorf("error while creating house record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateHouse,
		"successfully created house", nil)
	return nil

}

func (store PgressStore) UpdateHouse(house *model.House) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateHouse, "updating house data", *house)
	resp := store.DB.Save(house)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.UpdateHouse,
			"error while updating house data", resp.Error)
		return fmt.Errorf("error while updating house record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateHouse,
		"successfully updated house", nil)
	return nil
}

// DeleteHouse is used to delete record by given houseID
func (store PgressStore) DeleteHouse(houseID string) error {

	var house model.House
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteHouse, "deleting house data", houseID)
	if err := store.DB.First(&house, `"id" = '`+houseID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.DeleteHouse,
			"error while deleting house data", err)
		return fmt.Errorf("house not found for given id, ID = %v", houseID)
	}
	resp := store.DB.Delete(house)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting house record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteHouse,
		"successfully deleted house", nil)
	return nil
}