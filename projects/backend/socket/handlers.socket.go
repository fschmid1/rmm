package socket

import (
	"strconv"
	"time"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

	if err := c.BodyParser(event); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	switch event.Event {
	case "run":
		return RunCommand(c, *event)
	case "process-list":
		return GetDeviceProcessList(c, *event)
	case "process-kill":
		return DeviceKillProcess(c, *event)
	case "service-logs":
		return GetServiceLogs(c, *event)
	case "service-list":
		return GetServiceList(c, *event)
	case "service-status":
		return GetServiceStatus(c, *event)
	case "service-start":
		return StartService(c, *event)
	case "service-stop":
		return StopService(c, *event)
	case "service-restart":
		return RestartService(c, *event)
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
	userId := strconv.FormatFloat(c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user"].(map[string]interface{})["id"].(float64), 'f', 0, 64)
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
