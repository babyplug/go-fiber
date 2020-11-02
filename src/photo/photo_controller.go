package photo

import (
	"github.com/gofiber/fiber/v2"
)

func NewPhotoController(router fiber.Router) {
	SetupPhotoService()
	photoRouter := router.Group("/photos")

	// photoRouter.Use(middleware.TokenAuthMiddleware())

	photoRouter.Get("", handleGet)
}

func handleGet(c *fiber.Ctx) error {
	photos, err := FindAllPhoto()

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "Service error please try again",
			},
		)
	}
	return c.JSON(fiber.Map{
		"data": photos,
	})
}
