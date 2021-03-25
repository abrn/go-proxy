package server

import "proxy/network"

type HatchPetPacket struct {
	PetName  string
	PetType  int32
	ItemType int32
	Unknown  bool
}

func (h *HatchPetPacket) Read(p *network.GamePacket) {
	h.PetName = p.ReadString()
	h.ItemType = p.ReadInt32()
	h.PetType = p.ReadInt32()
	h.Unknown = p.ReadBool()
}

func (h HatchPetPacket) Write(p *network.GamePacket) {
	p.WriteString(h.PetName)
	p.WriteInt32(h.PetType)
	p.WriteInt32(h.ItemType)
	p.WriteBool(h.Unknown)
}
