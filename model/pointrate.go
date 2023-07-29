package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type PointRate struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" binding:"required" gorm:"not null"`
	Description string    `json:"description"`
	House       int       `json:"house" binding:"required" gorm:"not null"`
	Floor       int       `json:"floor" binding:"required" gorm:"not null"`
	Room        int       `json:"room" binding:"required" gorm:"not null"`
	User        int       `json:"user" binding:"required" gorm:"not null"`
	Sensor      int       `json:"sensor" binding:"required" gorm:"not null"`
	Actuator    int       `json:"actuator" binding:"required" gorm:"not null"`
	CreatedBy   string    `json:"created_by"  gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   string    `json:"updated_by"`
}
