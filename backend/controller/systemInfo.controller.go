package controller

import (
	"errors"

	"festech.de/rmm/backend/config"
	"festech.de/rmm/backend/models"
	"gorm.io/gorm"
)

func GetSystemInfoByMacAddress(macAddress string) models.SystemInfo {
	var systemInfo models.SystemInfo

	result := config.Database.Where("mac_address = ?", macAddress).First(&systemInfo)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return models.SystemInfo{}
	}

	return systemInfo
}
