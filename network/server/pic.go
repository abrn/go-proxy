package server

import "proxy/network"

type PicPacket struct {
	Data BitmapData
}

type BitmapData struct {
	Width  int32
	Height int32
	Data   []byte
}

func (pp *PicPacket) Read(p *network.Packet) {
	pp.Data.Width = p.ReadInt32()
	pp.Data.Height = p.ReadInt32()
	count := p.ReadInt16()
	pp.Data.Data = make([]byte, count)
	for i := 0; i < int(count); i++ {
		pp.Data.Data[i] = p.ReadByte()
	}
}
