package server

import "proxy/network"

type ExaltationBonusChangedPacket struct {
	ObjectType uint16 // the player class type
	Attack     byte
	Defense    byte
	Speed      byte
	Dexterity  byte
	Vitality   byte
	Wisdom     byte
	Life       byte
	Mana       byte
}

func (e *ExaltationBonusChangedPacket) Read(p *network.Packet) {
	e.ObjectType = p.ReadUInt16()
	e.Dexterity = p.ReadByte()
	e.Speed = p.ReadByte()
	e.Vitality = p.ReadByte()
	e.Wisdom = p.ReadByte()
	e.Defense = p.ReadByte()
	e.Attack = p.ReadByte()
	e.Mana = p.ReadByte()
	e.Life = p.ReadByte()
}

func (e ExaltationBonusChangedPacket) Write(p *network.Packet) {
	p.WriteUInt16(e.ObjectType)
	p.WriteByte(e.Dexterity)
	p.WriteByte(e.Speed)
	p.WriteByte(e.Vitality)
	p.WriteByte(e.Wisdom)
	p.WriteByte(e.Defense)
	p.WriteByte(e.Attack)
	p.WriteByte(e.Mana)
	p.WriteByte(e.Life)
}
