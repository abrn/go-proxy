package server

import "proxy/network"

type AllyShootPacket struct {
	BulletID      byte
	OwnerID       int32
	ContainerType int32
	Angle         float32
	Inspired      bool // whether the shot is affected by a Bard buff
}

func (a *AllyShootPacket) Read(p *network.Packet) {
	a.BulletID = p.ReadByte()
	a.OwnerID = p.ReadInt32()
	a.ContainerType = p.ReadInt32()
	a.Angle = p.ReadFloat()
	a.Inspired = p.ReadBool()
}

func (a AllyShootPacket) Write(p *network.Packet) {
	p.WriteByte(a.BulletID)
	p.WriteInt32(a.OwnerID)
	p.WriteInt32(a.ContainerType)
	p.WriteFloat(a.Angle)
	p.WriteBool(a.Inspired)
}
