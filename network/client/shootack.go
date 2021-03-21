package client

import "proxy/network"

type ShootAckPacket struct {
	Time    int32
	Unknown int16
}

func (s *ShootAckPacket) Read(p *network.Packet) {
	s.Time = p.ReadInt32()
	s.Unknown = p.ReadInt16()
}

func (s ShootAckPacket) Write(p *network.Packet) {
	p.WriteInt32(s.Time)
	p.WriteInt16(s.Unknown)
}
