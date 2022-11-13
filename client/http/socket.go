package http

import (
	"festech.de/rmm/client/models"
	"github.com/gorilla/websocket"
)

var SocketConn *websocket.Conn

func SocketSend(event string, data interface{}) {
	SocketConn.WriteJSON(models.SocketEvent{
		Event: event,
		Data: data,
	})
}