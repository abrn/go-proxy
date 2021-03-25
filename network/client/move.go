package client

import (
	"proxy/network"
	"proxy/network/data"
)

type MovePacket struct {
	TickID               int32
	TickTime             int32
	LastServerRealTimeMS uint32
	NewPosition          data.WorldPosData
	Records              []data.MoveRecord
}

func (m *MovePacket) Read(p *network.GamePacket) {
	m.TickID = p.ReadInt32()
	m.TickTime = p.ReadInt32()
	m.LastServerRealTimeMS = p.ReadUInt32()
	m.NewPosition = data.WorldPosData{}
	m.NewPosition.Read(p)
	records := p.ReadInt16()
	if records <= 0 {
		return
	}
	m.Records = make([]data.MoveRecord, records)
	for i := 0; i < int(records); i++ {
		m.Records[i] = data.MoveRecord{}
		m.Records[i].Read(p)
	}
}

func (m MovePacket) Write(p *network.GamePacket) {
	p.WriteInt32(m.TickID)
	p.WriteInt32(m.TickTime)
	p.WriteUInt32(m.LastServerRealTimeMS)
	m.NewPosition.Write(p)
	records := len(m.Records)
	p.WriteInt16(int16(records))
	if records <= 0 {
		return
	}
	for i := 0; i < records; i++ {
		m.Records[i].Write(p)
	}
}
