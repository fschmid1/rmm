package controller

import (
	"festech.de/rmm/backend/config"
)

func AddDeviceToUser(id string, token string) error {
	deviceToken, err := GetDeviceToken(token)
	if err != nil {
		return err
	}
	device, err := GetDeviceById(id)
	if err != nil {
		return err
	}
	result := map[string]interface{}{}
	config.Database.Table("user_devices").Where(&result).Where("user_id = ?", deviceToken.UserID).Where("device_id = ?", device.ID).First(&result)
	if len(result) == 0 {
		config.Database.Table("user_devices").Create(map[string]interface{}{"user_id": deviceToken.UserID, "device_id": device.ID})
	}
	return nil
}
