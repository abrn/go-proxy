package client

import "proxy/network"

type ClaimDailyRewardPacket struct {
	ClaimKey string
	Type     string
}

func (c *ClaimDailyRewardPacket) Read(p *network.Packet) {
	c.ClaimKey = p.ReadString()
	c.Type = p.ReadString()
}

func (c ClaimDailyRewardPacket) Write(p *network.Packet) {
	p.WriteString(c.ClaimKey)
	p.WriteString(c.Type)
}
