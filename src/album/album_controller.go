package album

import (
	"errors"
	"strconv"

	"github.com/babyplug/go-fiber/src/dto"
	"github.com/gofiber/fiber/v2"
)

func NewAlbumController(router fiber.Router) {
	SetupAlbumService()
	albumRouter := router.Group("/albums")

	// albumRouter.Use(middleware.TokenAuthMiddleware())

	albumRouter.Get("", getAllAlbum)
	albumRouter.Post("", createAlbum)
	albumRouter.Get("/:id", getAlbumByID)
	albumRouter.Put("/:id", updateAlbumByID)
	albumRouter.Delete("/:id", deleteAlbumByID)
}

func getAllAlbum(c *fiber.Ctx) error {
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

func createAlbum(c *fiber.Ctx) error {
	form := new(dto.AlbumRequestForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	album, err := CreateAlbum(form)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(album)
}

func convertParamsID(paramID string) (albumID int, err error) {
	if albumID, err = strconv.Atoi(paramID); err != nil {
		err = errors.New("Album id should be a number")
	}
	return
}

func getAlbumByID(c *fiber.Ctx) error {
	albumID, err := convertParamsID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	album, err := FindAlbumByID(albumID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(album)
}

func updateAlbumByID(c *fiber.Ctx) error {
	albumID, err := convertParamsID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	form := new(dto.AlbumRequestForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	album, err := UpdateAlbumByID(albumID, form)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(album)
}

func deleteAlbumByID(c *fiber.Ctx) error {
	albumID, err := convertParamsID(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err = DeleteAlbumByID(albumID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success"})
}
