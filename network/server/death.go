package server

import "proxy/network"

type DeathPacket struct {
	AccountID  string
	CharID     int32
	KilledBy   string
	ZombieType int32
	ZombieID   int32
	Unknown    int32
}

func (d *DeathPacket) Read(p *network.Packet) {
	d.AccountID = p.ReadString()
	d.CharID = p.ReadInt32()
	d.KilledBy = p.ReadString()
	d.ZombieType = p.ReadInt32()
	d.ZombieID = p.ReadInt32()
	d.Unknown = p.ReadInt32()
}

func (d DeathPacket) Write(p *network.Packet) {
	p.WriteString(d.AccountID)
	p.WriteInt32(d.CharID)
	p.WriteString(d.KilledBy)
	p.WriteInt32(d.ZombieType)
	p.WriteInt32(d.ZombieID)
	p.WriteInt32(d.Unknown)
}
