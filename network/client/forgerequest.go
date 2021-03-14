package client

import (
	"proxy/network"
	"proxy/network/data"
)

type ForgeRequestPacket struct {
	ObjectID int32
	Slots    []data.SlotObjectData
}

func (f *ForgeRequestPacket) Read(p *network.Packet) {
	f.ObjectID = p.ReadInt32()
	items := p.ReadInt32()
	f.Slots = make([]data.SlotObjectData, items)
	for i := 0; i < int(items); i++ {
		f.Slots[i] = data.SlotObjectData{}
		f.Slots[i].Read(p)
	}
}

func (f ForgeRequestPacket) Write(p *network.Packet) {
	p.WriteInt32(f.ObjectID)
	items := len(f.Slots)
	p.WriteInt32(int32(items))
	for i := 0; i < items; i++ {
		f.Slots[i].Write(p)
	}
}
