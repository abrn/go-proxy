package client

import "proxy/network"

type EscapePacket struct {}

func (e EscapePacket) Write(p *network.Packet) {}