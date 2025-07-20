package events

import "time"

type BaseEvent struct {
	Type             EventTyp  `bson:"type" json:"type"`
	Timestamp        time.Time `bson:"timestamp" json:"timestamp"`
	NodeId           string    `bson:"nodeId" json:"nodeId"`
	ValidatorAddress string    `bson:"validatorAddress" json:"validatorAddress"`
}

func (e *BaseEvent) SetNodeId(id string) {
	e.NodeId = id
}

func (e *BaseEvent) SetValidatorAddress(address string) {
	e.ValidatorAddress = address
}

func (e *BaseEvent) GetNodeId() string {
	return e.NodeId
}

func (e *BaseEvent) GetTimestamp() time.Time {
	return e.Timestamp
}
