package models

type SocketEvent struct {
	Id    string      `json:"id"`
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
