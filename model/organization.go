package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Organization struct {
	ID                       uuid.UUID `json:"id" gorm:"primarykey"`
	PackageID                  uuid.UUID `json:"package_id" binding:"required" gorm:"not null"`
	Name                       string    `json:"name" binding:"required" gorm:"not null"`
	Description              string    `json:"description"`
	PackageType                string    `json:"package_type" binding:"required" gorm:"not null"`
	Email                    string    `json:"email" binding:"required" gorm:"not null"`
	Password                   string    `json:"password" binding:"required" gorm:"not null" `
	AvailablePoint           int       `json:"available_point"`
	AvailableNumberOfRooms   int       `json:"available_number_of_rooms"`
	AvailableNumberOfFloors  int       `json:"available_number_of_floors"`
	AvailableNumberOfHouses  int       `json:"available_number_of_house"`
	AvailableNumberOfUsers   int       `json:"available_number_of_users"`
	AvailableNumberOfSensors int       `json:"available_number_of_sensors"`
	CreatedBy                string    `json:"created_by"`
	CreatedAt                time.Time `json:"created"`
	UpdatedAt                time.Time `json:"updated_at"`
	UpdatedBy                string    `json:"updated_by"`
}
