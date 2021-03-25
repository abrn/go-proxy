package client

import "proxy/network"

type PlayerTextPacket struct {
	Message string
}

func (pt *PlayerTextPacket) Read(p *network.GamePacket) {
	pt.Message = p.ReadString()
}

func (pt PlayerTextPacket) Write(p *network.GamePacket) {
	p.WriteString(pt.Message)
}
