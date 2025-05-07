package handlers

import (
	"context"
	"encoding/json"
	"family-control-backend/database"
	"family-control-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateDeviceCommand(c *fiber.Ctx) error {
	var cmd models.SubmitCommand
	if err := json.Unmarshal(c.Body(), &cmd); err != nil {
		log.Println("[CreateCommand] JSON parse error:", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid JSON",
		})
	}

	_, err := database.DB.Exec(
		context.Background(),
		`INSERT INTO device_commands (device_id, command, params)
		 VALUES ($1, $2, $3)`,
		cmd.DeviceID, cmd.Command, cmd.Params,
	)

	if err != nil {
		log.Println("[CreateCommand] DB error:", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create command",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Command created",
	})
}
