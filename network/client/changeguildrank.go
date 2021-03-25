package client

import "proxy/network"

type ChangeGuildRankPacket struct {
	Name string
	Rank GuildRankType
}

type GuildRankType byte

const (
	GuildInitiate GuildRankType = 0
	GuildMember   GuildRankType = 10
	GuildOfficer  GuildRankType = 20
	GuildLeader   GuildRankType = 30
	GuildFounder  GuildRankType = 40
)

func (c *ChangeGuildRankPacket) Read(p *network.GamePacket) {
	c.Name = p.ReadString()
	c.Rank = GuildRankType(p.ReadByte())
}

func (c ChangeGuildRankPacket) Write(p *network.GamePacket) {
	p.WriteString(c.Name)
	p.WriteByte(byte(c.Rank))
}
