package photo_metadata

import (
	"github.com/gofiber/fiber/v2"
)

func NewPhotoMetadataController(router fiber.Router) {
	SetupPhotoMetadataService()
	photoRouter := router.Group("/photo-metadata")

	// photoRouter.Use(middleware.TokenAuthMiddleware())

	photoRouter.Get("", handleGet)
}

func handleGet(c *fiber.Ctx) error {
	photoMetadata, err := FindAllPhotoMetadata()

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "Service error please try again",
			},
		)
	}
	return c.JSON(fiber.Map{
		"data": photoMetadata,
	})
}
