package client

import "proxy/network"

type ChangeGuildRankPacket struct {
	Name 	string
	Rank	int32
}

func (c *ChangeGuildRankPacket) Read(p *network.Packet) {
	c.Name = p.ReadString()
	c.Rank = p.ReadInt32()
}

func (c ChangeGuildRankPacket) Write(p *network.Packet) {
	p.WriteString(c.Name)
	p.WriteInt32(c.Rank)
}