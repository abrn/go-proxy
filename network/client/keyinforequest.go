package client

import "proxy/network"

type KeyInfoRequestPacket struct {
	ItemType int32
}

func (k *KeyInfoRequestPacket) Read(p *network.Packet) {
	k.ItemType = p.ReadInt32()
}

func (k KeyInfoRequestPacket) Write(p *network.Packet) {
	p.WriteInt32(k.ItemType)
}
