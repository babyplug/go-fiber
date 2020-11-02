package author

import (
	"github.com/babyplug/go-fiber/src/database"
	"github.com/babyplug/go-fiber/src/models"
)

var authorRepo *AuthorRepositoryImpl

func SetupAuthorService() {
	authorRepo = NewAuthorRepository(database.SqlDB)
}

func FindAllAuthor() ([]models.Author, error) {
	return authorRepo.FindAll()
}
