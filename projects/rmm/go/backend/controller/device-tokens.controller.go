package controller

import (
	"errors"

	"github.com/fes111/rmm/projects/rmm/go/backend/config"
	"github.com/fes111/rmm/projects/rmm/go/lib/models"
	"gorm.io/gorm"
)

func GetDeviceToken(token string) (models.DeviceToken, error) {
	var deviceToken models.DeviceToken
	result := config.Database.Where("token = ?", token).Find(&deviceToken)
	if result.Error != nil {
		return models.DeviceToken{}, errors.New("something went wrong")
	}
	if result.RowsAffected == 0 {
		return models.DeviceToken{}, errors.New("device-token not found")
	}

	return deviceToken, nil
}

func SetDeviceToken(id string, token string) error {
	deviceToken, err := GetDeviceToken(token)
	if err != nil {
		return err
	}
	deviceToken.DeviceID = id
	err = config.Database.Transaction(func(tx *gorm.DB) error {
		if result := tx.Updates(&deviceToken); result.Error != nil {
			return result.Error
		}
		return nil
	})
	return err
}
