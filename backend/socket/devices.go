package socket

import (
	"errors"
	"fmt"
	"time"

	"festech.de/rmm/backend/config"
	"festech.de/rmm/backend/handlers"
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

func registerDevice(data map[string]interface{}) (models.Device, error) {
	device := parseDevice(data)
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

func updateDevice(data map[string]interface{}) (models.Device, error) {
	device := parseDevice(data)
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

func parseDevice(data map[string]interface{}) models.Device {
	systemInfo := data["systemInfo"].(map[string]interface{})
	device := models.Device{
		DeviceID:     data["deviceID"].(string),
		Connected:    true,
		Name:         data["name"].(string),
		ID:           uint(data["id"].(float64)),
		CreatedAt:    parseDate(data["created_at"].(string)),
		UpdatedAt:    parseDate(data["updated_at"].(string)),
		SystemInfoId: uint(systemInfo["id"].(float64)),
		SystemInfo: models.SystemInfo{
			Os:         systemInfo["os"].(string),
			IP:         systemInfo["ip"].(string),
			MacAddress: systemInfo["macAddress"].(string),
			HostName:   systemInfo["hostName"].(string),
			Cores:      int(systemInfo["cores"].(float64)),
			Memory:     systemInfo["memory"].(string),
			Disk:       systemInfo["disk"].(string),
			ID:         uint(systemInfo["id"].(float64)),
		},
	}

	return device
}

func parseDate(date string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

func authDevice(id string) (models.Device, error) {
	device, err := GetDeviceByDeviceId(id)
	if err != nil {
		return models.Device{}, err
	}

	if result := config.Database.Save(&device); result.Error != nil {
		return models.Device{}, result.Error
	}

	return device, nil
}
