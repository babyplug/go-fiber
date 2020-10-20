package main

import (
	"github.com/babyplug/go-fiber/src/author"
	"github.com/babyplug/go-fiber/src/database"
	"github.com/gofiber/fiber/v2"
)

var (
	app = fiber.New()
)

func main() {
	database.NewMySQLConnect()

	api := app.Group("/api")

	author.NewAuthorController(api)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Error 404: Not found",
		})
	})

	app.Listen(":3000")
}
