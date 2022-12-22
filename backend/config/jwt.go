package config

import (
	"fmt"

	"festech.de/rmm/backend/models"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET string = Getenv("JWT_SECRET", "")
var SOCKET_JWT_SECRET string = Getenv("SOCKET_JWT_SECRET", "")

var JWT_CONFIG jwtware.Config = jwtware.Config{
	SigningKey: []byte(JWT_SECRET),
}

func GenerateDeviceJWT(device models.DeviceToken) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userID"] = device.UserID
	claims["name"] = device.Name
	tokenString, err := token.SignedString([]byte(SOCKET_JWT_SECRET))
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}

func VerifyUserJWT(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		fmt.Println(err)
		return false
	}
	if token.Valid {
		return true
	} else {
		return false
	}
}

func VerifyClientJWT(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SOCKET_JWT_SECRET), nil
	})
	if err != nil {
		fmt.Println(err)
		return false
	}
	if token.Valid {
		return true
	} else {
		return false
	}
}