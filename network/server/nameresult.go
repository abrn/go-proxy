package server

import "proxy/network"

type NameResultPacket struct {
	Success bool
	Message string
}

func (n *NameResultPacket) Read(p *network.GamePacket) {
	n.Success = p.ReadBool()
	n.Message = p.ReadString()
}

func (n NameResultPacket) Write(p *network.GamePacket) {
	p.WriteBool(n.Success)
	p.WriteString(n.Message)
}
