package controller

import (
	"errors"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/config"
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
