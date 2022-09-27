package controllers

import (
    "net/http"
    "time"
    "github.com/labstack/echo/v4"
	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("ghn-secret-auth-key")
var ACCESS_TOKEN = "access_token"

func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func VerifyToken(c echo.Context) bool {
	tokenString := c.Request().Header.Get(ACCESS_TOKEN)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})
	if err != nil {
		return false
	}
	if token.Valid {
		return true
	}
	return false
}

func GetJWT(c echo.Context) error {
	token, err := CreateJWT()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Creating JWT failed.")
	}
	return c.JSON(http.StatusOK, &echo.Map{ACCESS_TOKEN: token})
}
