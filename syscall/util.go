package syscall

import (
	"runtime"
	"strings"
)

type OperatingSys string

const (
	OSWindows OperatingSys = "Windows"
	OSDarwin  OperatingSys = "Mac OS"
	OSLinux   OperatingSys = "Linux"
	OSOther   OperatingSys = "Unknown"
)

var OperatingSystem OperatingSys

func SetOperatingSystem() {
	osname := runtime.GOOS
	switch {
	case strings.Contains(osname, "windows"):
		OperatingSystem = OSWindows
	case strings.Contains(osname, "darwin"):
		OperatingSystem = OSDarwin
	case strings.Contains(osname, "linux"):
		OperatingSystem = OSLinux
	default:
		OperatingSystem = OSOther

	}
}
