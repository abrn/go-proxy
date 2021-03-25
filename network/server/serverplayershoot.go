package server

import (
	"proxy/network"
	"proxy/network/data"
)

type ServerPlayerShootPacket struct {
	BulletID      byte
	OwnerID       int32
	ContainerType int32
	Angle         float32
	Damage        int16
	StartingPos   data.WorldPosData
}

func (s *ServerPlayerShootPacket) Read(p *network.GamePacket) {
	s.BulletID = p.ReadByte()
	s.OwnerID = p.ReadInt32()
	s.ContainerType = p.ReadInt32()
	// starting pos may be here
	s.StartingPos.Read(p)
	s.Angle = p.ReadFloat()
	s.Damage = p.ReadInt16()
	s.StartingPos = data.WorldPosData{}
}

func (s ServerPlayerShootPacket) Write(p *network.GamePacket) {
	p.WriteByte(s.BulletID)
	p.WriteInt32(s.OwnerID)
	p.WriteInt32(s.ContainerType)
	p.WriteFloat(s.Angle)
	p.WriteInt16(s.Damage)
	s.StartingPos.Write(p)
}
