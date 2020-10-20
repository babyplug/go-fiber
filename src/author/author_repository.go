package author

import (
	"gorm.io/gorm"
)

type Author struct {
	ID   uint
	Name string
	gorm.Model
}

type AuthorRepository interface {
	FindAll() ([]Author, error)
}

type AuthorRepositoryImpl struct {
	*gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepositoryImpl {
	return &AuthorRepositoryImpl{db}
}

func (authorRepo *AuthorRepositoryImpl) FindAll() ([]Author, error) {
	var authors []Author
	authorRepo.DB.Find(&authors)

	return authors, nil
}
