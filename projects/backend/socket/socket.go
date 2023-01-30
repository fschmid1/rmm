package socket

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Id           string
	Connection   *websocket.Conn
	User         bool
	Authenicated bool
}

var Clients = make(map[string]Client)
var register = make(chan Client)
var Broadcast = make(chan models.SocketEvent)
var unregister = make(chan Client)

var Results = make(map[string]chan interface{})

var UsageStreams = make(map[string]map[string]Client)
var ConnectionEvents = make(map[string]chan bool)

func runHub() {
	for {
		select {
		case client := <-register:
			if !client.User {
				SetDeviceConnected(client.Id, true)
				go NotfiyUserDeviceConnection(client.Id, true)
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
				go NotfiyUserDeviceConnection(client.Id, false)
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

func CreateConnectionEventChannel(id string) chan bool {
	if _, ok := ConnectionEvents[id]; ok {
		return ConnectionEvents[id]
	}
	connectionChannel := make(chan bool, 1)
	ConnectionEvents[id] = connectionChannel
	return connectionChannel
}

func SendMessage(id string, message interface{}) error {
	client, found := Clients[id]

	if !found {
		log.Println("Client not found")
		return errors.New("Client not found")
	}
	err := client.Connection.WriteJSON(message)
	return err
}

func RegisterWebsocketRoute(app *fiber.App) {
	go runHub()

	route := app.Group("/ws")

	route.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	route.Get("/user", websocket.New(func(c *websocket.Conn) {
		client := Client{
			Connection: c,
			User:       true,
		}
		defer func() {
			unregister <- client
			c.Close()
		}()

		time.AfterFunc(5*time.Second, func() {
			if !client.Authenicated {
				client.Connection.WriteJSON(models.SocketEvent{
					Event: "auth-fail",
				})
				client.Connection.Close()
			}
		})

		for {
			message := models.SocketEvent{}
			err := client.Connection.ReadJSON(&message)
			if err != nil {
				return
			}

			if message.Event == "auth" {
				verrify, _ := controller.VerifyUserJWT(message.Data.(map[string]interface{})["token"].(string))
				if verrify {
					client.Authenicated = true
					client.Id = message.Data.(map[string]interface{})["id"].(string)
					register <- client
					client.Connection.WriteJSON(models.SocketEvent{
						Event: "auth-success",
					})
				} else {
					client.Connection.WriteJSON(models.SocketEvent{
						Event: "auth-fail",
					})
					return
				}
			}
			if strings.HasPrefix(message.Event, "result-") {
				if channel, ok := Results[message.Event+client.Id]; ok {
					channel <- message.Data
					close(channel)
					delete(Results, message.Event+client.Id)
				}
			} else if message.Event == "usage" {
				if _, ok := UsageStreams[client.Id]; ok {
					for _, client := range UsageStreams[client.Id] {
						client.Connection.WriteJSON(message)
					}
				}
			}
		}
	}))

	route.Get("/client", websocket.New(func(c *websocket.Conn) {
		client := Client{
			Connection:   c,
			User:         false,
			Authenicated: false,
		}
		defer func() {
			unregister <- client
			c.Close()
		}()

		time.AfterFunc(5*time.Second, func() {
			if !client.Authenicated {
				client.Connection.WriteJSON(models.SocketEvent{
					Event: "auth-fail",
				})
				client.Connection.Close()
			}
		})

		for {
			message := models.SocketEvent{}
			err := client.Connection.ReadJSON(&message)
			if err != nil {
				return
			}
			if message.Event == "auth" {
				token := message.Data.(map[string]interface{})["token"].(string)
				verify := controller.VerifyClientJWT(token)
				if verify {
					client.Authenicated = true
					client.Id = message.Data.(map[string]interface{})["id"].(string)
					register <- client
					go controller.SetDeviceToken(client.Id, token)
					go controller.AddDeviceToUser(client.Id, token)
					client.Connection.WriteJSON(models.SocketEvent{
						Event: "auth-success",
					})
				} else {
					client.Connection.WriteJSON(models.SocketEvent{
						Event: "auth-fail",
					})
					return
				}
			}
			if !client.Authenicated {
				continue
			}

			if strings.HasPrefix(message.Event, "result-") {
				if channel, ok := Results[message.Event+client.Id]; ok {
					channel <- message
					close(channel)
					delete(Results, message.Event+client.Id)
				}
			} else if message.Event == "usage" {
				log.Println(UsageStreams, client.Id)
				if _, ok := UsageStreams[client.Id]; ok {
					log.Println("Sending usage to client")
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
