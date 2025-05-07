package models

import "encoding/json"

type DeviceCommand struct {
	Command string          `json:"command"`
	Params  json.RawMessage `json:"params"`
}

type SubmitCommand struct {
	DeviceID string          `json:"device_id"`
	Command  string          `json:"command"`
	Params   json.RawMessage `json:"params"`
}
