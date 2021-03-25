package server

import "proxy/network"

type UnlockNewSlotPacket struct {
	Type int32
}

func (u *UnlockNewSlotPacket) Read(p *network.GamePacket) {
	u.Type = p.ReadInt32()
}

func (u UnlockNewSlotPacket) Write(p *network.GamePacket) {
	p.WriteInt32(u.Type)
}
