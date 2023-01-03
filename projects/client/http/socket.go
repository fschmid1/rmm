package http

import (
	"github.com/fes111/rmm/libs/go/models"
	"github.com/gorilla/websocket"
)

var SocketConn *websocket.Conn

func SocketSend(event string, data interface{}) {
	SocketConn.WriteJSON(models.SocketEvent{
		Event: event,
		Data:  data,
	})
}
