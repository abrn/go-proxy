package server

import "proxy/network"

type TradeChangedPacket struct {
	OtherOffers []bool // the offers of your trade partner
}

func (t *TradeChangedPacket) Read(p *network.Packet) {
	count := p.ReadInt16()
	t.OtherOffers = make([]bool, count)
	for i := 0; i < int(count); i++ {
		t.OtherOffers[i] = p.ReadBool()
	}
}

func (t TradeChangedPacket) Write(p *network.Packet) {
	count := len(t.OtherOffers)
	p.WriteInt16(int16(count))
	if count > 0 {
		for i := 0; i < count; i++ {
			p.WriteBool(t.OtherOffers[i])
		}
	}
}