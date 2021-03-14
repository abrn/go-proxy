package client

import "proxy/network"

type EditAccountListPacket struct {
	AccountListID int32
	Add           bool
	ObjectID      int32
}

func (e *EditAccountListPacket) Read(p *network.Packet) {
	e.AccountListID = p.ReadInt32()
	e.Add = p.ReadBool()
	e.ObjectID = p.ReadInt32()
}

func (e EditAccountListPacket) Write(p *network.Packet) {
	p.WriteInt32(e.AccountListID)
	p.WriteBool(e.Add)
	p.WriteInt32(e.ObjectID)
}
