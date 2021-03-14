package server

import "proxy/network"

type TradeRequestedPacket struct {
	Name string
}

func (t *TradeRequestedPacket) Read(p *network.Packet) {
	t.Name = p.ReadString()
}

func (t TradeRequestedPacket) Write(p *network.Packet) {
	p.WriteString(t.Name)
}
