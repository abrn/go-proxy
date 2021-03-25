package server

import (
	"proxy/network"
	"proxy/network/data"
)

type GotoPacket struct {
	ObjectID int32
	Location data.WorldPosData
}

func (g *GotoPacket) Read(p *network.GamePacket) {
	g.ObjectID = p.ReadInt32()
	g.Location = data.WorldPosData{}
	g.Location.Read(p)
}

func (g GotoPacket) Write(p *network.GamePacket) {
	p.WriteInt32(g.ObjectID)
	g.Location.Write(p)
}
