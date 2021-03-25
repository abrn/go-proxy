package client

import "proxy/network"

type ResetDailyQuestsPacket struct {
	Unknown int32
}

func (r *ResetDailyQuestsPacket) Read(p *network.GamePacket) {
	r.Unknown = p.ReadInt32()
}

func (r ResetDailyQuestsPacket) Write(p *network.GamePacket) {
	p.WriteInt32(r.Unknown)
}
