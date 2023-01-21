package socket

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/fes111/rmm/libs/go/helpers"
	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/config"
	"github.com/fes111/rmm/projects/backend/controller"
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

func NotfiyUserDeviceConnection(id string, connected bool) {
	device, err := GetDeviceByDeviceId(id)
	if err != nil {
		return
	}

	result := []map[string]interface{}{}
	config.Database.Table("user_devices").Select("user_id").Where("device_id = ?", device.ID).Find(&result)
	connectionEvent := CreateConnectionEventChannel(device.DeviceID)
	var duration time.Duration
	if connected {
		duration = time.Second * 0
	} else {
		duration = time.Second * 10
	}
	timer := time.NewTimer(duration)
	select {
	case <-connectionEvent:
		if !helpers.IsClosed(connectionEvent) {
			timer.Stop()
			close(connectionEvent)
			delete(ConnectionEvents, device.DeviceID)
			return
		}
	case <-timer.C:
		if !connected {
			notifications := controller.GetDeviceNotificationsByDeviceID(device.ID)
			for _, userRaw := range notifications {
				user, err := controller.GetUserById(uint(userRaw["user_id"].(uint64)))
				if err != nil {
					return
				}
				if user.PushToken != "" {
					controller.SendMessage(user, fmt.Sprintf("Device %s is %s", device.Name, "disconnected"), "Device connection")
				}
			}
		}
		for _, userRaw := range result {
			userId := strconv.FormatUint(userRaw["user_id"].(uint64), 10)
			if client, ok := Clients[userId]; ok {
				SendMessage(client.Id, models.SocketEvent{
					Event: "device-connection",
					Data: map[string]interface{}{
						"id":        device.ID,
						"connected": connected,
					},
				})
			}
		}
		if !helpers.IsClosed(connectionEvent) {
			close(connectionEvent)
			delete(ConnectionEvents, device.DeviceID)
		}
		return
	}
}

func SetDeviceConnected(id string, connected bool) (bool, error) {
	device, err := GetDeviceByDeviceId(id)
	if err != nil {
		return false, errors.New("not found")
	}

	if result := config.Database.Model(&device).Update("connected", connected); result.Error != nil {
		return false, errors.New("db error")
	}
	if _, ok := ConnectionEvents[id]; ok {
		ConnectionEvents[id] <- connected
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
		log.Println(err)
	}
}

func getDeviceById(id string, mac bool) (models.Device, error) {
	var device models.Device
	var result *gorm.DB
	if mac {
		systemInfo := controller.GetSystemInfoByMacAddress(id)
		result = config.Database.Preload(clause.Associations).Where("system_info_id = ?", systemInfo.ID).First(&device)
	} else {
		result = config.Database.Preload(clause.Associations).Where("device_id = ?", id).First(&device)
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
	device, err = getDeviceById(device.DeviceID, false)
	return device, nil
}
