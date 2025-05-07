package handlers

import (
	"context"
	"family-control-backend/database"
	"family-control-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func PollDeviceCommand(c *fiber.Ctx) error {
	log.Println("[PollDeviceCommand] Hit")
	deviceID := c.Query("device_id")
	if deviceID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing device_id",
		})
	}

	row := database.DB.QueryRow(
		context.Background(),
		`SELECT id, command, params FROM device_commands
		 WHERE device_id = $1 AND executed = FALSE
		 ORDER BY created_at ASC LIMIT 1`,
		deviceID,
	)

	var (
		id      int
		command string
		params  []byte
	)

	err := row.Scan(&id, &command, &params)
	if err != nil {
		return c.Status(204).Send([]byte{}) // No content
	}

	// Tandai sebagai sudah dieksekusi
	_, _ = database.DB.Exec(
		context.Background(),
		`UPDATE device_commands SET executed = TRUE WHERE id = $1`,
		id,
	)

	return c.JSON(models.DeviceCommand{
		Command: command,
		Params:  params,
	})
}
