package client

import "proxy/network"

type QuestFetchAskPacket struct{}

func (q *QuestFetchAskPacket) Read(p *network.GamePacket) {}

func (q QuestFetchAskPacket) Write(p *network.GamePacket) {}
