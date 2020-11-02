package photo

import (
	"github.com/babyplug/go-fiber/src/database"
	"github.com/babyplug/go-fiber/src/models"
)

var photoRepo *PhotoRepositoryImpl

func SetupPhotoService() {
	photoRepo = NewPhotoRepository(database.SqlDB)
}

func FindAllPhoto() ([]models.Photo, error) {
	return photoRepo.FindAll()
}
