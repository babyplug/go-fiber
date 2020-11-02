package user

import (
	"errors"

	"github.com/babyplug/go-fiber/src/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)

	FindByID() (models.User, error)

	UpdateByID(userID int) (models.User, error)

	DeleteByID(userID int) (models.User, error)

	FindByUsername(username string) (models.User, error)
}

type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (userRepo *UserRepositoryImpl) FindAll() ([]models.User, error) {
	var users []models.User
	userRepo.DB.Find(&users)

	return users, nil
}

func (userRepo *UserRepositoryImpl) FindByID(userID int) (models.User, error) {
	var user models.User
	if err := userRepo.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return user, errors.New("User not found!")
	}

	return user, nil
}

func (userRepo *UserRepositoryImpl) DeleteByID(userID int) error {
	user, err := userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	userRepo.DB.Delete(user)
	return nil
}

func (userRepo *UserRepositoryImpl) FindByUsername(username string) (models.User, error) {
	var user models.User
	if err := userRepo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, errors.New("User not found!")
	}

	return user, nil
}
