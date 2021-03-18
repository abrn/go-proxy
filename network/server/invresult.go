package server

import "proxy/network"

type InvResultPacket struct {
	ResultCode int32
}

func (i *InvResultPacket) Read(p *network.Packet) {
	i.ResultCode = p.ReadInt32()
}

func (i InvResultPacket) Write(p *network.Packet) {
	p.WriteInt32(i.ResultCode)
}
