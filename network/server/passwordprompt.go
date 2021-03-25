package server

import "proxy/network"

type PasswordPromptPacket struct {
	CleanPasswordStatus int32
}

func (pp *PasswordPromptPacket) Read(p *network.GamePacket) {
	pp.CleanPasswordStatus = p.ReadInt32()
}

func (pp PasswordPromptPacket) Write(p *network.GamePacket) {
	p.WriteInt32(pp.CleanPasswordStatus)
}
