// +build windows

package syscall

import (
	"errors"
	"fmt"
	"github.com/winlabs/gowin32"
	"github.com/winlabs/gowin32/wrappers"
	"strings"
	"syscall"
)

var (
	LauncherRunning    bool = false
	GameRunning        bool = false
	CrashHandleRunning bool = false
)

type ProcHandles struct {
	LauncherPID    uint
	LauncherHandle syscall.Handle
	GamePID        uint
	GameHandle     syscall.Handle
	CrashPID       uint
}

func GrabRegistryData() {
	winHeight := make(chan uint32)
	winWidth := make(chan uint32)

	server, err := GetLastServer()
	height, err := GetWindowHeight()
	width, err := GetWindowWidth()

	select {
	case lastErr := <-err:
		fmt.Printf("Error getting registry key: %s\n", lastErr.Error())
	case LastServer := <-server:

	case height := <-winHeight:
		WinHeight = height
		fmt.Printf("Updated game window height: %d\n", height)
	}
}

func GetProcPIDs() (int, error) {
	procs, err := gowin32.GetProcesses()
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(procs); i++ {
		proc := procs[i]
		switch true {
		case proc.ExeFile == FileNameLauncher:
			LauncherPID = proc.ProcessID
			fmt.Printf("Found Exalt Launcher PID: %d - EXE: %s\n", proc.ProcessID, proc.ExeFile)
		case proc.ExeFile == FileNameGame:
			GamePID = proc.ProcessID
			fmt.Printf("Found Exalt Game PID: %d - EXE: %s\n", proc.ProcessID, proc.ExeFile)
		case proc.ExeFile == FileNameCrash:
			CrashPID = proc.ProcessID
			fmt.Printf("Found UnityCrashHandler PID: %d - EXE: %s\n\n", proc.ProcessID, proc.ExeFile)
		}
	}
	return 0, nil
}

// KillCrashHandler - check if the UnityCrashHandler proc is running and try kill it
func KillCrashHandler() (bool, error) {
	if Processes.CrashPID == 0 {
		return false, errors.New("crash handler not running (no PID)")
	}
	// todo: implement some fallbacks for errors
	running, err := gowin32.IsProcessRunning(Processes.CrashPID)
	if err != nil {
		return false, err
	} else if !running {
		return false, errors.New("crash handler not running (PID)")
	}
	killed := gowin32.KillProcess(Processes.CrashPID, 0)
	if killed != nil {
		return false, killed
	}
	return true, nil
}

// getCrashHandle - set the handle var for a running unitycrashhandler
func getCrashHandle() {

}

// getProcId
func getProcId(name string) (uint32, error) {

}
