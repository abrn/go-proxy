package client

import "proxy/network"

type ReskinPacket struct {
	SkinID int32
}

func (r *ReskinPacket) Read(p *network.GamePacket) {
	r.SkinID = p.ReadInt32()
}

func (r ReskinPacket) Write(p *network.GamePacket) {
	p.WriteInt32(r.SkinID)
}
