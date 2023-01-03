package controller

import (
	"errors"
	"time"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/config"
	"gorm.io/gorm/clause"
)

func GetDeviceById(id string) (models.Device, error) {
	var device models.Device
	result := config.Database.Preload(clause.Associations).Where("device_id = ?", id).Find(&device)
	if result.Error != nil {
		return models.Device{}, errors.New("something went wrong")
	}
	if result.RowsAffected == 0 {
		return models.Device{}, errors.New("device not found")
	}

	return device, nil
}

func ParseDevice(data map[string]interface{}) models.Device {
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
			Os:          systemInfo["os"].(string),
			IP:          systemInfo["ip"].(string),
			MacAddress:  systemInfo["macAddress"].(string),
			HostName:    systemInfo["hostName"].(string),
			Cores:       int(systemInfo["cores"].(float64)),
			MemoryTotal: systemInfo["memoryTotal"].(float64),
			MemoryUsed:  systemInfo["memoryUsed"].(float64),
			DiskTotal:   systemInfo["diskTotal"].(float64),
			DiskUsed:    systemInfo["diskUsed"].(float64),
			CPU:         systemInfo["cpu"].(string),
			GPU:         systemInfo["gpu"].(string),
			ID:          uint(systemInfo["id"].(float64)),
		},
	}

	return device
}

func parseDate(date string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}
