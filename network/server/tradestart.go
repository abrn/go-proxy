package server

import (
	"proxy/network"
	"proxy/network/data"
)

type TradeStartPacket struct {
	MyItems 	[]data.TradeItem
	OtherName 	string // the trade partner username
	OtherItems 	[]data.TradeItem
	ObjectID 	int32 // the clients object ID
}

func (t *TradeStartPacket) Read(p *network.Packet) {
	myCount := p.ReadInt16()
	if myCount > 0 {
		t.MyItems = make([]data.TradeItem, myCount)
		for i := 0; i < int(myCount); i++ {
			t.MyItems[i] = data.TradeItem{}
			t.MyItems[i].Read(p)
		}
	}
	t.OtherName = p.ReadString()
	otherCount := p.ReadInt16()
	if otherCount > 0 {
		t.OtherItems = make([]data.TradeItem, otherCount)
		for i := 0; i < int(otherCount); i++ {
			t.OtherItems[i] = data.TradeItem{}
			t.OtherItems[i].Read(p)
		}
	}
	t.ObjectID = p.ReadInt32()
}

func (t TradeStartPacket) Write(p *network.Packet) {
	myCount := len(t.MyItems)
	p.WriteInt16(int16(myCount))
	if myCount > 0 {
		for i := 0; i < myCount; i++ {
			t.MyItems[i].Write(p)
		}
	}
	p.WriteString(t.OtherName)
	otherCount := len(t.OtherItems)
	p.WriteInt16(int16(otherCount))
	if otherCount > 0 {
		for i := 0; i < otherCount; i++ {
			t.OtherItems[i].Write(p)
		}
	}
	p.WriteInt32(t.ObjectID)
}