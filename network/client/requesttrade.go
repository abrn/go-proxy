package client

import "proxy/network"

type RequestTradePacket struct {
	Username string // of the player to request a trade with
}

func (r *RequestTradePacket) Read(p *network.Packet) {
	r.Username = p.ReadString()
}

func (r RequestTradePacket) Write(p *network.Packet) {
	p.WriteString(r.Username)
}
