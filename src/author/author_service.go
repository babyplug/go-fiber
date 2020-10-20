package author

import (
	"github.com/babyplug/go-fiber/src/database"
)

var authorRepo *AuthorRepositoryImpl

func SetupAuthorService() {
	authorRepo = NewAuthorRepository(database.SqlDB)
}

func FindAllAuthor() ([]Author, error) {
	return authorRepo.FindAll()
}
