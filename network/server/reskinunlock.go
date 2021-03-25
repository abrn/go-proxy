package server

import "proxy/network"

type ReskinUnlockPacket struct {
	SkinID    int32
	IsPetSkin int32
}

func (r *ReskinUnlockPacket) Read(p *network.GamePacket) {
	r.SkinID = p.ReadInt32()
	r.IsPetSkin = p.ReadInt32()
}

func (r ReskinUnlockPacket) Write(p *network.GamePacket) {
	p.WriteInt32(r.SkinID)
	p.WriteInt32(r.IsPetSkin)
}
