package events

import "time"

type Event interface {
	SetNodeId(id string)
	SetValidatorAddress(address string)
	GetNodeId() string
	GetTimestamp() time.Time
	GetValidatorAddress() string
}

type HasHeight interface {
	GetHeight() uint64
}

type HasRound interface {
	GetRound() uint64
}

type HasHeightAndRound interface {
	HasHeight
	HasRound
}

type ConsensusEvent interface {
	Event
	HasHeightAndRound
}

type HasRawBytes interface {
	GetRawBytes() []byte
}
