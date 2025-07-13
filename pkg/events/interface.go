package events

import "time"

type Event interface {
	SetNodeId(id string)
	SetValidatorAddress(address string)
	GetNodeId() string
	GetTimestamp() time.Time
}
