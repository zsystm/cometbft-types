package events

import (
	"github.com/zsystm/cometbft-types/pkg/core"
	"time"
)

type RecipientInfo struct {
	RecipientPeer   string `bson:"recipientPeer"`
	RecipientPeerId string `bson:"recipientPeerId"`
}

type SourceInfo struct {
	SourcePeer   string `bson:"sourcePeer"`
	SourcePeerId string `bson:"sourcePeerId"`
}

type StepChangeInfo struct {
	CurrHeight uint64 `bson:"currentHeight"`
	CurrRound  uint64 `bson:"currentRound"`
	CurrStep   string `bson:"currentStep"`
	TargetStep string `bson:"targetStep"`
}

type EventEnteringNewRound struct {
	BaseEvent  `bson:",inline"`
	Height     uint64 `bson:"height"`
	Round      uint64 `bson:"round"`
	Proposer   string `bson:"proposer"`
	PrevHeight uint64 `bson:"previousHeight"`
	PrevRound  uint64 `bson:"previousRound"`
	PrevStep   string `bson:"previousStep"`
}

type EventProposeStep struct {
	BaseEvent `bson:",inline"`
	Height    uint64 `bson:"height"`
	Round     uint64 `bson:"round"`
	Proposer  string `bson:"proposer"`
	IsOurTurn bool   `bson:"isOurTurn"`
}

type EventEnteringPrevoteStep struct {
	BaseEvent      `bson:",inline"`
	StepChangeInfo `bson:",inline"`
}

type EventEnteringPrecommitStep struct {
	BaseEvent      `bson:",inline"`
	StepChangeInfo `bson:",inline"`
}

type EventEnteringPrevoteWaitStep struct {
	BaseEvent      `bson:",inline"`
	StepChangeInfo `bson:",inline"`
}

type EventEnteringPrecommitWaitStep struct {
	BaseEvent      `bson:",inline"`
	StepChangeInfo `bson:",inline"`
}

type EventEnteringCommitStep struct {
	BaseEvent      `bson:",inline"`
	StepChangeInfo `bson:",inline"`
}

type EventEnteringWaitStep struct {
	BaseEvent      `bson:",inline"`
	StepChangeInfo `bson:",inline"`
}

type EventSendVote struct {
	BaseEvent     `bson:",inline"`
	Vote          *core.Vote `bson:"vote"`
	RecipientInfo `bson:",inline"`
}

type EventSendNewRoundStep struct {
	BaseEvent             `bson:",inline"`
	Height                int64  `bson:"height"`
	Round                 int32  `bson:"round"`
	Step                  string `bson:"step"`
	SecondsSinceStartTime int64  `bson:"secondsSinceStartTime"`
	LastCommitRound       int32  `bson:"lastCommitRound"`
	RecipientInfo         `bson:",inline"`
}

type EventSendNewValidBlock struct {
	BaseEvent          `bson:",inline"`
	Height             int64              `bson:"height"`
	Round              int32              `bson:"round"`
	BlockPartSetHeader core.PartSetHeader `bson:"blockPartSetHeader"`
	BlockParts         core.BitArray      `bson:"blockParts"`
	IsCommit           bool               `bson:"isCommit"`
	RecipientInfo      `bson:",inline"`
}

type EventSendProposal struct {
	BaseEvent     `bson:",inline"`
	Type          string       `bson:"type"`
	Height        int64        `bson:"height"`
	Round         int32        `bson:"round"`
	PolRound      int32        `bson:"polRound"`
	BlockID       core.BlockID `bson:"blockId"`
	Timestamp     time.Time    `bson:"timestamp"`
	Signature     string       `bson:"signature"`
	RecipientInfo `bson:",inline"`
}

type EventSendProposalPOL struct {
	BaseEvent        `bson:",inline"`
	Height           int64         `bson:"height"`
	ProposalPolRound int32         `bson:"proposalPolRound"`
	ProposalPol      core.BitArray `bson:"proposalPol"`
	RecipientInfo    `bson:",inline"`
}

type EventSendBlockPart struct {
	BaseEvent     `bson:",inline"`
	Height        int64     `bson:"height"`
	Round         int32     `bson:"round"`
	Part          core.Part `bson:"part"`
	RecipientInfo `bson:",inline"`
}

type EventSendHasVote struct {
	BaseEvent     `bson:",inline"`
	Height        int64  `bson:"height"`
	Round         int32  `bson:"round"`
	Type          string `bson:"type"`
	Index         int32  `bson:"index"`
	RecipientInfo `bson:",inline"`
}

