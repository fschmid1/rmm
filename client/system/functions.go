package system

import (
	"fmt"

	"festech.de/rmm/client/vars"
)

func KillProcess(name string) string {
	if !vars.Configuration.AllowKill {
		return "Process kill is not allowed on this device"
	}
	if IsLinux() {
		return Run("pkill " + name)
	}
	return ""
}

func GetProcessList() string {
	if !vars.Configuration.AllowProcessList {
		return "Process list is not allowed on this device"
	}
	if IsLinux() {
		return Run("ps aux")
	}
	return ""
}

func GetServiceList() string {
	if !vars.Configuration.AllowServiceList {
		return "Service list is not allowed on this device"
	}
	if IsLinux() {
		return Run("service --status-all")
	}
	return ""
}

func GetServiceLogs(service string) string {
	if !vars.Configuration.AllowServiceLogs {
		return "Service logs is not allowed on this device"
	}
	if IsLinux() {
		return Run(fmt.Sprintf("journalctl -u %s -e --no-pager", service))
	}
	return ""
}

func GetServiceStatus(service string) string {
	if !vars.Configuration.AllowServiceStatus {
		return "Service status is not allowed on this device"
	}
	if IsLinux() {
		return Run(fmt.Sprintf("service %s status", service))
	}
	return ""
}

func StartService(service string) string {
	if !vars.Configuration.AllowServiceStart {
		return "Service start is not allowed on this device"
	}
	if IsLinux() {
		return Run(fmt.Sprintf("service %s start", service))
	}
	return ""
}

func StopService(service string) string {
	if !vars.Configuration.AllowServiceStop {
		return "Service stop is not allowed on this device"
	}
	if IsLinux() {
		return Run(fmt.Sprintf("service %s stop", service))
	}
	return ""
}

func RestartService(service string) string {
	if !vars.Configuration.AllowServiceRestart {
		return "Service restart is not allowed on this device"
	}
	if IsLinux() {
		return Run(fmt.Sprintf("service %s restart", service))
	}
	return ""
}

func Reboot() string {
	if !vars.Configuration.AllowReboot {
		return "Reboot is not allowed on this device"
	}
	if IsLinux() {
		return Run("reboot")
	}
	return ""
}

func Shutdown() string {
	if !vars.Configuration.AllowShutdown {
		return "Shutdown is not allowed on this device"
	}
	if IsLinux() {
		return Run("shutdown -h 10")
	}
	return ""
}
