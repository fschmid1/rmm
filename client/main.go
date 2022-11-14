package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"festech.de/rmm/client/config"
	"festech.de/rmm/client/http"
	"festech.de/rmm/client/models"
	"festech.de/rmm/client/system"
	"github.com/gorilla/websocket"
)

var interrupt = make(chan os.Signal, 1)
var isInterrupt = false

func main() {
	flag.Parse()
	log.SetFlags(0)

	signal.Notify(interrupt, os.Interrupt)

	config.ReadConfiguration()
	system.GetMacAddress()
	u := config.WsUrl + fmt.Sprintf("%s?token=%s", config.Device.DeviceID, "123")
	connectWebsocket(u)
}

func connectWebsocket(url string) {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		tryReconnect(url)
	}
	fmt.Println("Connected to server")
	defer c.Close()

	http.SocketConn = c
	done := make(chan struct{})
	go system.SendUsage()

	go func() {
		var msg models.SocketEvent
		defer close(done)
		for {
			err := c.ReadJSON(&msg)
			if err != nil {
				fmt.Println(err)
				tryReconnect(url)
				return
			}
			switch msg.Event {
			case "usage-start":
				system.StartStopUsageStream <- true
			case "usage-stop":
				system.StartStopUsageStream <- false
			case "run":
				if config.Configuration.AllowRun {
					c.WriteJSON(models.SocketEvent{
						Event: "result-run",
						Data:  system.Run(fmt.Sprintf("%v", msg.Data)),
					})
				} else {
					c.WriteJSON(models.SocketEvent{
						Event: "result-process-list",
						Data:  "Run is not allowed on this device",
					})
				}
			case "shutdown":
				if config.Configuration.AllowShutdown {
					system.Run("ls")
				}
			case "reboot":
				if config.Configuration.AllowReboot {
					system.Run("ls")
				}
			case "process-list":
				if config.Configuration.AllowProcessList {
					c.WriteJSON(models.SocketEvent{
						Event: "result-process-list",
						Data:  system.GetProcessList(),
					})
				} else {
					c.WriteJSON(models.SocketEvent{
						Event: "result-process-list",
						Data:  "Process list is not allowed on this device",
					})
				}
			case "process-kill":
				if config.Configuration.AllowKill {
					c.WriteJSON(models.SocketEvent{
						Event: "result-process-kill",
						Data:  system.KillProcess(msg.Data.(string)),
					})
				} else {
					c.WriteJSON(models.SocketEvent{
						Event: "result-process-kill",
						Data:  "Process kill is not allowed on this device",
					})
				}
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			isInterrupt = true
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				os.Exit(0)
			}
		}
	}
}

func tryReconnect(url string) {
	fmt.Println("trying to reconnect to server")
	if isInterrupt {
		return
	}
	close(system.EndUsageStream)
	time.Sleep(time.Second * 5)
	connectWebsocket(url)
}
