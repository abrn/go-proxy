package server

import (
	"proxy/network"
	"proxy/network/data"
)

type UpdatePacket struct {
	Tiles      []data.GroundTileData
	NewObjects []data.ObjectData
	Drops      []int32
}

func (u *UpdatePacket) Read(p *network.GamePacket) {
	tiles := int(p.ReadCompressed())
	if tiles > 0 {
		u.Tiles = make([]data.GroundTileData, tiles)
		for i := 0; i < tiles; i++ {
			u.Tiles[i] = data.GroundTileData{}
			u.Tiles[i].Read(p)
		}
	}
	objs := int(p.ReadCompressed())
	if objs > 0 {
		u.NewObjects = make([]data.ObjectData, objs)
		for i := 0; i < objs; i++ {
			u.NewObjects[i] = data.ObjectData{}
			u.NewObjects[i].Read(p)
		}
	}
	drops := int(p.ReadCompressed())
	if drops > 0 {
		u.Drops = make([]int32, drops)
		for i := 0; i < drops; i++ {
			u.Drops[i] = p.ReadCompressed()
		}
	}
}

func (u UpdatePacket) Write(p *network.GamePacket) {
	tiles := len(u.Tiles)
	p.WriteCompressed(int32(tiles))
	if tiles > 0 {
		for i := 0; i < tiles; i++ {
			u.Tiles[i].Write(p)
		}
	}
	objs := len(u.NewObjects)
	p.WriteCompressed(int32(objs))
	if objs > 0 {
		for i := 0; i < objs; i++ {
			u.NewObjects[i].Write(p)
		}
	}
	drops := len(u.Drops)
	p.WriteCompressed(int32(drops))
	if drops > 0 {
		for i := 0; i < drops; i++ {
			p.WriteCompressed(u.Drops[i])
		}
	}
}
