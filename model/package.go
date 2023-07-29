package model

import (
	"time"

	"github.com/google/uuid"
)

type Package struct {
	ID               uuid.UUID     `json:"id" gorm:"primaryKey"`
	PointRateID      uuid.UUID     `json:"point_rate_Id"`
	Name             string        `json:"name" binding:"required" gorm:"not null"`
	Description      string        `json:"description"`
	Price            int           `json:"price" binding:"required" gorm:"not null"`
	Duration         time.Duration `json:"Duration" binding:"required" gorm:"not null"`
	Limit            string        `json:"limit"`
	NumberOfRooms    int           `json:"number_of_room"`
	NumberOfHouse    int           `json:"number_of_house"`
	NumberOfFloor    int           `json:"number_of_floor"`
	NumberOfUser     int           `json:"number_of_user"`
	NumberOfActuator int           `json:"number_of_actuator"`
	NumberOfSensor   int           `json:"number_of_sensor"`
	Points           int           `json:"point"`
	Benefits         string        `json:"benefits"`
	Availibility     string        `json:"availibility"`
	Features         string        `json:"features"`
	Status           bool          `json:"status" binding:"required" gorm:"not null"`
	CreatedAt        time.Time     `json:"created_at" gorm:"not null"`
	CreatedBy        string        `json:"created_by" binding:"required" gorm:"not null"`
	UpdatedAt        time.Time     `json:"updated_at"`
	UpdatedBy        string        `json:"updated_by"`
}
