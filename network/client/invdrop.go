package client

import (
	"proxy/network"
	"proxy/network/data"
)

type InvDropPacket struct {
	Slot    data.SlotObjectData
	Unknown bool
}

func (i *InvDropPacket) Read(p *network.GamePacket) {
	i.Slot = data.SlotObjectData{}
	i.Slot.Read(p)
	i.Unknown = p.ReadBool()
}

func (i InvDropPacket) Write(p *network.GamePacket) {
	i.Slot.Write(p)
	p.WriteBool(i.Unknown)
}
