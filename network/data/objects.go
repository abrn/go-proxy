package data

import "proxy/network"

// SlotObjectData - represents a slot in a player inventory or a container
type SlotObjectData struct {
	ObjectID   int32
	SlotID     int32
	ObjectType int32
}

func (s *SlotObjectData) Read(p *network.Packet) {
	s.ObjectID = p.ReadInt32()
	s.SlotID = p.ReadInt32()
	s.ObjectType = p.ReadInt32()
}

func (s SlotObjectData) Write(p *network.Packet) {
	p.WriteInt32(s.ObjectID)
	p.WriteInt32(s.SlotID)
	p.WriteInt32(s.ObjectType)
}

// TradeItem - included in most trade packets to represent an item involved in the trade
type TradeItem struct {
	ItemID    int32
	SlotType  int32
	Tradeable bool
	Included  bool
}

func (t *TradeItem) Read(p *network.Packet) {
	t.ItemID = p.ReadInt32()
	t.SlotType = p.ReadInt32()
	t.Tradeable = p.ReadBool()
	t.Included = p.ReadBool()
}

func (t TradeItem) Write(p *network.Packet) {
	p.WriteInt32(t.ItemID)
	p.WriteInt32(t.SlotType)
	p.WriteBool(t.Tradeable)
	p.WriteBool(t.Included)
}

// QuestData - information about a daily quest item
type QuestData struct {
	ID           string
	Name         string
	Description  string
	Expiration   string
	Category     int32
	Requirements []int32 // item IDs required to complete the quest
	Rewards      []int32 // item IDs of the rewards for completing
	Completed    bool
	ItemOfChoice bool
	Repeatable   bool
}

func (q *QuestData) Read(p *network.Packet) {
	q.ID = p.ReadString()
	q.Name = p.ReadString()
	q.Description = p.ReadString()
	q.Expiration = p.ReadString()
	q.Category = p.ReadInt32()
	reqs := p.ReadInt16()
	if reqs > 0 {
		q.Requirements = make([]int32, reqs)
		for i := 0; i < int(reqs); i++ {
			q.Requirements[i] = p.ReadInt32()
		}
	}
	rewards := p.ReadInt16()
	if rewards > 0 {
		q.Rewards = make([]int32, rewards)
		for i := 0; i < int(rewards); i++ {
			q.Rewards[i] = p.ReadInt32()
		}
	}
	q.Completed = p.ReadBool()
	q.ItemOfChoice = p.ReadBool()
	q.Repeatable = p.ReadBool()
}

func (q QuestData) Write(p *network.Packet) {
	p.WriteString(q.ID)
	p.WriteString(q.Name)
	p.WriteString(q.Description)
	p.WriteString(q.Expiration)
	p.WriteInt32(q.Category)
	reqs := len(q.Requirements)
	p.WriteInt16(int16(reqs))
	if reqs > 0 {
		for i := 0; i < reqs; i++ {
			p.WriteInt32(q.Requirements[i])
		}
	}
	rewards := len(q.Rewards)
	p.WriteInt16(int16(rewards))
	if rewards > 0 {
		for i := 0; i < rewards; i++ {
			p.WriteInt32(q.Rewards[i])
		}
	}
	p.WriteBool(q.Completed)
	p.WriteBool(q.ItemOfChoice)
	p.WriteBool(q.Repeatable)
}
