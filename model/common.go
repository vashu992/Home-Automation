package model

import (
	"time"

)

type CommonParameters struct {
	Name       string  `json:"name"`
	StartTime  time.Time `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`

}

// ErrorResponce struct

type ErrorResponce struct {
	Message  string   `json:"message"`

}

// SuccessResponce struct
type SuccessResponce struct {
	Message    string   `json:"message"`
}