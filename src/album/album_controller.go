package album

import (
	"github.com/gofiber/fiber/v2"
)

func NewAlbumController(router fiber.Router) {
	SetupAlbumService()
	albumRouter := router.Group("/albums")

	// albumRouter.Use(middleware.TokenAuthMiddleware())

	albumRouter.Get("", handleGet)
}

func handleGet(c *fiber.Ctx) error {
	albums, err := FindAllAlbum()

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "Service error please try again",
			},
		)
	}
	return c.JSON(fiber.Map{
		"data": albums,
	})
}
