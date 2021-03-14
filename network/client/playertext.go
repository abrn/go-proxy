package client

import "proxy/network"

type PlayerTextPacket struct {
	Message string
}

func (pt *PlayerTextPacket) Read(p *network.Packet) {
	pt.Message = p.ReadString()
}

func (pt PlayerTextPacket) Write(p *network.Packet) {
	p.WriteString(pt.Message)
}
