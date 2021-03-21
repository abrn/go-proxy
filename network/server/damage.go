package server

import (
	"proxy/network"
	"proxy/network/data"
)

type DamagePacket struct {
	TargetID    int32
	Effects     []data.ConditionEffect
	Damage      uint16
	DamageType  DamageType
	Killed      bool
	ArmorPierce bool
	Laser       bool
	BulletID    byte
	ObjectID    int32
}

type DamageType byte

const (
	DamageTypeKilled      DamageType = 1
	DamageTypeArmorPierce DamageType = 2
	DamageTypeLaser       DamageType = 4
)

func (d *DamagePacket) Read(p *network.Packet) {
	d.TargetID = p.ReadInt32()
	effects := p.ReadByte()
	if effects > 0 {
		d.Effects = make([]data.ConditionEffect, effects)
		for i := 0; i < int(effects); i++ {
			d.Effects[i] = data.ConditionEffect{}
			d.Effects[i].Read(p)
		}
	}
	d.Damage = p.ReadUInt16()
	d.DamageType = DamageType(p.ReadByte())
	d.Killed = d.DamageType&DamageTypeKilled != 0
	d.ArmorPierce = d.DamageType&DamageTypeArmorPierce != 0
	d.Laser = d.DamageType&DamageTypeLaser != 0
	d.BulletID = p.ReadByte()
	d.ObjectID = p.ReadInt32()
}

func (d DamagePacket) Write(p *network.Packet) {
	p.WriteInt32(d.TargetID)
	effects := len(d.Effects)
	if effects > 0 {
		for i := 0; i < effects; i++ {
			d.Effects[i].Write(p)
		}
	}
	p.WriteUInt16(d.Damage)
	if d.Killed {
		p.WriteByte(byte(DamageTypeKilled))
	}
	if d.ArmorPierce {
		p.WriteByte(byte(DamageTypeArmorPierce))
	}
	if d.Laser {
		p.WriteByte(byte(DamageTypeLaser))
	}
	p.WriteByte(d.BulletID)
	p.WriteInt32(d.ObjectID)
}
