package photo_metadata

import (
	"github.com/babyplug/go-fiber/src/models"
	"gorm.io/gorm"
)

type PhotoMetadataRepository interface {
	FindAll() ([]models.PhotoMetadata, error)
}

type PhotoMetadataRepositoryImpl struct {
	*gorm.DB
}

func NewPhotoMetadataRepository(db *gorm.DB) *PhotoMetadataRepositoryImpl {
	return &PhotoMetadataRepositoryImpl{db}
}

func (authorRepo *PhotoMetadataRepositoryImpl) FindAll() ([]models.PhotoMetadata, error) {
	var photos []models.PhotoMetadata
	authorRepo.DB.Find(&photos)

	return photos, nil
}
