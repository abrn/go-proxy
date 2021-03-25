package client

import "proxy/network"

type CreateGuildPacket struct {
	GuildName string
}

func (c *CreateGuildPacket) Read(p *network.GamePacket) {
	c.GuildName = p.ReadString()
}

func (c CreateGuildPacket) Write(p *network.GamePacket) {
	p.WriteString(c.GuildName)
}
