package handlers

import (
	"context"
	"family-control-backend/database"
	"family-control-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UploadContacts(c *fiber.Ctx) error {
	type Payload struct {
		DeviceID string           `json:"device_id"`
		Contacts []models.Contact `json:"contacts"`
	}

	log.Println("[UploadContacts] Hit")

	var payload Payload
	if err := c.BodyParser(&payload); err != nil {
		log.Println("Failed to parse:", err)
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid payload"})
	}

	for _, contact := range payload.Contacts {
		_, err := database.DB.Exec(
			context.Background(),
			`INSERT INTO device_contacts (device_id, name, number)
			 VALUES ($1, $2, $3)
			 ON CONFLICT (device_id, name, number) DO NOTHING`,
			payload.DeviceID, contact.Name, contact.Number,
		)
		if err != nil {
			log.Println("Insert error:", err)
		}
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Contacts stored"})
}
