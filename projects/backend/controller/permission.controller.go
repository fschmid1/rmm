package controller

import (
	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/config"
	"gorm.io/gorm/clause"
)

func GetDevicePermissions(id uint) ([]models.DevicePermissions, error) {
	permissions := []models.DevicePermissions{}
	err := config.Database.Preload(clause.Associations).Where("device_id = ?", id).Find(&permissions).Error
	return permissions, err
}

func GetDevicePermissionsByUserId(deviceId uint, id uint64) (models.DevicePermissions, error) {
	permissions := models.DevicePermissions{}
	err := config.Database.Where("user_id = ?", id).Where("device_id = ?", deviceId).First(&permissions).Error
	return permissions, err
}

func GetDevicePermissionByID(id uint64) (models.DevicePermissions, error) {
	permission := models.DevicePermissions{}
	err := config.Database.Preload(clause.Associations).Where("id = ?", id).Find(&permission).Error
	return permission, err
}

func UpdateDevicePermissions(permissions models.DevicePermissions) (models.DevicePermissions, error) {
	permission, err := GetDevicePermissionByID(permissions.ID)
	if permission.ID == 0 {
		err = config.Database.Create(&permissions).Error
	} else {
		err = config.Database.Model(&permission).Updates(map[string]interface{}{
			"Run":               permissions.Run,
			"Kill":              permissions.Kill,
			"Reboot":            permissions.Reboot,
			"Shutdown":          permissions.Shutdown,
			"ChangePermissions": permissions.ChangePermissions,
			"ProcessList":       permissions.ProcessList,
			"ServiceList":       permissions.ServiceList,
			"ServiceStart":      permissions.ServiceStart,
			"ServiceStop":       permissions.ServiceStop,
			"ServiceRestart":    permissions.ServiceRestart,
			"ServiceLogs":       permissions.ServiceLogs,
			"ServiceStatus":     permissions.ServiceStatus,
		}).Error
	}
	return permission, err
}

func DeleteDevicePermissions(id uint) error {
	err := config.Database.Where("id = ?", id).Delete(models.DevicePermissions{}).Error
	return err
}
