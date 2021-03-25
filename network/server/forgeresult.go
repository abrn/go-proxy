package server

import (
	"proxy/network"
	"proxy/network/data"
)

type ForgeResultPacket struct {
	Success bool
	Items   []data.SlotObjectData
}

func (f *ForgeResultPacket) Read(p *network.GamePacket) {
	f.Success = p.ReadBool()
	count := p.ReadInt16()
	f.Items = make([]data.SlotObjectData, count)
	for i := 0; i < int(count); i++ {
		f.Items[i] = data.SlotObjectData{}
		f.Items[i].Read(p)
	}
}

func (f ForgeResultPacket) Write(p *network.GamePacket) {
	p.WriteBool(f.Success)
	count := len(f.Items)
	if count <= 0 {
		return
	}
	for i := 0; i < count; i++ {
		f.Items[i].Write(p)
	}
}
