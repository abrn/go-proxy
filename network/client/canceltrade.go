package client

import "proxy/network"

type CancelTradePacket struct{}

func (c *CancelTradePacket) Read(p *network.GamePacket) {}

func (c CancelTradePacket) Write(p *network.GamePacket) {}
