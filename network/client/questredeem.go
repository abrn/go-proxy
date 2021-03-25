package client

import (
	"proxy/network"
	"proxy/network/data"
)

type QuestRedeemPacket struct {
	UnknownOne string
	QuestID    int32
	Slots      []data.SlotObjectData
}

func (q *QuestRedeemPacket) Read(p *network.GamePacket) {
	q.UnknownOne = p.ReadString()
	q.QuestID = p.ReadInt32()
	items := p.ReadInt16() // todo: check if correct int type
	if items > 0 {
		q.Slots = make([]data.SlotObjectData, items)
		for i := 0; i < int(items); i++ {
			q.Slots[i] = data.SlotObjectData{}
			q.Slots[i].Read(p)
		}
	}
}

func (q QuestRedeemPacket) Write(p *network.GamePacket) {
	p.WriteString(q.UnknownOne)
	p.WriteInt32(q.QuestID)
	items := len(q.Slots)
	if items > 0 {
		p.WriteInt16(int16(items))
		for i := 0; i < items; i++ {
			q.Slots[i].Write(p)
		}
	}
}
