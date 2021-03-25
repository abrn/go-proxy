package server

import "proxy/network"

type InvitedToGuildPacket struct {
	Username  string
	GuildName string
}

func (i *InvitedToGuildPacket) Read(p *network.GamePacket) {
	i.Username = p.ReadString()
	i.GuildName = p.ReadString()
}

func (i InvitedToGuildPacket) Write(p *network.GamePacket) {
	p.WriteString(i.Username)
	p.WriteString(i.GuildName)
}
