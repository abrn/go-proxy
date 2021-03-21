// +build windows

package syscall

import (
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
	RegSweepTime time.Duration   = 5 // how often the registry should be scanned in seconds
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

func GetLastServer() string {
	// todo: parse this server into a GameServer
	if RegData.LastServer != "" {
		return RegData.LastServer
	}
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
func grabSubKeyInt(keyName string) (uint32, error) {
	val, err := gowin32.GetRegValueDWORD(RegRootKey, RegSubKey, keyName)
	if err != nil {
		return -1, err
	}

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
