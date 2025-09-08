package events

import (
	"time"

	"github.com/zsystm/cometbft-types/pkg/core"
)

// P2pMessageStatus represents the status of a P2P message
type P2pMessageStatus string

const (
	P2pMsgStatusSent      P2pMessageStatus = "sent"
	P2pMsgStatusReceived  P2pMessageStatus = "received"
	P2pMsgStatusConfirmed P2pMessageStatus = "confirmed"
)

// P2pInfo contains metadata shared by all P2P message events.
type P2pInfo struct {
	Status          P2pMessageStatus `json:"status" bson:"status"`
	SenderPeerId    string           `json:"senderPeerId" bson:"senderPeerId"`
	RecipientPeerId string           `json:"recipientPeerId" bson:"recipientPeerId"`
	SentTime        time.Time        `json:"sentTime" bson:"sentTime"`
	ReceivedTime    time.Time        `json:"receivedTime" bson:"receivedTime"`
	Latency         time.Duration    `json:"latency" bson:"latency"`
}

// Universal P2P events that contain both send and receive information
// These events are only created when we have matching send/receive pairs

// EventP2pVote represents a confirmed vote message exchange
type EventP2pVote struct {
	BaseEvent `bson:",inline"`
	Vote      *core.Vote `json:"vote" bson:"vote"`
	P2pInfo   `json:",inline" bson:",inline"`
}

func (e *EventP2pVote) GetHeight() uint64 {
	if e.Vote != nil {
		return e.Vote.Height
	}
	return 0
}

func (e *EventP2pVote) GetRound() uint64 {
	if e.Vote != nil {
		return e.Vote.Round
	}
	return 0
}

// EventP2pBlockPart represents a confirmed block part message exchange
type EventP2pBlockPart struct {
	BaseEvent `bson:",inline"`
	Height    int64     `json:"height" bson:"height"`
	Round     int32     `json:"round" bson:"round"`
	Part      core.Part `json:"part" bson:"part"`
	P2pInfo   `json:",inline" bson:",inline"`
}

func (e *EventP2pBlockPart) GetHeight() uint64 {
	return uint64(e.Height)
}

func (e *EventP2pBlockPart) GetRound() uint64 {
	return uint64(e.Round)
}

// EventP2pProposal represents a confirmed proposal message exchange
type EventP2pProposal struct {
	BaseEvent `bson:",inline"`
	Height    int64        `json:"height" bson:"height"`
	Round     int32        `json:"round" bson:"round"`
	PolRound  int32        `json:"polRound" bson:"polRound"`
	BlockID   core.BlockID `json:"blockId" bson:"blockId"`
	Timestamp time.Time    `json:"proposalTimestamp" bson:"proposalTimestamp"`
	Signature string       `json:"signature" bson:"signature"`
	P2pInfo   `json:",inline" bson:",inline"`
}

func (e *EventP2pProposal) GetHeight() uint64 {
	return uint64(e.Height)
}

func (e *EventP2pProposal) GetRound() uint64 {
	return uint64(e.Round)
}

// EventP2pProposalPOL represents a confirmed proposal POL message exchange
type EventP2pProposalPOL struct {
	BaseEvent        `bson:",inline"`
	Height           int64         `json:"height" bson:"height"`
	ProposalPolRound int32         `json:"proposalPolRound" bson:"proposalPolRound"`
	ProposalPol      core.BitArray `json:"proposalPol" bson:"proposalPol"`
	P2pInfo          `json:",inline" bson:",inline"`
}

func (e *EventP2pProposalPOL) GetHeight() uint64 {
	return uint64(e.Height)
}

// EventP2pNewRoundStep represents a confirmed new round step message exchange
type EventP2pNewRoundStep struct {
	BaseEvent             `bson:",inline"`
	Height                int64  `json:"height" bson:"height"`
	Round                 int32  `json:"round" bson:"round"`
	Step                  string `json:"step" bson:"step"`
	SecondsSinceStartTime int32  `json:"secondsSinceStartTime" bson:"secondsSinceStartTime"`
	LastCommitRound       int32  `json:"lastCommitRound" bson:"lastCommitRound"`
	P2pInfo               `json:",inline" bson:",inline"`
}

func (e *EventP2pNewRoundStep) GetHeight() uint64 {
	return uint64(e.Height)
}

func (e *EventP2pNewRoundStep) GetRound() uint64 {
	return uint64(e.Round)
}

// EventP2pHasVote represents a confirmed has vote message exchange
type EventP2pHasVote struct {
	BaseEvent `bson:",inline"`
	Type      string `json:"type" bson:"type"`
	Height    int64  `json:"height" bson:"height"`
	Round     int32  `json:"round" bson:"round"`
	Index     int32  `json:"index" bson:"index"`
	P2pInfo   `json:",inline" bson:",inline"`
}

func (e *EventP2pHasVote) GetHeight() uint64 {
	return uint64(e.Height)
}

func (e *EventP2pHasVote) GetRound() uint64 {
	return uint64(e.Round)
}

// EventP2pVoteSetMaj23 represents a confirmed vote set majority message exchange
type EventP2pVoteSetMaj23 struct {
	BaseEvent `bson:",inline"`
	Type      string        `json:"type" bson:"type"`
	Height    int64         `json:"height" bson:"height"`
	Round     int32         `json:"round" bson:"round"`
	BlockID   *core.BlockID `json:"blockId" bson:"blockId"`
	P2pInfo   `json:",inline" bson:",inline"`
}

func (e *EventP2pVoteSetMaj23) GetHeight() uint64 {
	return uint64(e.Height)
}

func (e *EventP2pVoteSetMaj23) GetRound() uint64 {
	return uint64(e.Round)
}

// EventP2pVoteSetBits represents a confirmed vote set bits message exchange
type EventP2pVoteSetBits struct {
	BaseEvent `bson:",inline"`
	Type      string        `json:"type" bson:"type"`
	Height    int64         `json:"height" bson:"height"`
	Round     int32         `json:"round" bson:"round"`
	BlockID   *core.BlockID `json:"blockId" bson:"blockId"`
	Votes     core.BitArray `json:"votes" bson:"votes"`
	P2pInfo   `json:",inline" bson:",inline"`
}

func (e *EventP2pVoteSetBits) GetHeight() uint64 {
	return uint64(e.Height)
}

func (e *EventP2pVoteSetBits) GetRound() uint64 {
	return uint64(e.Round)
}
