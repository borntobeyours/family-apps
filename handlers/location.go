package handlers

import (
	"context"
	"family-control-backend/database"
	"family-control-backend/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SubmitDeviceLocation(c *fiber.Ctx) error {
	deviceID := c.Query("device_id")
	if deviceID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing device_id",
		})
	}

	var loc models.DeviceLocation
	if err := c.BodyParser(&loc); err != nil {
		log.Println("[Location] JSON parse error:", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid JSON",
		})
	}

	_, err := database.DB.Exec(
		context.Background(),
		`INSERT INTO device_locations (device_id, latitude, longitude, timestamp)
		 VALUES ($1, $2, $3, $4)`,
		deviceID, loc.Latitude, loc.Longitude, time.Now(),
	)

	if err != nil {
		log.Println("[Location] DB error:", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to store location",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Location stored",
	})
}
