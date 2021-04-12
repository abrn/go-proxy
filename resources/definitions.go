package resources

import "encoding/xml"

/* // // CONSTANTS //  // // */

const (
	ResourceURL    string = "https://rotmg.network"
	AppspotURLprod string = "https://realmofthemadgod.com"
	AppspotURLtest string = "https://test.realmofthemadgod.com"

	UnityUserAgent string = "UnityPlayer/2019.4.9f1 (UnityWebRequest/1.0, libcurl/7.52.0-DEV)"
	UnityXVersion  string = "2019.4.9f1"
)

/* // // ACCOUNT // // // // */

// Account - represents a game account
type Account struct {
	Email string
	Password string
	Username string
	AccountID string
	Verified bool
	Banned bool
	Muted int32
	LastServer string
	TotalFame int32
	ForgeFire int32
	PetYardType byte
}

/* // // CHARACTER // // // */

// Character - represents a game character
type Character struct {

}

// CharactersURL - the appspot endpoint used to grab the character list, pets and vault
const CharactersURL string = "/char/list?muledump=true"

type AccountXML struct {
}

/* // // PET // // // // // */



/* // // ACCESSTOKEN // // */

type AccessTokenXML struct {
	Token string `xml:"AccessToken"`
	Timestamp string `xml:"AccessTokenTimestamp"`
	Expiration string `xml:"AccessTokenExpiration"`
}

// VerifyAccessTokenRequest - the POST data needed to verify an access token
type VerifyAccessTokenRequest struct {
	clientToken string
	accessToken string
	gameNet string
	gameNetUserID string
	playPlatform string
}

// VerifyAccessTokenURL - the appspot endpoint used to verify access tokens
const VerifyAccessTokenURL string = "/account/verifyAccessTokenClient"

/* // // SERVERS // // // */

// GameServer - info parsed from the appspot to represent a game server
type GameServer struct {
	Name     string
	Hostname string
	Usage    float32
}

// ServersRequest - the POST data needed to request the server list
type ServersRequest struct {
	guid          string
	password      string
	gameNet       string
	gameNetUserID string
	playPlatform  string
}

// ServersXML - struct for marshalling the servers response
type ServersXML struct {
	XMLName xml.Name    `xml:"Servers"`
	Servers []ServerXML `xml:"Server"`
}

// ServerXML - struct for marshalling the servers response
type ServerXML struct {
	XMLName xml.Name	`xml:"Server"`
	Name	string `xml:"Name"`
	DNS 	string `xml:"DNS"`
	Usage string `xml:"Usage"`
}

// ServersURL - the appspot endpoint to grab servers from
const ServersURL string = "/account/servers"








