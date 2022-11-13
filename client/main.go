package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"festech.de/rmm/client/config"
	"festech.de/rmm/client/http"
	"festech.de/rmm/client/models"
	"festech.de/rmm/client/system"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var interrupt = make(chan os.Signal, 1)
var isInterrupt = false

func main() {
	flag.Parse()
	log.SetFlags(0)

	signal.Notify(interrupt, os.Interrupt)

	config.ReadConfiguration()
	system.GetMacAddress()
	u := url.URL{Scheme: "ws", Host: *addr, Path: fmt.Sprintf("/ws/%s", config.Device.DeviceID), RawQuery: "token=123"}
	connectWebsocket(u.String())
}

func connectWebsocket(url string) {
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		tryReconnect(url)
	}
	fmt.Println("Connected to server")
	defer c.Close()
	
	http.SocketConn = c;
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
				case "start-usage":
					system.StartStopUsageStream <- true
				case "stop-usage":
					system.StartStopUsageStream <- false
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
