package server

import "proxy/network"

type ExaltationRedeemMessagePacket struct {
	ResultCode int32
}

func (e *ExaltationRedeemMessagePacket) Read(p *network.Packet) {
	e.ResultCode = p.ReadInt32()
}

func (e ExaltationRedeemMessagePacket) Write(p *network.Packet) {
	p.WriteInt32(e.ResultCode)
}
