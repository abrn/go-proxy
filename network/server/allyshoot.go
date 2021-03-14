package server

import "proxy/network"

type AllyShootPacket struct {
	BulletID      byte
	OwnerID       int32
	ContainerType int32
	Angle         float32
	Inspired      bool // whether the shot is affected by a Bard buff
}

func (c *AllyShootPacket) Read(p *network.Packet) {
	c.BulletID = p.ReadByte()
	c.OwnerID = p.ReadInt32()
	c.ContainerType = p.ReadInt32()
	c.Angle = p.ReadFloat()
	c.Inspired = p.ReadBool()
}

func (c AllyShootPacket) Write(p *network.Packet) {
	p.WriteByte(c.BulletID)
	p.WriteInt32(c.OwnerID)
	p.WriteInt32(c.ContainerType)
	p.WriteFloat(c.Angle)
	p.WriteBool(c.Inspired)
}
