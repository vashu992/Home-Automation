package pgress

import (
	"fmt"

	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/gorm"
)

func (store PgressStore) GetPackages() (*[]model.Package, error) {
	var pkgs []model.Package
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackages, "reading all pkg data ", nil)
	resp := store.DB.Find(&pkgs)
	if resp.Error != nil {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackages,
			"error while reading all pkg data ", resp.Error)
		return &pkgs, fmt.Errorf("error while fetching pkg record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackages,
		"returining all pkg data ", pkgs)
	return &pkgs, nil

}

func (store PgressStore) GetPackagesByFilter(filter map[string]string) (*[]model.Package, error) {
	var pkg []model.Package
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackageByFilter,
		"reading pkg data from DB based on filter", filter)
	var query *gorm.DB
	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == filter[model.StartDate] || key == filter[model.EndDate] {
			continue
		}
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackageByFilter,
			"filters key", key+"value = "+value)
		query = store.DB.Where(fmt.Sprintf("%s = ?", key), value)
	}
	store.setLimitAndPage(filter, query)
	store.setDateRangeFilter(filter, query)
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackageByFilter,
		"generated query ", query)
	err := query.Find(&pkg).Error
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetPackageByFilter,
			"error while reading pkg data ", err)
		return &pkg, fmt.Errorf("error while fetching pkg record from DB for given id, err = %v", err)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackageByFilter,
		"returning pkg data ", pkg)
	return &pkg, nil
}

func (store PgressStore) GetPackage(pkgID string) (*model.Package, error) {
	var pkg model.Package
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackage, "reading pkg data", nil)
	resp := store.DB.Find(&pkg, `"id" = '`+pkgID+`'`)
	if resp.Error != nil {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackage,
			"error while reading pkg data", resp.Error)
		return &pkg, fmt.Errorf("error while fetching pkg record from DB for given id , err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetPackage,
		"returning pkg data ", pkg)
	return &pkg, nil

}
func (store PgressStore) CreatePackage(pkg *model.Package) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreatePackage, "creating pkg data ", *pkg)
	resp := store.DB.Create(pkg)
	if resp.Error != nil {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreatePackage,
			"error while creating pkg data ", resp.Error)
		return fmt.Errorf("error while creating pkg record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreatePackage,
		"successfully created pkg", nil)
	return nil
}

func (store PgressStore) UpdatePackage(pkg *model.Package) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdatePackage, "updating pkg data", *pkg)
	resp := store.DB.Save(pkg)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.UpdatePackage,
			"error while updating pkg data", resp.Error)
		return fmt.Errorf("error while updating pkg record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdatePackage,
		"successfully updated pkg", nil)
	return nil

}

// DeletePackage is used to delete record by given pkgID
func (store PgressStore) DeletePackage(pkgID string) error {

	var pkg model.Package
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeletePackage, "deleting pkg data", pkgID)
	if err := store.DB.First(&pkg, `"id" = '`+pkgID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.DeletePackage,
			"error while deleting pkg data", err)
		return fmt.Errorf("pkg not found for given id, ID = %v", pkgID)
	}
	resp := store.DB.Delete(pkg)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting pkg record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeletePackage,
		"successfully deleted pkg", nil)
	return nil
}
