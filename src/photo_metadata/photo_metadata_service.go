package photo_metadata

import (
	"github.com/babyplug/go-fiber/src/database"
	"github.com/babyplug/go-fiber/src/models"
)

var photoMetadataRepo *PhotoMetadataRepositoryImpl

func SetupPhotoMetadataService() {
	photoMetadataRepo = NewPhotoMetadataRepository(database.SqlDB)
}

func FindAllPhotoMetadata() ([]models.PhotoMetadata, error) {
	return photoMetadataRepo.FindAll()
}
