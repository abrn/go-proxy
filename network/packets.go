package network

import (
	"encoding/binary"
	"math"
	"proxy/log"
)

// Packet - a packet parsed from a GamePacket which can be
// hooked and modified/blocked before leaving the proxy
type Packet struct {
	Type   PacketType // the packet type mapping to its struct object
	ID     byte       // the games ID value of the packet
	Block  bool       // whether the packet should be blocked before sent/received
	Hooks  []string   // the list of modules hooking this packet
	Packet GamePacket // the underlying GamePacket object
}

// GamePacket - a raw game packet which can be sent
// and received from the game connection freely
type GamePacket struct {
	Index int    // the byte position of the reader
	Size  uint   // the size of the packet in bytes
	ID    byte   // the games ID value of the packet
	Data  []byte // the actual packet data in bytes, including the header
}

// CreateGamePacket - create a new game packet based on 5 header bytes
func CreateGamePacket(head []byte) *GamePacket {
	// discard nil packets
	if head == nil {
		return nil
	}
	// a game packet header is always 5 bytes
	if len(head) != 5 {
		log.Logger.Warn("Creating packet: got %d bytes instead of 5 in header\n", len(head))
		return nil
	}
	// read length from bytes 0-4 and ID from 5
	packet := new(GamePacket)
	packet.Index = 5 // 5 is where the data starts, since the header has been parsed
	packet.Size = uint(binary.BigEndian.Uint32(head[0:4]))
	packet.ID = head[4]
	// packets should never be >= 64KB
	if packet.Size < 0xFFFF {
		packet.Data = make([]byte, packet.Size+5)
	} else {
		log.Logger.Warn("Creating packet: large packet length in header: %d bytes\n", packet.Size)
		// create a 1MB buffer, should never happen unless theres a net issue
		packet.Data = make([]byte, 0x100000)
	}
	// copy the passed header bytes to the final packet
	packet.Data[0] = head[0]
	packet.Data[1] = head[1]
	packet.Data[2] = head[2]
	packet.Data[3] = head[3]
	packet.Data[4] = head[4]
	return packet
}

type PacketType byte

const (
	FAILURE                PacketType = 0
	HELLO                  PacketType = 1
	CLAIMDAILYREWARD       PacketType = 3
	DELETEPET              PacketType = 4
	REQUESTTRADE           PacketType = 5
	QUESTFETCHRESPONSE     PacketType = 6
	JOINGUILD              PacketType = 7
	PING                   PacketType = 8
	NEWTICK                PacketType = 9
	PLAYERTEXT             PacketType = 10
	USEITEM                PacketType = 11
	SERVERPLAYERSHOOT      PacketType = 12
	SHOWEFFECT             PacketType = 13
	TRADEACCEPTED          PacketType = 14
	GUILDREMOVE            PacketType = 15
	PETUPGRADEREQUEST      PacketType = 16
	ENTERARENA             PacketType = 17 // unused
	GOTO                   PacketType = 18
	INVSWAP                PacketType = 19
	OTHERHIT               PacketType = 20
	NAMERESULT             PacketType = 21
	BUYRESULT              PacketType = 22
	HATCHPET               PacketType = 23
	ACTIVEPETUPDATEREQUEST PacketType = 24
	ENEMYHIT               PacketType = 25
	GUILDRESULT            PacketType = 26
	EDITACCOUNTLIST        PacketType = 27
	TRADECHANGED           PacketType = 28
	_29                    PacketType = 29 // unused
	PLAYERSHOOT            PacketType = 30
	PONG                   PacketType = 31
	_32                    PacketType = 32 // unused
	RESKINPET              PacketType = 33
	TRADEDONE              PacketType = 34
	ENEMYSHOOT             PacketType = 35
	ACCEPTTRADE            PacketType = 36
	CHANGEGUILDRANK        PacketType = 37
	PLAYSOUND              PacketType = 38
	VERIFYEMAIL            PacketType = 39 // unused
	SQUAREHIT              PacketType = 40
	NEWABILITYUNLOCKED     PacketType = 41
	MOVE                   PacketType = 42
	_43                    PacketType = 43 // unused
	TEXT                   PacketType = 44
	RECONNECT              PacketType = 45
)

