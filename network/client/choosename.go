package client

import "proxy/network"

type ChooseNamePacket struct {
	Username string
}

func (c *ChooseNamePacket) Read(p *network.Packet) {
	c.Username = p.ReadString()
}

func (c ChooseNamePacket) Write(p *network.Packet) {
	p.WriteString(c.Username)
}
