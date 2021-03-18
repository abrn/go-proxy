package client

import "proxy/network"

type CancelTradePacket struct{}

func (c *CancelTradePacket) Read(p *network.Packet) {}

func (c CancelTradePacket) Write(p *network.Packet) {}
