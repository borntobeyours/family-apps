package handlers

import (
	"context"
	"family-control-backend/database"
	"family-control-backend/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RegisterDevice(c *fiber.Ctx) error {
	log.Println("[RegisterDevice] Hit endpoint")

	var device models.Device
	if err := c.BodyParser(&device); err != nil {
		log.Println("[RegisterDevice] Failed to parse body:", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
		})
	}

	log.Printf("[RegisterDevice] Received: ID=%s, Model=%s, Android=%s\n",
		device.DeviceID, device.Model, device.AndroidVersion)

	query := `
		INSERT INTO devices (id, model, android_version, last_seen)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE
		SET last_seen = EXCLUDED.last_seen;
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		device.DeviceID,
		device.Model,
		device.AndroidVersion,
		time.Now(),
	)

	if err != nil {
		log.Println("[RegisterDevice] DB Error:", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to register device",
		})
	}

	log.Println("[RegisterDevice] Success")
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Device registered",
	})
}
