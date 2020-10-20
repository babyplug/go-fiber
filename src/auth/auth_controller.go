package auth

import (
	"github.com/babyplug/go-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

func NewAuthController(router fiber.Router) {
	SetupAuthService()

	router.Post("/login", handleLogin)
}

func handleLogin(c *fiber.Ctx) error {
	credentials := new(dto.Credentials)

	if err := c.BodyParser(credentials); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := Login(credentials)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	response := dto.CredentialsResponse{
		Prefix: "Bearer",
		Token:  token,
	}

	return c.JSON(response)
}
