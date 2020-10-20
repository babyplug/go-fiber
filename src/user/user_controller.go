package user

import (
	"strconv"

	"github.com/babyplug/go-fiber/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewUserController(router fiber.Router) {
	SetupUserService()
	userRouter := router.Group("/users")

	userRouter.Use(middleware.TokenAuthMiddleware())

	userRouter.Get("", getAllUser)
	userRouter.Get("/:id", getUserByID)
}

func getAllUser(c *fiber.Ctx) error {
	users, err := FindAllUser()

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "Service error please try again",
			},
		)
	}
	return c.JSON(fiber.Map{
		"data": users,
	})
}

func getUserByID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "User id should be a number",
			},
		)
	}

	user, err := FindByID(userID)
	if err != nil {
		return c.Status(404).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(user)
}
