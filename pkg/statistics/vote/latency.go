package vote

import (
	"github.com/zsystm/cometbft-types/pkg/core"
	"time"
)

type VoteMsgStatus string

const (
	// Sender has sent the message, but it has not yet been received by the recipient.
	VoteMsgStatusSent VoteMsgStatus = "sent"
	// Sender has sent the message, and it has been received by the recipient, but not yet confirmed.
	VoteMsgStatusReceived VoteMsgStatus = "received"
	// Sender has sent the message, and it has been received by the recipient.
	VoteMsgStatusConfirmed VoteMsgStatus = "confirmed"
)

type VoteLatency struct {
	Status          VoteMsgStatus `bson:"status"`
	Vote            *core.Vote    `bson:"vote"`
	SenderPeerId    string        `bson:"senderPeerId"`
	RecipientPeerId string        `bson:"recipientPeerId"`
	SentTime        time.Time     `bson:"sentTime"`
	ReceivedTime    time.Time     `bson:"receivedTime"`
	ConfirmedTime   time.Time     `bson:"confirmedTime"`
	Latency         time.Duration `bson:"latency"`
}
