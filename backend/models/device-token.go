package models

type DeviceToken struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	DeviceID string `json:"deviceID" gorm:"uniqueIndex; not null"`
	UserID   uint   `json:"userID" gorm:"uniqueIndex; not null"`
	Token    string `json:"token" gorm:"not null"`
}
