package socket

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"festech.de/rmm/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Id         string
	Connection *websocket.Conn
	User       bool
}

var Clients = make(map[string]Client)
var register = make(chan Client)
var Broadcast = make(chan models.SocketEvent)
var unregister = make(chan Client)

var Results = make(map[string]chan interface{})

var UsageStreams = make(map[string]map[string]Client)

func runHub() {
	for {
		select {
		case client := <-register:
			if !client.User {
				SetDeviceConnected(client.Id, true)
			}
			Clients[client.Id] = client

		case message := <-Broadcast:
			for clientID := range Clients {
				client := Clients[clientID]
				if err := client.Connection.WriteJSON(message); err != nil {
					log.Println("write error:", err)

					client.Connection.WriteMessage(websocket.CloseMessage, []byte{})
					client.Connection.Close()
					unregister <- client
				}
			}

		case client := <-unregister:
			if !client.User {
				SetDeviceConnected(client.Id, false)
			}
			delete(Clients, client.Id)
			for key, stream := range UsageStreams {
				delete(stream, client.Id)
				if len(stream) == 0 {
					SendMessage(key, models.SocketEvent{
						Event: "usage-stop",
					})
				}
			}
		}
	}
}

func CreateResultChannel(event string, id string) chan interface{} {
	resultChannel := make(chan interface{})
	Results["result-"+event+id] = resultChannel
	return resultChannel
}

func SendMessage(id string, message interface{}) error {
	client, found := Clients[id]

	if !found {
		return errors.New("Client not found")
	}
	err := client.Connection.WriteJSON(message)

	return err
}

func RegisterWebsocketRoute(app *fiber.App) {
	go runHub()

	route := app.Group("/ws")

	route.Use(func(c *fiber.Ctx) error {
		//TODO Authentication
		if websocket.IsWebSocketUpgrade(c) && c.Query("token") != "" {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	route.Get("/user/:id", websocket.New(func(c *websocket.Conn) {
		client := Client{
			Id:         c.Params("id"),
			Connection: c,
			User:       true,
		}
		defer func() {
			unregister <- client
			c.Close()
		}()

		register <- client

		for {
			message := models.SocketEvent{}
			err := client.Connection.ReadJSON(&message)
			if err != nil {
				return
			}
			if strings.HasPrefix(message.Event, "result-") {
				if channel, ok := Results[message.Event+client.Id]; ok {
					channel <- message.Data
					close(channel)
					delete(Results, message.Event+client.Id)
				}
			} else if message.Event == "usage" {
				fmt.Println(message)
				if _, ok := UsageStreams[client.Id]; ok {
					for _, client := range UsageStreams[client.Id] {
						client.Connection.WriteJSON(message)
					}
				}
			}
		}
	}))

	route.Get("/client/:id", websocket.New(func(c *websocket.Conn) {
		client := Client{
			Id:         c.Params("id"),
			Connection: c,
			User:       false,
		}
		defer func() {
			unregister <- client
			c.Close()
		}()

		register <- client

		for {
			message := models.SocketEvent{}
			err := client.Connection.ReadJSON(&message)
			if err != nil {
				return
			}
			if strings.HasPrefix(message.Event, "result-") {
				if channel, ok := Results[message.Event+client.Id]; ok {
					channel <- message.Data
					close(channel)
					delete(Results, message.Event+client.Id)
				}
			} else if message.Event == "usage" {
				fmt.Println(message)
				if _, ok := UsageStreams[client.Id]; ok {
					for _, client := range UsageStreams[client.Id] {
						client.Connection.WriteJSON(message)
					}
				}
			} else if strings.HasPrefix(message.Event, "devices-") {
				HandleDeviceEvent(client.Id, message)
			}
		}
	}))
}
