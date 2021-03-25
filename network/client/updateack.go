package client

import "proxy/network"

type UpdateAckPacket struct{}

func (c *UpdateAckPacket) Read(p *network.GamePacket) {}

func (c UpdateAckPacket) Write(p *network.GamePacket) {}
