package pgress

import (
	"fmt"

	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/gorm"
)

func (store PgressStore) GetPointRates() (*[]model.PointRate, error) {
	var pointRates []model.PointRate
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPointRates, "reading all pointRate data ", nil)
	resp := store.DB.Find(&pointRates)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetPointRates,
			"error while reading all pointRate data ", resp.Error)
		return &pointRates, fmt.Errorf("error while fetching pointRates record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPointRates,
		"returning all pointRate data ", pointRates)
	return &pointRates, nil

}

func (store PgressStore) GetPointRatesByFilter(filter map[string]string) (*[]model.PointRate, error) {
	var pointRate []model.PointRate
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPointRateByFilter,
		"reading poinRate data from DB based on filter ", filter)
	var query *gorm.DB
	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == filter[model.StartDate] || key == filter[model.EndDate] {
			continue
		}
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackageByFilter,
			"filter key ", key+"value = "+value)
		query = store.DB.Where(fmt.Sprintf("%s = ?", key), value)
	}
	store.setLimitAndPage(filter, query)
	store.setDateRangeFilter(filter, query)
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPointRateByFilter,
		"generated query ", query)
	err := query.Find(&pointRate).Error
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetPointRateByFilter,
			"error while reading pointRate data ", err)
		return &pointRate, fmt.Errorf("error while fetching pointRate record from DB for given id, err = %v", err)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPointRateByFilter,
		"returning pointRate data ", pointRate)
	return &pointRate, nil
}

func (store PgressStore) GetPointRate(pointRateID string) (*model.PointRate, error) {
	var pointRate model.PointRate
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPointRate, "reading pointRate data ", nil)
	resp := store.DB.Find(&pointRate, `"id" = '`+pointRateID+`'`)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetPointRate,
			"error while reading pointRate data ", resp.Error)
		return &pointRate, fmt.Errorf(" error while reading pointRate record from DB for given id, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPointRate,
		"returning pointRate data ", pointRate)
	return &pointRate, nil
}

func (store PgressStore) CreatePointRate(pointRate *model.PointRate) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreatePointRate, "creating pointRate data ", *pointRate)
	resp := store.DB.Create(pointRate)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.CreatePointRate,
			"error while creating pointRate data ", resp.Error)
		return fmt.Errorf("error while creating pointRate record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreatePointRate,
	"successfully created pointRate", nil)
	return nil

}

func (store PgressStore) UpdatePointRate(pointRate *model.PointRate) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdatePointRate, "updating pointRate data ", *pointRate)
	resp := store.DB.Save(pointRate)
	if resp.Error != nil {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdatePointRate,
		"error while updating pointRate data ",resp.Error)
		return fmt.Errorf("error while updating pointRate record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdatePointRate,
	"successfully updated pointRate", nil)
	return nil

}

// DeletePointRate is used to delete record by given pointRateID 
func (store PgressStore) DeletePointRate(pointRateID string) error {

	var poinRate model.PointRate
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeletePointRate, "deleting pointRate data ",pointRateID)
	if err := store.DB.First(&poinRate, `"id" = '`+pointRateID+`'`).Error; err != nil {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeletePointRate, 
		"error while deleting pointRate data ", err)
		return fmt.Errorf("pointRate not found for given id, ID = %v", pointRateID)
	}
	resp := store.DB.Delete(poinRate)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting pointRate record from DB, err = %v ", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeletePointRate,
	"successfully deleted pointRate ", nil)
	return nil

}
