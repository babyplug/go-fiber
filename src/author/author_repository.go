package author

import (
	"github.com/babyplug/go-fiber/src/models"
	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll() ([]models.Author, error)
}

type AuthorRepositoryImpl struct {
	*gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepositoryImpl {
	return &AuthorRepositoryImpl{db}
}

func (authorRepo *AuthorRepositoryImpl) FindAll() ([]models.Author, error) {
	var authors []models.Author
	authorRepo.DB.Find(&authors)

	return authors, nil
}
