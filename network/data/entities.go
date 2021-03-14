package data

import "proxy/network"

// ObjectData - the primary data of an in-game entity
type ObjectData struct {
	Type   uint16
	Status ObjectStatusData
}

func (o *ObjectData) Read(p *network.Packet) {
	o.Type = p.ReadUInt16()
	o.Status = ObjectStatusData{}
	o.Status.Read(p)
}

func (o ObjectData) Write(p *network.Packet) {
	p.WriteUInt16(o.Type)
	o.Status.Write(p)
}

// ObjectStatusData - the current status of an in-game entity
type ObjectStatusData struct {
	ObjectID int32
	Position WorldPosData
	Stats    []StatData // todo: FIX STATDATA
}

func (o *ObjectStatusData) Read(p *network.Packet) {
	o.ObjectID = p.ReadCompressed()
	o.Position = WorldPosData{}
	o.Position.Read(p)
	statCount := p.ReadCompressed()
	o.Stats = make([]StatData, statCount)
	if statCount <= 0 {
		return
	}
	for i := 0; i < int(statCount); i++ {
		o.Stats[i] = StatData{}
		o.Stats[i].Read(p)
	}
}

func (o ObjectStatusData) Write(p *network.Packet) {
	p.WriteCompressed(o.ObjectID)
	o.Position.Write(p)
	statCount := len(o.Stats)
	p.WriteCompressed(int32(statCount))
	if statCount <= 0 {
		return
	}
	for i := 0; i < statCount; i++ {
		o.Stats[i].Write(p)
	}
}