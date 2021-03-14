package client

import (
	"proxy/network"
	"proxy/network/data"
)

type UseItemPacket struct {
	Time       int32
	SlotObject data.SlotObjectData
	Position   data.WorldPosData
	UseType    byte
}

func (u *UseItemPacket) Read(p *network.Packet) {
	u.Time = p.ReadInt32()
	u.SlotObject = data.SlotObjectData{}
	u.SlotObject.Read(p)
	u.Position = data.WorldPosData{}
	u.Position.Read(p)
	u.UseType = p.ReadByte()
}

func (u UseItemPacket) Write(p *network.Packet) {
	p.WriteInt32(u.Time)
	u.SlotObject.Write(p)
	u.Position.Write(p)
	p.WriteByte(u.UseType)
}
