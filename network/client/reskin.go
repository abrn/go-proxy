package client

import "proxy/network"

type ReskinPacket struct {
	SkinID int32
}

func (r *ReskinPacket) Read(p *network.Packet) {
	r.SkinID = p.ReadInt32()
}

func (r ReskinPacket) Write(p *network.Packet) {
	p.WriteInt32(r.SkinID)
}
