package events

import "time"

type Event interface {
	SetValidatorAddress(address string)
	GetValidatorAddress() string
	GetEventType() EventTyp
	SetNodeId(id string)
	GetNodeId() string
	GetTimestamp() time.Time
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
