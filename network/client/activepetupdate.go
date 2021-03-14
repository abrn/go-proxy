package client

import (
	"proxy/network"
)

type ActivePetUpdatePacket struct {
	CommandType 	byte
	InstanceID 		int32
}

func (a *ActivePetUpdatePacket) Read(p *network.Packet) {
	a.CommandType = p.ReadByte()
	a.InstanceID = p.ReadInt32()
}

func (a ActivePetUpdatePacket) Write(p *network.Packet) {
	p.WriteByte(a.CommandType)
	p.WriteInt32(a.InstanceID)
}