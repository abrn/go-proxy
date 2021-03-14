package server

import "proxy/network"

type ReskinUnlockPacket struct {
	SkinID    int32
	IsPetSkin int32
}

func (r *ReskinUnlockPacket) Read(p *network.Packet) {
	r.SkinID = p.ReadInt32()
	r.IsPetSkin = p.ReadInt32()
}

func (r ReskinUnlockPacket) Write(p *network.Packet) {
	p.WriteInt32(r.SkinID)
	p.WriteInt32(r.IsPetSkin)
}
