package server

import "proxy/network"

type PetYardUpdatePacket struct {
	Type int32
}

func (py *PetYardUpdatePacket) Read(p *network.Packet) {
	py.Type = p.ReadInt32()
}

func (py PetYardUpdatePacket) Write(p *network.Packet) {
	p.WriteInt32(py.Type)
}
