package client

import "proxy/network"

type GuildRemovePacket struct {
	Username string // of the player to remove
}

func (g *GuildRemovePacket) Read(p *network.GamePacket) {
	g.Username = p.ReadString()
}

func (g GuildRemovePacket) Write(p *network.GamePacket) {
	p.WriteString(g.Username)
}
