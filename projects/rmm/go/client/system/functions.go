package system

import (
	"fmt"

	"github.com/fes111/rmm/projects/rmm/go/client/vars"
)

func KillProcess(name string) (string, string) {
	if !vars.Configuration.Allow.Kill {
		return "", "Process kill is not allowed on this device"
	}
	if IsLinux() {
		return Run("pkill " + name), ""
	}
	return "", ""
}

func GetProcessList() (string, string) {
	if !vars.Configuration.Allow.ProcessList {
		return "", "Process list is not allowed on this device"
	}
	if IsLinux() {
		return Run("ps aux"), ""
	}
	return "", ""
}

func GetServiceList() (string, string) {
	if !vars.Configuration.Allow.ServiceList {
		return "", "Service list is not allowed on this device"
	}
	if IsLinux() {
		return Run("service --status-all"), ""
	}
	return "", ""
}

func GetServiceLogs(service string) (string, string) {
	if !vars.Configuration.Allow.ServiceLogs {
		return "", "Service logs are not allowed on this device"
	}
	if IsLinux() {
		return Run(fmt.Sprintf("journalctl -u %s -e --no-pager", service)), ""
	}
	return "", ""
}

func GetServiceStatus(service string) (string, string) {
	if !vars.Configuration.Allow.ServiceStatus {
		return "", "Service status is not allowed on this device"
	}
	if IsLinux() {
		return Run(fmt.Sprintf("service %s status", service)), ""
	}
	return "", ""
}

func StartService(service string) (string, string) {
	if !vars.Configuration.Allow.ServiceStart {
		return "", "Service start is not allowed on this device"
	}
	if IsLinux() {
		return Run(fmt.Sprintf("systemctl start %s", service)), ""
	}
	return "", ""
}

func StopService(service string) (string, string) {
	if !vars.Configuration.Allow.ServiceStop {
		return "", "Service stop is not allowed on this device"
	}
	if IsLinux() {
		return Run(fmt.Sprintf("systemctl stop %s", service)), ""
	}
	return "", ""
}

func RestartService(service string) (string, string) {
	if !vars.Configuration.Allow.ServiceRestart {
		return "", "Service restart is not allowed on this device"
	}
	if IsLinux() {
		return Run(fmt.Sprintf("systemctl restart %s", service)), ""
	}
	return "", ""
}

func Reboot() (string, string) {
	if !vars.Configuration.Allow.Reboot {
		return "", "Reboot is not allowed on this device"
	}
	if IsLinux() {
		return Run("reboot"), ""
	}
	return "", ""
}

func Shutdown() (string, string) {
	if !vars.Configuration.Allow.Shutdown {
		return "", "Shutdown is not allowed on this device"
	}
	if IsLinux() {
		return Run("shutdown -h 10"), ""
	}
	return "", ""
}
