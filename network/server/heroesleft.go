package server

import "proxy/network"

type HeroesLeftPacket struct {
	Heroes int32
}

func (h *HeroesLeftPacket) Read(p *network.GamePacket) {
	h.Heroes = p.ReadInt32()
}

func (h HeroesLeftPacket) Write(p *network.GamePacket) {
	p.WriteInt32(h.Heroes)
}
