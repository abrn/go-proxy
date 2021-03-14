package client

import (
	"proxy/network"
)

type LoadPacket struct {
	CharID    int32
	FromArena bool
}

func (l *LoadPacket) Read(p *network.Packet) {
	l.CharID = p.ReadInt32()
	l.FromArena = p.ReadBool()
}

func (l LoadPacket) Write(p *network.Packet) {
	p.WriteInt32(l.CharID)
	p.WriteBool(l.FromArena)
}