type EventSendVoteSetMaj23 struct {
	BaseEvent     `bson:",inline"`
	Height        int64        `bson:"height"`
	Round         int32        `bson:"round"`
	Type          string       `bson:"type"`
	BlockID       core.BlockID `bson:"blockId"`
	RecipientInfo `bson:",inline"`
}

type EventSendVoteSetBits struct {
	BaseEvent     `bson:",inline"`
	Height        int64         `bson:"height"`
	Round         int32         `bson:"round"`
	Type          string        `bson:"type"`
	BlockID       core.BlockID  `bson:"blockId"`
	Votes         core.BitArray `bson:"bits"`
	RecipientInfo `bson:",inline"`
}

type EventSendHasProposalBlockPart struct {
	BaseEvent     `bson:",inline"`
	Height        int64 `bson:"height"`
	Round         int32 `bson:"round"`
	Index         int32 `bson:"index"`
	RecipientInfo `bson:",inline"`
}

type EventReceivePacketNewRoundStep struct {
	BaseEvent             `bson:",inline"`
	Height                int64  `bson:"height"`
	Round                 int32  `bson:"round"`
	Step                  string `bson:"step"`
	SecondsSinceStartTime int64  `bson:"secondsSinceStartTime"`
	LastCommitRound       int32  `bson:"lastCommitRound"`
	SourceInfo            `bson:",inline"`
}

type EventReceivePacketNewValidBlock struct {
	BaseEvent          `bson:",inline"`
	Height             int64              `bson:"height"`
	Round              int32              `bson:"round"`
	BlockPartSetHeader core.PartSetHeader `bson:"blockPartSetHeader"`
	BlockParts         core.BitArray      `bson:"blockParts"`
	IsCommit           bool               `bson:"isCommit"`
	SourceInfo         `bson:",inline"`
}

type EventReceivePacketProposal struct {
	BaseEvent  `bson:",inline"`
	Type       string       `bson:"type"`
	Height     int64        `bson:"height"`
	Round      int32        `bson:"round"`
	PolRound   int32        `bson:"polRound"`
	BlockID    core.BlockID `bson:"blockId"`
	Timestamp  time.Time    `bson:"timestamp"`
	Signature  string       `bson:"signature"`
	SourceInfo `bson:",inline"`
}

type EventReceivePacketProposalPOL struct {
	BaseEvent        `bson:",inline"`
	Height           int64         `bson:"height"`
	ProposalPolRound int32         `bson:"proposalPolRound"`
	ProposalPol      core.BitArray `bson:"proposalPol"`
	SourceInfo       `bson:",inline"`
}

type EventReceivePacketBlockPart struct {
	BaseEvent  `bson:",inline"`
	Height     int64     `bson:"height"`
	Round      int32     `bson:"round"`
	Part       core.Part `bson:"part"`
	SourceInfo `bson:",inline"`
}

type EventReceivePacketHasVote struct {
	BaseEvent  `bson:",inline"`
	Height     int64  `bson:"height"`
	Round      int32  `bson:"round"`
	Type       string `bson:"type"`
	Index      int32  `bson:"index"`
	SourceInfo `bson:",inline"`
}

type EventReceivePacketVote struct {
	BaseEvent  `bson:",inline"`
	Vote       *core.Vote `bson:"vote"`
	SourceInfo `bson:",inline"`
}

type EventReceivePacketVoteSetMaj23 struct {
	BaseEvent  `bson:",inline"`
	Height     int64        `bson:"height"`
	Round      int32        `bson:"round"`
	Type       string       `bson:"type"`
	BlockID    core.BlockID `bson:"blockId"`
	SourceInfo `bson:",inline"`
}

type EventReceivePacketVoteSetBits struct {
	BaseEvent  `bson:",inline"`
	Height     int64         `bson:"height"`
	Round      int32         `bson:"round"`
	Type       string        `bson:"type"`
	BlockID    core.BlockID  `bson:"blockId"`
	Votes      core.BitArray `bson:"bits"`
	SourceInfo `bson:",inline"`
}

type EventReceivePacketHasProposalBlockPart struct {
	BaseEvent  `bson:",inline"`
	Height     int64 `bson:"height"`
	Round      int32 `bson:"round"`
	Index      int32 `bson:"index"`
	SourceInfo `bson:",inline"`
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
