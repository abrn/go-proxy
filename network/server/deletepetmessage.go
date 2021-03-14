package server

import "proxy/network"

type DeletePetMessagePacket struct {
	PetID int32
}

func (d *DeletePetMessagePacket) Read(p *network.Packet) {
	d.PetID = p.ReadInt32()
}

func (d DeletePetMessagePacket) Write(p *network.Packet) {
	p.WriteInt32(d.PetID)
}
