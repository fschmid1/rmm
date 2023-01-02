package vars

import "github.com/fes111/rmm/projects/rmm/go/lib/models"

var Device models.Device
var Configuration models.Configuration

var RestUrl string
var WsUrl string

type handler struct {
	Handlers map[string]func(models.SocketEvent)
	Onces    map[string]bool
}

func NewHandler() handler {
	h := handler{}
	h.Handlers = make(map[string]func(models.SocketEvent))
	h.Onces = make(map[string]bool)
	return h
}

func (h *handler) On(event string, f func(models.SocketEvent)) {
	h.Handlers[event] = f
}

func (h *handler) Once(event string, f func(models.SocketEvent)) {
	h.Handlers[event] = f
	h.Onces[event] = true
}

var Handlers = NewHandler()
var Queue = make(chan models.SocketEvent, 100)
