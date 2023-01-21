package main

import (
	"flag"
	"log"
	"os"

	"github.com/fes111/rmm/projects/client/config"
	"github.com/fes111/rmm/projects/client/socket"
	"github.com/fes111/rmm/projects/client/vars"
)

var interrupt = make(chan os.Signal, 1)
var isInterrupt = false

func main() {
	flag.Parse()
	log.SetFlags(0)

	config.ReadConfiguration()
	config.SetupDevice()
	if vars.Configuration.Token == "" {
		log.Printf("No token found, please create one first and add it to the config file\n")
		os.Exit(1)
	}
	if vars.Device.DeviceID == "" {
		vars.Device.DeviceID = "new"
	}
	socket.ConnectWebsocket()
}
