package photo

import (
	"github.com/babyplug/go-fiber/src/models"
	"gorm.io/gorm"
)

type PhotoRepository interface {
	FindAll() ([]models.Photo, error)
}

type PhotoRepositoryImpl struct {
	*gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepositoryImpl {
	return &PhotoRepositoryImpl{db}
}

func (authorRepo *PhotoRepositoryImpl) FindAll() ([]models.Photo, error) {
	var photos []models.Photo
	authorRepo.DB.Find(&photos)

	return photos, nil
}
