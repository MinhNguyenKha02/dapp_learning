package handlers

import (
	"dapp_learning/config"
	"dapp_learning/internal/services"

	"github.com/gofiber/fiber/v2"
)

func UploadFiles(c *fiber.Ctx) error {
	// Get uploaded files
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No files uploaded",
		})
	}
	files := form.File["files"]

	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No files selected",
		})
	}

	config, err := config.LoadGGDriveConfig()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load Google Drive config: " + err.Error(),
		})
	}

	ggDriveService, err := services.NewGGDriveService(config)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to initialize Google Drive service: " + err.Error(),
		})
	}

	// Upload all files and collect URLs
	urls := make([]string, 0, len(files))
	errors := make([]string, 0)

	for _, file := range files {
		url, err := ggDriveService.UploadFile(file)
		if err != nil {
			errors = append(errors, file.Filename+": "+err.Error())
			continue
		}
		urls = append(urls, url)
	}

	response := fiber.Map{
		"uploaded_files": len(urls),
		"urls":           urls,
	}

	if len(errors) > 0 {
		response["errors"] = errors
	}

	return c.JSON(response)
}
