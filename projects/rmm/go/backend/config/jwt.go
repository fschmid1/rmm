package config

import (
	"github.com/fes111/rmm/libs/go/helpers"
	jwtware "github.com/gofiber/jwt/v3"
)

var JWT_SECRET string = helpers.Getenv("JWT_SECRET", "")
var SOCKET_JWT_SECRET string = helpers.Getenv("SOCKET_JWT_SECRET", "")

var JWT_CONFIG jwtware.Config = jwtware.Config{
	SigningKey: []byte(JWT_SECRET),
}
