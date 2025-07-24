package events

import (
	"github.com/zsystm/cometbft-types/pkg/core"
	"time"
)

type RecipientInfo struct {
	RecipientPeer   string `bson:"recipientPeer" json:"recipientPeer"`
	RecipientPeerId string `bson:"recipientPeerId" json:"recipientPeerId"`
}

type SourceInfo struct {
	SourcePeer   string `bson:"sourcePeer" json:"sourcePeer"`
	SourcePeerId string `bson:"sourcePeerId" json:"sourcePeerId"`
}

type StepChangeInfo struct {
	CurrHeight uint64 `bson:"currentHeight" json:"currentHeight"`
	CurrRound  uint64 `bson:"currentRound" json:"currentRound"`
	CurrStep   string `bson:"currentStep" json:"currentStep"`
	TargetStep string `bson:"targetStep" json:"targetStep"`
}

type EventEnteringNewRound struct {
	BaseEvent  `bson:",inline" json:",inline"`
	Height     uint64 `bson:"height" json:"height"`
	Round      uint64 `bson:"round" json:"round"`
	Proposer   string `bson:"proposer" json:"proposer"`
	PrevHeight uint64 `bson:"previousHeight" json:"previousHeight"`
	PrevRound  uint64 `bson:"previousRound" json:"previousRound"`
	PrevStep   string `bson:"previousStep" json:"previousStep"`
}

type EventProposeStep struct {
	BaseEvent `bson:",inline" json:",inline"`
	Height    uint64 `bson:"height" json:"height"`
	Round     uint64 `bson:"round" json:"round"`
	Proposer  string `bson:"proposer" json:"proposer"`
	IsOurTurn bool   `bson:"isOurTurn" json:"isOurTurn"`
}

type EventEnteringPrevoteStep struct {
	BaseEvent      `bson:",inline" json:",inline"`
	StepChangeInfo `bson:",inline" json:",inline"`
}

type EventEnteringPrecommitStep struct {
	BaseEvent      `bson:",inline" json:",inline"`
	StepChangeInfo `bson:",inline" json:",inline"`
}

type EventEnteringPrevoteWaitStep struct {
	BaseEvent      `bson:",inline" json:",inline"`
	StepChangeInfo `bson:",inline" json:",inline"`
}

type EventEnteringPrecommitWaitStep struct {
	BaseEvent      `bson:",inline" json:",inline"`
	StepChangeInfo `bson:",inline" json:",inline"`
}

type EventEnteringCommitStep struct {
	BaseEvent      `bson:",inline" json:",inline"`
	StepChangeInfo `bson:",inline" json:",inline"`
}

type EventEnteringWaitStep struct {
	BaseEvent      `bson:",inline" json:",inline"`
	StepChangeInfo `bson:",inline" json:",inline"`
}

type EventSendVote struct {
	BaseEvent     `bson:",inline" json:",inline"`
	Vote          *core.Vote `bson:"vote" json:"vote"`
	RecipientInfo `bson:",inline" json:",inline"`
}

type EventSendNewRoundStep struct {
	BaseEvent             `bson:",inline" json:",inline"`
	Height                int64  `bson:"height" json:"height"`
	Round                 int32  `bson:"round" json:"round"`
	Step                  string `bson:"step" json:"step"`
	SecondsSinceStartTime int64  `bson:"secondsSinceStartTime" json:"secondsSinceStartTime"`
	LastCommitRound       int32  `bson:"lastCommitRound" json:"lastCommitRound"`
	RecipientInfo         `bson:",inline" json:",inline"`
}

type EventSendNewValidBlock struct {
	BaseEvent          `bson:",inline" json:",inline"`
	Height             int64              `bson:"height" json:"height"`
	Round              int32              `bson:"round" json:"round"`
	BlockPartSetHeader core.PartSetHeader `bson:"blockPartSetHeader" json:"blockPartSetHeader"`
	BlockParts         core.BitArray      `bson:"blockParts" json:"blockParts"`
	IsCommit           bool               `bson:"isCommit" json:"isCommit"`
	RecipientInfo      `bson:",inline" json:",inline"`
}

