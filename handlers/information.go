package handlers

import (
	"context"
	"encoding/json"
	"family-control-backend/database"
	"family-control-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UploadDeviceInformation(c *fiber.Ctx) error {
	log.Println("[UploadDeviceInformation] Hit")

	var payload models.DeviceInformation
	if err := c.BodyParser(&payload); err != nil {
		log.Println("Parse error:", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid payload",
		})
	}

	// 🔍 Log data mentah yang dikirim mobile
	log.Printf("Raw payload: device_id=%s, information=%+v\n", payload.DeviceID, payload.Info)

	jsonData, err := json.Marshal(payload.Info)
	if err != nil {
		log.Println("JSON encode error:", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to encode info",
		})
	}

	// 🔍 Log hasil jsonData yang akan dikirim ke DB
	log.Printf("Encoded JSON to DB: %s\n", string(jsonData))

	_, err = database.DB.Exec(
		context.Background(),
		`INSERT INTO device_information (device_id, data)
		 VALUES ($1, $2)`,
		payload.DeviceID, jsonData,
	)

	if err != nil {
		log.Println("Insert error:", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Database insert failed",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Information stored",
	})
}
