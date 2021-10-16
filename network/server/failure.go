package server

import "proxy/network"

type FailurePacket struct {
	ID      int32
	Message string
}

type FailureCode int32

const (
	FailureBadVersion   FailureCode = 4
	FailureBadKey       FailureCode = 5
	FailureBadTeleport  FailureCode = 6
	FailureEmailNeeded  FailureCode = 7
	FailureUnknown1     FailureCode = 8
	FailureTeleCooldown FailureCode = 9
	FailureWrongServer  FailureCode = 10
	FailureServerFull   FailureCode = 14
	FailureServerQueue  FailureCode = 15
	FailureUnknown2     FailureCode = 16
	FailureUnknown3     FailureCode = 1
)

type ProtocolErrorCode int32

const (
	ProtoInvalidMove      ProtocolErrorCode = 5
	ProtoInvalidPong      ProtocolErrorCode = 9
	ProtoInvalidSerial    ProtocolErrorCode = 10
	ProtoInvalidUpdateAck ProtocolErrorCode = 11
	ProtoInvalidHello     ProtocolErrorCode = 15
	ProtoIgnoredAck       ProtocolErrorCode = 21
	ProtoTooManyPackets   ProtocolErrorCode = 42
	ProtoTooManyEntities  ProtocolErrorCode = 48
	ProtoRateLimited      ProtocolErrorCode = 64
)

func (c *FailurePacket) Read(p *network.GamePacket) {
	c.ID = p.ReadInt32()
	c.Message = p.ReadString()
}

func (c FailurePacket) Write(p *network.GamePacket) {
	p.WriteInt32(c.ID)
	p.WriteString(c.Message)
}
