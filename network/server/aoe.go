package server

import (
	"proxy/network"
	"proxy/network/data"
)

type AoEPacket struct {
	Position       data.WorldPosData
	Radius         float32
	Damage         uint16
	Effects        data.ConditionEffect
	EffectDuration float32
	OriginType     int16
	Color          int32
	ArmorPierce    bool
}

func (a *AoEPacket) Read(p *network.Packet) {
	a.Position = data.WorldPosData{}
	a.Position.Read(p)
	a.Radius = p.ReadFloat()
	a.Damage = p.ReadUInt16()
	a.Effects = data.ConditionEffect{}
	a.Effects.Read(p)
	a.EffectDuration = p.ReadFloat()
	a.OriginType = p.ReadInt16()
	a.Color = p.ReadInt32()
	a.ArmorPierce = p.ReadBool()
}

func (a AoEPacket) Write(p *network.Packet) {
	a.Position.Write(p)
	p.WriteFloat(a.Radius)
	p.WriteUInt16(a.Damage)
	a.Effects.Write(p)
	p.WriteFloat(a.EffectDuration)
	p.WriteInt16(a.OriginType)
	p.WriteInt32(a.Color)
	p.WriteBool(a.ArmorPierce)
}
