package socket

import (
	"errors"
	"log"

	"festech.de/rmm/backend/handlers"
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

func SendMessage(id string, message string) error {
	client, found := Clients[id]

	if !found {
		return errors.New("Client not found")
	}
	client.Connection.WriteMessage(websocket.TextMessage, []byte(message))

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
			messageType, message, err := client.Connection.ReadMessage()
			if err != nil {
				return
			}

			if messageType == websocket.TextMessage {
				Broadcast <- string(message)
			}
		}
	}))
}
