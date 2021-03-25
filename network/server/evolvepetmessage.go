package server

import "proxy/network"

type EvolvePetMessagePacket struct {
	PetID       int32
	InitialSkin int32
	FinalSkin   int32
}

func (h *EvolvePetMessagePacket) Read(p *network.GamePacket) {
	h.PetID = p.ReadInt32()
	h.InitialSkin = p.ReadInt32()
	h.FinalSkin = p.ReadInt32()
}

func (h EvolvePetMessagePacket) Write(p *network.GamePacket) {
	p.WriteInt32(h.PetID)
	p.WriteInt32(h.InitialSkin)
	p.WriteInt32(h.FinalSkin)
}
