package server

import "proxy/network"

type FailurePacket struct {
	ID 		int32
	Message string
}

type FailureCode int32

const (
	FailureBadVersion	FailureCode = 4
	FailureBadKey 		FailureCode = 5
	FailureBadTeleport	FailureCode = 6
	FailureEmailNeeded	FailureCode = 7
	FailureTeleCooldown	FailureCode = 9
	FailureWrongServcer	FailureCode = 10
	FailureServerFull 	FailureCode = 11
	FailureServerQueue	FailureCode = 15
)

func (c *FailurePacket) Read(p *network.Packet) {
	c.ID = p.ReadInt32()
	c.Message = p.ReadString()
}

func (c FailurePacket) Write(p *network.Packet) {
	p.WriteInt32(c.ID)
	p.WriteString(c.Message)
}