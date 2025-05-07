package handlers

import (
	"context"
	"encoding/json"
	"family-control-backend/database"
	"family-control-backend/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SubmitAppUsage(c *fiber.Ctx) error {
	deviceID := c.Query("device_id")
	if deviceID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing device_id",
		})
	}

	var usages []models.AppUsage
	if err := json.Unmarshal(c.Body(), &usages); err != nil {
		log.Println("Invalid JSON:", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid JSON format",
		})
	}

	for _, usage := range usages {
		_, err := database.DB.Exec(
			context.Background(),
			`INSERT INTO app_usages (device_id, package_name, duration_seconds, timestamp)
			 VALUES ($1, $2, $3, $4)`,
			deviceID, usage.PackageName, usage.DurationSeconds, time.Now(),
		)
		if err != nil {
			log.Println("Insert error:", err)
			return c.Status(500).JSON(fiber.Map{
				"status":  "error",
				"message": "Failed to store usage",
			})
		}
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Usage data stored",
	})
}
