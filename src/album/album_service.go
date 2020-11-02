package album

import (
	"github.com/babyplug/go-fiber/src/database"
	"github.com/babyplug/go-fiber/src/models"
)

var albumRepository *AlbumRepositoryImpl

func SetupAlbumService() {
	albumRepository = NewAlbumRepository(database.SqlDB)
}

func FindAllAlbum() ([]models.Album, error) {
	return albumRepository.FindAll()
}
