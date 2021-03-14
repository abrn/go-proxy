package server

import "proxy/network"

type NewAbilityMessagePacket struct {
	Type int32
}

func (n *NewAbilityMessagePacket) Read(p *network.Packet) {
	n.Type = p.ReadInt32()
}

func (n NewAbilityMessagePacket) Write(p *network.Packet) {
	p.WriteInt32(n.Type)
}
