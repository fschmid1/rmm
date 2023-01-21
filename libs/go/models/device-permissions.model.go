package models

type DevicePermissions struct {
	ID                uint64 `gorm:"primary_key" json:"id"`
	DeviceID          uint   `json:"deviceID" gorm:"uniqueIndex"`
	UserID            uint64 `json:"userID" gorm:"not null"`
	Run               bool   `json:"run" gorm:"default:false"`
	Shutdown          bool   `json:"shutdown" gorm:"default:false"`
	Reboot            bool   `json:"reboot" gorm:"default:false"`
	ProcessList       bool   `json:"processList" gorm:"default:false"`
	ServiceList       bool   `json:"serviceList" gorm:"default:false"`
	ServiceLogs       bool   `json:"serviceLogs" gorm:"default:false"`
	ServiceStart      bool   `json:"serviceStart" gorm:"default:false"`
	ServiceStop       bool   `json:"serviceStop" gorm:"default:false"`
	ServiceRestart    bool   `json:"serviceRestart" gorm:"default:false"`
	ServiceStatus     bool   `json:"serviceStatus" gorm:"default:false"`
	Kill              bool   `json:"kill" gorm:"default:false"`
	ChangePermissions bool   `json:"changePermissions" gorm:"default:false"`
}
