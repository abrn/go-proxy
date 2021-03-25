package server

import "proxy/network"

type NewCharacterInfoPacket struct {
	CharXML string
}

func (n *NewCharacterInfoPacket) Read(p *network.GamePacket) {
	n.CharXML = p.ReadString()
}

func (n NewCharacterInfoPacket) Write(p *network.GamePacket) {
	p.WriteString(n.CharXML)
}
