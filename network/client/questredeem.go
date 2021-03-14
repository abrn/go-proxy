package client

import (
	"proxy/network"
	"proxy/network/data"
)

type QuestRedeemPacket struct {
	Slot data.SlotObjectData
}

func (q *QuestRedeemPacket) Read(p *network.Packet) {
	q.Slot = data.SlotObjectData{}
	q.Slot.Read(p)
}

func (q QuestRedeemPacket) Write(p *network.Packet) {
	q.Slot.Write(p)
}
