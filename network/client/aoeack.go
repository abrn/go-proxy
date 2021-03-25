package client

import (
	"proxy/network"
	"proxy/network/data"
)

type AoEAckPacket struct {
	Time     int32
	Position data.WorldPosData
}

func (a *AoEAckPacket) Read(p *network.GamePacket) {
	a.Time = p.ReadInt32()
	a.Position = data.WorldPosData{}
	a.Position.Read(p)
}

func (a AoEAckPacket) Write(p *network.GamePacket) {
	p.WriteInt32(a.Time)
	a.Position.Write(p)
}
