package client

import "proxy/network"

type KeyInfoRequestPacket struct {
	ItemType int32
}

func (k *KeyInfoRequestPacket) Read(p *network.GamePacket) {
	k.ItemType = p.ReadInt32()
}

func (k KeyInfoRequestPacket) Write(p *network.GamePacket) {
	p.WriteInt32(k.ItemType)
}
