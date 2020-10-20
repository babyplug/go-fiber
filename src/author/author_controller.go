package author

import (
	"github.com/gofiber/fiber/v2"
)

func NewAuthorController(router fiber.Router) {
	SetupAuthorService()
	authorRouter := router.Group("/authors")

	authorRouter.Get("", handleGet)
}

func handleGet(c *fiber.Ctx) error {
	authors, err := FindAllAuthor()

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "Service error please try again",
			},
		)
	}
	return c.JSON(fiber.Map{
		"data": authors,
	})
}
