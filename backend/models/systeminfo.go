package models

type SystemInfo struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	Os          string  `json:"os" gorm:"not null"`
	IP          string  `json:"ip" gorm:"not null"`
	MacAddress  string  `json:"macAddress" gorm:"uniqueIndex; not null"`
	HostName    string  `json:"hostName" gorm:"not null"`
	Cores       int     `json:"cores" gorm:"not null"`
	GPU         string  `json:"gpu" gorm:"not null"`
	CPU         string  `json:"cpu" gorm:"not null"`
	MemoryTotal float64 `json:"memoryTotal" gorm:"not null"`
	MemoryUsed  float64 `json:"memoryUsed" gorm:"not null"`
	DiskTotal   float64 `json:"diskTotal" gorm:"not null"`
	DiskUsed    float64 `json:"diskUsed" gorm:"not null"`
}
