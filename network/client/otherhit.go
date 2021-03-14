package client

import "proxy/network"

type OtherHitPacket struct {
	Time     int32
	BulletID byte
	ObjectID int32
	TargetID int32
}

func (c *OtherHitPacket) Read(p *network.Packet) {
	c.Time = p.ReadInt32()
	c.BulletID = p.ReadByte()
	c.ObjectID = p.ReadInt32()
	c.TargetID = p.ReadInt32()
}

func (c OtherHitPacket) Write(p *network.Packet) {
	p.WriteInt32(c.Time)
	p.WriteByte(c.BulletID)
	p.WriteInt32(c.ObjectID)
	p.WriteInt32(c.TargetID)
}
