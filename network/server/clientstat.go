package server

import "proxy/network"

type ClientStatPacket struct {
	Name  string
	Value int32
}

func (c *ClientStatPacket) Read(p *network.Packet) {
	c.Name = p.ReadString()
	c.Value = p.ReadInt32()
}

func (c ClientStatPacket) Write(p *network.Packet) {
	p.WriteString(c.Name)
	p.WriteInt32(c.Value)
}
