package server

import "proxy/network"

type CreateSuccessPacket struct {
	ObjectID int32
	CharID   int32
	Unknown  string // probably clientToken
}

func (c *CreateSuccessPacket) Read(p *network.Packet) {
	c.ObjectID = p.ReadInt32()
	c.CharID = p.ReadInt32()
	c.Unknown = p.ReadString()
}

func (c CreateSuccessPacket) Write(p *network.Packet) {
	p.WriteInt32(c.ObjectID)
	p.WriteInt32(c.CharID)
	p.WriteString(c.Unknown)
}
