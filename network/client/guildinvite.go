package client

import "proxy/network"

type GuildInvitePacket struct {
	Username string // of the player to invite
}

func (g *GuildInvitePacket) Read(p *network.Packet) {
	g.Username = p.ReadString()
}

func (g GuildInvitePacket) Write(p *network.Packet) {
	p.WriteString(g.Username)
}
