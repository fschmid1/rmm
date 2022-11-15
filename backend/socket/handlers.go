package socket

import (
	"time"

	"festech.de/rmm/backend/models"
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
		case data := <-resultChannel:
			return c.JSON(models.SocketEvent{
				Event: "result-process-list",
				Data:  data,
				Id:    event.Id,
			})
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
		case data := <-resultChannel:
			return c.JSON(models.SocketEvent{
				Event: "result-process-kill",
				Data:  data,
				Id:    event.Id,
			})
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
		case data := <-resultChannel:
			return c.JSON(models.SocketEvent{
				Event: "result-shutdown",
				Data:  data,
				Id:    event.Id,
			})
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
		case data := <-resultChannel:
			return c.JSON(models.SocketEvent{
				Event: "result-reboot",
				Data:  data,
				Id:    event.Id,
			})
		case <-timeChan:
			return c.Status(500).SendString("Something went wrong")
		}
	}
}

func FunctionsHandler(c *fiber.Ctx) error {
	event := new(models.SocketEvent)

	if err := c.BodyParser(event); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	switch event.Event {
	case "process-list":
		return GetDeviceProcessList(c, *event)
	case "process-kill":
		return DeviceKillProcess(c, *event)
	case "usage-start":
		return StartUsageStream(c, *event)
	case "usage-stop":
		return StopUsageStream(c, *event)
	case "shutdown":
		return ShutdownDevice(c, *event)
	case "reboot":
		return RebootDevice(c, *event)
	default:
		return c.Status(400).SendString("Function not valid")
	}
}

func StartUsageStream(c *fiber.Ctx, event models.SocketEvent) error {
	err := SendMessage(event.Id, models.SocketEvent{
		Event: "usage-start",
	})
	if err != nil {
		return c.SendStatus(500)
	}
	if client, ok := Clients[c.GetReqHeaders()["X-Auth-User"]]; ok {
		if UsageStreams[event.Id] == nil {
			UsageStreams[event.Id] = make(map[string]Client)
		}
		UsageStreams[event.Id][client.Id] = client
	}
	
	return c.SendString("Started usage stream")
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
	return c.SendString("Stoped usage stream")
}
