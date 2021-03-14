package server

import "proxy/network"

type MapInfoPacket struct {
	Width       int32
	Height      int32
	Name        string
	DisplayName string
	RealmName   string
	Difficulty  int32
	FP          uint32
	Background  int32
	AllowTP     bool
	ShowDisplay bool
	MaxPlayers  int16
	OpenedTime  uint32
	//ServerVer 	string
}

func (m *MapInfoPacket) Read(p *network.Packet) {
	m.Width = p.ReadInt32()
	m.Height = p.ReadInt32()
	m.Name = p.ReadString()
	m.DisplayName = p.ReadString()
	m.RealmName = p.ReadString()
	m.Difficulty = p.ReadInt32()
	m.FP = p.ReadUInt32()
	m.Background = p.ReadInt32()
	m.AllowTP = p.ReadBool()
	m.ShowDisplay = p.ReadBool()
	m.MaxPlayers = p.ReadInt16()
	m.OpenedTime = p.ReadUInt32()
}

// todo: MAPINFO add write function
