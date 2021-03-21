package server

import "proxy/network"

type ReconnectPacket struct {
	MapName   string
	Hostname  string
	Port      uint16
	GameID    int32
	KeyTime   int32
	Key       []byte
	FromArena bool
	Unknown   int32
}

func (r *ReconnectPacket) Read(p *network.Packet) {
	r.MapName = p.ReadString()
	r.Hostname = p.ReadString()
	r.Port = p.ReadUInt16()
	r.GameID = p.ReadInt32()
	r.KeyTime = p.ReadInt32()
	keyLen := p.ReadInt16()
	r.Key = make([]byte, keyLen)
	for i := 0; i < int(keyLen); i++ {
		r.Key[i] = p.ReadByte()
	}
	r.FromArena = p.ReadBool()
	r.Unknown = p.ReadInt32()
}

func (r ReconnectPacket) Write(p *network.Packet) {
	p.WriteString(r.MapName)
	p.WriteString(r.Hostname)
	p.WriteUInt16(r.Port)
	p.WriteInt32(r.GameID)
	p.WriteInt32(r.KeyTime)
	keyLen := len(r.Key)
	p.WriteInt16(int16(keyLen))
	for i := 0; i < keyLen; i++ {
		p.WriteByte(r.Key[i])
	}
	p.WriteBool(r.FromArena)
	p.WriteInt32(r.Unknown)
}
