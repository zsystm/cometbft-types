package core

import "time"

type Proposal struct {
	Height    uint64    `bson:"height" json:"height"`
	Round     uint64    `bson:"round" json:"round"`
	PolRound  int64     `bson:"polRound" json:"polRound"`
	BlockID   BlockID   `bson:"blockId" json:"blockId"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	Signature string    `bson:"signature" json:"signature"`
}

func (p *Proposal) GetHeight() uint64 {
	return p.Height
}

func (p *Proposal) GetRound() uint64 {
	return p.Round
}
