package server

import "proxy/network"

type PasswordPromptPacket struct {
	CleanPasswordStatus int32
}

func (pp *PasswordPromptPacket) Read(p *network.Packet) {
	pp.CleanPasswordStatus = p.ReadInt32()
}

func (pp PasswordPromptPacket) Write(p *network.Packet) {
	p.WriteInt32(pp.CleanPasswordStatus)
}
