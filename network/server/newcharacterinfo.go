package server

import "proxy/network"

type NewCharacterInfoPacket struct {
	CharXML string
}

func (n *NewCharacterInfoPacket) Read(p *network.Packet) {
	n.CharXML = p.ReadString()
}

func (n NewCharacterInfoPacket) Write(p *network.Packet) {
	p.WriteString(n.CharXML)
}