type EventSendProposal struct {
	BaseEvent     `bson:",inline" json:",inline"`
	Type          string       `bson:"type" json:"type"`
	Height        int64        `bson:"height" json:"height"`
	Round         int32        `bson:"round" json:"round"`
	PolRound      int32        `bson:"polRound" json:"polRound"`
	BlockID       core.BlockID `bson:"blockId" json:"blockId"`
	Timestamp     time.Time    `bson:"timestamp" json:"timestamp"`
	Signature     string       `bson:"signature" json:"signature"`
	RecipientInfo `bson:",inline" json:",inline"`
}

type EventSendProposalPOL struct {
	BaseEvent        `bson:",inline" json:",inline"`
	Height           int64         `bson:"height" json:"height"`
	ProposalPolRound int32         `bson:"proposalPolRound" json:"proposalPolRound"`
	ProposalPol      core.BitArray `bson:"proposalPol" json:"proposalPol"`
	RecipientInfo    `bson:",inline" json:",inline"`
}

type EventSendBlockPart struct {
	BaseEvent     `bson:",inline" json:",inline"`
	Height        int64     `bson:"height" json:"height"`
	Round         int32     `bson:"round" json:"round"`
	Part          core.Part `bson:"part" json:"part"`
	RecipientInfo `bson:",inline" json:",inline"`
}

type EventSendHasVote struct {
	BaseEvent     `bson:",inline" json:",inline"`
	Height        int64  `bson:"height" json:"height"`
	Round         int32  `bson:"round" json:"round"`
	Type          string `bson:"type" json:"type"`
	Index         int32  `bson:"index" json:"index"`
	RecipientInfo `bson:",inline" json:",inline"`
}

type EventSendVoteSetMaj23 struct {
	BaseEvent     `bson:",inline" json:",inline"`
	Height        int64        `bson:"height" json:"height"`
	Round         int32        `bson:"round" json:"round"`
	Type          string       `bson:"type" json:"type"`
	BlockID       core.BlockID `bson:"blockId" json:"blockId"`
	RecipientInfo `bson:",inline" json:",inline"`
}

type EventSendVoteSetBits struct {
	BaseEvent     `bson:",inline" json:",inline"`
	Height        int64         `bson:"height" json:"height"`
	Round         int32         `bson:"round" json:"round"`
	Type          string        `bson:"type" json:"type"`
	BlockID       core.BlockID  `bson:"blockId" json:"blockId"`
	Votes         core.BitArray `bson:"bits" json:"votes"`
	RecipientInfo `bson:",inline" json:",inline"`
}

type EventSendHasProposalBlockPart struct {
	BaseEvent     `bson:",inline" json:",inline"`
	Height        int64 `bson:"height" json:"height"`
	Round         int32 `bson:"round" json:"round"`
	Index         int32 `bson:"index" json:"index"`
	RecipientInfo `bson:",inline" json:",inline"`
}

type EventReceivePacketNewRoundStep struct {
	BaseEvent             `bson:",inline" json:",inline"`
	Height                int64  `bson:"height" json:"height"`
	Round                 int32  `bson:"round" json:"round"`
	Step                  string `bson:"step" json:"step"`
	SecondsSinceStartTime int64  `bson:"secondsSinceStartTime" json:"secondsSinceStartTime"`
	LastCommitRound       int32  `bson:"lastCommitRound" json:"lastCommitRound"`
	SourceInfo            `bson:",inline" json:",inline"`
}

type EventReceivePacketNewValidBlock struct {
	BaseEvent          `bson:",inline" json:",inline"`
	Height             int64              `bson:"height" json:"height"`
	Round              int32              `bson:"round" json:"round"`
	BlockPartSetHeader core.PartSetHeader `bson:"blockPartSetHeader" json:"blockPartSetHeader"`
	BlockParts         core.BitArray      `bson:"blockParts" json:"blockParts"`
	IsCommit           bool               `bson:"isCommit" json:"isCommit"`
	SourceInfo         `bson:",inline" json:",inline"`
}

