package server

import (
	"proxy/network"
	"proxy/network/data"
)

type ShowEffectPacket struct {
	Type      byte
	TargetID  int32
	StartPos  data.WorldPosData
	EndPos    data.WorldPosData
	Color     int32
	Duration  float32
	EffectBit EffectBitmask
}

type EffectBitmask byte

const (
	ColorBit  EffectBitmask = 1
	StartPosX EffectBitmask = 2
	StartPosY EffectBitmask = 4
	EndPosX   EffectBitmask = 8
	EndPosY   EffectBitmask = 16
	Duration  EffectBitmask = 32
	TargetID  EffectBitmask = 64
)

func (s *ShowEffectPacket) Read(p *network.Packet) {
	s.Type = p.ReadByte()
	s.StartPos = data.WorldPosData{}
	s.EndPos = data.WorldPosData{}
	s.EffectBit = EffectBitmask(p.ReadByte())
	if s.EffectBit&TargetID != 0 {
		s.TargetID = p.ReadInt32()
	}
	if s.EffectBit&StartPosX != 0 {
		s.StartPos.X = p.ReadFloat()
	}
	if s.EffectBit&StartPosY != 0 {
		s.StartPos.Y = p.ReadFloat()
	}
	if s.EffectBit&EndPosX != 0 {
		s.EndPos.X = p.ReadFloat()
	}
	if s.EffectBit&EndPosY != 0 {
		s.EndPos.Y = p.ReadFloat()
	}
	if s.EffectBit&ColorBit != 0 {
		s.Color = p.ReadInt32()
	} else {
		s.Color = 4294967295
	}
	if s.EffectBit&Duration != 0 {
		s.Duration = p.ReadFloat()
	} else {
		s.Duration = 1.0
	}
}

// todo: SHOWEFFECT add write function
