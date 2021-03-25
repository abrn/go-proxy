package server

import "proxy/network"

type ActivePetPacket struct {
	InstanceID int32
}

func (a *ActivePetPacket) Read(p *network.GamePacket) {
	a.InstanceID = p.ReadInt32()
}

func (a ActivePetPacket) Write(p *network.GamePacket) {
	p.WriteInt32(a.InstanceID)
}
