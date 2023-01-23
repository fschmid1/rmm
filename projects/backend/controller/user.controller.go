package controller

import (
	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/config"
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
	devicePermissions, _ := GetDevicePermissionsByUserId(device.ID, deviceToken.UserID)
	if devicePermissions.ID == 0 {
		UpdateDevicePermissions(models.DevicePermissions{
			ID:     0,
			UserID: deviceToken.UserID, DeviceID: device.ID, Run: true, Shutdown: true, Reboot: true,
			ProcessList:       true,
			ServiceList:       true,
			ServiceStart:      true,
			ServiceStop:       true,
			ServiceRestart:    true,
			ServiceLogs:       true,
			ServiceStatus:     true,
			Kill:              true,
			ChangePermissions: true,
		})
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.Database.Find(&users).Error
	return users, err
}

func UpdateUser(user models.User) error {
	err := config.Database.Save(&user).Error
	return err
}

func GetUserById(id uint) (models.User, error) {
	user := models.User{}
	err := config.Database.First(&user, id).Error
	return user, err
}
