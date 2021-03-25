package server

import "proxy/network"

type NotificationPacket struct {
	ObjectID int32
	Message  string
	Color    int32
}

func (n *NotificationPacket) Read(p *network.GamePacket) {
	n.ObjectID = p.ReadInt32()
	n.Message = p.ReadString()
	n.Color = p.ReadInt32()
}

func (n NotificationPacket) Write(p *network.GamePacket) {
	p.WriteInt32(n.ObjectID)
	p.WriteString(n.Message)
	p.WriteInt32(n.Color)
}
