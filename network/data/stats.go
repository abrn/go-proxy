package data

import "proxy/network"

type ConditionEffect struct {
	Condition byte
}

func (c *ConditionEffect) Read(p *network.Packet) {
	c.Condition = p.ReadByte()
}

func (c ConditionEffect) Write(p *network.Packet) {
	p.WriteByte(c.Condition)
}
