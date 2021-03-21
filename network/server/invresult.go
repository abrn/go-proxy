package server

import (
	"proxy/network"
	"proxy/network/data"
)

type InvResultPacket struct {
	Success bool
	OldSlot data.SlotObjectData
	NewSlot data.SlotObjectData
}

func (i *InvResultPacket) Read(p *network.Packet) {
	i.Success = p.ReadBool()
	i.OldSlot = data.SlotObjectData{}
	i.OldSlot.Read(p)
	i.NewSlot = data.SlotObjectData{}
	i.NewSlot.Read(p)
}

func (i InvResultPacket) Write(p *network.Packet) {
	p.WriteBool(i.Success)
	i.OldSlot.Write(p)
	i.NewSlot.Write(p)
}
