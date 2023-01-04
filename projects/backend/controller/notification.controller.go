package controller

import (
	"github.com/fes111/rmm/projects/backend/config"
)

func ToggleDeviceNotification(userId uint, deviceId uint) {
	result := map[string]interface{}{}
	config.Database.Table("device_notifications").Where("user_id = ?", userId).Where("device_id = ?", deviceId).Find(&result)
	if len(result) == 0 {
		config.Database.Table("device_notifications").Create(map[string]interface{}{"user_id": userId, "device_id": deviceId})
	} else {
		config.Database.Table("device_notifications").Where("user_id = ?", userId).Where("device_id = ?", deviceId).Delete(map[string]uint{"user_id": userId, "device_id": deviceId})
	}
}

func GetDeviceNotificationsByUserID(userId uint) []map[string]interface{} {
	result := []map[string]interface{}{}
	config.Database.Table("device_notifications").Where("user_id = ?", userId).Find(&result)
	return result
}

func GetDeviceNotificationsByDeviceID(deviceID uint) []map[string]interface{} {
	result := []map[string]interface{}{}
	config.Database.Table("device_notifications").Where("device_id = ?", deviceID).Find(&result)
	return result
}
