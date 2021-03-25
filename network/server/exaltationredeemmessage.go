package server

import "proxy/network"

type ExaltationRedeemMessagePacket struct {
	Success bool
	Message string
}

func (e *ExaltationRedeemMessagePacket) Read(p *network.GamePacket) {
	e.Success = p.ReadBool()
	e.Message = p.ReadString()
}

func (e ExaltationRedeemMessagePacket) Write(p *network.GamePacket) {
	p.WriteBool(e.Success)
	p.WriteString(e.Message)
}
