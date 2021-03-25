package data

import "proxy/network"

type ConditionEffect struct {
	Condition byte
}

func (c *ConditionEffect) Read(p *network.GamePacket) {
	c.Condition = p.ReadByte()
}

func (c ConditionEffect) Write(p *network.GamePacket) {
	p.WriteByte(c.Condition)
}
