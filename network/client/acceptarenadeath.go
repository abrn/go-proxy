package client

import "proxy/network"

type AcceptArenaDeathPacket struct {}

func (c *AcceptArenaDeathPacket) Read(p *network.Packet) {}

func (c AcceptArenaDeathPacket) Write(p *network.Packet) {}