package server

import "proxy/network"

type PetYardUpdatePacket struct {
	Type int32
}

func (py *PetYardUpdatePacket) Read(p *network.GamePacket) {
	py.Type = p.ReadInt32()
}

func (py PetYardUpdatePacket) Write(p *network.GamePacket) {
	p.WriteInt32(py.Type)
}
