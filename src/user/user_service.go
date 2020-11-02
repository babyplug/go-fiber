package user

import (
	"github.com/babyplug/go-fiber/src/database"
	"github.com/babyplug/go-fiber/src/models"
)

var userRepo *UserRepositoryImpl

func SetupUserService() {
	userRepo = NewUserRepository(database.SqlDB)
}

func FindAllUser() ([]models.User, error) {
	return userRepo.FindAll()
}

func FindByID(userID int) (models.User, error) {
	return userRepo.FindByID(userID)
}

func FindByUsername(username string) (models.User, error) {
	return userRepo.FindByUsername(username)
}
