package events

import "time"

type BaseEvent struct {
	Type             EventTyp  `bson:"type"`
	Timestamp        time.Time `bson:"timestamp"`
	NodeId           string    `bson:"nodeId"`
	ValidatorAddress string    `bson:"validatorAddress"`
}

func (e *BaseEvent) SetNodeId(id string) {
	e.NodeId = id
}

func (e *BaseEvent) SetValidatorAddress(address string) {
	e.ValidatorAddress = address
}
