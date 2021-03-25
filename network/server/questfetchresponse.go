package server

import (
	"proxy/network"
	"proxy/network/data"
)

type QuestFetchResponsePacket struct {
	Quests  []data.QuestData
	Unknown int32
}

func (q *QuestFetchResponsePacket) Read(p *network.GamePacket) {
	count := p.ReadInt16() // todo: check if this is the right int type
	if count > 0 {
		q.Quests = make([]data.QuestData, count)
		for i := 0; i < int(count); i++ {
			q.Quests[i] = data.QuestData{}
			q.Quests[i].Read(p)
		}
	}
	q.Unknown = p.ReadInt32()
}

func (q QuestFetchResponsePacket) Write(p *network.GamePacket) {
	count := len(q.Quests)
	p.WriteInt16(int16(count))
	if count > 0 {
		for i := 0; i < count; i++ {
			q.Quests[i].Write(p)
		}
	}
	p.WriteInt32(q.Unknown)
}
