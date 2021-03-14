package client

import "proxy/network"

type UpdateAckPacket struct {}

func (c *UpdateAckPacket) Read(p *network.Packet) {}

func (c UpdateAckPacket) Write(p *network.Packet) {}