package user

import (
	"github.com/babyplug/go-fiber/src/database"
)

var userRepo *UserRepositoryImpl

func SetupUserService() {
	userRepo = NewUserRepository(database.SqlDB)
}

func FindAllUser() ([]User, error) {
	return userRepo.FindAll()
}

func FindByID(userID int) (User, error) {
	return userRepo.FindByID(userID)
}

func FindByUsername(username string) (User, error) {
	return userRepo.FindByUsername(username)
}
