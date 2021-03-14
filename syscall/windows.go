package syscall

import (
	"errors"
	"fmt"
	"github.com/winlabs/gowin32"
	"github.com/winlabs/gowin32/wrappers"
	"strings"
	"syscall"
)

const (
	RegRootKey 			gowin32.RegRoot = wrappers.HKEY_CURRENT_USER
	RegSubKey			string = "Software\\DECA Live Operations GmbH\\RotMGExalt"

	FileNameLauncher 	string = "RotMG Exalt Launcher.exe"
	FileNameGame 		string = "RotMG Exalt.exe"
	FileNameCrash 		string = "UnityCrashHandler64.exe"
)

var (
	LauncherRunning		bool = false
	GameRunning 		bool = false
	CrashHandleRunning	bool = false

	Processes 			ProcHandles
	RegData 			RegistryData
	RegKeys 			[]string
)

type ProcHandles struct {
	LauncherPID 		uint
	LauncherHandle		syscall.Handle
	GamePID 			uint
	GameHandle 			syscall.Handle
	CrashPID 			uint
}

type RegistryData struct {
	LastServer 			string
	BestServer 			string
	GUID 				string
	WinHeight 			uint32
	WinWidth 			uint32
	FullScreen 			bool
}

func GrabRegistryData() {
	winHeight := make(chan uint32)
	winWidth := make(chan uint32)

	server, err := GetLastServer()
	height, err := GetWindowHeight()
	width, err := GetWindowWidth()

	select {
	case lastErr := <- err:
		fmt.Printf("Error getting registry key: %s\n", lastErr.Error())
		case LastServer := <- server:

		case height := <- winHeight:
			WinHeight = height
			fmt.Printf("Updated game window height: %d\n", height)
	}
}

func GetLastServer() (chan string, chan error) {
	keyName := "preferredServer_h3991771845"
	resChan := make(chan string)
	errChan := make(chan error)
	go func() {

	}()
	key, err := gowin32.OpenRegKey(RegRootKey, RegSubKey, false)
	if err != nil {
		return "", err
	}
	bin, err := key.GetValueBinary(keyName)
	if err != nil {
		return "", err
	}
	return string(bin), err
}

func GetWindowHeight() (chan uint32, chan error) {
	resChan := make(chan uint32)
	errChan := make(chan error)
	go func() {
		val, err := gowin32.GetRegValueDWORD(RegRootKey, RegSubKey, "screenHeight_h4096606934")
		if err != nil {
			errChan <- err
		}
		resChan <- val
	}()
	return resChan, errChan
}

func GetWindowWidth() (chan uint32, chan error) {
	resChan := make(chan uint32)
	errChan := make(chan error)
	go func() {
		val, err := gowin32.GetRegValueDWORD(RegRootKey, RegSubKey, "screeWidth_h3938008705")
		if err != nil {
			errChan <- err
		}
		resChan <- val
	}()
	return resChan, errChan
}

func GetExaltGUID() (string, error) {
	keyName := "guid_h2087642266"
	key, err := gowin32.OpenRegKey(RegRootKey, RegSubKey, false)
	if err != nil {
		return "", err
	}
	bin, err := key.GetValueBinary(keyName)
	if err != nil {
		return "", err
	}
	return string(bin), err
}

// getSubKeys - grab every key name under the exalt root key
func getSubKeys() (chan []string, chan error) {
	resChan := make(chan []string)
	errChan := make(chan error)
	go func() {
		key, err := gowin32.OpenRegKey(RegRootKey, RegSubKey, false)
		if err != nil {
			errChan <- err
		} else {
			keys, err := key.GetSubKeys()
			if err != nil {
				errChan <- err
			}
			RegKeys = keys
			resChan <- keys
		}
	}()
	return resChan, errChan
}

// parseSubKeys - parse every exalt registry key name and grab only the ones we need
func parseSubKeys(keys []string) {
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		switch true {
		case strings.HasPrefix(k ,"preferredServer"):
			server, err := grabSubKeyStr(k)
		case strings.HasPrefix(k, "screenHeight"):
			height, err := grabSubKeyInt(k)
		case strings.HasPrefix(k, "screeWidth"):
			width, err := grabSubKeyStr(k)
		case strings.HasPrefix(k, "guid"):
			guid, err := grabSubKeyStr(k)
		}
	}

}

// grabSubKeyInt - grab a single integer subkey value from the exalt registry root
func grabSubKeyInt(keyName string) (uint32, error) {

}

// grabSubKeyStr - grab a single string subkey (binary first) value from the exalt registry root
func grabSubKeyStr(keyName string) (string, error) {
	key, err := gowin32.OpenRegKey(RegRootKey, RegSubKey, false)
	if err != nil {
		return "", err
	}
	bin, err := key.GetValueBinary(keyName)
	if err != nil {
		return "", err
	}
	return string(bin), err
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

// KillCrashHandle - check if the UnityCrashHandler proc is running and try kill it
func KillCrashHandle() (bool, error) {
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