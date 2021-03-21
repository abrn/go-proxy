package client

import "proxy/network"

type ChangeAllyShootPacket struct {
	Unknown int32
}

func (c *ChangeAllyShootPacket) Read(p *network.Packet) {
	c.Unknown = p.ReadInt32()
}

func (c ChangeAllyShootPacket) Write(p *network.Packet) {
	p.WriteInt32(c.Unknown)
}
