package client

import "proxy/network"

type ClaimDailyRewardPacket struct {
	ClaimKey string
	Type     string
}

func (c *ClaimDailyRewardPacket) Read(p *network.GamePacket) {
	c.ClaimKey = p.ReadString()
	c.Type = p.ReadString()
}

func (c ClaimDailyRewardPacket) Write(p *network.GamePacket) {
	p.WriteString(c.ClaimKey)
	p.WriteString(c.Type)
}
