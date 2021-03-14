package client

import "proxy/network"

type ProjectileAckPacket struct {
	OwnerID int32
	Unknown int16
}

func (c *ProjectileAckPacket) Read(p *network.Packet) {
	c.OwnerID = p.ReadInt32()
	c.Unknown = p.ReadInt16()
}

func (c ProjectileAckPacket) Write(p *network.Packet) {
	p.WriteInt32(c.OwnerID)
	p.WriteInt16(c.Unknown)
}

// todo: PROJECTILACK add pretty print function