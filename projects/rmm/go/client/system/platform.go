package system

import (
	"os"
	"runtime"
)

func FileOrFolderExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func IsMac() bool {
	return runtime.GOOS == "mac"
}

func CreateConfigurationPath() string {
	if IsLinux() || IsMac() {
		if !FileOrFolderExists("/etc/fes-rmm") {
			os.Mkdir("/etc/fes-rmm", 0775)
		}
		return "/etc/fes-rmm/"
	}
	return ""
}
