package client

import "proxy/network"

// todo: REDEEMEXALTATIONREWARD test this
type RedeemExaltationRewardPacket struct{}

func (r *RedeemExaltationRewardPacket) Read(p *network.GamePacket) {}

func (r RedeemExaltationRewardPacket) Write(p *network.GamePacket) {}
