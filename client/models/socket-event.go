package models

type SocketEvent struct  {
	Event string `json:"event"`
	Data interface{} `json:"data"`
}