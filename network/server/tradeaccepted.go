package server

import "proxy/network"

type TradeAcceptedPacket struct {
	MyOffers    []bool
	OtherOffers []bool
}

func (t *TradeAcceptedPacket) Read(p *network.Packet) {
	myCount := p.ReadInt16()
	t.MyOffers = make([]bool, myCount)
	for i := 0; i < int(myCount); i++ {
		t.MyOffers[i] = p.ReadBool()
	}
	otherCount := p.ReadInt16()
	t.OtherOffers = make([]bool, otherCount)
	for i := 0; i < int(otherCount); i++ {
		t.OtherOffers[i] = p.ReadBool()
	}
}

func (t TradeAcceptedPacket) Write(p *network.Packet) {
	myCount := len(t.MyOffers)
	p.WriteInt16(int16(myCount))
	if myCount > 0 {
		for i := 0; i < myCount; i++ {
			p.WriteBool(t.MyOffers[i])
		}
	}
	otherCount := len(t.OtherOffers)
	p.WriteInt16(int16(otherCount))
	if otherCount > 0 {
		for i := 0; i < otherCount; i++ {
			p.WriteBool(t.OtherOffers[i])
		}
	}
}
