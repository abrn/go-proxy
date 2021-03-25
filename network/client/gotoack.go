package client

import "proxy/network"

type GotoAckPacket struct {
	Time int32
}

func (g *GotoAckPacket) Read(p *network.GamePacket) {
	g.Time = p.ReadInt32()
}

func (g GotoAckPacket) Write(p *network.GamePacket) {
	p.WriteInt32(g.Time)
}
