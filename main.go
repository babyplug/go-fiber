package main

import (
	"fmt"

	"github.com/babyplug/go-fiber/src/auth"
	"github.com/babyplug/go-fiber/src/author"
	"github.com/babyplug/go-fiber/src/database"
	"github.com/babyplug/go-fiber/src/photo"
	"github.com/babyplug/go-fiber/src/user"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (
	app = fiber.New()
)

func main() {
	setupViper()
	database.NewMySQLConnect()

	api := app.Group("/api")

	author.NewAuthorController(api)
	user.NewUserController(api)
	auth.NewAuthController(api)
	photo.NewPhotoController(api)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Error 404: Not found",
		})
	})

	app.Listen(":3000")
}

func setupViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
