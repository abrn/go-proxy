package data

import "proxy/network"

// GroundTileData - represents a tile in the map with an X,Y coordinate and tile object type
type GroundTileData struct {
	X    int16
	Y    int16
	Type uint16
}

func (t *GroundTileData) Read(p *network.GamePacket) {
	t.X = p.ReadInt16()
	t.Y = p.ReadInt16()
	t.Type = p.ReadUInt16()
}

func (t GroundTileData) Write(p *network.GamePacket) {
	p.WriteInt16(t.X)
	p.WriteInt16(t.Y)
	p.WriteUInt16(t.Type)
}

// MapTile - extends GroundTileData to include more information useful for mods
type MapTile struct {
	Data       GroundTileData
	Occupied   bool
	OccupiedBy int
	Damage     bool
	LastDamage int
	Sink       bool
}

// MoveRecord - compiled together in the MOVE packet for a time-series list of client positions
type MoveRecord struct {
	Time int32
	X    float32
	Y    float32
}

func (mv *MoveRecord) Read(p *network.GamePacket) {
	mv.Time = p.ReadInt32()
	mv.X = p.ReadFloat()
	mv.Y = p.ReadFloat()
}

func (mv MoveRecord) Write(p *network.GamePacket) {
	p.WriteInt32(mv.Time)
	p.WriteFloat(mv.X)
	p.WriteFloat(mv.Y)
}

// WorldPosData - used in packets to represent the positions of entities/objects on the map
type WorldPosData struct {
	X float32
	Y float32
}

func (wp *WorldPosData) Read(p *network.GamePacket) {
	wp.X = p.ReadFloat()
	wp.X = p.ReadFloat()
}

func (wp WorldPosData) Write(p *network.GamePacket) {
	p.WriteFloat(wp.X)
	p.WriteFloat(wp.Y)
}
