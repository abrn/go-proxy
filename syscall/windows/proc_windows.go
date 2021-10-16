// +build windows

package windows

import (
	"errors"
	"fmt"
	"github.com/winlabs/gowin32"
	"proxy/log"
	"syscall"
	"time"
)

const (
	FileNameLauncher string = "RotMG Exalt Launcher.exe"
	FileNameGame     string = "RotMG Exalt.exe"
	FileNameCrash    string = "UnityCrashHandler64.exe"

	PROCESS_ALL_ACCESS  = syscall.STANDARD_RIGHTS_REQUIRED | syscall.SYNCHRONIZE | 0xfff // win32 opts we need for control
	MEM_COMMIT          = 0x001000                                                       // win32 opt for writing memory
	MEM_RESERVE         = 0x002000                                                       // win32 opt for creating new memory
	STILL_RUNNING       = 259                                                            // win32 status code for process == running
	PARENT_PROCESS_NAME = "cosmos.exe"                                                   // the name of the parent process used to
)

var (
	// LicenseVersion - the type of license the user has - either "free" or "premium"
	//                  ** keep this set as a string instead of 0/1 since you can bitflip that and crack **
	//                  ** use built in functions to check whether the license is premium instead **
	LicenseVersion string = "free"
	// MonitorRunning - goroutine to check if the launcher, game or crash handler are running and repeat
	//                  ** depending on what injection method we use, we might not even need to make the **
	//                  ** user close their game or launcher, but it might be needed for DRM **
	//                  ##todo: work out if we can inject and fulfill our requirements without making the user
	//                     quit the launcher or even game process
	MonitorRunning  bool = false // the monitor is in a running state, scanning for processes
	MonitorComplete bool = false // the monitor has started the launcher, stopped checking and started the control process
	// ControlRunning - goroutine to start the library injection process, opening the game/launcher
	//           ** tries to open the launcher or game as a child process, depending on whether the user used the faster **
	//           ** launch option or account manager login ** #todo: add faster launch and account manager options
	ControlRunning  bool = false // we have started trying to control processes and inject the dll
	ControlComplete bool = false // we have control of the launcher or the game itself
)

// ProcessStatus - the name of a process and whether it's currently running
type ProcessStatus struct {
	Name    string
	Running bool
}

// Process - represents a process we need to control and includes all the data we need to do it
type Process struct {
	Status     ProcessStatus
	Info       gowin32.ProcessInfo
	ModuleInfo gowin32.ModuleInfo
}

// CurrentProcesses - contains all running system processes and our handle to the launcher, game and crash reporter
type CurrentProcesses struct {
	All           *[]gowin32.ProcessInfo
	Game          *Process // our handle to the game process
	Launcher      *Process // our handle to the launcher process
	CrashReporter uint     // the PID of any running unity crash handler
}

// StartWinMonitor - this is the starting function for our windows process management
//      we will call the win32 process list to see if the game, launcher or Unity crash
//      handler are running
func StartWinMonitor() {
	MonitorTicker := time.NewTicker(2 * time.Second)
	StopSignal := make(chan struct{})
	log.Logger.Info("Starting WinProcessMonitor..\n")
	StartWinHooks()
	go func() {
		for {
			select {
			case <-MonitorTicker.C:
				MonitorRunning = true
				log.Logger.Trace("WinProcessMonitor: ticked\n")
				WinMonitorCheck()
			case <-StopSignal:
				MonitorRunning = false
				log.Logger.Trace("WinProcessMonitor: received stop signal\n")
				MonitorTicker.Stop()
				return
			}
		}
	}()
}

func StartWinHooks() {
	checkProcs := true
	procs, err := GetProcessList()
	if err != nil {
		log.Logger.Debug("Error when trying to scan system processes: %s\n", err)
		log.Logger.Trace("Retrying...\n")
		checkProcs = false
	}
	if checkProcs {
		flag := 0
		for i := 0; i < len(procs); i++ {
			switch {
			case procs[i].ExeFile == FileNameGame:
				sendRunningWarning("Exalt")
				flag++
			case procs[i].ExeFile == FileNameLauncher:
				sendRunningWarning("Exalt Launcher")
				flag++
			case procs[i].ExeFile == FileNameCrash:
				//
				flag++
			}
		}
		if flag == 0 {
			// start the exalt launcher process injected with a dll
			MonitorComplete = true
			kernel32 := syscall.MustLoadDLL("kernel32.dll")
		}
	} else {

	}
}

func WinMonitorCheck() {
	switch {
	case !MonitorRunning:
		StartWinMonitor()
	case MonitorRunning && !MonitorComplete:
		StartWinHooks()
	case MonitorRunning && MonitorComplete:
		// check we still have a handle on the file
		// check the crash handler is still killed
		return
	}
}

func GetProcessList() ([]gowin32.ProcessInfo, error) {
	procs, err := gowin32.GetProcesses()
	if err != nil {
		return nil, err
	}
	return procs, nil
}

func GetProcPIDs() (int, error) {
	procs, err := gowin32.GetProcesses()
	if err != nil {
		return 0, err
	}
	for i := 0; i < len(procs); i++ {
		proc := procs[i]
		switch true {
		case proc.ExeFile == windows.FileNameLauncher:
			LauncherPID = proc.ProcessID
			fmt.Printf("Found Exalt Launcher PID: %d - EXE: %s\n", proc.ProcessID, proc.ExeFile)
		case proc.ExeFile == windows.FileNameGame:
			GamePID = proc.ProcessID
			fmt.Printf("Found Exalt Game PID: %d - EXE: %s\n", proc.ProcessID, proc.ExeFile)
		case proc.ExeFile == windows.FileNameCrash:
			CrashPID = proc.ProcessID
			fmt.Printf("Found UnityCrashHandler PID: %d - EXE: %s\n\n", proc.ProcessID, proc.ExeFile)
		}
	}
	return 0, nil
}

func KillCrashReporter() {
	// if Processes.CrashPID == 0 {
	//     return false, errors.New("crash handler not running (no PID)")
	// }
	// // todo: implement some fallbacks for errors
	// running, err := gowin32.IsProcessRunning(Processes.CrashPID)
	// if err != nil {
	//     return false, err
	// } else if !running {
	//     return false, errors.New("crash handler not running (PID)")
	// }
	// killed := gowin32.KillProcess(Processes.CrashPID, 0)
	// if killed != nil {
	//     return false, killed
	// }
	// return true, nil
}

func IsLauncherRunning() bool {
	if RunningProcs == nil {
		return false
	}
}

func IsGameRunning() bool {

}

func scanProccessList() {

}

func sendRunningWarning(msg string) {
	log.Logger.Warn("Detected %s already running, please close it and let the proxy launch the game for you\n")
	log.Logger.Warn("Checking again in 5 seconds...\n")
}
