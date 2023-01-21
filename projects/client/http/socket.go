package http

import (
	"github.com/fes111/rmm/libs/go/models"
	"github.com/recws-org/recws"
)

var SocketConn recws.RecConn

func SocketSend(event string, data interface{}) {
	SocketConn.WriteJSON(models.SocketEvent{
		Event: event,
		Data:  data,
	})
}
