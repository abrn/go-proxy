package resources

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"proxy/log"
	"strings"
)

const (
	resourceURL    string = "https://rotmg.network"
	appspotURLprod string = "https://realmofthemadgod.com"
	appspotURLtest string = "https://test.realmofthemadgod.com"

	unityUserAgent string = "UnityPlayer/2019.4.9f1 (UnityWebRequest/1.0, libcurl/7.52.0-DEV)"
	unityXVersion  string = "2019.4.9f1"
)

var serverCache *ServerCache

func GetServers() *ServerCache {
	if serverCache != nil {
		return serverCache
	}
	// attempt to load the servers from cache
	cache, err := loadServers()
	if err != nil {
		log.Logger.Warn("Could not find a server cache file.. downloading server list\n")
	} else {
		log.Logger.Info("Loaded cached server list!\n")
		return cache
	}
	// download a new server list

	// todo: check if we found a guid + pass from exalt

}

// GetExecutablePath - returns the path where the running binary is located
func GetExecutablePath() string {
	ex, _ := os.Executable()
	path := filepath.Dir(ex)
	return path
}

// UnityGET - send a GET request that mimics the unity client
func UnityGET(res string) (string, error) {
	client := &http.Client{}
	request := makeRequest(res, "GET", nil, true)

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// UnityGET - send a POST request that mimics the unity client
func UnityPOST(res string, data *url.Values) (string, error) {
	client := &http.Client{}
	request := makeRequest(res, "POST", data, true)

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// loadServers - try loading the server cache file
func loadServers() (*ServerCache, error) {
	path := GetExecutablePath()
	f, err := ioutil.ReadFile(filepath.Join(path, "resources", "cache.servers"))
	if err != nil {
		return nil, err
	}
	cache := &ServerCache{}
	if err := proto.Unmarshal(f, cache); err != nil {
		return nil, err
	}
	return cache, nil
}

// grabServers - request the server endpoint and try to parse the list
func grabServers() ([]GameServer, error) {
	values := &url.Values{}
	values.Set("guid", "NQPWXHHQFCDRDIFJMMXDCDHJZM@ggnetwork.xyz") //todo set guid and pass
	values.Set("password", "RowJow100998")

	resp, err := UnityPOST(path, values)

	var serversResponse struct {
		TagName xml.Name `xml:"Servers"`
	}
	var singleServer struct {
		TagName xml.Name `xml:"Server"`
	}
}

// makeRequest - create a default http request with the option for Unity headers
func makeRequest(path string, method string, values *url.Values, unity bool) *http.Request {
	// todo: check if prod or testing
	uri, _ := url.ParseRequestURI(appspotURLprod)
	uri.Path = path

	var request *http.Request
	if method == "POST" {
		var buf *strings.Reader
		if values != nil {
			buf = strings.NewReader(values.Encode())
		}
		request, _ = http.NewRequest(http.MethodPost, uri.String(), buf)
	} else {
		request, _ = http.NewRequest(http.MethodGet, uri.String(), nil)
	}
	// add the headers to mimic the Unity client
	if unity {
		request.Header.Add("User-Agent", unityUserAgent)
		request.Header.Add("X-Unity-Version", unityXVersion)
	}
	return request
}

func parseAppspotResponse(resp http.Response) (result bool, err error) {
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("bad request")
	}
	ba, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, nil
	}
	str := string(ba)
	if len(str) < 11 {
		return false, errors.New("bad response: " + str)
	}
	if strings.Contains(str, "<Error>") {

	}
}
