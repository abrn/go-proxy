package server

import "proxy/network"

type DeathPacket struct {
	AccountID  string
	CharID     int32
	KilledBy   string
	ZombieType int32
	ZombieID   int32
	UnknownOne []UnknownDeathType
	UnknownTwo string
}

type UnknownDeathType struct {
	String string
	IntOne int32
	IntTwo int32
}

func (d *DeathPacket) Read(p *network.Packet) {
	d.AccountID = p.ReadString()
	d.CharID = p.ReadInt32()
	d.KilledBy = p.ReadString()
	d.ZombieType = p.ReadInt32()
	d.ZombieID = p.ReadInt32()
	ucount := p.ReadInt16()
	if ucount > 0 {
		d.UnknownOne = make([]UnknownDeathType, ucount)
		for i := 0; i < int(ucount); i++ {
			d.UnknownOne[i] = UnknownDeathType{}
			d.UnknownOne[i].String = p.ReadString()
			d.UnknownOne[i].IntOne = p.ReadInt32()
			d.UnknownOne[i].IntTwo = p.ReadInt32()
		}
	}
	d.UnknownTwo = p.ReadString()
}

// todo: fix writer after unknown struct reversed
func (d DeathPacket) Write(p *network.Packet) {
	p.WriteString(d.AccountID)
	p.WriteInt32(d.CharID)
	p.WriteString(d.KilledBy)
	p.WriteInt32(d.ZombieType)
	p.WriteInt32(d.ZombieID)
	//p.WriteInt32(d.Unknown)
}
