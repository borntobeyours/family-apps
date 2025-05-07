package models

type AppUsage struct {
	PackageName     string `json:"package_name"`
	DurationSeconds int    `json:"duration_seconds"`
}
