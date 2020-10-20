package user

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Email     *string
	Age       uint8
	Username  string
	Password  string `json:"-"`
	FirstName string
	LastName  string
	gorm.Model
}

type UserRepository interface {
	FindAll() ([]User, error)

	FindByID() (User, error)

	UpdateByID(userID int) (User, error)

	DeleteByID(userID int) (User, error)

	FindByUsername(username string) (User, error)
}

type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (userRepo *UserRepositoryImpl) FindAll() ([]User, error) {
	var users []User
	userRepo.DB.Find(&users)

	return users, nil
}

func (userRepo *UserRepositoryImpl) FindByID(userID int) (User, error) {
	var user User
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

func (userRepo *UserRepositoryImpl) FindByUsername(username string) (User, error) {
	var user User
	if err := userRepo.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return user, errors.New("User not found!")
	}

	return user, nil
}
