package client

import (
	"proxy/network"
	"proxy/network/data"
)

type InvDropPacket struct {
	Slot    data.SlotObjectData
	Unknown bool
}

func (i *InvDropPacket) Read(p *network.Packet) {
	i.Slot = data.SlotObjectData{}
	i.Slot.Read(p)
	i.Unknown = p.ReadBool()
}

func (i InvDropPacket) Write(p *network.Packet) {
	i.Slot.Write(p)
	p.WriteBool(i.Unknown)
}
