package server

import "proxy/network"

type NameResultPacket struct {
	Success bool
	Message string
}

func (n *NameResultPacket) Read(p *network.Packet) {
	n.Success = p.ReadBool()
	n.Message = p.ReadString()
}

func (n NameResultPacket) Write(p *network.Packet) {
	p.WriteBool(n.Success)
	p.WriteString(n.Message)
}