package client

import (
	"proxy/network"
	"proxy/network/data"
)

type PlayerShootPacket struct {
	Time          int32
	BulletID      byte
	ContainerType int16
	Position      data.WorldPosData
	Angle         float32
	SpeedMult     int16
	LifeMult      int16
	Unknown       byte
}

func (s *PlayerShootPacket) Read(p *network.Packet) {
	s.Time = p.ReadInt32()
	s.BulletID = p.ReadByte()
	s.ContainerType = p.ReadInt16()
	s.Position = data.WorldPosData{}
	s.Position.Read(p)
	s.Angle = p.ReadFloat()
	s.SpeedMult = p.ReadInt16()
	s.LifeMult = p.ReadInt16()
	s.Unknown = p.ReadByte()
}

func (s PlayerShootPacket) Write(p *network.Packet) {
	p.WriteInt32(s.Time)
	p.WriteByte(s.BulletID)
	p.WriteInt16(s.ContainerType)
	s.Position.Write(p)
	p.WriteFloat(s.Angle)
	p.WriteInt16(s.SpeedMult)
	p.WriteInt16(s.LifeMult)
	p.WriteByte(s.Unknown)
}
