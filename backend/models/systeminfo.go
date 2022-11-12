package models

type SystemInfo struct {
	ID     uint   `gorm:"primary_key" json:"id"`
	Os     string `json:"os" gorm:"not null"`
	Cores  int    `json:"cores" gorm:"not null"`
	Memory string `json:"memory" gorm:"not null"`
}
