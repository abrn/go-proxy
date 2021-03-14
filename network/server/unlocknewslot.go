package server

import "proxy/network"

type UnlockNewSlotPacket struct {
	Type int32
}

func (u *UnlockNewSlotPacket) Read(p *network.Packet) {
	u.Type = p.ReadInt32()
}

func (u UnlockNewSlotPacket) Write(p *network.Packet) {
	p.WriteInt32(u.Type)
}
