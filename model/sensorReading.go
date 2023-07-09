package model

import (

	"time"

	"github.com/google/uuid"


)

type SensorReading struct {
	ID   uuid.UUID  `json:"id"gorm:"primarykey"`
	SensorID uuid.UUID `json:"sensor_id"gorm:`
	Reading  string `json:"reading"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}