package pgress

import (
	"github.com/vashu992/Home-Automation/model"
	"github.com/vashu992/Home-Automation/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgressStore struct {
	DB *gorm.DB
}

func (store *PgressStore) NewStore() {
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.NewStore, "Setting DataBase", nil)

	db, err := gorm.Open(postgres.Open(model.DSN), &gorm.Config{})
	if err != nil {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.NewStore,
			"error while connecting database", err)
		panic(err)
	}

	err = db.AutoMigrate(
		model.Actuator{},
		model.Floor{},
		model.House{},
		model.Organization{},
		model.Package{},
		model.PointRate{},
		model.Room{},
		model.Sensor{},
		model.SensorReading{},
		model.User{},
	)
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.NewStore,
			"error while automigrating database", err)
		panic(err)
	}
	store.DB = db
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.NewStore, "setting dataBase completed .... ", nil)

}
