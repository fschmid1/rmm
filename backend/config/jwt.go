package config

import jwtware "github.com/gofiber/jwt/v3"

var JWT_SECRET string = Getenv("JWT_SECRET", "")

var JWT_CONFIG jwtware.Config = jwtware.Config{
	SigningKey: []byte(JWT_SECRET),
}
