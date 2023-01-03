package models

import "time"

type Device struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Name         string     `json:"name" gorm:"not null"`
	DeviceID     string     `json:"deviceID" gorm:"uniqueIndex; not null"`
	Connected    bool       `json:"connected" gorm:"default:false"`
	SystemInfo   SystemInfo `json:"systemInfo" gorm:"constraint:OnDelete:CASCADE;"`
	SystemInfoId uint       `json:"systemInfoId"`
}
