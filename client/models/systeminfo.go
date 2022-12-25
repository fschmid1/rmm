package models

type SystemInfo struct {
	ID          uint    `json:"id"`
	Os          string  `json:"os"`
	IP          string  `json:"ip"`
	GPU         string  `json:"gpu"`
	MacAddress  string  `json:"macAddress"`
	HostName    string  `json:"hostName"`
	CPU         string  `json:"cpu"`
	Cores       int     `json:"cores"`
	MemoryTotal float64 `json:"memoryTotal"`
	MemoryUsed  float64 `json:"memoryUsed"`
	DiskTotal   float64 `json:"diskTotal"`
	DiskUsed    float64 `json:"diskUsed"`
}
