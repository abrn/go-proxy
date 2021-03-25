package resources

import "encoding/xml"

/* // // ACCOUNT // // */

// Account - represents a game account
type Account struct {
	Email string
	Password string
	AccountID string

}

/* // // CHARACTER // // */

// Character - represents a game character
type Character struct {

}

type AccountXML struct {
}

/* // // SERVERS // // */

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

// ServersURL - the appspot endpoint to grab servers from
const ServersURL string = "/account/servers"
