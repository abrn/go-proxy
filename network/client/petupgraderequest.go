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
	SlotObject   data.SlotObjectData
	CurrencyType byte
}

func (pu *PetUpgradeRequestPacket) Read(p *network.Packet) {
	pu.TransType = p.ReadByte()
	pu.PIDOne = p.ReadInt32()
	pu.PIDTwo = p.ReadInt32()
	pu.ObjectID = p.ReadInt32()
	pu.SlotObject = data.SlotObjectData{}
	pu.SlotObject.Read(p)
	pu.CurrencyType = p.ReadByte()
}

func (pu PetUpgradeRequestPacket) Write(p *network.Packet) {
	p.WriteByte(pu.TransType)
	p.WriteInt32(pu.PIDOne)
	p.WriteInt32(pu.PIDTwo)
	p.WriteInt32(pu.ObjectID)
	pu.SlotObject.Write(p)
	p.WriteByte(pu.CurrencyType)
}
