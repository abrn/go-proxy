package server

import "proxy/network"

type KeyInfoResponsePacket struct {
	Name    string
	Message string
	Creator string
}

func (k *KeyInfoResponsePacket) Read(p *network.Packet) {
	k.Name = p.ReadString()
	k.Message = p.ReadString()
	k.Creator = p.ReadString()
}

func (k KeyInfoResponsePacket) Write(p *network.Packet) {
	p.WriteString(k.Name)
	p.WriteString(k.Message)
	p.WriteString(k.Creator)
}
