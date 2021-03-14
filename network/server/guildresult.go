package server

import "proxy/network"

type GuildResultPacket struct {
	Success bool
	Message string
}

func (g *GuildResultPacket) Read(p *network.Packet) {
	g.Success = p.ReadBool()
	g.Message = p.ReadString()
}

func (g GuildResultPacket) Write(p *network.Packet) {
	p.WriteBool(g.Success)
	p.WriteString(g.Message)
}
