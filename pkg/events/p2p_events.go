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

// EventP2pBlockPart represents a confirmed block part message exchange
type EventP2pBlockPart struct {
	BaseEvent `bson:",inline"`
	Height    int64     `json:"height" bson:"height"`
	Round     int32     `json:"round" bson:"round"`
	Part      core.Part `json:"part" bson:"part"`
	P2pInfo   `json:",inline" bson:",inline"`
}

// EventP2pProposal represents a confirmed proposal message exchange
type EventP2pProposal struct {
	BaseEvent `bson:",inline"`
	Type      string       `json:"type" bson:"type"`
	Height    int64        `json:"height" bson:"height"`
	Round     int32        `json:"round" bson:"round"`
	PolRound  int32        `json:"polRound" bson:"polRound"`
	BlockID   core.BlockID `json:"blockId" bson:"blockId"`
	Timestamp time.Time    `json:"proposalTimestamp" bson:"proposalTimestamp"`
	Signature string       `json:"signature" bson:"signature"`
	P2pInfo   `json:",inline" bson:",inline"`
}

// EventP2pProposalPOL represents a confirmed proposal POL message exchange
type EventP2pProposalPOL struct {
	BaseEvent        `bson:",inline"`
	Height           int64         `json:"height" bson:"height"`
	ProposalPolRound int32         `json:"proposalPolRound" bson:"proposalPolRound"`
	ProposalPol      core.BitArray `json:"proposalPol" bson:"proposalPol"`
	P2pInfo          `json:",inline" bson:",inline"`
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

// EventP2pHasVote represents a confirmed has vote message exchange
type EventP2pHasVote struct {
	BaseEvent `bson:",inline"`
	Height    int64  `json:"height" bson:"height"`
	Round     int32  `json:"round" bson:"round"`
	Type      string `json:"type" bson:"type"`
	Index     int32  `json:"index" bson:"index"`
	P2pInfo   `json:",inline" bson:",inline"`
}

// EventP2pVoteSetMaj23 represents a confirmed vote set majority message exchange
type EventP2pVoteSetMaj23 struct {
	BaseEvent `bson:",inline"`
	Height    int64         `json:"height" bson:"height"`
	Round     int32         `json:"round" bson:"round"`
	Type      string        `json:"type" bson:"type"`
	BlockID   *core.BlockID `json:"blockId" bson:"blockId"`
	P2pInfo   `json:",inline" bson:",inline"`
}

// EventP2pVoteSetBits represents a confirmed vote set bits message exchange
type EventP2pVoteSetBits struct {
	BaseEvent `bson:",inline"`
	Height    int64         `json:"height" bson:"height"`
	Round     int32         `json:"round" bson:"round"`
	Type      string        `json:"type" bson:"type"`
	BlockID   *core.BlockID `json:"blockId" bson:"blockId"`
	Votes     core.BitArray `json:"votes" bson:"votes"`
	P2pInfo   `json:",inline" bson:",inline"`
}
