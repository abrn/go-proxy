package server

import "proxy/network"

type GlobalNotificationPacket struct {
	TypeID  GlobalNotificationType
	Message string
}

type GlobalNotificationType int32

/**
Constants for notification types which will be converted on read, write and printing
*/
const (
	NotificationXP              GlobalNotificationType = 0
	NotificationDamage          GlobalNotificationType = 1
	NotificationFame            GlobalNotificationType = 2
	NotificationLevelUp         GlobalNotificationType = 3
	NotificationQuestComplete   GlobalNotificationType = 4
	NotificationClassUnlocked   GlobalNotificationType = 5
	NotificationConditionEffect GlobalNotificationType = 6
	NotificationNormal          GlobalNotificationType = 7
	NotificationExaltation      GlobalNotificationType = 8
)

func (g *GlobalNotificationPacket) Read(p *network.GamePacket) {
	g.TypeID = GlobalNotificationType(p.ReadInt32())
	g.Message = p.ReadString()
}

func (g GlobalNotificationPacket) Write(p *network.GamePacket) {
	p.WriteInt32(int32(g.TypeID))
	p.WriteString(g.Message)
}
