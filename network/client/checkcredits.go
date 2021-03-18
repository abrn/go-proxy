package client

import "proxy/network"

type CheckCreditsPacket struct{}

func (c *CheckCreditsPacket) Read(p *network.Packet) {}

func (c CheckCreditsPacket) Write(p *network.Packet) {}
