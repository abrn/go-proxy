package server

import "proxy/network"

type HatchPetPacket struct {
	PetName  string
	PetType  int32
	ItemType int32
}

func (h *HatchPetPacket) Read(p *network.Packet) {
	h.PetName = p.ReadString()
	h.ItemType = p.ReadInt32()
	h.PetType = p.ReadInt32()
}

func (h HatchPetPacket) Write(p *network.Packet) {
	p.WriteString(h.PetName)
	p.WriteInt32(h.PetType)
	p.WriteInt32(h.ItemType)
}
