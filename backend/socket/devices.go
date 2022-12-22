package socket

import (
	"errors"
	"fmt"

	"festech.de/rmm/backend/config"
	"festech.de/rmm/backend/controller"
	"festech.de/rmm/backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DeviceQuery struct {
	Id  string `json:"id"`
	Mac bool   `json:"mac"`
}

func GetDeviceByDeviceId(id string) (models.Device, error) {
	var device models.Device

	result := config.Database.Where("device_id = ?", id).First(&device)

	if result.RowsAffected == 0 {
		return models.Device{}, errors.New("not found")
	}

	return device, nil
}

func SetDeviceConnected(id string, connected bool) (bool, error) {
	device, err := GetDeviceByDeviceId(id)
	if err != nil {
		return false, errors.New("not found")
	}

	if result := config.Database.Model(&device).Update("connected", connected); result.Error != nil {
		return false, errors.New("db error")
	}
	return true, nil
}

func HandleDeviceEvent(clientId string, message models.SocketEvent) {
	var response_data interface{}
	var err error = nil
	switch message.Event {
	case "devices-register":
		response_data, err = registerDevice(message.Data.(map[string]interface{}))
	case "devices-get":
		response_data, err = getDeviceById(message.Data.(map[string]interface{})["id"].(string), message.Data.(map[string]interface{})["mac"].(bool))
	case "devices-update":
		response_data, err = updateDevice(message.Data.(map[string]interface{}))
	}

	response_message := models.SocketEvent{
		Event: "response-" + message.Event,
		Data:  response_data,
		Id:    clientId,
	}
	if err != nil {
		response_message.Error = err.Error()
	}
	err = SendMessage(clientId, response_message)
	if err != nil {
		fmt.Println(err)
	}
}

func getDeviceById(id string, mac bool) (models.Device, error) {
	var device models.Device
	var result *gorm.DB
	if mac {
		systemInfo := controller.GetSystemInfoByMacAddress(id)
		result = config.Database.Preload(clause.Associations).Where("system_info_id = ?", systemInfo.ID).First(&device)
	} else {
		result = config.Database.Preload(clause.Associations).Find(&device, id)
	}
	if result.Error != nil {
		return models.Device{}, errors.New("something went wrong")
	}
	if result.RowsAffected == 0 {
		return models.Device{}, errors.New("device not found")
	}

	return device, nil
}

func registerDevice(data map[string]interface{}) (models.Device, error) {
	device := controller.ParseDevice(data)
	if controller.GetSystemInfoByMacAddress(device.SystemInfo.MacAddress).ID > 0 {
		return models.Device{}, errors.New("device with this mac address is already registered")
	}

	err := config.Database.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(&device); result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return models.Device{}, err
	}

	return device, nil
}

func updateDevice(data map[string]interface{}) (models.Device, error) {
	device := controller.ParseDevice(data)
	err := config.Database.Transaction(func(tx *gorm.DB) error {
		if result := tx.Updates(&device); result.Error != nil {
			return result.Error
		}
		if result := tx.Updates(&device.SystemInfo); result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return models.Device{}, err
	}

	return device, nil
}