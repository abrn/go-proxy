package client

import "proxy/network"

type GuildRemovePacket struct {
	Username string // of the player to remove
}

func (g *GuildRemovePacket) Read(p *network.Packet) {
	g.Username = p.ReadString()
}

func (g GuildRemovePacket) Write(p *network.Packet) {
	p.WriteString(g.Username)
}
