package client

import "proxy/network"

type HelloPacket struct {
	BuildVersion  string
	GameId        int32
	AccessToken   string
	KeyTime       uint32
	Key           []byte
	MapJSON       string
	EntryTag      string
	GameNet       string
	GameNetUserID string
	PlayPlatform  string
	PlatformToken string
	//UserToken     string
	ClientToken   string
}

func (h *HelloPacket) Read(p *network.GamePacket) {
	h.BuildVersion = p.ReadString()
	h.GameId = p.ReadInt32()
	h.AccessToken = p.ReadString()
	h.KeyTime = p.ReadUInt32()
	keySize := p.ReadInt16()
	h.Key = p.ReadBytes(int(keySize))
	h.MapJSON = p.ReadUTFString()
	h.EntryTag = p.ReadString()
	h.GameNet = p.ReadString()
	h.GameNetUserID = p.ReadString()
	h.PlayPlatform = p.ReadString()
	h.PlatformToken = p.ReadString()
	//h.UserToken = p.ReadString()
	h.ClientToken = p.ReadString()
}

func (h HelloPacket) Write(p *network.GamePacket) {
	p.WriteString(h.BuildVersion)
	p.WriteInt32(h.GameId)
	p.WriteString(h.AccessToken)
	p.WriteUInt32(h.KeyTime)
	p.WriteInt16(int16(len(h.Key)))
	for i := 0; i < len(h.Key); i++ {
		p.WriteByte(h.Key[i])
	}
	p.WriteUTFString(h.MapJSON)
	p.WriteString(h.EntryTag)
	p.WriteString(h.GameNet)
	p.WriteString(h.GameNetUserID)
	p.WriteString(h.PlayPlatform)
	p.WriteString(h.PlatformToken)
	//p.WriteString(h.UserToken)
	p.WriteString(h.ClientToken)
}
