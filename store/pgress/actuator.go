package pgress

import (
	"fmt"

	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/gorm"
)

func (store PgressStore) GetActuators() (*[]model.Actuator, error) {
	var actuators []model.Actuator
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetActuators, "reading all actuator data ", nil)
	resp := store.DB.Find(&actuators)
	if resp.Error != nil {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetActuators,
			"error while reading all actuators data ", resp.Error)
		return &actuators, fmt.Errorf("error while fetching actuators record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetActuators,
		"returning all actuator data ", actuators)
	return &actuators, nil
}

func (store PgressStore) GetActuatorByFilter(filter map[string]string) (*[]model.Actuator, error) {
	var actuator []model.Actuator
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetActuatorByFilter,
		"reading actuator data from DB based on filter ", filter)
	var query *gorm.DB
	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == filter[model.StartDate] || key == filter[model.EndDate] {
			continue
		}
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetActuatorByFilter,
			"filter key ", key+" value = "+value)
		query = store.DB.Where(fmt.Sprintf("%s = ?", key), value)

	}
	store.setLimitAndPage(filter, query)
	store.setDateRangeFilter(filter, query)
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetActuatorByFilter,
		"generated query ", query)
	err := query.Find(&actuator).Error
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetActuatorByFilter,
			"eror while reading actuator data ", err)
		return &actuator, fmt.Errorf("error while fetching actuator record from DB for given id, err = %v", err)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetActuatorByFilter,
		"returning actuator data ", actuator)
	return &actuator, nil

}

func (store PgressStore) GetActuator(actuatorID string) (*model.Actuator, error) {
	var actuator model.Actuator
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetActuator, "reading actuator data ", nil)
	resp := store.DB.Find(&actuator, `"id" = '`+actuatorID+`'`)
	if resp.Error != nil {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetActuator,
			"error while reading actuator data ", resp.Error)
		return &actuator, fmt.Errorf("error while fetching actuator record from DB for given id, err = %v ", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetActuator,
		"returning actuator data ", actuator)
	return &actuator, nil

}

func (store PgressStore) CreateActuator(actuator *model.Actuator) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateActuator, "creating actuator data ", *actuator)
	resp := store.DB.Create(actuator)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.CreateActuator,
			"error while creating actuator data ", resp.Error)
		return fmt.Errorf("error while creating actuator record, err = %v", resp.Error)

	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateActuator,
		"successfully created actuator ", nil)
	return nil

}

func (store PgressStore) UpdateActuator(actuator *model.Actuator) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateActuator, "updating actuator data ", *actuator)
	resp := store.DB.Save(actuator)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.UpdateActuator,
			"error while updating actuator data ", resp.Error)
		return fmt.Errorf("error while updating actuator record, err = %v ", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateActuator,
		"successfully updated actuator ", nil)
	return nil

}

// DeleteActuator is used to delete record by given actuatorID
func (store PgressStore) DeleteActuator(actuatorID string) error {

	var actuator model.Actuator
	util.Log(model.LogLevelError, model.PgressPackageLevel, model.DeleteActuator, "deleting actuator data ", actuatorID)
	if err := store.DB.First(&actuator, `"id" = '`+actuatorID+`'`).Error; err != nil {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteActuator,
			"error while deleting actuator data ", err)
		return fmt.Errorf("actuator not found for gives id, ID = %v ", actuatorID)
	}
	resp := store.DB.Delete(actuator)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting actuator record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteActuator,
		"successfully deleted actuator ", nil)
	return nil
}
