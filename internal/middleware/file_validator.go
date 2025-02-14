package middleware

import (
	"path/filepath"
	"strings"

	"dapp_learning/internal/utils"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
)

func ValidateFileUpload() fiber.Handler {
	return func(c *fiber.Ctx) error {
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

		// Validate each file
		for _, file := range files {
			// Check file extension
			ext := strings.ToLower(filepath.Ext(file.Filename))
			if !utils.AllowedExtensions[ext] {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "File '" + file.Filename + "' not allowed. Supported formats: Video (MP4, AVI, MPEG, WebM...), Audio (MP3, WAV, OGG, M4A...), Documents (PDF, Word, Excel, PowerPoint, OpenDocument...)",
				})
			}

			// Open and check MIME type
			f, err := file.Open()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to process file: " + file.Filename,
				})
			}
			defer f.Close()

			mtype, err := mimetype.DetectReader(f)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to detect file type: " + file.Filename,
				})
			}

			if !utils.AllowedMimeTypes[mtype.String()] {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid file type for " + file.Filename + ": " + mtype.String(),
				})
			}

			// Reset file pointer
			if seeker, ok := f.(interface {
				Seek(int64, int) (int64, error)
			}); ok {
				_, err = seeker.Seek(0, 0)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"error": "Failed to process file: " + file.Filename,
					})
				}
			}
		}

		return c.Next()
	}
}
