package client

import "proxy/network"

type AcceptArenaDeathPacket struct{}

func (c *AcceptArenaDeathPacket) Read(p *network.GamePacket) {}

func (c AcceptArenaDeathPacket) Write(p *network.GamePacket) {}
