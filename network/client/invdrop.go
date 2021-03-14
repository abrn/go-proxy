package client

import (
	"proxy/network"
	"proxy/network/data"
)

type InvDropPacket struct {
	Slot data.SlotObjectData
}

func (i *InvDropPacket) Read(p *network.Packet) {
	i.Slot = data.SlotObjectData{}
	i.Slot.Read(p)
}

func (i InvDropPacket) Write(p *network.Packet) {
	i.Slot.Write(p)
}
