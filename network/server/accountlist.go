package server

import "proxy/network"

type AccountListPacket struct {
	AccountListID int32
	AccountIDs    []string
	LockAction    int32
}

func (c *AccountListPacket) Read(p *network.GamePacket) {
	c.AccountListID = p.ReadInt32()
	count := p.ReadInt16()
	c.AccountIDs = make([]string, count)
	for i := 0; i < int(count); i++ {
		c.AccountIDs[i] = p.ReadString()
	}
	c.LockAction = p.ReadInt32()
}

func (c AccountListPacket) Write(p *network.GamePacket) {
	p.WriteInt32(c.AccountListID)
	count := len(c.AccountIDs)
	p.WriteInt16(int16(count))
	if count > 0 {
		for i := 0; i < count; i++ {
			p.WriteString(c.AccountIDs[i])
		}
	}
	p.WriteInt32(c.LockAction)
}
