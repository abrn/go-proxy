package client

import "proxy/network"

type GuildInvitePacket struct {
	Username string // of the player to invite
}

func (g *GuildInvitePacket) Read(p *network.GamePacket) {
	g.Username = p.ReadString()
}

func (g GuildInvitePacket) Write(p *network.GamePacket) {
	p.WriteString(g.Username)
}
