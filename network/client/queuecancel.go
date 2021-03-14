package client

import "proxy/network"

// todo: QUEUECANCEL test this, ID 113
type QueueCancelPacket struct {}

func (q *QueueCancelPacket) Read(p *network.Packet) {}

func (q QueueCancelPacket) Write(p *network.Packet) {}