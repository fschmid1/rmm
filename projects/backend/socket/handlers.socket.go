package socket

import (
	"strconv"
	"time"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/controller"
	"github.com/gofiber/fiber/v2"
)

func GetDeviceProcessList(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("process-list", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "process-list",
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 5).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func RunCommand(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("run", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "run",
		Data:  event.Data,
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 5).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func DeviceKillProcess(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("process-kill", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "process-kill",
		Data:  event.Data,
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 5).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func ShutdownDevice(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("shutdown", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "shutdown",
		Data:  event.Data,
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 5).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func GetServiceLogs(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("service-logs", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "service-logs",
		Data:  event.Data,
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 5).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func GetServiceStatus(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("service-status", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "service-status",
		Data:  event.Data,
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 5).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func GetServiceList(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("service-list", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "service-list",
		Data:  event.Data,
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 30).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func StartService(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("service-start", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "service-start",
		Data:  event.Data,
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 5).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func StopService(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("service-stop", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "service-stop",
		Data:  event.Data,
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 5).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func RestartService(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("service-restart", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "service-restart",
		Data:  event.Data,
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 5).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func RebootDevice(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("reboot", event.Id)
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "reboot",
		Data:  event.Data,
	})
	if err != nil {
		return c.SendStatus(500)
	}
	timeChan := time.NewTimer(time.Second * 5).C
	for {
		select {
		case msg := <-resultChannel:
			return c.JSON(msg)
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func FunctionsHandler(c *fiber.Ctx) error {
	event := new(models.SocketEvent)
	userId := c.Locals("user").(models.User).ID

	if err := c.BodyParser(event); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	device, err := GetDeviceByDeviceId(event.Id)
	if err != nil {
		return c.SendStatus(403)
	}
	permission, err := controller.GetDevicePermissionsByUserId(device.ID, uint64(userId))
	if err != nil {
		return c.SendStatus(403)
	}
	if event.Event == "run" && permission.Run {
		return RunCommand(c, *event)
	}
	if event.Event == "process-list" && permission.ProcessList {
		return GetDeviceProcessList(c, *event)
	}
	if event.Event == "process-kill" && permission.Kill {
		return DeviceKillProcess(c, *event)
	}
	if event.Event == "service-logs" && permission.ServiceLogs {
		return GetServiceLogs(c, *event)
	}
	if event.Event == "service-list" && permission.ServiceList {
		return GetServiceList(c, *event)
	}
	if event.Event == "service-status" && permission.ServiceStatus {
		return GetServiceStatus(c, *event)
	}
	if event.Event == "service-start" && permission.ServiceStart {
		return StartService(c, *event)
	}
	if event.Event == "service-stop" && permission.ServiceStop {
		return StopService(c, *event)
	}
	if event.Event == "service-restart" && permission.ServiceRestart {
		return RestartService(c, *event)
	}
	if event.Event == "usage-start" {
		return StartUsageStream(c, *event)
	}
	if event.Event == "usage-stop" {
		return StopUsageStream(c, *event)
	}
	if event.Event == "shutdown" && permission.Shutdown {
		return ShutdownDevice(c, *event)
	}
	if event.Event == "reboot" && permission.Reboot {
		return RebootDevice(c, *event)
	}
	return c.SendStatus(403)
}

func StartUsageStream(c *fiber.Ctx, event models.SocketEvent) error {
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "usage-start",
	})
	if err != nil {
		return c.SendStatus(500)
	}
	userId := strconv.FormatUint(uint64(c.Locals("user").(models.User).ID), 10)
	if client, ok := Clients[userId]; ok {
		if UsageStreams[event.Id] == nil {
			UsageStreams[event.Id] = make(map[string]Client)
		}
		UsageStreams[event.Id][client.Id] = client
	}

	return c.JSON(models.SocketEvent{
		Event: "usage-start",
		Data:  "Started usage stream",
		Id:    event.Id,
	})
}

func StopUsageStream(c *fiber.Ctx, event models.SocketEvent) error {
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "usage-stop",
	})
	if err != nil {
		return c.SendStatus(500)
	}
	if _, ok := UsageStreams[event.Id]; ok {
		UsageStreams[event.Id] = make(map[string]Client)
	}
	return c.JSON(models.SocketEvent{
		Event: "usage-stop",
		Data:  "Stopped usage stream",
		Id:    event.Id,
	})
}
