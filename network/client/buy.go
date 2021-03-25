package client

import "proxy/network"

type BuyPacket struct {
	ObjectID int32
	Quantity int32
}

func (b *BuyPacket) Read(p *network.GamePacket) {
	b.ObjectID = p.ReadInt32()
	b.Quantity = p.ReadInt32()
}

func (b BuyPacket) Write(p *network.GamePacket) {
	p.WriteInt32(b.ObjectID)
	p.WriteInt32(b.Quantity)
}
