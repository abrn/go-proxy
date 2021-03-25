package server

import "proxy/network"

type QuestObjectIDPacket struct {
	ObjectID   int32
	HealthBars []int32
}

func (q *QuestObjectIDPacket) Read(p *network.GamePacket) {
	q.ObjectID = p.ReadInt32()
	bars := p.ReadCompressed()
	q.HealthBars = make([]int32, bars)
	for i := 0; i < int(bars); i++ {
		q.HealthBars[i] = p.ReadCompressed()
	}
}

func (q QuestObjectIDPacket) Write(p *network.GamePacket) {
	p.WriteInt32(q.ObjectID)
	bars := len(q.HealthBars)
	p.WriteCompressed(int32(bars))
	if bars <= 0 {
		return
	}
	for i := 0; i < bars; i++ {
		p.WriteCompressed(q.HealthBars[i])
	}
}
