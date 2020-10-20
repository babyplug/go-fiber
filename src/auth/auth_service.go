package auth

import (
	"errors"
	"time"

	"github.com/babyplug/go-fiber/src/dto"
	"github.com/babyplug/go-fiber/src/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func SetupAuthService() {
	user.SetupUserService()
}

func createToken(userid uint) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(viper.GetString("app.secret")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Login(credentials *dto.Credentials) (string, error) {
	user, err := user.FindByUsername(credentials.Username)
	if err != nil {
		return "", errors.New("Username or password is not correct")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return "", errors.New("Username or password is not correct")
	}

	token, err := createToken(user.ID)
	if err != nil {
		return "", errors.New("Server error please contact admin")
	}

	return token, nil
}
