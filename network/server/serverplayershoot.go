package server

import (
	"proxy/network"
	"proxy/network/data"
)

type ServerPlayerShootPacket struct {
	BulletID      byte
	OwnerID       int32
	ContainerType int32
	StartingPos   data.WorldPosData
	Angle         float32
	Damage        int16
}

func (s *ServerPlayerShootPacket) Read(p *network.Packet) {
	s.BulletID = p.ReadByte()
	s.OwnerID = p.ReadInt32()
	s.ContainerType = p.ReadInt32()
	s.StartingPos = data.WorldPosData{}
	s.StartingPos.Read(p)
	s.Angle = p.ReadFloat()
	s.Damage = p.ReadInt16()
}

func (s ServerPlayerShootPacket) Write(p *network.Packet) {
	p.WriteByte(s.BulletID)
	p.WriteInt32(s.OwnerID)
	p.WriteInt32(s.ContainerType)
	s.StartingPos.Write(p)
	p.WriteFloat(s.Angle)
	p.WriteInt16(s.Damage)
}
