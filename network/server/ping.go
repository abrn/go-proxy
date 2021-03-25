package server

import "proxy/network"

type PingPacket struct {
	Serial int32
}

func (pp *PingPacket) Read(p *network.GamePacket) {
	pp.Serial = p.ReadInt32()
}

func (pp PingPacket) Write(p *network.GamePacket) {
	p.WriteInt32(pp.Serial)
}
