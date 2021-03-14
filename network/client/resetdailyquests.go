package client

import "proxy/network"

type ResetDailyQuestsPacket struct{}

func (c *ResetDailyQuestsPacket) Read(p *network.Packet) {}

func (c ResetDailyQuestsPacket) Write(p *network.Packet) {}
