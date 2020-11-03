package album

import (
	"github.com/babyplug/go-fiber/src/database"
	"github.com/babyplug/go-fiber/src/dto"
	"github.com/babyplug/go-fiber/src/models"
)

var albumRepository *AlbumRepositoryImpl

func SetupAlbumService() {
	albumRepository = NewAlbumRepository(database.SqlDB)
}

func FindAllAlbum() ([]models.Album, error) {
	return albumRepository.FindAll()
}

func CreateAlbum(form *dto.AlbumRequestForm) (album models.Album, err error) {
	album.Name = form.Name
	err = albumRepository.Save(&album)
	if err != nil {
		return album, err
	}

	return album, nil
}

func FindAlbumByID(albumID int) (models.Album, error) {
	return albumRepository.FindAlbumByID(albumID)
}

func UpdateAlbumByID(albumID int, form *dto.AlbumRequestForm) (album models.Album, err error) {
	album, err = FindAlbumByID(albumID)
	if err != nil {
		return
	}
	album.Name = form.Name
	err = albumRepository.Save(&album)
	return
}

func DeleteAlbumByID(albumID int) (err error) {
	album, err := FindAlbumByID(albumID)
	if err != nil {
		return
	}
	err = albumRepository.Delete(&album)
	return
}
