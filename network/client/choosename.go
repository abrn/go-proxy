package client

import "proxy/network"

type ChooseNamePacket struct {
	Username string
}

func (c *ChooseNamePacket) Read(p *network.GamePacket) {
	c.Username = p.ReadString()
}

func (c ChooseNamePacket) Write(p *network.GamePacket) {
	p.WriteString(c.Username)
}
