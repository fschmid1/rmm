package socket

import (
	"time"

	"festech.de/rmm/backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetDeviceProcessList(c *fiber.Ctx, event models.SocketEvent) error {
	resultChannel := CreateResultChannel("process-list", event.Id)
	SendMessage(event.Id, models.SocketEvent{
		Event: "process-list",
	})
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
	SendMessage(event.Id, models.SocketEvent{
		Event: "process-kill",
		Data:  event.Data,
	})
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
	default:
		return c.Status(400).SendString("Function not valid")
	}
}

func StartUsageStream(c *fiber.Ctx, event models.SocketEvent) error {
	SendMessage(event.Id, models.SocketEvent{
		Event: "usage-start",
	})
	return c.SendString("Started usage stream")
}

func StopUsageStream(c *fiber.Ctx, event models.SocketEvent) error {
	SendMessage(event.Id, models.SocketEvent{
		Event: "usage-stop",
	})
	return c.SendString("Stoped usage stream")
}
