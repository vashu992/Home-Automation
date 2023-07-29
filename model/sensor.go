package model

import (
	"time"

	"github.com/google/uuid"
)

type Sensor struct {
	ID          uuid.UUID     `json:"id" gorm:"primarykey"`
	RoomID      uuid.UUID     `json:"room_id" binding:"required" gorm:"not null"`
	HouseID     uuid.UUID     `json:"house_id" binding:"required" gorm:"not null"`
	FloorID     uuid.UUID     `json:"floor_id" binding:"required" gorm:"not null"`
	Name        string        `json:"name" binding:"required" gorm:"not null"`
	Status      string        `json:"status" binding:"required" gorm:"not null"`
	UpdatedBy   string        `json:"updated_by"`
	CreatedBy   string        `json:"created_by" binding:"required" gorm:"not null"`
	Reading     string        `json:"reading"`
	Unit        string        `json:"unit"`
	MinReading  string        `json:"min_reading"`
	MaxReading  string        `json:"max_reading"`
	RefreshRate time.Duration `json:"refresh_rate" binding:"required" gorm:"not null"`
	StartTime   time.Time     `json:"start_time"`
	EndTime     time.Time     `json:"end_time"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}
