package client

import "proxy/network"

type SquareHitPacket struct {
	Time     int32
	BulletID byte
	ObjectID int32
}

func (s *SquareHitPacket) Read(p *network.Packet) {
	s.Time = p.ReadInt32()
	s.BulletID = p.ReadByte()
	s.ObjectID = p.ReadInt32()
}

func (s SquareHitPacket) Write(p *network.Packet) {
	p.WriteInt32(s.Time)
	p.WriteByte(s.BulletID)
	p.WriteInt32(s.ObjectID)
}
