package server

import (
	"proxy/network"
	"proxy/network/data"
)

type ShowEffectPacket struct {
	Type     byte
	TargetID int32
	StartPos data.WorldPosData
	EndPos   data.WorldPosData
	Color    int32
	Duration float32
}

func (s *ShowEffectPacket) Read(p *network.Packet) {
	s.Type = p.ReadByte()
	s.StartPos = data.WorldPosData{}
	s.EndPos = data.WorldPosData{}
	bit := p.ReadByte()
	if bit & 64 != 0 {
		s.TargetID = p.ReadInt32()
	}
	if bit & 2 != 0 {
		s.StartPos.X = p.ReadFloat()
	}
	if bit & 4 != 0 {
		s.StartPos.Y = p.ReadFloat()
	}
	if bit & 8 != 0 {
		s.EndPos.X = p.ReadFloat()
	}
	if bit & 16 != 0 {
		s.EndPos.Y = p.ReadFloat()
	}
	if bit & 1 != 0 {
		s.Color = p.ReadInt32()
	} else {
		s.Color = 4294967295
	}
	if bit & 32 != 0 {
		s.Duration = p.ReadFloat()
	} else {
		s.Duration = 1.0
	}
}

// todo: SHOWEFFECT add write function