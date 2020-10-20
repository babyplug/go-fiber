package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

func TokenAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := tokenValid(c.Request())
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(err.Error())
		}
		return c.Next()
	}
}

func tokenValid(r *fasthttp.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func verifyToken(r *fasthttp.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("app.secret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func extractToken(r *fasthttp.Request) string {
	bearToken := string(r.Header.Peek("Authorization"))
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
