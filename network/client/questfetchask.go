package client

import "proxy/network"

type QuestFetchAskPacket struct {}

func (q *QuestFetchAskPacket) Read(p *network.Packet) {}

func (q QuestFetchAskPacket) Write(p *network.Packet) {}