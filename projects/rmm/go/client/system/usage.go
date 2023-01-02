package system

import (
	"strconv"
	"strings"
	"time"

	"github.com/fes111/rmm/projects/rmm/go/client/vars"
	"github.com/fes111/rmm/projects/rmm/go/lib/models"

	"github.com/fes111/rmm/projects/rmm/go/client/http"
)

var StartStopUsageStream chan bool
var EndUsageStream chan bool

func GetUsage() models.Usage {
	usage := models.Usage{}
	if IsLinux() {
		mem, _ := strconv.ParseFloat(strings.ReplaceAll(Run("free -m | awk '$1 == \"Mem:\" {print $3}'"), "\n", ""), 64)
		cpu, _ := strconv.ParseFloat(strings.ReplaceAll(Run("top -bn1 | grep '^%Cpu' | cut -c 9- | xargs | awk '{printf(\"%.1f\"), $1 + $3}'"), ",", "."), 64)
		usage = models.Usage{
			CPU:    cpu,
			Memory: (mem / 1024) / vars.Device.SystemInfo.MemoryTotal * 100,
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
		case <-EndUsageStream:
			return
		default:
			if sending {
				http.SocketSend("usage", GetUsage())
				time.Sleep(time.Second)
			}
		}
	}
}
