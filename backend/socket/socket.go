package socket

import (
	"errors"
	"log"
	"strings"

	"festech.de/rmm/backend/handlers"
	"festech.de/rmm/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Id         string
	Connection *websocket.Conn
}

var Clients = make(map[string]Client)
var register = make(chan Client)
var Broadcast = make(chan string)
var unregister = make(chan Client)

var Results = make(map[string]chan interface{})

func runHub() {
	for {
		select {
		case client := <-register:
			handlers.SetDeviceConnected(client.Id, true)
			Clients[client.Id] = client

		case message := <-Broadcast:
			for clientID := range Clients {
				client := Clients[clientID]
				if err := client.Connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Println("write error:", err)

					client.Connection.WriteMessage(websocket.CloseMessage, []byte{})
					client.Connection.Close()
					unregister <- client
				}
			}

		case client := <-unregister:
			handlers.SetDeviceConnected(client.Id, false)
			delete(Clients, client.Id)
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
	client.Connection.WriteJSON(message)

	return nil
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

	route.Get("/:id", websocket.New(func(c *websocket.Conn) {
		client := Client{
			Id:         c.Params("id"),
			Connection: c,
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
			}
		}
	}))
}
