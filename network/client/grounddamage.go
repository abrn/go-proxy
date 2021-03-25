package client

import (
	"proxy/network"
	"proxy/network/data"
)

type GroundDamagePacket struct {
	Time     int32
	Position data.WorldPosData
}

func (g *GroundDamagePacket) Read(p *network.GamePacket) {
	g.Time = p.ReadInt32()
	g.Position = data.WorldPosData{}
	g.Position.Read(p)
}

func (g GroundDamagePacket) Write(p *network.GamePacket) {
	p.WriteInt32(g.Time)
	g.Position.Write(p)
}
