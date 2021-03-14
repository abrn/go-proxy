package server

import "proxy/network"

type QuestRedeemResponsePacket struct {
	Success bool
	Message string
}

func (c *QuestRedeemResponsePacket) Read(p *network.Packet) {
	c.Success = p.ReadBool()
	c.Message = p.ReadString()
}

func (c QuestRedeemResponsePacket) Write(p *network.Packet) {
	p.WriteBool(c.Success)
	p.WriteString(c.Message)
}
