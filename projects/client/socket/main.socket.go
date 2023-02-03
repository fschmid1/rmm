package socket

import (
	"context"
	"fmt"
	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/client/system"
	"github.com/fes111/rmm/projects/client/vars"
	"github.com/recws-org/recws"
	"log"
	"strings"
	"time"
)

var SocketConn recws.RecConn

func SocketSend(event string, data interface{}) {
	SocketConn.WriteJSON(models.SocketEvent{
		Event: event,
		Data:  data,
	})
}

func ConnectWebsocket() {
	url := vars.WsUrl
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	firstConnection := true
	ws := recws.RecConn{
		KeepAliveTimeout: 0,
		RecIntvlMin:      time.Second * 5,
		RecIntvlFactor:   1,
	}

	ws.Dial(url, nil)
	go SendUsage()

	for {
		select {
		case <-ctx.Done():
			go ws.Close()
			return
		default:
			if !ws.IsConnected() && !firstConnection {
				StartStopUsageStream <- false
				fmt.Println("Disconnected from server")
				firstConnection = true
			}
			if ws.IsConnected() {
				if firstConnection {
					fmt.Println("Connected to server")
					SocketConn = ws
					ws.WriteJSON(models.SocketEvent{
						Event: "auth",
						Data:  map[string]string{"token": vars.Configuration.Token, "id": vars.Device.DeviceID},
					})
				}
				firstConnection = false
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
				case "auth-success":
					fmt.Println("Auth success")
					go func() {
						for item := range vars.Queue {
							err := ws.WriteJSON(item)
							if err != nil {
								log.Println(err)
							}
						}
					}()
				case "usage-start":
					StartStopUsageStream <- true
				case "usage-stop":
					StartStopUsageStream <- false
				case "run":
					if vars.Configuration.Allow.Run {
						data := system.Run(msg.Data.(string), time.Second*5)
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
