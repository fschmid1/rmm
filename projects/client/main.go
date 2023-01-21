package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/client/config"
	"github.com/fes111/rmm/projects/client/http"
	"github.com/fes111/rmm/projects/client/system"
	"github.com/fes111/rmm/projects/client/vars"
	"github.com/recws-org/recws"
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
	u := vars.WsUrl + fmt.Sprintf("%s?token=%s", vars.Device.DeviceID, vars.Configuration.Token)
	connectWebsocket(u)
}

func connectWebsocket(url string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	firstConnecton := false
	ws := recws.RecConn{
		KeepAliveTimeout: 0,
		RecIntvlMin:      time.Second * 5,
		RecIntvlFactor:   1,
	}

	ws.Dial(url, nil)
	http.SocketConn = ws

	go system.SendUsage()

	for {
		select {
		case <-ctx.Done():
			go ws.Close()
			return
		default:
			if ws.IsConnected() {
				if !firstConnecton {
					fmt.Println("Connected to server")
					go func() {
						for item := range vars.Queue {
							err := ws.WriteJSON(item)
							if err != nil {
								log.Println(err)
							}
							return
						}
					}()
				}
				firstConnecton = true
			}
			if !ws.IsConnected() && firstConnecton {
				fmt.Println("Disconnected from server")
				firstConnecton = false
			}
			if ws.IsConnected() {
				var msg models.SocketEvent
				err := ws.ReadJSON(&msg)
				if err != nil {
					log.Println(err)
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
						fmt.Println(data)
						ws.WriteJSON(models.SocketEvent{
							Event: "result-run",
							Data:  data,
						})
					} else {
						ws.WriteJSON(models.SocketEvent{
							Event: "result-run",
							Data:  "",
							Error: "Run is not allowed on this device",
						})
					}
				case "shutdown":
					data, err := system.Shutdown()
					ws.WriteJSON(models.SocketEvent{
						Event: "result-shutdown",
						Data:  data,
						Error: err,
					})
				case "reboot":
					data, err := system.Reboot()
					ws.WriteJSON(models.SocketEvent{
						Event: "result-reboot",
						Data:  data,
						Error: err,
					})
				case "process-list":
					data, err := system.GetProcessList()
					ws.WriteJSON(models.SocketEvent{
						Event: "result-process-list",
						Data:  data,
						Error: err,
					})
				case "service-list":
					data, err := system.GetServiceList()
					ws.WriteJSON(models.SocketEvent{
						Event: "result-service-list",
						Data:  data,
						Error: err,
					})
				case "service-logs":
					data, err := system.GetServiceLogs(msg.Data.(string))
					ws.WriteJSON(models.SocketEvent{
						Event: "result-service-logs",
						Data:  data,
						Error: err,
					})
				case "service-start":
					data, err := system.StartService(msg.Data.(string))
					ws.WriteJSON(models.SocketEvent{
						Event: "result-service-start",
						Data:  data,
						Error: err,
					})
				case "service-stop":
					data, err := system.StopService(msg.Data.(string))
					ws.WriteJSON(models.SocketEvent{
						Event: "result-service-stop",
						Data:  data,
						Error: err,
					})
				case "service-restart":
					data, err := system.RestartService(msg.Data.(string))
					ws.WriteJSON(models.SocketEvent{
						Event: "result-service-restart",
						Data:  data,
						Error: err,
					})
				case "service-status":
					data, err := system.GetServiceStatus(msg.Data.(string))
					ws.WriteJSON(models.SocketEvent{
						Event: "result-service-status",
						Data:  data,
						Error: err,
					})
				case "process-kill":
					data, err := system.KillProcess(msg.Data.(string))
					ws.WriteJSON(models.SocketEvent{
						Event: "result-process-kill",
						Data:  data,
						Error: err,
					})
				}
			}
		}

	}
}
