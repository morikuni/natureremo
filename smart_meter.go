package natureremo

import (
	"time"
)

type SmartMeter struct {
	ECHONETLiteProperties []*ECHONETLiteProperty `json:"echonetlite_properties"`
}

type ECHONETLiteProperty struct {
	Name      string    `json:"name"`
	EPC       uint8     `json:"epc"`
	Val       string    `json:"val"`
	UpdatedAt time.Time `json:"updated_at"`
}
