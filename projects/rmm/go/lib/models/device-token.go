package models

type DeviceToken struct {
	ID       uint64 `gorm:"primary_key" json:"id"`
	DeviceID string `json:"deviceID" gorm:"uniqueIndex"`
	UserID   uint64 `json:"userID" gorm:"not null"`
	Token    string `json:"token" gorm:"not null"`
	Name     string `json:"name" gorm:"not null"`
}
