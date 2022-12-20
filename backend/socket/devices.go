package socket

import (
	"errors"

	"festech.de/rmm/backend/config"
	"festech.de/rmm/backend/handlers"
	"festech.de/rmm/backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DeviceQuery struct {
	Id string `json:"id"`
	Mac bool `json:"mac"`
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
	var err error
	switch message.Event {
	case "devices-register":
		response_data, err = registerDevice(message.Data.(models.Device))
	case "devices-get":
		response_data, err = getDeviceById(message.Data.(DeviceQuery).Id, message.Data.(DeviceQuery).Mac)
	case "devices-update":
		response_data, err = updateDevice(message.Data.(models.Device))
	}
	if err != nil {
		response_data = err.Error()
	}
	
	response_message := models.SocketEvent{
		Event: "devices-response",
		Data: response_data,
		Id: clientId,
	}

	SendMessage(clientId, response_message)
}

func getDeviceById(id string, mac bool)(models.Device, error) {
	var device models.Device
	var result *gorm.DB
	if mac  {
		systemInfo := handlers.GetSystemInfoByMacAddress(id)
		result = config.Database.Preload(clause.Associations).Where("system_info_id = ?", systemInfo.ID).First(&device)
	} else {
		result = config.Database.Preload(clause.Associations).Find(&device, id)
	}
	if result.Error != nil {
		return models.Device{}, errors.New("Something went wrong")
	}
	if result.RowsAffected == 0 {
		return models.Device{}, errors.New("Device not found")
	}

	return device, nil
}

func registerDevice(device models.Device) (models.Device, error) {
	if handlers.GetSystemInfoByMacAddress(device.SystemInfo.MacAddress).ID > 0 {
		return models.Device{}, errors.New("Device with this mac address is already registered")
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

func updateDevice(device models.Device) (models.Device, error) {
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