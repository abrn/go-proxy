package server

import "proxy/network"

type GuildResultPacket struct {
	Success bool
	Message string
}

func (g *GuildResultPacket) Read(p *network.GamePacket) {
	g.Success = p.ReadBool()
	g.Message = p.ReadString()
}

func (g GuildResultPacket) Write(p *network.GamePacket) {
	p.WriteBool(g.Success)
	p.WriteString(g.Message)
}
