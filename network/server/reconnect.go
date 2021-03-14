package server

import "proxy/network"

type ReconnectPacket struct {
	MapName   string
	Hostname  string
	Port      int32
	GameID    int32
	KeyTime   int32
	FromArena bool
	Key       []byte
}

func (r *ReconnectPacket) Read(p *network.Packet) {
	r.MapName = p.ReadString()
	r.Hostname = p.ReadString()
	r.Port = p.ReadInt32()
	r.GameID = p.ReadInt32()
	r.KeyTime = p.ReadInt32()
	r.FromArena = p.ReadBool()
	keyLen := p.ReadInt16()
	r.Key = make([]byte, keyLen)
	for i := 0; i < int(keyLen); i++ {
		r.Key[i] = p.ReadByte()
	}
}

func (r ReconnectPacket) Write(p *network.Packet) {
	p.WriteString(r.MapName)
	p.WriteString(r.Hostname)
	p.WriteInt32(r.Port)
	p.WriteInt32(r.GameID)
	p.WriteInt32(r.KeyTime)
	p.WriteBool(r.FromArena)
	keyLen := len(r.Key)
	p.WriteInt16(int16(keyLen))
	for i := 0; i < keyLen; i++ {
		p.WriteByte(r.Key[i])
	}
}
