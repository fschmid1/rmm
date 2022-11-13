package models

type SystemInfo struct {
	ID         uint   `json:"id"`
	Os         string `json:"os"`
	MacAddress string `json:"macAddress"`
	HostName   string `json:"hostName"`
	Cores      int    `json:"cores"`
	Memory     string `json:"memory"`
	Disk       string `json:"disk"`
}
