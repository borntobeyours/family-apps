package handlers

import (
	"context"
	"family-control-backend/database"
	"family-control-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UploadSms(c *fiber.Ctx) error {
	type Payload struct {
		DeviceID string              `json:"device_id"`
		Sms      []models.SmsMessage `json:"sms"`
	}

	log.Println("[UploadSms] Hit")

	var payload Payload
	if err := c.BodyParser(&payload); err != nil {
		log.Println("Failed to parse body:", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid payload",
		})
	}

	for _, sms := range payload.Sms {
		_, err := database.DB.Exec(
			context.Background(),
			`INSERT INTO device_sms (device_id, address, body, date, type)
			 VALUES ($1, $2, $3, $4, $5)
			 ON CONFLICT (device_id, address, body, date, type) DO NOTHING`,
			payload.DeviceID, sms.Address, sms.Body, sms.Date, sms.Type,
		)
		if err != nil {
			log.Println("Insert error:", err)
		}
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "SMS processed",
	})
}
