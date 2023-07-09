package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Room struct {
	ID                 uuid.UUID `json:"id" gorm:"primarykey"`
	HouseId            uuid.UUID `json:"house_id"binding:"required"gorm:"not null"`
	FloorID            uuid.UUID `json:"fllor_id" binding:"required"gorm:"not null"`
	Description        string    `json"description"`
	Area               string    `json:"area"`
	Occupancy          string    `json:"occupancy"`
	HeatingEnabled     bool      `json:"heating_enabled"`
	HeatingType        string    `json:"heating_type"`
	CoolingEnabled     bool      `json:"cooling_enabaled"`
	CoolingType        string    `json:"cooling_type"`
	VentilationEnabled bool      `json:"ventilation_enabled"`
	VentilationType    string    `json:"ventilation_type"`
	Temperature        string    `json:"temperature"`
	Humidity           string    `json:"humidity"`
	LightLevel         string    `json:"light_level"`
	CO2Level           string    `json:"co2_level"`
	CreatedAt          time.Time `json:"created_at"`
	CreatedBy          string    `json:"created_by"`
	UpdatedAt          time.Time `json:"updated_at"`
	UpdatedBy          string    `json:"updated_By"`
}
