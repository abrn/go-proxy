package server

import "proxy/network"

type ClientStatPacket struct {
	Name  string
	Value int32
}

func (c *ClientStatPacket) Read(p *network.GamePacket) {
	c.Name = p.ReadString()
	c.Value = p.ReadInt32()
}

func (c ClientStatPacket) Write(p *network.GamePacket) {
	p.WriteString(c.Name)
	p.WriteInt32(c.Value)
}
