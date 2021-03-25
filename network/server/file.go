package server

import "proxy/network"

type FilePacket struct {
	Name string
	Data []byte
}

func (f *FilePacket) Read(p *network.GamePacket) {
	f.Name = p.ReadString()
	count := p.ReadUInt32()
	f.Data = make([]byte, count)
	for i := 0; i < int(count); i++ {
		f.Data[i] = p.ReadByte()
	}
}

func (f FilePacket) Write(p *network.GamePacket) {
	p.WriteString(f.Name)
	count := len(f.Data)
	for i := 0; i < count; i++ {
		p.WriteByte(f.Data[i])
	}
}
