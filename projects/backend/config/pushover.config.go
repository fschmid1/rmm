package config

import (
	"github.com/fes111/rmm/libs/go/helpers"
	"github.com/gregdel/pushover"
)

var apiToken string = helpers.Getenv("PUSHOVER_APITOKEN", "")
var Pusher *pushover.Pushover = pushover.New(apiToken)
