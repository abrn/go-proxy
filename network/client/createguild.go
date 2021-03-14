package client

import "proxy/network"

type CreateGuildPacket struct {
	GuildName string
}

func (c *CreateGuildPacket) Read(p *network.Packet) {
	c.GuildName = p.ReadString()
}

func (c CreateGuildPacket) Write(p *network.Packet) {
	p.WriteString(c.GuildName)
}