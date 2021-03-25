package client

import "proxy/network"

type ChangePetSkinPacket struct {
	PetID    int32
	SkinType int32
	Currency int32
}

func (c *ChangePetSkinPacket) Read(p *network.GamePacket) {
	c.PetID = p.ReadInt32()
	c.SkinType = p.ReadInt32()
	c.Currency = p.ReadInt32()
}

func (c ChangePetSkinPacket) Write(p *network.GamePacket) {
	p.WriteInt32(c.PetID)
	p.WriteInt32(c.SkinType)
	p.WriteInt32(c.Currency)
}
