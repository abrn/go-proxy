package client

import "proxy/network"

type UsePortalPacket struct {
	ObjectID int32
}

func (u *UsePortalPacket) Read(p *network.GamePacket) {
	u.ObjectID = p.ReadInt32()
}

func (u UsePortalPacket) Write(p *network.GamePacket) {
	p.WriteInt32(u.ObjectID)
}