// Advance - move the buffer reader position
// forward and return the number of bytes moved
func (p *GamePacket) Advance(amount int) int {
	p.Index += amount
	return amount
}

// Backstep - safely move the buffer reader backwards
// and return the number of bytes moved
func (p *GamePacket) Backstep(amount int) int {
	left := p.Index - amount
	if left < 0 {
		p.Index = 0
		return left
	}
	return left
}

func (p *GamePacket) ReadString() string {
	n := int(p.ReadUInt16())
	if n == 0 {
		return ""
	}
	var str []byte
	str = p.Data[p.Index : p.Index+n]
	p.Advance(n)
	return string(str)
}

func (p *GamePacket) WriteString(s string) {
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
func (p *GamePacket) ReadUTFString() string {
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
func (p *GamePacket) WriteUTFString(s string) {
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
func (p *GamePacket) ReadBool() bool {
	if p.ReadByte() == 1 {
		return true
	}
	return false
}

// WriteBool - write a boolean as a single byte being either 1 or 0
func (p *GamePacket) WriteBool(b bool) {
	if b == true {
		p.WriteByte(1)
	} else {
		p.WriteByte(0)
	}
}

// ReadFloat - read a 32bit floating point
func (p *GamePacket) ReadFloat() float32 {
	return math.Float32frombits(p.ReadUInt32())
}

// WriteFloat - write a 32bit floating point
func (p *GamePacket) WriteFloat(f float32) {
	binary.BigEndian.PutUint32(p.Data[p.Index:p.Index+p.Advance(4)], math.Float32bits(f))
}

// ReadInt16 - read a 16bit integer
func (p *GamePacket) ReadInt16() int16 {
	return int16(binary.BigEndian.Uint16(p.Data[p.Index : p.Index+p.Advance(2)]))
}

// WriteInt16 - write an unsigned 16bit int
func (p *GamePacket) WriteInt16(i int16) {
	binary.BigEndian.PutUint16(p.Data[p.Index:p.Index+p.Advance(2)], uint16(i))
}

// ReadUInt16 - read an unsigned 16bit int
func (p *GamePacket) ReadUInt16() uint16 {
	return binary.BigEndian.Uint16(p.Data[p.Index : p.Index+p.Advance(2)])
}

// WriteUInt16 - write an unsigned 16bit int
func (p *GamePacket) WriteUInt16(i uint16) {
	binary.BigEndian.PutUint16(p.Data[p.Index:p.Index+p.Advance(2)], i)
}

// ReadInt32 - read a 32bit integer
func (p *GamePacket) ReadInt32() int32 {
	i := int32(binary.BigEndian.Uint32(p.Data[p.Index : p.Index+4]))
	p.Advance(4)
	return i
}

// WriteInt32 - write a 32bit integer
func (p *GamePacket) WriteInt32(i int32) {
	binary.BigEndian.PutUint32(p.Data[p.Index:p.Index+p.Advance(4)], uint32(i))
}

// ReadUInt32 - read an unsigned 32bit int
func (p *GamePacket) ReadUInt32() uint32 {
	i := binary.BigEndian.Uint32(p.Data[p.Index : p.Index+4])
	p.Advance(4)
	return i
}

// WriteUInt32 - write an unsigned 32bit int
func (p *GamePacket) WriteUInt32(i uint32) {
	binary.BigEndian.PutUint32(p.Data[p.Index:p.Index+p.Advance(4)], i)
}

// ReadByte - read a single byte from the buffer
func (p *GamePacket) ReadByte() byte {
	b := p.Data[p.Index : p.Index+1][0]
	p.Advance(1)
	return b
}

// WriteByte - write a single byte to the buffer
func (p *GamePacket) WriteByte(b byte) {
	p.Data[p.Index] = b
	p.Advance(1)
}

// ReadBytes - read a certain amount of bytes
func (p *GamePacket) ReadBytes(amount int) []byte {
	return p.Data[p.Index : p.Index+p.Advance(amount)]
}

// ReadCompressed - read a compressed integer from the buffer
func (p *GamePacket) ReadCompressed() int32 {
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
func (p *GamePacket) WriteCompressed(i int32) {
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
