package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadPhoto(c *fiber.Ctx) error {
	deviceID := c.Query("device_id")
	if deviceID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing device_id",
		})
	}

	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing photo file",
		})
	}

	// Buat folder berdasarkan device_id
	deviceDir := filepath.Join("uploads", deviceID)
	err = os.MkdirAll(deviceDir, os.ModePerm)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create device folder",
		})
	}

	// Simpan file dengan nama unik
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("photo_%d_%s", timestamp, file.Filename)
	savePath := filepath.Join(deviceDir, filename)

	err = c.SaveFile(file, savePath)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to save file",
		})
	}

	fmt.Printf("[UploadPhoto] Saved: %s\n", savePath)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Photo uploaded",
		"path":    fmt.Sprintf("/uploads/%s/%s", deviceID, filename),
	})
}
