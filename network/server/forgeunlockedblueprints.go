package server

import "proxy/network"

type ForgeUnlockedBlueprintsPacket struct {
	Blueprints []int32 // an array of blueprint item IDs
}

func (f *ForgeUnlockedBlueprintsPacket) Read(p *network.Packet) {
	count := p.ReadByte()
	f.Blueprints = make([]int32, count)
	if count <= 0 {
		return
	}
	for i := 0; i < int(count); i++ {
		f.Blueprints[i] = p.ReadCompressed()
	}
}

func (f ForgeUnlockedBlueprintsPacket) Write(p *network.Packet) {
	count := len(f.Blueprints)
	p.WriteByte(byte(count))
	if count <= 0 {
		return
	}
	for i := 0; i < count; i++ {
		p.WriteCompressed(f.Blueprints[i])
	}
}
