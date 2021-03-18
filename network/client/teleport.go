package client

import "proxy/network"

type TeleportPacket struct {
	ObjectID int32
}

func (t *TeleportPacket) Read(p *network.Packet) {
	t.ObjectID = p.ReadInt32()
}

func (t TeleportPacket) Write(p *network.Packet) {
	p.WriteInt32(t.ObjectID)
}
