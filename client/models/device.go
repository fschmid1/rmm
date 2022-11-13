package models

import (
	"time"
)

type Device struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Name         string     `json:"name"`
	DeviceID     string     `json:"deviceID"`
	Connected    bool       `json:"connected"`
	SystemInfo   SystemInfo `json:"systemInfo"`
	SystemInfoId uint       `json:"systemInfoId"`
}
