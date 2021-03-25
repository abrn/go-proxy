package server

import "proxy/network"

type TradeRequestedPacket struct {
	Name string
}

func (t *TradeRequestedPacket) Read(p *network.GamePacket) {
	t.Name = p.ReadString()
}

func (t TradeRequestedPacket) Write(p *network.GamePacket) {
	p.WriteString(t.Name)
}
