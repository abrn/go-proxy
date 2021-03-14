package client

import "proxy/network"

// todo: REDEEMEXALTATIONREWARD test this
type RedeemExaltationRewardPacket struct {}

func (r *RedeemExaltationRewardPacket) Read(p *network.Packet) {}

func (r RedeemExaltationRewardPacket) Write(p *network.Packet) {}