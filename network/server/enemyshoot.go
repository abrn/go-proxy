package server

import (
	"proxy/network"
	"proxy/network/data"
)

type EnemyShoot struct {
	BulletID   byte
	OwnerID    int32
	BulletType byte
	Location   data.WorldPosData
	Angle      float32
	Damage     int16
	NumShots   byte
	AngleInc   float32
}

func (s *EnemyShoot) Read(p *network.Packet) {
	s.BulletID = p.ReadByte()
	s.OwnerID = p.ReadInt32()
	s.BulletType = p.ReadByte()
	s.Location = data.WorldPosData{}
	s.Location.Read(p)
	s.Angle = p.ReadFloat()
	s.Damage = p.ReadInt16()
	if len(p.Data[p.Index:]) > 0 {
		s.NumShots = p.ReadByte()
		s.AngleInc = p.ReadFloat()
	} else {
		s.NumShots = 1
		s.AngleInc = 0.0
	}
}

func (s EnemyShoot) Write(p *network.Packet) {
	p.WriteByte(s.BulletID)
	p.WriteInt32(s.OwnerID)
	p.WriteByte(s.BulletType)
	s.Location.Write(p)
	p.WriteFloat(s.Angle)
	p.WriteInt16(s.Damage)
	if s.NumShots != 1 || s.AngleInc != 0.0 {
		p.WriteByte(s.NumShots)
		p.WriteFloat(s.AngleInc)
	}
}
