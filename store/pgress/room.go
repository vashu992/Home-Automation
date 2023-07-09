package pgress

import (
	"fmt"

	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/gorm"
)

func (store PgressStore) GetRooms() (*[]model.Room, error) {
	var rooms []model.Room
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetRooms, "reading all room data ", nil)
	resp := store.DB.Find(&rooms)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetRooms,
			"error while reading all rooms data ", resp.Error)
		return &rooms, fmt.Errorf(" error while fetching rooms record from DB, err = %v ", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetRooms,
		"returning all room data ", rooms)
	return &rooms, nil
}

func (store PgressStore) GetRoomsByFilter(filter map[string]string) (*[]model.Room, error) {
	var room []model.Room
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetRoomByFilter,
		"reading room data from DB based on filter", filter)
	var query *gorm.DB
	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == filter[model.StartDate] || key == filter[model.EndDate] {
			continue
		}
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetRoomByFilter,
			"filter key ", key+"value = "+value)
		query = store.DB.Where(fmt.Sprintf("%s = ?", key), value)
	}
	store.setLimitAndPage(filter, query)
	store.setDateRangeFilter(filter, query)
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetRoomByFilter,
		"generated query ", query)
	err := query.Find(&room).Error
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetRoomByFilter,
			"error while reading room data ", err)
		return &room, fmt.Errorf(" error while fetching room record from DB for given id, err = %v", err)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetRoomByFilter,
		"returning room data ", room)
	return &room, nil
}

func (store PgressStore) GetRoom(roomID string) (*model.Room, error) {
	var room model.Room
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetRoom, "reading room data ", nil)
	resp := store.DB.Find(&room, `"id" = '`+roomID+`'`)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetRoom,
			"error while reading room data ", resp.Error)
		return &room, fmt.Errorf("error while fetching room record from DB for given id, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetRoom,
		"returning room data ", room)
	return &room, nil
}

func (store PgressStore) CreateRoom(room *model.Room) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateRoom, "creating room data ", *room)
	resp := store.DB.Create(room)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.CreateRoom,
			"error while creating room data ", resp.Error)
		return fmt.Errorf("error while creating room record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateRoom,
		"successfully created room ", room)
		return nil
}

func (store PgressStore) UpdateRoom(room *model.Room) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateRoom, "updating room data ", *room )
	resp := store.DB.Save(room)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.UpdateRoom, 
		"error while updating room data ", resp.Error)
		return fmt.Errorf("error while updating room record, err = %v ", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateRoom,
	"successfully updated room ", nil)
	return nil

}

// DeleteRoom  is used to delete record by given roomID 
func ( store PgressStore) DeleteRoom(roomID string) error {

	var room model.Room
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteRoom, "deleting room data ", roomID)
	if err := store.DB.First(&room, `"id" '`+roomID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.DeleteRoom, 
		"error while deleting room data ", err)
		return fmt.Errorf("room not found for given id, ID = %v", roomID)
	}
	resp := store.DB.Delete(room)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting room record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteRoom, 
	"successfully deleted room ", nil)
	return nil
}