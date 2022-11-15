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
	"festech.de/rmm/client/vars"
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
	u := vars.WsUrl + fmt.Sprintf("%s?token=%s", vars.Device.DeviceID, "123")
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
				if vars.Configuration.AllowRun {
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
				c.WriteJSON(models.SocketEvent{
					Event: "result-shutdown",
					Data:  system.Shutdown(),
				})
			case "reboot":
				c.WriteJSON(models.SocketEvent{
					Event: "result-reboot",
					Data:  system.Reboot(),
				})
			case "process-list":
				c.WriteJSON(models.SocketEvent{
					Event: "result-process-list",
					Data:  system.GetProcessList(),
				})
			case "service-list":
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-list",
					Data:  system.GetServiceList(),
				})
			case "service-logs":
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-logs",
					Data:  system.GetServiceLogs(msg.Data.(string)),
				})
			case "service-start":
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-start",
					Data:  system.StartService(msg.Data.(string)),
				})
			case "service-stop":
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-stop",
					Data:  system.StartService(msg.Data.(string)),
				})
			case "service-restart":
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-restart",
					Data:  system.RestartService(msg.Data.(string)),
				})
			case "service-status":
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-status",
					Data:  system.GetServiceStatus(msg.Data.(string)),
				})
			case "process-kill":
				c.WriteJSON(models.SocketEvent{
					Event: "result-process-kill",
					Data:  system.KillProcess(msg.Data.(string)),
				})
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

func IsClosed(ch <-chan bool) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

func tryReconnect(url string) {
	fmt.Println("trying to reconnect to server")
	if isInterrupt {
		return
	}
	if system.EndUsageStream != nil && !IsClosed(system.EndUsageStream) {
		close(system.EndUsageStream)
	}
	time.Sleep(time.Second * 5)
	connectWebsocket(url)
}
