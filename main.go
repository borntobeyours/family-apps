package main

import (
	"family-control-backend/database"
	"family-control-backend/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	app := fiber.New()

	app.Post("/api/device/register", handlers.RegisterDevice)
	app.Post("/api/device/usage", handlers.SubmitAppUsage)
	app.Post("/api/device/location", handlers.SubmitDeviceLocation)
	app.Post("/api/device/command", handlers.CreateDeviceCommand)
	app.Get("/api/device/command", handlers.PollDeviceCommand)
	app.Post("/api/device/upload_photo", handlers.UploadPhoto)
	app.Post("/api/device/installed_apps", handlers.StoreInstalledApps)
	app.Post("/api/device/upload_gallery_image", handlers.UploadGalleryImage)
	app.Post("/api/device/upload_sms", handlers.UploadSms)
	app.Post("/api/device/upload_contacts", handlers.UploadContacts)
	app.Post("/api/device/information", handlers.UploadDeviceInformation)

	log.Fatal(app.Listen("0.0.0.0:8080"))
}
