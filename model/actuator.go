package model

import (
	"time"

	uuid "github.com/google/uuid"
)

type Actuator struct {
	Id                 uuid.UUID `json:"id" gorm:"primarykey"`
	RoomsID            uuid.UUID `json:"rooms_id" binding:"required" gorm:"not null"`
	HouseID            uuid.UUID `json:"house_id" binding:"required" gorm:"not null"`
	FloorID            uuid.UUID `json:"floor_id" binding:"required" gorm:"not null"`
	Name               string    `json:"name" binding:"required" gorm:"not null"`
	Description        string    `json:"description"`
	Area               string    `json:"area"`
	WattConsuption     string    `json:"watt_consuption"`
	Occupancy          string    `json:"occupancy"`
	HeatingEnabled     string    `json:"heating_enabled"`
	HeatingType        string    `json:"heating_type"`
	CoolingEnabled     bool      `json:"cooling_enabled"`
	CoolingType        string    `json:"cooling_type"`
	VentilationEnabled bool      `json:"ventilation_enabled"`
	VentilationType    string    `json:"ventilation_type"`
	Temperature        string    `json:"temperature"`
	Humidity           string    `json:"humidity"`
	LightLevel         string    `json:"light_level"`
	CO2Level           string    `json:"co2_level"`
	CreatedAt          time.Time `json:"craeted_at" gorm:"not null"`
	CreatedBy          string    `json:"created_by" binding:"required" gorm:"not null"`
	UpdatedAt          time.Time `json:"updated_at"`
	UpdatedBy          string    `json:"updated_by"`
}
