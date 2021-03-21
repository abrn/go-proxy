package server

import "proxy/network"

type ExaltationRedeemMessagePacket struct {
	Success bool
	Message string
}

func (e *ExaltationRedeemMessagePacket) Read(p *network.Packet) {
	e.Success = p.ReadBool()
	e.Message = p.ReadString()
}

func (e ExaltationRedeemMessagePacket) Write(p *network.Packet) {
	p.WriteBool(e.Success)
	p.WriteString(e.Message)
}
