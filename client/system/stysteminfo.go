package system

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Run(cmd string) string {
	out := exec.Command("bash", "-c", cmd)
	stdout, _ := out.Output()
	return string(stdout)
}

func GetProcessList() string {
	if IsLinux() {
		return Run("ps aux")
	}
	return ""
}

func KillProcess(name string) string {
	if IsLinux() {
		return Run("pkill " + name)
	}
	return ""
}

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}
	return hostname
}

func GetMacAddress() string {
	ifas, err := net.Interfaces()
	if err != nil {
		return ""
	}
	return ifas[1].HardwareAddr.String()
}

func GetOS() string {
	var os string
	if IsLinux() {
		os = strings.ReplaceAll(Run("uname -osri"), "\n", "")
	}
	return os
}

func GetCores() int {
	var cores int
	var _ error
	if IsLinux() {
		cores, _ = strconv.Atoi(strings.ReplaceAll(Run("grep \"^processor\" /proc/cpuinfo | wc -l"), "\n", ""))
	}
	return cores
}

func GetMemory() string {
	var memory string
	var _ error
	if IsLinux() {
		temp, _ := strconv.Atoi(strings.ReplaceAll(Run("free -m | awk '$1 == \"Mem:\" {print $2}'"), "\n", ""))
		memory = fmt.Sprintf("%dGB", (temp / 1000))
	}
	return memory
}

func GetDisk() string {
	var disk string
	var _ error
	if IsLinux() {
		disk = strings.ReplaceAll(Run("df -Bg | grep '^/dev/' | grep -v '/boot$' | awk '{ft += $2} END {print ft}'"), "\n", "") + "GB"
	}
	return disk
}
