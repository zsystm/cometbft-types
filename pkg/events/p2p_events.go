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
	SenderPeerId    string           `json:"sender_peer_id" bson:"sender_peer_id"`
	RecipientPeerId string           `json:"recipient_peer_id" bson:"recipient_peer_id"`
	SentTime        time.Time        `json:"sent_time" bson:"sent_time"`
	ReceivedTime    time.Time        `json:"received_time" bson:"received_time"`
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
	PolRound  int32        `json:"pol_round" bson:"pol_round"`
	BlockID   core.BlockID `json:"block_id" bson:"block_id"`
	Timestamp time.Time    `json:"proposal_timestamp" bson:"proposal_timestamp"`
	Signature string       `json:"signature" bson:"signature"`
	P2pInfo   `json:",inline" bson:",inline"`
}

// EventP2pProposalPOL represents a confirmed proposal POL message exchange
type EventP2pProposalPOL struct {
	BaseEvent        `bson:",inline"`
	Height           int64         `json:"height" bson:"height"`
	ProposalPolRound int32         `json:"proposal_pol_round" bson:"proposal_pol_round"`
	ProposalPol      core.BitArray `json:"proposal_pol" bson:"proposal_pol"`
	P2pInfo          `json:",inline" bson:",inline"`
}

// EventP2pNewRoundStep represents a confirmed new round step message exchange
type EventP2pNewRoundStep struct {
	BaseEvent             `bson:",inline"`
	Height                int64  `json:"height" bson:"height"`
	Round                 int32  `json:"round" bson:"round"`
	Step                  string `json:"step" bson:"step"`
	SecondsSinceStartTime int32  `json:"seconds_since_start_time" bson:"seconds_since_start_time"`
	LastCommitRound       int32  `json:"last_commit_round" bson:"last_commit_round"`
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
	BlockID   *core.BlockID `json:"block_id" bson:"block_id"`
	P2pInfo   `json:",inline" bson:",inline"`
}

// EventP2pVoteSetBits represents a confirmed vote set bits message exchange
type EventP2pVoteSetBits struct {
	BaseEvent `bson:",inline"`
	Height    int64         `json:"height" bson:"height"`
	Round     int32         `json:"round" bson:"round"`
	Type      string        `json:"type" bson:"type"`
	BlockID   *core.BlockID `json:"block_id" bson:"block_id"`
	Votes     core.BitArray `json:"votes" bson:"votes"`
	P2pInfo   `json:",inline" bson:",inline"`
}
