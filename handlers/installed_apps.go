package handlers

import (
	"context"
	"family-control-backend/database"
	"family-control-backend/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func StoreInstalledApps(c *fiber.Ctx) error {
	log.Println("[StoreInstalledApps] Hit")

	deviceID := c.Query("device_id")
	if deviceID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing device_id",
		})
	}

	var apps []models.InstalledApp

	raw := c.Body()
	log.Println("Raw body:", string(raw))

	if err := c.BodyParser(&apps); err != nil {
		log.Println("Failed to parse body:", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	if len(apps) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Empty app list",
		})
	}

	for _, app := range apps {
		_, err := database.DB.Exec(
			context.Background(),
			`INSERT INTO installed_apps (device_id, package_name, app_name) VALUES ($1, $2, $3)`,
			deviceID, app.PackageName, app.AppName,
		)
		if err != nil {
			log.Println("Insert failed for", app.PackageName, ":", err)
		}
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Apps stored",
	})
}
