package client

import (
	"proxy/network"
	"proxy/network/data"
)

type ForgeRequestPacket struct {
	ObjectID int32
	Slots    []data.SlotObjectData
}

func (f *ForgeRequestPacket) Read(p *network.GamePacket) {
	f.ObjectID = p.ReadInt32()
	items := p.ReadInt32()
	if items > 0 {
		f.Slots = make([]data.SlotObjectData, items)
		for i := 0; i < int(items); i++ {
			f.Slots[i] = data.SlotObjectData{}
			f.Slots[i].Read(p)
		}
	}
}

func (f ForgeRequestPacket) Write(p *network.GamePacket) {
	p.WriteInt32(f.ObjectID)
	items := len(f.Slots)
	if items > 0 {
		p.WriteInt32(int32(items))
		for i := 0; i < items; i++ {
			f.Slots[i].Write(p)
		}
	}
}
