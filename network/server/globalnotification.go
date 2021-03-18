package server

import "proxy/network"

type GlobalNotificationPacket struct {
	TypeID  int32
	Message string
}

func (g *GlobalNotificationPacket) Read(p *network.Packet) {
	g.TypeID = p.ReadInt32()
	g.Message = p.ReadString()
}

func (g GlobalNotificationPacket) Write(p *network.Packet) {
	p.WriteInt32(g.TypeID)
	p.WriteString(g.Message)
}
