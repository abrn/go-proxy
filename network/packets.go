package network

import (
	"encoding/binary"
	"math"
)

type Packet struct {
	Index  int
	Length uint
	ID     int
	Data   []byte
}

func (p *Packet) Advance(amount int) int {
	p.Index += amount
	return amount
}

func (p *Packet) ReadString() string {
	n := int(p.ReadUInt16())
	if n == 0 {
		return ""
	}
	var str []byte
	str = p.Data[p.Index : p.Index+n]
	p.Advance(n)
	return string(str)
}

func (p *Packet) WriteString(s string) {
	if s == "" {
		p.WriteUInt16(uint16(0))
		return
	}
	p.WriteUInt16(uint16(len(s)))
	for i := range s {
		p.WriteByte(s[i])
	}
}

// ReadUTFString - read a 16bit UTF string from the buffer
func (p *Packet) ReadUTFString() string {
	n := int(p.ReadUInt32())
	if n == 0 {
		return ""
	}
	var str []byte
	str = p.Data[p.Index : p.Index+n]
	p.Advance(n)
	return string(str)
}

// WriteUTFString - write a 16bit UTF string to the buffer
func (p *Packet) WriteUTFString(s string) {
	if s == "" {
		p.WriteUInt32(0)
		return
	}
	p.WriteUInt32(uint32(len(s)))
	for i := range s {
		p.WriteByte(s[i])
	}
}

// ReadBool - read a single byte to a boolean
func (p *Packet) ReadBool() bool {
	if p.ReadByte() == 1 {
		return true
	}
	return false
}

// WriteBool - write a boolean as a single byte being either 1 or 0
func (p *Packet) WriteBool(b bool) {
	if b == true {
		p.WriteByte(1)
	} else {
		p.WriteByte(0)
	}
}

// ReadFloat - read a 32bit floating point
func (p *Packet) ReadFloat() float32 {
	return math.Float32frombits(p.ReadUInt32())
}

// WriteFloat - write a 32bit floating point
func (p *Packet) WriteFloat(f float32) {
	binary.BigEndian.PutUint32(p.Data[p.Index:p.Index+p.Advance(4)], math.Float32bits(f))
}

// ReadInt16 - read a 16bit integer
func (p *Packet) ReadInt16() int16 {
	return int16(binary.BigEndian.Uint16(p.Data[p.Index : p.Index+p.Advance(2)]))
}

// WriteInt16 - write an unsigned 16bit int
func (p *Packet) WriteInt16(i int16) {
	binary.BigEndian.PutUint16(p.Data[p.Index:p.Index+p.Advance(2)], uint16(i))
}

// ReadUInt16 - read an unsigned 16bit int
func (p *Packet) ReadUInt16() uint16 {
	return binary.BigEndian.Uint16(p.Data[p.Index : p.Index+p.Advance(2)])
}

// WriteUInt16 - write an unsigned 16bit int
func (p *Packet) WriteUInt16(i uint16) {
	binary.BigEndian.PutUint16(p.Data[p.Index:p.Index+p.Advance(2)], i)
}

// ReadInt32 - read a 32bit integer
func (p *Packet) ReadInt32() int32 {
	i := int32(binary.BigEndian.Uint32(p.Data[p.Index : p.Index+4]))
	p.Advance(4)
	return i
}

// WriteInt32 - write a 32bit integer
func (p *Packet) WriteInt32(i int32) {
	binary.BigEndian.PutUint32(p.Data[p.Index:p.Index+p.Advance(4)], uint32(i))
}

// ReadUInt32 - read an unsigned 32bit int
func (p *Packet) ReadUInt32() uint32 {
	i := binary.BigEndian.Uint32(p.Data[p.Index : p.Index+4])
	p.Advance(4)
	return i
}

// WriteUInt32 - write an unsigned 32bit int
func (p *Packet) WriteUInt32(i uint32) {
	binary.BigEndian.PutUint32(p.Data[p.Index:p.Index+p.Advance(4)], i)
}

// ReadByte - read a single byte from the buffer
func (p *Packet) ReadByte() byte {
	b := p.Data[p.Index : p.Index+1][0]
	p.Advance(1)
	return b
}

// WriteByte - write a single byte to the buffer
func (p *Packet) WriteByte(b byte) {
	p.Data[p.Index] = b
	p.Advance(1)
}

// ReadBytes - read a certain amount of bytes
func (p *Packet) ReadBytes(amount int) []byte {
	return p.Data[p.Index : p.Index+p.Advance(amount)]
}

// ReadCompressed - read a compressed integer from the buffer
func (p *Packet) ReadCompressed() int32 {
	var value uint32 = 0
	data := uint32(p.ReadByte())
	negative := true
	if data&64 == 0 {
		negative = false
	}
	var mask uint32 = 6
	value = uint32(data) & 63
	for data&128 > 0 {
		data = uint32(p.ReadByte())
		value = value | (data&127)<<mask
		mask = mask + 7
	}
	if negative == true {
		value = -value
	}
	return int32(value)
}

// WriteCompressed - write a compressed integer to the buffer
func (p *Packet) WriteCompressed(i int32) {
	negative := i < 0
	if negative {
		i = -i
	}
	val := uint32(i)
	b := byte(val & 63)
	if negative {
		b |= 64
	}
	// todo: fix this
	val >>= 6
	positive := val > 0
	if positive {
		val |= 128
	}
	p.WriteByte(b)
	p.Advance(1)
	for val > 0 {
		b = byte(val & 127)
		val >>= 7
		if val > 0 {
			b |= 128
		}
		p.WriteByte(b)
		p.Advance(1)
	}
}
