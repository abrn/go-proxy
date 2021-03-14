package server

import "proxy/network"

type PingPacket struct {
	Serial int32
}

func (pp *PingPacket) Read(p *network.Packet) {
	pp.Serial = p.ReadInt32()
}

func (pp PingPacket) Write(p *network.Packet) {
	p.WriteInt32(pp.Serial)
}
