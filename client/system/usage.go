package system

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"festech.de/rmm/client/models"

	"festech.de/rmm/client/http"
)

var StartStopUsageStream chan bool
var EndUsageStream chan bool

func GetUsage() models.Usage {
	usage := models.Usage{}
	if IsLinux() {
		mem, _ := strconv.Atoi(strings.ReplaceAll(Run("free -m | awk '$1 == \"Mem:\" {print $3}'"), "\n", ""))
		disk, _ := strconv.Atoi(strings.ReplaceAll(Run("df -Bm | grep '^/dev/' | grep -v '/boot$' | awk '{ut += $3} END {print ut}'"), "\n", ""))
		usage = models.Usage{
			CPU: Run("top -bn1 | grep '^%Cpu' | cut -c 9- | xargs | awk '{printf(\"%.1f%%\"), $1 + $3}'"),
			Memory: fmt.Sprintf("%dGB", mem / 1000),
			Disk: fmt.Sprintf("%dGB", disk / 1000),
		}
	}
	return usage
}

func SendUsage() {
	sending := false
	StartStopUsageStream = make(chan bool)
	EndUsageStream = make(chan bool)
	defer close(StartStopUsageStream)
	for {
		select {
			case value := <-StartStopUsageStream:
				sending = value
			case <- EndUsageStream:
				return
			default:
				if sending {
					http.SocketSend("usage", GetUsage())
					time.Sleep(time.Second)
				}
		}
	}
}