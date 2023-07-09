package pgress 

import (
	"fmt"

	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/gorm"
)

func (store PgressStore) GetSensors() (*[]model.Sensor, error) {
	var sensors []model.Sensor
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensors, "reading all sensor data", nil)
	resp := store.DB.Find(&sensors)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetSensors,
			"error while reading all sensors data", resp.Error)
		return &sensors, fmt.Errorf("error while fetching sensors record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensors,
		"returning all sensor data", sensors)
	return &sensors, nil
}

func (store PgressStore) GetSensor(sensorID string) (*model.Sensor, error) {
	var sensor model.Sensor
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensor, "reading sensor data", nil)
	resp := store.DB.Find(&sensor, `"id" = '`+sensorID+`'`)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetSensor,
			"error while reading sensor data", resp.Error)
		return &sensor, fmt.Errorf("error while fetching sensor record from DB for given id, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensor,
		"returning sensor data", sensor)
	return &sensor, nil
}

func (store PgressStore) GetSensorsByFilter(filter map[string]string) (*[]model.Sensor, error) {
	var sensor []model.Sensor
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorByFilter,
		"reading sensor data from db based on filter", filter)
	var query *gorm.DB
	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == filter[model.StartDate] || key == filter[model.EndDate] {
			continue
		}
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorByFilter,
			"filters key", key+" value = "+value)
		query = store.DB.Where(fmt.Sprintf("%s = ?", key), value)
	}
	store.setLimitAndPage(filter, query)
	store.setDateRangeFilter(filter, query)
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorByFilter,
		"generated query ", query)
	err := query.Find(&sensor).Error
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetSensorByFilter,
			"error while reading sensor data", err)
		return &sensor, fmt.Errorf("error while fetching sensor record from DB for given id, err = %v", err)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorByFilter,
		"returning sensor data", sensor)
	return &sensor, nil
}

func (store PgressStore) CreateSensor(sensor *model.Sensor) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateSensor, "creating sensor data", *sensor)
	resp := store.DB.Create(sensor)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.CreateSensor,
			"error while creating sensor data", resp.Error)
		return fmt.Errorf("error while creating sensor record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateSensor,
		"successfully created sensor", nil)
	return nil
}

func (store PgressStore) UpdateSensor(sensor *model.Sensor) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateSensor, "updating sensor data", *sensor)
	resp := store.DB.Save(sensor)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.UpdateSensor,
			"error while updating sensor data", resp.Error)
		return fmt.Errorf("error while updating sensor record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateSensor,
		"successfully updated sensor", nil)
	return nil
}

// DeleteSensor is used to delete record by given sensorID
func (store PgressStore) DeleteSensor(sensorID string) error {

	var sensor model.Sensor
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteSensor, "deleting sensor data", sensorID)
	if err := store.DB.First(&sensor, `"id" = '`+sensorID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.DeleteSensor,
			"error while deleting sensor data", err)
		return fmt.Errorf("sensor not found for given id, ID = %v", sensorID)
	}
	resp := store.DB.Delete(sensor)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting sensor record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteSensor,
		"successfully deleted sensor", nil)
	return nil
}