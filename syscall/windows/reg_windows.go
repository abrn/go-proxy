// +build windows

package windows

import (
	"fmt"
	"github.com/winlabs/gowin32"
	"github.com/winlabs/gowin32/wrappers"
	"strings"
	"time"
)

type RegistryData struct {
	LastServer  string // last used server
	BestServer  string // server with lowest RTT
	GUID        string // saved account email
	Password    string // saved account password
	WinHeight   uint32 // exalt window height
	WinWidth    uint32 // exalt window width
	FullScreen  bool   // whether exalt is in fullscreen mode
	CharacterID int16  // last used character ID
}

const (
	RegRootKey   gowin32.RegRoot = wrappers.HKEY_CURRENT_USER
	RegSubKey    string          = "Software\\DECA Live Operations GmbH\\RotMGExalt"
)

var (
	RegData RegistryData // current registry data
	RegKeys []string     // all exalt registry key names (randomized)
)

func StartRegScan() {
	// get all registry key names

	// iterate over key names to find the ones we need
	// set these keys to exported variables in this file

	// if we can't get a key
	go func() {
		for {
			time.Sleep(RegSweepTime * time.Second)
		}
	}()
}

func GrabRegistryData() {
	winHeight := make(chan uint32)
	winWidth := make(chan uint32)

	server := GetLastServer()
	height := GetWindowHeight()
	width := GetWindowWidth()

	select {
	case lastErr := <-err:
		fmt.Printf("Error getting registry key: %s\n", lastErr.Error())
	case LastServer := <-server:

	case height := <-winHeight:
		WinHeight = height
		fmt.Printf("Updated game window height: %d\n", height)
	}
}

func GetLastServer() string {
	// todo: parse this server into a GameServer
	if RegData.LastServer != "" {
		return RegData.LastServer
	}
	key, err := gowin32.OpenRegKey(RegRootKey, RegSubKey, false)
	if err != nil {
		return ""
	}
	bin, err := key.GetValueBinary(keyName)
	if err != nil {
		return ""
	}
	return string(bin)
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
//				   ** depending on the hardware and timing, each registry key has a random **
// 				   ** prefix, so we need to loop through them all and check if the name of **
//				   ** the registry key contains what we need **
// @params keys = a slice of all registry key names under the exalt registry root
func parseSubKeys(keys []string) {
	for i := 0; i < len(keys); i++ {
		k := keys[i]
		switch true {
		case strings.HasPrefix(k, "preferredServer"):
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
// 				   ** returned as an unsigned 32bit int (DWORD) no matter how large **
// @params keyName = the name of the registry key to grab
func grabSubKeyInt(keyName string) (uint32, error) {
	val, err := gowin32.GetRegValueDWORD(RegRootKey, RegSubKey, keyName)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// grabSubKeyStr - grab a single string subkey value from the exalt registry root
//				   ** exalt strings are stored as binary so we have to grab the value **
// 				   ** as a binary first, then convert it back to a string **
// @params keyName = the name of the registry key to grab
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
