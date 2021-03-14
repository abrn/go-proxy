package client

import "proxy/network"

type GotoAckPacket struct {
	Time int32
}

func (g *GotoAckPacket) Read(p *network.Packet) {
	g.Time = p.ReadInt32()
}

func (g GotoAckPacket) Write(p *network.Packet) {
	p.WriteInt32(g.Time)
}
