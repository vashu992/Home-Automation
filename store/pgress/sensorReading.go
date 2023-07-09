package pgress

import(
	"fmt"

	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/gorm"
)

func (store PgressStore) GetSensorReadings() (*[]model.SensorReading, error) {
	var sensors []model.SensorReading
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorReadings, "reading all sensorReading data", nil)
	resp := store.DB.Find(&sensors)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetSensorReadings,
			"error while reading all sensors data", resp.Error)
		return &sensors, fmt.Errorf("error while fetching sensors record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorReadings,
		"returning all sensorReading data", sensors)
	return &sensors, nil
}

func (store PgressStore) GetSensorReading(sensorID string) (*model.SensorReading, error) {
	var sensorReading model.SensorReading
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorReading, "reading sensorReading data", nil)
	resp := store.DB.Find(&sensorReading, `"id" = '`+sensorID+`'`)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetSensorReading,
			"error while reading sensorReading data", resp.Error)
		return &sensorReading, fmt.Errorf("error while fetching sensorReading record from DB for given id, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorReading,
		"returning sensorReading data", sensorReading)
	return &sensorReading, nil
}

func (store PgressStore) GetSensorReadingsByFilter(filter map[string]string) (*[]model.SensorReading, error) {
	var sensorReading []model.SensorReading
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorReadingByFilter,
		"reading sensorReading data from db based on filter", filter)
	var query *gorm.DB
	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == filter[model.StartDate] || key == filter[model.EndDate] {
			continue
		}
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorReadingByFilter,
			"filters key", key+" value = "+value)
		query = store.DB.Where(fmt.Sprintf("%s = ?", key), value)
	}
	store.setLimitAndPage(filter, query)
	store.setDateRangeFilter(filter, query)
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorReadingByFilter,
		"generated query ", query)
	err := query.Find(&sensorReading).Error
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetSensorReadingByFilter,
			"error while reading sensorReading data", err)
		return &sensorReading, fmt.Errorf("error while fetching sensorReading record from DB for given id, err = %v", err)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetSensorReadingByFilter,
		"returning sensorReading data", sensorReading)
	return &sensorReading, nil
}

func (store PgressStore) CreateSensorReading(sensorReading *model.SensorReading) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateSensorReading, "creating sensorReading data", *sensorReading)
	resp := store.DB.Create(sensorReading)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.CreateSensorReading,
			"error while creating sensorReading data", resp.Error)
		return fmt.Errorf("error while creating sensorReading record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateSensorReading,
		"successfully created sensorReading", nil)
	return nil
}

func (store PgressStore) UpdateSensorReading(sensorReading *model.SensorReading) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateSensorReading, "updating sensorReading data", *sensorReading)
	resp := store.DB.Save(sensorReading)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.UpdateSensorReading,
			"error while updating sensorReading data", resp.Error)
		return fmt.Errorf("error while updating sensorReading record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateSensorReading,
		"successfully updated sensorReading", nil)
	return nil
}

// DeleteSensorReading is used to delete record by given sensorID
func (store PgressStore) DeleteSensorReading(sensorID string) error {

	var sensorReading model.SensorReading
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteSensorReading, "deleting sensorReading data", sensorID)
	if err := store.DB.First(&sensorReading, `"id" = '`+sensorID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.DeleteSensorReading,
			"error while deleting sensorReading data", err)
		return fmt.Errorf("sensorReading not found for given id, ID = %v", sensorID)
	}
	resp := store.DB.Delete(sensorReading)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting sensorReading record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteSensorReading,
		"successfully deleted sensorReading", nil)
	return nil
}