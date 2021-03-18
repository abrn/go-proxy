package client

import "proxy/network"

// todo: changeallyshoot - check if this is 121

type ChangeAllyShootPacket struct {
}

func (c *ChangeAllyShootPacket) Read(p *network.Packet) {}

func (c ChangeAllyShootPacket) Write(p *network.Packet) {}
