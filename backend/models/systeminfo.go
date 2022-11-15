package models

type SystemInfo struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Os         string `json:"os" gorm:"not null"`
	IP         string `json:"ip" gorm:"not null"`
	MacAddress string `json:"macAddress" gorm:"uniqueIndex; not null"`
	HostName   string `json:"hostName" gorm:"not null"`
	Cores      int    `json:"cores" gorm:"not null"`
	Memory     string `json:"memory" gorm:"not null"`
	Disk       string `json:"disk" gorm:"not null"`
}
