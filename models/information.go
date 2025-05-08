package models

type DeviceInformation struct {
	DeviceID string                 `json:"device_id"`
	Info     map[string]interface{} `json:"information"`
}
