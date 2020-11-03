package album

import (
	"errors"

	"github.com/babyplug/go-fiber/src/models"
	"gorm.io/gorm"
)

type AlbumRepository interface {
	FindAll() ([]models.Album, error)
}

type AlbumRepositoryImpl struct {
	DB *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) *AlbumRepositoryImpl {
	return &AlbumRepositoryImpl{db}
}

func (repo *AlbumRepositoryImpl) FindAll() ([]models.Album, error) {
	var albums []models.Album
	repo.DB.Find(&albums)

	return albums, nil
}

func (repo *AlbumRepositoryImpl) Save(dto *models.Album) (err error) {
	if dto.ID == 0 {
		err = repo.DB.Create(dto).Error
	} else {
		err = repo.DB.Save(dto).Error
	}
	return
}

func (repo *AlbumRepositoryImpl) FindAlbumByID(albumID int) (albumDTO models.Album, err error) {
	err = repo.DB.Where("id = ?", albumID).First(&albumDTO).Error
	return
}

func (repo *AlbumRepositoryImpl) Delete(dto *models.Album) (err error) {
	if dto == nil {
		err = errors.New("Dto should not empty")
	}
	err = repo.DB.Delete(dto).Error
	return
}
