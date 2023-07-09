package pgress

import (
	"fmt"

	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/gorm"
)

func (store PgressStore) GetUsers() (*[]model.User, error) {
	var users []model.User
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUsers, "reading all user data", nil)
	resp := store.DB.Find(&users)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetUsers,
			"error while reading all users data", resp.Error)
		return &users, fmt.Errorf("error while fetching users record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUsers,
		"returning all user data server", users)
	return &users, nil
}

func (store PgressStore) GetUser(userID string) (*model.User, error) {
	var user model.User
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUser,
		"reading user data from db", nil)
	resp := store.DB.Find(&user, `"id" = '`+userID+`'`)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetUser,
			"error while reading user data", resp.Error)
		return &user, fmt.Errorf("error while fetching user record from DB for given id, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUser,
		"returning user data", user)
	return &user, nil
}

func (store PgressStore) GetUsersByFilter(filter map[string]string) (*[]model.User, error) {
	var user []model.User
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUserByFilter,
		"reading user data from db based on filter", filter)
	var query *gorm.DB
	for key, value := range filter {
		if key == model.PageNumber || key == model.DataPerPage || key == filter[model.StartDate] || key == filter[model.EndDate]{
			continue
		}
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUserByFilter,
			"filters key", key + " value = "+ value)
		query = store.DB.Where(fmt.Sprintf("%s = ?", key), value)
	}
	store.setLimitAndPage(filter, query)
	store.setDateRangeFilter(filter , query)
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUserByFilter,
		"generated query ", query)
	err := query.Find(&user).Error
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetUserByFilter,
			"error while reading user data", err)
		return &user, fmt.Errorf("error while fetching user record from DB for given id, err = %v", err)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUserByFilter,
		"returning user data", user)
	return &user, nil
}

func (store PgressStore) CreateUser(user *model.User) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateUser, "creating user data", *user)
	resp := store.DB.Create(user)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.CreateUser,
			"error while creating user data", resp.Error)
		return fmt.Errorf("error while creating user record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateUser,
		"successfully created user", nil)
	return nil
}

func (store PgressStore) SignUp(user *model.User) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateUser, "creating user data", *user)
	resp := store.DB.Create(user)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.CreateUser,
			"error while creating user data", resp.Error)
		return fmt.Errorf("error while creating user record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.CreateUser,
		"successfully created user", nil)
	return nil
}

func (store PgressStore) SingIn(userSignIn model.UserSignIn) (*model.User, error) {
	var user model.User
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUser,
		"reading user data from db based on email", userSignIn)
	resp := store.DB.Where("email = ? AND password = ?", userSignIn.Email, userSignIn.Password).First(&user)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.GetUser,
			"error while reading user data", resp.Error)
		return &user, fmt.Errorf("error while fetching user record from DB for given id, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.GetUser,
		"returning user data", user)
	return &user, nil
}

func (store PgressStore) UpdateUser(user *model.User) error {

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateUser, "updating user data", *user)
	resp := store.DB.Save(user)
	if resp.Error != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.UpdateUser,
			"error while updating user data", resp.Error)
		return fmt.Errorf("error while updating user record, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.UpdateUser,
		"successfully updated user", nil)
	return nil
}

// DeleteUser is used to delete record by given userID
func (store PgressStore) DeleteUser(userID string) error {

	var user model.User
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteUser, "deleting user data", userID)
	if err := store.DB.First(&user, `"id" = '`+userID+`'`).Error; err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.DeleteUser,
			"error while deleting user data", err)
		return fmt.Errorf("user not found for given id, ID = %v", userID)
	}
	resp := store.DB.Delete(user)
	if resp.Error != nil {
		return fmt.Errorf("error while deleting user record from DB, err = %v", resp.Error)
	}
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.DeleteUser,
		"successfully deleted user", nil)
	return nil
}
