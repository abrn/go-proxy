package client

import "proxy/network"

type SetConditionPacket struct {
	Effect   byte
	Duration float32
}

func (s *SetConditionPacket) Read(p *network.GamePacket) {
	s.Effect = p.ReadByte()
	s.Duration = p.ReadFloat()
}

func (s SetConditionPacket) Write(p *network.GamePacket) {
	p.WriteByte(s.Effect)
	p.WriteFloat(s.Duration)
}
