package handlers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func UploadGalleryImage(c *fiber.Ctx) error {
	deviceID := c.FormValue("device_id")
	if deviceID == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing device_id",
		})
	}

	fileHeader, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing image file",
		})
	}

	// Ensure upload directory exists
	dir := fmt.Sprintf("uploads/%s/gallery", deviceID)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to create upload directory",
		})
	}

	// Save file
	dstPath := filepath.Join(dir, fileHeader.Filename)
	src, err := fileHeader.Open()
	if err != nil {
		return c.Status(500).SendString("Failed to open file")
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return c.Status(500).SendString("Failed to create destination file")
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return c.Status(500).SendString("Failed to save file")
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Image uploaded",
	})
}
