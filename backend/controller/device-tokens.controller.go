package controller

import (
	"errors"

	"festech.de/rmm/backend/config"
	"festech.de/rmm/backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetDeviceToken(token string) (models.DeviceToken, error) {
	var deviceToken models.DeviceToken
	result := config.Database.Preload(clause.Associations).Find(&deviceToken).Where("token = ?", token)
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
