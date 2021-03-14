package server

import (
	"proxy/network"
	"proxy/network/data"
)

type DamagePacket struct {
	TargetID    int32
	Effect      data.ConditionEffect
	Damage      uint16
	Killed      bool
	ArmorPierce bool
	BulletID    byte
	ObjectID    int32
}

func (d *DamagePacket) Read(p *network.Packet) {
	d.TargetID = p.ReadInt32()
	d.Effect = data.ConditionEffect{}
	d.Effect.Read(p)
	d.Damage = p.ReadUInt16()
	d.Killed = p.ReadBool()
	d.ArmorPierce = p.ReadBool()
	d.BulletID = p.ReadByte()
	d.ObjectID = p.ReadInt32()
}

func (d DamagePacket) Write(p *network.Packet) {
	p.WriteInt32(d.TargetID)
	d.Effect.Write(p)
	p.WriteUInt16(d.Damage)
	p.WriteBool(d.Killed)
	p.WriteBool(d.ArmorPierce)
	p.WriteByte(d.BulletID)
	p.WriteInt32(d.ObjectID)
}
