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
	MaxPlayers  int32 // todo: confirm this is right
	AllowTP     bool
	ShowDisplay bool
	ServerVer   string // confirm this is right
	OpenedTime  uint32 // confirm this is right
}

func (m *MapInfoPacket) Read(p *network.GamePacket) {
	m.Width = p.ReadInt32()
	m.Height = p.ReadInt32()
	m.Name = p.ReadString()
	m.DisplayName = p.ReadString()
	m.RealmName = p.ReadString()
	m.Difficulty = p.ReadInt32()
	m.FP = p.ReadUInt32()
	m.Background = p.ReadInt32()
	m.MaxPlayers = p.ReadInt32()
	m.AllowTP = p.ReadBool()
	m.ShowDisplay = p.ReadBool()
	m.ServerVer = p.ReadString()
	m.OpenedTime = p.ReadUInt32()
}

func (m MapInfoPacket) Write(p *network.GamePacket) {
	p.WriteInt32(m.Width)
	p.WriteInt32(m.Height)
	p.WriteString(m.Name)
	p.WriteString(m.DisplayName)
	p.WriteString(m.RealmName)
	p.WriteInt32(m.Difficulty)
	p.WriteUInt32(m.FP)
	p.WriteInt32(m.Background)
	p.WriteInt32(m.MaxPlayers)
	p.WriteBool(m.AllowTP)
	p.WriteBool(m.ShowDisplay)
	p.WriteString(m.ServerVer)
	p.WriteUInt32(m.OpenedTime)
}
