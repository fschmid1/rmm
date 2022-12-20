package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"festech.de/rmm/client/models"
	"festech.de/rmm/client/system"
	"festech.de/rmm/client/vars"
	"github.com/google/uuid"
)

var configPath = system.CreateConfigurationPath() + "config"
var devicePath = system.CreateConfigurationPath() + "device"

func RegisterDevice() {
	id := uuid.New().String()

	systemInfo := models.SystemInfo{
		HostName:   system.GetHostName(),
		Os:         system.GetOS(),
		IP:         system.GetIP(),
		MacAddress: system.GetMacAddress(),
		Cores:      system.GetCores(),
		Memory:     system.GetMemory(),
		Disk:       system.GetDisk(),
	}
	device := models.Device{
		DeviceID:   id,
		Name:       system.GetHostName(),
		SystemInfo: systemInfo,
	}
	vars.Handlers.Once("response-devices-register", func(event models.SocketEvent) {
		if event.Error == "Devcice with this mac address is already registered" {
			vars.Handlers.Once("response-devices-get", func(event models.SocketEvent) {
				device = event.Data.(models.Device)
				WriteConfiguration(devicePath, device)
			})
			vars.Queue <- models.SocketEvent{
				Event: "devices-get",
				Data:  id,
			}
		} else {
			device = event.Data.(models.Device)
			WriteConfiguration(devicePath, device)
		}
	})
	vars.Queue <- models.SocketEvent{
		Event: "devices-register",
		Data:  device,
	}
}

func UpdateSystemInfo() {
	vars.Device.SystemInfo = models.SystemInfo{
		HostName:   system.GetHostName(),
		Os:         system.GetOS(),
		IP:         system.GetIP(),
		MacAddress: system.GetMacAddress(),
		Cores:      system.GetCores(),
		Memory:     system.GetMemory(),
		Disk:       system.GetDisk(),
		ID:         vars.Device.SystemInfo.ID,
	}
	vars.Handlers.Once("response-devices-update", func(event models.SocketEvent) {
		if event.Error != "" {
			fmt.Println(event.Error)
			return
		}
		vars.Device = event.Data.(models.Device)
		WriteConfiguration(devicePath, vars.Device)
	})
	vars.Queue <- models.SocketEvent{
		Event: "devices-update",
		Data:  vars.Device,
	}
}

func CreateConfiguration() {
	config := models.Configuration{}
	WriteConfiguration(configPath, config)
	fmt.Printf("Configuration files was created under '%s'\nAfter configuration pls restart the client\n", configPath)
	os.Exit(0)
}

func WriteConfiguration(path string, data interface{}) {
	file, _ := json.MarshalIndent(data, "", " ")
	ioutil.WriteFile(path, file, 0775)
}

func createUrls() {
	var secure = ""
	if vars.Configuration.Secure {
		secure = "s"
	}
	vars.RestUrl = fmt.Sprintf("http%s://%s:%s/", secure, vars.Configuration.Host, vars.Configuration.Port)
	vars.WsUrl = fmt.Sprintf("ws%s://%s:%s/ws/client/", secure, vars.Configuration.Host, vars.Configuration.Port)
}

func ReadConfiguration() {
	firstTime := false
	if !system.FileOrFolderExists(configPath) {
		CreateConfiguration()
	}
	file, _ := ioutil.ReadFile(configPath)
	config := models.Configuration{}
	json.Unmarshal(file, &config)
	vars.Configuration = config
	WriteConfiguration(configPath, config)
	createUrls()
	if !system.FileOrFolderExists(devicePath) {
		RegisterDevice()
		firstTime = true
	}
	file, _ = ioutil.ReadFile(devicePath)
	device := models.Device{}
	json.Unmarshal(file, &device)

	vars.Device = device
	if !firstTime {
		UpdateSystemInfo()
	}
}
