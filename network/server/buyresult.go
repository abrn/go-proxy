package server

import "proxy/network"

type BuyResultPacket struct {
	ResultCode int32
	Message    string
}

func (b *BuyResultPacket) Read(p *network.Packet) {
	b.ResultCode = p.ReadInt32()
	b.Message = p.ReadString()
}

func (b BuyResultPacket) Write(p *network.Packet) {
	p.WriteInt32(b.ResultCode)
	p.WriteString(b.Message)
}
