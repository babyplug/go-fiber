package album

import (
	"github.com/babyplug/go-fiber/src/models"
	"gorm.io/gorm"
)

type AlbumRepository interface {
	FindAll() ([]models.Album, error)
}

type AlbumRepositoryImpl struct {
	*gorm.DB
}

func NewAlbumRepository(db *gorm.DB) *AlbumRepositoryImpl {
	return &AlbumRepositoryImpl{db}
}

func (album *AlbumRepositoryImpl) FindAll() ([]models.Album, error) {
	var albums []models.Album
	album.DB.Find(&albums)

	return albums, nil
}
