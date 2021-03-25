package server

import "proxy/network"

type QuestRedeemResponsePacket struct {
	Success bool
	Message string
}

func (c *QuestRedeemResponsePacket) Read(p *network.GamePacket) {
	c.Success = p.ReadBool()
	c.Message = p.ReadString()
}

func (c QuestRedeemResponsePacket) Write(p *network.GamePacket) {
	p.WriteBool(c.Success)
	p.WriteString(c.Message)
}
