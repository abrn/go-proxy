package client

import "proxy/network"

type TeleportPacket struct {
	ObjectID int32
}

func (t *TeleportPacket) Read(p *network.GamePacket) {
	t.ObjectID = p.ReadInt32()
}

func (t TeleportPacket) Write(p *network.GamePacket) {
	p.WriteInt32(t.ObjectID)
}
