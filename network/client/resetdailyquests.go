package client

import "proxy/network"

type ResetDailyQuestsPacket struct {
	Unknown int32
}

func (r *ResetDailyQuestsPacket) Read(p *network.Packet) {
	r.Unknown = p.ReadInt32()
}

func (r ResetDailyQuestsPacket) Write(p *network.Packet) {
	p.WriteInt32(r.Unknown)
}
