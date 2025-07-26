package core

import "time"

type Vote struct {
	Type               string    `bson:"type" json:"type"`
	Height             uint64    `bson:"height" json:"height"`
	Round              uint64    `bson:"round" json:"round"`
	BlockId            BlockID   `bson:"blockId" json:"blockId"`
	Timestamp          time.Time `bson:"timestamp" json:"timestamp"`
	ValidatorAddress   string    `bson:"validatorAddress" json:"validatorAddress"`
	ValidatorIndex     uint64    `bson:"validatorIndex" json:"validatorIndex"`
	Signature          string    `bson:"signature" json:"signature"`
	Extension          string    `bson:"extension" json:"extension"`
	ExtensionSignature string    `bson:"extensionSignature" json:"extensionSignature"`
}

func (v *Vote) GetHeight() uint64 {
	return v.Height
}

func (v *Vote) GetRound() uint64 {
	return v.Round
}
