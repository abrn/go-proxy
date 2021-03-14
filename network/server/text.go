package server

import "proxy/network"

type TextPacket struct {
	Name         string
	ObjectID     int32
	NumStars     int32
	BubbleTime   int16
	Recipient    string
	Message      string
	CleanMessage string
	Supporter    bool
	StarBg       int32
}

func (t *TextPacket) Read(p *network.Packet) {
	t.Name = p.ReadString()
	t.ObjectID = p.ReadInt32()
	t.NumStars = p.ReadInt32()
	t.BubbleTime = p.ReadInt16()
	t.Recipient = p.ReadString()
	t.Message = p.ReadString()
	t.CleanMessage = p.ReadString()
	t.Supporter = p.ReadBool()
	t.StarBg = p.ReadInt32()
}

func (t TextPacket) Write(p *network.Packet) {
	p.WriteString(t.Name)
	p.WriteInt32(t.ObjectID)
	p.WriteInt32(t.NumStars)
	p.WriteInt16(t.BubbleTime)
	p.WriteString(t.Recipient)
	p.WriteString(t.Message)
	p.WriteString(t.CleanMessage)
	p.WriteBool(t.Supporter)
	p.WriteInt32(t.StarBg)
}
