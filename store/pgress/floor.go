package pgress

import (
	"fmt"

	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/gorm"
)

func (store PgressStore) GetFloors() (*[]model.Floor, error) {
	var floors []model.Floor
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetFloors, "reading all floor data", nil)
	resp := store.DB.Find(&floors)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetFloors,
			"error while reading all floors data", resp.Error)
		return &floors, fmt.Errorf("error while fetching floors record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetFloors,
		"returning all floor data", floors)
	return &floors, nil
}

func (store PgressStore) GetFloorsByFilter(filter map[string]string) (*[]model.Floor, error) {
	var floor []model.Floor
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetFloorByFilter,
		"reading floor data from DB based on filter", filter)
	var query *gorm.DB
	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == filter[model.StartDate] || key == filter[model.EndDate] {
			continue
		}
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetFloorByFilter,
			"filters key", key+" value = "+value)
		query = store.DB.Where(fmt.Sprintf("%s = ?", key), value)
	}
	store.setLimitAndPage(filter, query)
	store.setDateRangeFilter(filter, query)
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetFloorByFilter,
		"generated query ", query)
	err := query.Find(&floor).Error
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetFloorByFilter,
			"error while reading floor data", err)
		return &floor, fmt.Errorf("error while fetching floor record from DB for given id, err = %v", err)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetFloorByFilter,
		"returning floor data", floor)
	return &floor, nil
}

func (store PgressStore) GetFloor(floorID string) (*model.Floor, error) {
	var floor model.Floor
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetFloor, "reading floor data", nil)
	resp := store.DB.Find(&floor, `"id" = '`+floorID+`'`)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetFloor,
			"error while reading floor data", resp.Error)
		return &floor, fmt.Errorf("error while fetching floor record from DB for given id, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetFloor,
		"returning floor data", floor)
	return &floor, nil
}

func (store PgressStore) CreateFloor(floor *model.Floor) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateFloor, "creating floor data", *floor)
	resp := store.DB.Create(floor)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.CreateFloor,
			"error while creating floor data", resp.Error)
		return fmt.Errorf("error while creating floor record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateFloor,
		"successfully created floor", nil)
	return nil
}

func (store PgressStore) UpdateFloor(floor *model.Floor) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateFloor, "updating floor data", *floor)
	resp := store.DB.Save(floor)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.UpdateFloor,
			"error while updating floor data", resp.Error)
		return fmt.Errorf("error while updating floor record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateFloor,
		"successfully updated floor", nil)
	return nil
}

// DeleteFloor is used to delete record by given floorID
func (store PgressStore) DeleteFloor(floorID string) error {

	var floor model.Floor
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteFloor, "deleting floor data", floorID)
	if err := store.DB.First(&floor, `"id" = '`+floorID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.DeleteFloor,
			"error while deleting floor data", err)
		return fmt.Errorf("floor not found for given id, ID = %v", floorID)
	}
	resp := store.DB.Delete(floor)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting floor record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteFloor,
		"successfully deleted floor", nil)
	return nil
}
