package server

import "proxy/network"

type NewTickPacket struct {
	TickID           int32
	TickTime         int32  // last client tick time
	ServerRealTimeMS uint32 // last server tick time in milliseconds
	ServerLastRTTMS  uint16 // last server packet round-trip time in milliseconds
}

func (n *NewTickPacket) Read(p *network.Packet) {
	n.TickID = p.ReadInt32()
	n.TickTime = p.ReadInt32()
	n.ServerRealTimeMS = p.ReadUInt32()
	n.ServerLastRTTMS = p.ReadUInt16()
}

func (n NewTickPacket) Write(p *network.Packet) {
	p.WriteInt32(n.TickID)
	p.WriteInt32(n.TickTime)
	p.WriteUInt32(n.ServerRealTimeMS)
	p.WriteUInt16(n.ServerLastRTTMS)
}
