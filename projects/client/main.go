package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/fes111/rmm/libs/go/helpers"
	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/client/config"
	"github.com/fes111/rmm/projects/client/http"
	"github.com/fes111/rmm/projects/client/system"
	"github.com/fes111/rmm/projects/client/vars"
	"github.com/gorilla/websocket"
)

var interrupt = make(chan os.Signal, 1)
var isInterrupt = false

func main() {
	flag.Parse()
	log.SetFlags(0)

	signal.Notify(interrupt, os.Interrupt)

	config.ReadConfiguration()
	config.SetupDevice()
	if vars.Configuration.Token == "" {
		log.Printf("No token found, please create one first and add it to the config file\n")
		os.Exit(1)
	}
	u := vars.WsUrl + fmt.Sprintf("%s?token=%s", vars.Device.DeviceID, vars.Configuration.Token)
	connectWebsocket(u)
}

func connectWebsocket(url string) {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		tryReconnect(url)
	}
	log.Println("Connected to server")
	defer c.Close()

	http.SocketConn = c
	done := make(chan struct{})
	go system.SendUsage()
	go func() {
		time.Sleep(time.Minute * 15)
		c.Close()
	}()

	go func() {
		var msg models.SocketEvent
		defer close(done)
		go func() {
			for item := range vars.Queue {
				err := c.WriteJSON(item)
				if err != nil {
					log.Println(err)
					tryReconnect(url)
					return
				}
			}
		}()
		for {
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Println(err)
				tryReconnect(url)
				return
			}
			if strings.HasPrefix(msg.Event, "response-") {
				if handler, ok := vars.Handlers.Handlers[msg.Event]; ok {
					handler(msg)
					if _, ok := vars.Handlers.Onces[msg.Event]; ok {
						delete(vars.Handlers.Handlers, msg.Event)
						delete(vars.Handlers.Onces, msg.Event)
					}
				}
				continue
			}
			switch msg.Event {
			case "usage-start":
				system.StartStopUsageStream <- true
			case "usage-stop":
				system.StartStopUsageStream <- false
			case "run":
				if vars.Configuration.Allow.Run {
					data := system.Run(fmt.Sprintf("%v", msg.Data))
					c.WriteJSON(models.SocketEvent{
						Event: "result-run",
						Data:  data,
					})
				} else {
					c.WriteJSON(models.SocketEvent{
						Event: "result-run",
						Data:  "",
						Error: "Run is not allowed on this device",
					})
				}
			case "shutdown":
				data, err := system.Shutdown()
				c.WriteJSON(models.SocketEvent{
					Event: "result-shutdown",
					Data:  data,
					Error: err,
				})
			case "reboot":
				data, err := system.Reboot()
				c.WriteJSON(models.SocketEvent{
					Event: "result-reboot",
					Data:  data,
					Error: err,
				})
			case "process-list":
				data, err := system.GetProcessList()
				c.WriteJSON(models.SocketEvent{
					Event: "result-process-list",
					Data:  data,
					Error: err,
				})
			case "service-list":
				data, err := system.GetServiceList()
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-list",
					Data:  data,
					Error: err,
				})
			case "service-logs":
				data, err := system.GetServiceLogs(msg.Data.(string))
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-logs",
					Data:  data,
					Error: err,
				})
			case "service-start":
				data, err := system.StartService(msg.Data.(string))
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-start",
					Data:  data,
					Error: err,
				})
			case "service-stop":
				data, err := system.StopService(msg.Data.(string))
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-stop",
					Data:  data,
					Error: err,
				})
			case "service-restart":
				data, err := system.RestartService(msg.Data.(string))
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-restart",
					Data:  data,
					Error: err,
				})
			case "service-status":
				data, err := system.GetServiceStatus(msg.Data.(string))
				c.WriteJSON(models.SocketEvent{
					Event: "result-service-status",
					Data:  data,
					Error: err,
				})
			case "process-kill":
				data, err := system.KillProcess(msg.Data.(string))
				c.WriteJSON(models.SocketEvent{
					Event: "result-process-kill",
					Data:  data,
					Error: err,
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

func tryReconnect(url string) {
	log.Println("trying to reconnect to server")
	if isInterrupt {
		return
	}
	if system.EndUsageStream != nil && !helpers.IsClosed(system.EndUsageStream) {
		close(system.EndUsageStream)
	}
	time.Sleep(time.Second * 5)
	connectWebsocket(url)
}
