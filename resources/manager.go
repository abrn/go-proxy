package resources

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
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

type GameServer struct {
	Name     string
	Hostname string
	Usage    float32
}

func GetServers() *ServerCache {
	if serverCache != (&ServerCache{}) {
		return serverCache
	}

}

// UnityGET - send a GET request that mimics the unity client
func UnityGET(res string) (string, error) {
	client := &http.Client{}
	request := makeUnityRequest(res, "GET", nil)

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
}

// UnityGET - send a POST request that mimics the unity client
func UnityPOST(res string, data *url.Values) (string, error) {
	client := &http.Client{}
	request := makeUnityRequest(res, "POST", data)

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
}

// grabServers - request the server endpoint and try to parse the list
func grabServers() ([]GameServer, error) {
	path := "/account/servers"
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

// makeUnityRequest - create a default unity-like request
func makeUnityRequest(path string, method string, values *url.Values) *http.Request {
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
	request.Header.Add("User-Agent", unityUserAgent)
	request.Header.Add("X-Unity-Version", unityXVersion)
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
