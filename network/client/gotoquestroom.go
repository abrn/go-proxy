package client

import "proxy/network"

type GotoQuestRoomPacket struct{}

func (c *GotoQuestRoomPacket) Read(p *network.Packet) {}

func (c GotoQuestRoomPacket) Write(p *network.Packet) {}
