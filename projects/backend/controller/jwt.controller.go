package controller

import (
	"log"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/config"
	"github.com/golang-jwt/jwt"
)

func GenerateDeviceJWT(device models.DeviceToken) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userID"] = device.UserID
	claims["name"] = device.Name
	tokenString, err := token.SignedString([]byte(config.SOCKET_JWT_SECRET))
	if err != nil {
		return "Signing Error", err
	}

	return tokenString, nil
}

func VerifyUserJWT(tokenString string) (bool, models.User) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		log.Println(err)
		return false, models.User{}
	}
	if token.Valid {
		userMap := token.Claims.(jwt.MapClaims)["user"].(map[string]interface{})
		return true, models.User{
			ID:        uint(userMap["id"].(float64)),
			Name:      userMap["name"].(string),
			Email:     userMap["email"].(string),
			PushToken: userMap["pushToken"].(string),
		}

	} else {
		return false, models.User{}
	}
}

func VerifyClientJWT(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SOCKET_JWT_SECRET), nil
	})
	if err != nil {
		log.Println(err)
		return false
	}
	if token.Valid {
		_, err := GetDeviceToken(tokenString)
		return err == nil
	} else {
		return false
	}
}
