package server

import "proxy/network"

type PlaySoundPacket struct {
	OwnerID int32
	SoundID byte
}

func (ps *PlaySoundPacket) Read(p *network.GamePacket) {
	ps.OwnerID = p.ReadInt32()
	ps.SoundID = p.ReadByte()
}

func (ps PlaySoundPacket) Write(p *network.GamePacket) {
	p.WriteInt32(ps.OwnerID)
	p.WriteByte(ps.SoundID)
}
