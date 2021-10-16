package syscall

import (
	"os"
	"proxy/log"
	"proxy/syscall/windows"
	"runtime"
	"strings"
)

type OperatingSys string

const (
	OSWindows OperatingSys = "Windows"
	OSDarwin  OperatingSys = "Mac OS"
)

var OperatingSystem OperatingSys

func StartSysHooks() {
	SetOperatingSystem()
	log.Logger.Info("Setting up %s launcher and game injection\n", OperatingSystem)
	if OperatingSystem == OSWindows {
		windows.StartWinMonitor()
	} else {
		// start mac stuff here
	}
}

// SetOperatingSystem - detect and set the machine OS
func SetOperatingSystem() {
	osname := runtime.GOOS
	switch {
	case strings.Contains(osname, "windows"):
		OperatingSystem = OSWindows
		log.Logger.Debug("Detected Windows OS with build %s\n", osname)
	case strings.Contains(osname, "darwin"):
		log.Logger.Debug("Detected Mac OS with build %s\n", osname)
		OperatingSystem = OSDarwin
	default:
		log.Logger.Error("Sorry, operating system '%s' is not supported... exiting\n", osname)
		os.Exit(1)
	}
}
