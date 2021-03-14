package client

import "proxy/network"

type EnemyHitPacket struct {
	Time     int32
	BulletID byte
	TargetID int32
	Killed   bool
}

func (e *EnemyHitPacket) Read(p *network.Packet) {
	e.Time = p.ReadInt32()
	e.BulletID = p.ReadByte()
	e.TargetID = p.ReadInt32()
	e.Killed = p.ReadBool()
}

func (e EnemyHitPacket) Write(p *network.Packet) {
	p.WriteInt32(e.Time)
	p.WriteByte(e.BulletID)
	p.WriteInt32(e.TargetID)
	p.WriteBool(e.Killed)
}
