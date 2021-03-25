package server

import (
	"proxy/network"
	"proxy/network/data"
)

type NewTickPacket struct {
	TickID           int32
	TickTime         int32  // last client tick time
	ServerRealTimeMS uint32 // last server tick time in milliseconds
	ServerLastRTTMS  uint16 // last server packet round-trip time in milliseconds
	Statuses         []data.ObjectStatusData
}

func (n *NewTickPacket) Read(p *network.GamePacket) {
	n.TickID = p.ReadInt32()
	n.TickTime = p.ReadInt32()
	n.ServerRealTimeMS = p.ReadUInt32()
	n.ServerLastRTTMS = p.ReadUInt16()
	statCount := p.ReadInt16()
	if statCount <= 0 {
		return
	}
	n.Statuses = make([]data.ObjectStatusData, statCount)
	for i := 0; i < int(statCount); i++ {
		n.Statuses[i] = data.ObjectStatusData{}
		n.Statuses[i].Read(p)
	}
}

func (n NewTickPacket) Write(p *network.GamePacket) {
	p.WriteInt32(n.TickID)
	p.WriteInt32(n.TickTime)
	p.WriteUInt32(n.ServerRealTimeMS)
	p.WriteUInt16(n.ServerLastRTTMS)
	statCount := len(n.Statuses)
	if statCount <= 0 {
		return
	}
	for i := 0; i < statCount; i++ {
		n.Statuses[i].Write(p)
	}
}