type EventReceivePacketProposal struct {
	BaseEvent  `bson:",inline" json:",inline"`
	Type       string       `bson:"type" json:"type"`
	Height     int64        `bson:"height" json:"height"`
	Round      int32        `bson:"round" json:"round"`
	PolRound   int32        `bson:"polRound" json:"polRound"`
	BlockID    core.BlockID `bson:"blockId" json:"blockId"`
	Timestamp  time.Time    `bson:"timestamp" json:"timestamp"`
	Signature  string       `bson:"signature" json:"signature"`
	SourceInfo `bson:",inline" json:",inline"`
}

type EventReceivePacketProposalPOL struct {
	BaseEvent        `bson:",inline" json:",inline"`
	Height           int64         `bson:"height" json:"height"`
	ProposalPolRound int32         `bson:"proposalPolRound" json:"proposalPolRound"`
	ProposalPol      core.BitArray `bson:"proposalPol" json:"proposalPol"`
	SourceInfo       `bson:",inline" json:",inline"`
}

type EventReceivePacketBlockPart struct {
	BaseEvent  `bson:",inline" json:",inline"`
	Height     int64     `bson:"height" json:"height"`
	Round      int32     `bson:"round" json:"round"`
	Part       core.Part `bson:"part" json:"part"`
	SourceInfo `bson:",inline" json:",inline"`
}

type EventReceivePacketHasVote struct {
	BaseEvent  `bson:",inline" json:",inline"`
	Height     int64  `bson:"height" json:"height"`
	Round      int32  `bson:"round" json:"round"`
	Type       string `bson:"type" json:"type"`
	Index      int32  `bson:"index" json:"index"`
	SourceInfo `bson:",inline" json:",inline"`
}

type EventReceivePacketVote struct {
	BaseEvent  `bson:",inline" json:",inline"`
	Vote       *core.Vote `bson:"vote" json:"vote"`
	SourceInfo `bson:",inline" json:",inline"`
}

type EventReceivePacketVoteSetMaj23 struct {
	BaseEvent  `bson:",inline" json:",inline"`
	Height     int64        `bson:"height" json:"height"`
	Round      int32        `bson:"round" json:"round"`
	Type       string       `bson:"type" json:"type"`
	BlockID    core.BlockID `bson:"blockId" json:"blockId"`
	SourceInfo `bson:",inline" json:",inline"`
}

type EventReceivePacketVoteSetBits struct {
	BaseEvent  `bson:",inline" json:",inline"`
	Height     int64         `bson:"height" json:"height"`
	Round      int32         `bson:"round" json:"round"`
	Type       string        `bson:"type" json:"type"`
	BlockID    core.BlockID  `bson:"blockId" json:"blockId"`
	Votes      core.BitArray `bson:"bits" json:"votes"`
	SourceInfo `bson:",inline" json:",inline"`
}

type EventReceivePacketHasProposalBlockPart struct {
	BaseEvent  `bson:",inline" json:",inline"`
	Height     int64 `bson:"height" json:"height"`
	Round      int32 `bson:"round" json:"round"`
	Index      int32 `bson:"index" json:"index"`
	SourceInfo `bson:",inline" json:",inline"`
}
type EventReceivedProposal struct {
	BaseEvent `bson:",inline" json:",inline"`
	Proposal  *core.Proposal `bson:"proposal" json:"proposal"`
	Proposer  string         `bson:"proposer" json:"proposer"`
}

type EventReceivedCompleteProposalBlock struct {
	BaseEvent `bson:",inline" json:",inline"`
	Hash      string `bson:"hash" json:"hash"`
	Height    uint64 `bson:"height" json:"height"`
}

type EventScheduledTimeout struct {
	BaseEvent `bson:",inline" json:",inline"`
	Duration  string `bson:"duration" json:"duration"`
	Height    uint64 `bson:"height" json:"height"`
	Step      string `bson:"step" json:"step"`
}
