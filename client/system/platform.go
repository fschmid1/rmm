package system

import (
	"fmt"
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
	return runtime.GOOS == "linux"
}

func CreateConfigurationPath() string {
	if IsLinux() {
		home, _ := os.UserHomeDir()
		if !FileOrFolderExists(fmt.Sprintf("%s/.fes-rmm", home)) {
			os.Mkdir(fmt.Sprintf("%s/.fes-rmm", home), 0775)
		}
		return fmt.Sprintf("%s/.fes-rmm/", home)
	}
	return ""
}
