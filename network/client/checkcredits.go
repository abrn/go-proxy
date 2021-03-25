package client

import "proxy/network"

type CheckCreditsPacket struct{}

func (c *CheckCreditsPacket) Read(p *network.GamePacket) {}

func (c CheckCreditsPacket) Write(p *network.GamePacket) {}
