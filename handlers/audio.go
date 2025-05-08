package handlers

import (
	"context"
	"family-control-backend/database"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadAudio(c *fiber.Ctx) error {
	log.Println("[UploadAudio] Hit")

	deviceID := c.FormValue("device_id")
	if deviceID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing device_id",
		})
	}

	file, err := c.FormFile("audio")
	if err != nil {
		log.Println("File receive error:", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Audio file not received",
		})
	}

	// Create directory: /uploads/device_id/
	dir := fmt.Sprintf("uploads/%s", deviceID)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Println("Failed to create dir:", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to prepare storage"})
	}

	// Save the file
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("audio_%s.3gp", timestamp)
	filePath := filepath.Join(dir, filename)

	src, err := file.Open()
	if err != nil {
		log.Println("Open file error:", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to open file"})
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		log.Println("Create file error:", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to save file"})
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		log.Println("Copy error:", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to write file"})
	}

	// Optionally, store metadata
	_, _ = database.DB.Exec(
		context.Background(),
		`INSERT INTO device_audio (device_id, file_path, created_at) VALUES ($1, $2, NOW())`,
		deviceID, filePath,
	)

	log.Printf("Audio saved: %s\n", filePath)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Audio uploaded",
		"path":    filePath,
	})
}
