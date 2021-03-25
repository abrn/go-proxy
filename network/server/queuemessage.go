package server

import "proxy/network"

// QueueMessagePacket - received when the client joins a server with a queue or their position changes
type QueueMessagePacket struct {
	CurrentPosition uint16
	MaxPosition     uint16
}

func (q *QueueMessagePacket) Read(p *network.GamePacket) {
	q.CurrentPosition = p.ReadUInt16()
	q.MaxPosition = p.ReadUInt16()
}

func (q QueueMessagePacket) Write(p *network.GamePacket) {
	p.WriteUInt16(q.CurrentPosition)
	p.WriteUInt16(q.MaxPosition)
}
