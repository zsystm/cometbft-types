package core

import "time"

type Vote struct {
	Type               string    `bson:"type"`
	Height             uint64    `bson:"height"`
	Round              uint64    `bson:"round"`
	BlockId            BlockID   `bson:"blockId"`
	Timestamp          time.Time `bson:"timestamp"`
	ValidatorAddress   string    `bson:"validatorAddress"`
	ValidatorIndex     uint64    `bson:"validatorIndex"`
	Signature          string    `bson:"signature"`
	Extension          string    `bson:"extension"`
	ExtensionSignature string    `bson:"extensionSignature"`
}
