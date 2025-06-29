package events

import "github.com/zsystm/cometbft-types/pkg/core"

type EventEnteringNewRound struct {
	BaseEvent  `bson:",inline"`
	Height     uint64 `bson:"height"`
	Round      uint64 `bson:"round"`
	Proposer   string `bson:"proposer"`
	PrevHeight uint64 `bson:"previousHeight"`
	PrevRound  uint64 `bson:"previousRound"`
	PrevStep   string `bson:"previousStep"`
}

type EventStepChange struct {
	BaseEvent  `bson:",inline"`
	CurrHeight uint64 `bson:"currentHeight"`
	CurrRound  uint64 `bson:"currentRound"`
	CurrStep   string `bson:"currentStep"`
	TargetStep string `bson:"targetStep"`
}

type EventSendVote struct {
	BaseEvent       `bson:",inline"`
	Vote            *core.Vote `bson:"vote"`
	RecipientPeer   string     `bson:"recipientPeer"`
	RecipientPeerId string     `bson:"recipientPeerId"`
}

type EventReceiveVote struct {
	BaseEvent    `bson:",inline"`
	Vote         *core.Vote `bson:"vote"`
	SourcePeer   string     `bson:"sourcePeer"`
	SourcePeerId string     `bson:"sourcePeerId"`
}

type EventReceivedProposal struct {
	BaseEvent `bson:",inline"`
	Proposal  *core.Proposal `bson:"proposal"`
	Proposer  string         `bson:"proposer"`
}

type EventReceivedCompleteProposalBlock struct {
	BaseEvent `bson:",inline"`
	Hash      string `bson:"hash"`
	Height    uint64 `bson:"height"`
}

type EventScheduledTimeout struct {
	BaseEvent `bson:",inline"`
	Duration  string `bson:"duration"`
	Height    uint64 `bson:"height"`
	Step      string `bson:"step"`
}
