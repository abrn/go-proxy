package client

import (
	"proxy/network"
	"proxy/network/data"
)

type InvSwapPacket struct {
	Time     int32
	Position data.WorldPosData
	OldSlot  data.SlotObjectData
	NewSlot  data.SlotObjectData
}

func (s *InvSwapPacket) Read(p *network.Packet) {
	s.Time = p.ReadInt32()
	s.Position = data.WorldPosData{}
	s.Position.Read(p)
	s.OldSlot = data.SlotObjectData{}
	s.NewSlot = data.SlotObjectData{}
	s.OldSlot.Read(p)
	s.NewSlot.Read(p)
}

func (s InvSwapPacket) Write(p *network.Packet) {
	p.WriteInt32(s.Time)
	s.Position.Write(p)
	s.OldSlot.Write(p)
	s.NewSlot.Write(p)
}
