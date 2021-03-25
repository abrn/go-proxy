package server

import "proxy/network"

type DeletePetMessagePacket struct {
	PetID int32
}

func (d *DeletePetMessagePacket) Read(p *network.GamePacket) {
	d.PetID = p.ReadInt32()
}

func (d DeletePetMessagePacket) Write(p *network.GamePacket) {
	p.WriteInt32(d.PetID)
}
