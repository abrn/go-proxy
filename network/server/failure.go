package server

import "proxy/network"

type FailurePacket struct {
	ID      int32
	Message string
}

func (c *FailurePacket) Read(p *network.Packet) {
	c.ID = p.ReadInt32()
	c.Message = p.ReadString()
}

func (c FailurePacket) Write(p *network.Packet) {
	p.WriteInt32(c.ID)
	p.WriteString(c.Message)
}
