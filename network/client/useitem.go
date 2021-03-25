package client

import (
	"proxy/network"
	"proxy/network/data"
)

type UseItemPacket struct {
	Time       int32
	SlotObject data.SlotObjectData
	Position   data.WorldPosData
	UseType    UseItemType
}

type UseItemType byte

const (
	UseTypeDefault UseItemType = 0 // using a consumable
	UseTypeStart   UseItemType = 1 // starting an ability
	UseTypeEnd     UseItemType = 2 // ending an ability (i.e ninja star)
)

func (u *UseItemPacket) Read(p *network.GamePacket) {
	u.Time = p.ReadInt32()
	u.SlotObject = data.SlotObjectData{}
	u.SlotObject.Read(p)
	u.Position = data.WorldPosData{}
	u.Position.Read(p)
	u.UseType = UseItemType(p.ReadByte())
}

func (u UseItemPacket) Write(p *network.GamePacket) {
	p.WriteInt32(u.Time)
	u.SlotObject.Write(p)
	u.Position.Write(p)
	p.WriteByte(byte(u.UseType))
}
