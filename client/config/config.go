package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"festech.de/rmm/client/http"
	"festech.de/rmm/client/models"
	"festech.de/rmm/client/system"
	"github.com/google/uuid"
)

var Device models.Device
var Configuration models.Configuration

var RestUrl string
var WsUrl string

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

	status, textBody, _ := http.Post(RestUrl+"devices", device, &device)
	if status == 400 && textBody == "Device with this mac address is already registered" {
		http.Get(fmt.Sprintf(RestUrl+"devices/%s?mac=1", device.SystemInfo.MacAddress), &device)
	}
	WriteConfiguration(devicePath, device)
}

func UpdateSystemInfo() {
	Device.SystemInfo = models.SystemInfo{
		HostName:   system.GetHostName(),
		Os:         system.GetOS(),
		IP:         system.GetIP(),
		MacAddress: system.GetMacAddress(),
		Cores:      system.GetCores(),
		Memory:     system.GetMemory(),
		Disk:       system.GetDisk(),
		ID:         Device.SystemInfo.ID,
	}
	status, textBody, _ := http.Patch(RestUrl+"devices", Device, &Device)
	if status == 400 && textBody == "Device with this mac address is already registered" {
		http.Get(fmt.Sprintf(RestUrl+"devices/%s?mac=1", Device.SystemInfo.MacAddress), &Device)
	}
	WriteConfiguration(devicePath, Device)
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
	if Configuration.Secure {
		secure = "s"
	}
	RestUrl = fmt.Sprintf("http%s://%s:%s/", secure, Configuration.Host, Configuration.Port)
	WsUrl = fmt.Sprintf("ws%s://%s:%s/ws/client/", secure, Configuration.Host, Configuration.Port)
}

func ReadConfiguration() {
	firstTime := false
	if !system.FileOrFolderExists(configPath) {
		CreateConfiguration()
	}
	file, _ := ioutil.ReadFile(configPath)
	config := models.Configuration{}
	json.Unmarshal(file, &config)
	Configuration = config
	WriteConfiguration(configPath, config)
	createUrls()
	if !system.FileOrFolderExists(devicePath) {
		RegisterDevice()
		firstTime = true
	}
	file, _ = ioutil.ReadFile(devicePath)
	device := models.Device{}
	json.Unmarshal(file, &device)

	Device = device
	if !firstTime {
		UpdateSystemInfo()
	}
}
