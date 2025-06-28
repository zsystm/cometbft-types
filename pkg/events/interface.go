package events

type Event interface {
	SetNodeId(id string)
	SetValidatorAddress(address string)
}
