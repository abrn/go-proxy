package server

import (
	"proxy/network"
	"proxy/network/data"
)

type QuestFetchResponsePacket struct {
	Quests []data.QuestData
}

func (q *QuestFetchResponsePacket) Read(p *network.Packet) {
	count := p.ReadInt16() // todo: check if this is the right int type
	if count <= 0 {
		return
	}
	q.Quests = make([]data.QuestData, count)
	for i := 0; i < int(count); i++ {
		q.Quests[i] = data.QuestData{}
		q.Quests[i].Read(p)
	}
}

func (q QuestFetchResponsePacket) Write(p *network.Packet) {
	count := len(q.Quests)
	p.WriteInt16(int16(count))
	if count <= 0 {
		return
	}
	for i := 0; i < count; i++ {
		q.Quests[i].Write(p)
	}
}
