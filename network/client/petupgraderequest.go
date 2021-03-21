package client

import (
	"proxy/network"
	"proxy/network/data"
)

type PetUpgradeRequestPacket struct {
	TransType    byte
	PIDOne       int32
	PIDTwo       int32
	ObjectID     int32
	Slots        []data.SlotObjectData
	CurrencyType byte
}

func (pu *PetUpgradeRequestPacket) Read(p *network.Packet) {
	pu.TransType = p.ReadByte()
	pu.PIDOne = p.ReadInt32()
	pu.PIDTwo = p.ReadInt32()
	pu.ObjectID = p.ReadInt32()
	items := p.ReadInt16() // todo: check if correct int type
	if items > 0 {
		pu.Slots = make([]data.SlotObjectData, items)
		for i := 0; i < int(items); i++ {
			pu.Slots[i] = data.SlotObjectData{}
			pu.Slots[i].Read(p)
		}
	}
	pu.CurrencyType = p.ReadByte()
}

func (pu PetUpgradeRequestPacket) Write(p *network.Packet) {
	p.WriteByte(pu.TransType)
	p.WriteInt32(pu.PIDOne)
	p.WriteInt32(pu.PIDTwo)
	p.WriteInt32(pu.ObjectID)
	items := len(pu.Slots)
	if items > 0 {
		p.WriteInt16(int16(items))
		for i := 0; i < items; i++ {
			pu.Slots[i].Write(p)
		}
	}
	p.WriteByte(pu.CurrencyType)
}
