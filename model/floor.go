package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Floor struct {
	ID                 uuid.UUID `json:"id" gorm:"primarykey" `
	HouseID            uuid.UUID `json:"house_id" binding:"required" gorm:"not null"`
	Name               string    `json:"name" binding:"required" gorm:"not null"`
	Description        string    `json:"area"`
	Occupancy          string    `json:"occupancy"`
	HeatingEnabled     bool      `json:"heating_enabled"`
	HeatingType        string    `json:"heating_type"`
	CoolingEnabled     bool      `json:"cooling_enabled"`
	CoolingType        string    `json:"cooling_Type"`
	VentilationEnabled bool      `json:"ventilation_enabled"`
	VentilationType    string    `json:"ventilation_type"`
	Temperature        string    `json:"temperature"`
	Humidity           string    `json:"humidity"`
	LightLevel         string    `json:"light_level"`
	CO2Level           string    `json:"co2_level"`
	CreatedAt          time.Time `json:"created_at"`
	CreatedBy          string    `json:"created_by"`
	UpdatedAt          time.Time `json:"updated_at"`
	UpdatedBy          string    `json:"updated_by"`
}
