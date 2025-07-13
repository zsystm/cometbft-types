package events

import "time"

type Event interface {
	SetNodeId(id string)
	SetValidatorAddress(address string)
	GetTimestamp() time.Time
}
