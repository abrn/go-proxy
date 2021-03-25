package server

import "proxy/network"

type VaultContentPacket struct {
	Unknown           bool // 1 or 0, unknown what this is for
	VaultObjectID     int32
	GiftObjectID      int32
	PotionObjectID    int32
	VaultContent      []int32 // each content value is an array of item IDs
	GiftContent       []int32
	PotionContent     []int32
	VaultUpgradeCost  int16
	PotionUpgradeCost int16
	MaxPotionCount    int16
	NextPotionCount   int16 // the max potion slots after the next upgrade
}

func (v *VaultContentPacket) Read(p *network.GamePacket) {
	v.Unknown = p.ReadBool()
	v.VaultObjectID = p.ReadCompressed()
	v.GiftObjectID = p.ReadCompressed()
	v.PotionObjectID = p.ReadCompressed()
	vaultCount := p.ReadCompressed()
	if vaultCount > 0 {
		v.VaultContent = make([]int32, vaultCount)
		for i := 0; i < int(vaultCount); i++ {
			v.VaultContent[i] = p.ReadCompressed()
		}
	}
	giftCount := p.ReadCompressed()
	if giftCount > 0 {
		v.GiftContent = make([]int32, giftCount)
		for i := 0; i < int(giftCount); i++ {
			v.GiftContent[i] = p.ReadCompressed()
		}
	}
	potionCount := p.ReadCompressed()
	if potionCount > 0 {
		v.PotionContent = make([]int32, potionCount)
		for i := 0; i < int(potionCount); i++ {
			v.PotionContent[i] = p.ReadCompressed()
		}
	}
	v.VaultUpgradeCost = p.ReadInt16()
	v.PotionUpgradeCost = p.ReadInt16()
	v.MaxPotionCount = p.ReadInt16()
	v.NextPotionCount = p.ReadInt16()
}

func (v VaultContentPacket) Write(p *network.GamePacket) {
	p.WriteBool(v.Unknown)
	p.WriteCompressed(v.VaultObjectID)
	p.WriteCompressed(v.GiftObjectID)
	p.WriteCompressed(v.PotionObjectID)
	vaultCount := len(v.VaultContent)
	p.WriteCompressed(int32(vaultCount))
	if vaultCount > 0 {
		for i := 0; i < vaultCount; i++ {
			p.WriteCompressed(v.VaultContent[i])
		}
	}
	giftCount := len(v.GiftContent)
	p.WriteCompressed(int32(giftCount))
	if giftCount > 0 {
		for i := 0; i < giftCount; i++ {
			p.WriteCompressed(v.GiftContent[i])
		}
	}
	potionCount := len(v.PotionContent)
	p.WriteCompressed(int32(potionCount))
	if potionCount > 0 {
		for i := 0; i < potionCount; i++ {
			p.WriteCompressed(v.PotionContent[i])
		}
	}
	p.WriteInt16(v.VaultUpgradeCost)
	p.WriteInt16(v.PotionUpgradeCost)
	p.WriteInt16(v.MaxPotionCount)
	p.WriteInt16(v.NextPotionCount)
}
