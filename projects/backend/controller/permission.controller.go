package controller

import (
	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/config"
)

func GetDevicePermissions(id uint) ([]models.DevicePermissions, error) {
	permissions := []models.DevicePermissions{}
	err := config.Database.Where("device_id = ?", id).Find(&permissions).Error
	return permissions, err
}

func GetDevicePermissionsByUserId(deviceId uint, id uint64) ([]models.DevicePermissions, error) {
	permissions := []models.DevicePermissions{}
	err := config.Database.Where("user_id = ?", id).Where("device_id = ?", deviceId).Find(&permissions).Error
	return permissions, err
}

func UpdateDevicePermissions(permissions models.DevicePermissions) (models.DevicePermissions, error) {
	var err error = nil
	if config.Database.Model(&permissions).Where("id = ?", permissions.ID).Updates(permissions).RowsAffected == 0 {
		err = config.Database.Create(&permissions).Error
	}
	return permissions, err
}

func DeleteDevicePermissions(id uint) error {
	err := config.Database.Where("id = ?", id).Delete(models.DevicePermissions{}).Error
	return err
}
