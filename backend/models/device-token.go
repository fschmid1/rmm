package models

type DeviceToken struct {
	ID       uint `gorm:"primary_key" json:"id"`
	DeviceID uint `json:"deviceID" gorm:"uniqueIndex; not null"`
	UserID   uint `json:"userID" gorm:"uniqueIndex; not null"`
}
