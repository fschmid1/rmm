package system

import (
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func Run(cmd string) string {
	out := exec.Command("bash", "-c", cmd)
	timer := time.AfterFunc(time.Second*4, func() {
		out.Process.Kill()
	})

	defer timer.Stop()
	stdout, _ := out.CombinedOutput()
	return string(stdout)
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

func GetIP() string {
	var ip string
	if IsLinux() {
		ip = strings.SplitN(Run("hostname -I"), " ", 2)[0]
	}
	return ip
}

func GetCores() int {
	var cores int
	var _ error
	if IsLinux() {
		cores, _ = strconv.Atoi(strings.ReplaceAll(Run("grep \"^processor\" /proc/cpuinfo | wc -l"), "\n", ""))
	}
	return cores
}

func GetMemoryUsed() float64 {
	var memory float64
	var _ error
	if IsLinux() {
		temp, _ := strconv.ParseFloat(strings.ReplaceAll(Run("free -m | awk '$1 == \"Mem:\" {print $3}'"), "\n", ""), 64)
		memory = temp / 1024
	}
	return memory
}

func GetMemoryTotal() float64 {
	var memory float64
	var _ error
	if IsLinux() {
		temp, _ := strconv.ParseFloat(strings.ReplaceAll(Run("free -m | awk '$1 == \"Mem:\" {print $2}'"), "\n", ""), 64)
		memory = temp / 1024
	}
	return memory
}

func GetDiskTotal() float64 {
	var disk float64
	var _ error
	if IsLinux() {
		temp, _ := strconv.ParseFloat(strings.ReplaceAll(Run("df -Bg | grep '^/dev/' | grep -v '/boot$' | awk '{ft += $2} END {print ft}'"), "\n", ""), 64)
		disk = temp
	}
	return disk
}

func GetDiskUsed() float64 {
	var disk float64
	var _ error
	if IsLinux() {
		temp, _ := strconv.ParseFloat(strings.ReplaceAll(Run("df -Bg | grep '^/dev/' | grep -v '/boot$' | awk '{ft += $3} END {print ft}'"), "\n", ""), 64)
		disk = temp
	}
	return disk
}

func GetCPU() string {
	var cpu string
	if IsLinux() {
		cpu = strings.ReplaceAll(Run("cat /proc/cpuinfo | grep 'model name' | uniq | awk -F: '{print $2}'"), "\n", "")
	}
	return cpu
}

func GetGPU() string {
	var gpu string
	if IsLinux() {
		gpu = strings.ReplaceAll(Run("lspci | grep VGA | awk -F: '{print $3}'"), "\n", "")
	}
	return gpu
}
