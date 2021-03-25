package client

import "proxy/network"

type JoinGuildPacket struct {
	Name string
}

func (j *JoinGuildPacket) Read(p *network.GamePacket) {
	j.Name = p.ReadString()
}

func (j JoinGuildPacket) Write(p *network.GamePacket) {
	p.WriteString(j.Name)
}
