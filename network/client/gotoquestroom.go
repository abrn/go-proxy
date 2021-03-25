package client

import "proxy/network"

type GotoQuestRoomPacket struct{}

func (c *GotoQuestRoomPacket) Read(p *network.GamePacket) {}

func (c GotoQuestRoomPacket) Write(p *network.GamePacket) {}
