package client

import "proxy/network"

type JoinGuildPacket struct {
	Name string
}

func (j *JoinGuildPacket) Read(p *network.Packet) {
	j.Name = p.ReadString()
}

func (j JoinGuildPacket) Write(p *network.Packet) {
	p.WriteString(j.Name)
}
