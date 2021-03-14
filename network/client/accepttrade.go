package client

import "proxy/network"

type AcceptTradePacket struct {
	MyOffers    []bool
	OtherOffers []bool
}

func (c *AcceptTradePacket) Read(p *network.Packet) {
	myCount := p.ReadInt16()
	c.MyOffers = make([]bool, myCount)
	for i := 0; i < int(myCount); i++ {
		c.MyOffers[i] = p.ReadBool()
	}
	otherCount := p.ReadInt16()
	c.OtherOffers = make([]bool, otherCount)
	for i := 0; i < int(otherCount); i++ {
		c.OtherOffers[i] = p.ReadBool()
	}
}

func (c AcceptTradePacket) Write(p *network.Packet) {
	myCount := len(c.MyOffers)
	p.WriteInt16(int16(myCount))
	if myCount > 0 {
		for i := 0; i < myCount; i++ {
			p.WriteBool(c.MyOffers[i])
		}
	}
	otherCount := len(c.OtherOffers)
	p.WriteInt16(int16(otherCount))
	if otherCount > 0 {
		for i := 0; i < otherCount; i++ {
			p.WriteBool(c.OtherOffers[i])
		}
	}
}
