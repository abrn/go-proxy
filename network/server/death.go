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

// todo: DEATH add write function
