package models

type Device struct {
	DeviceID       string `json:"device_id"`
	Model          string `json:"model"`
	AndroidVersion string `json:"android_version"`
}
