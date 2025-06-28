package core

import "time"

type Proposal struct {
	Height    uint64    `bson:"height"`
	Round     uint64    `bson:"round"`
	PolRound  int64     `bson:"polRound"`
	BlockID   BlockID   `bson:"blockIid"`
	Timestamp time.Time `bson:"timestamp"`
	Signature string    `bson:"signature"`
}
